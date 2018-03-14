// Copyright 2018 Matthieu CORNUT-CHAUVINC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package adyen

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/mcornut/go-adyen/constants"
	log "github.com/molotovtv/go-logger"
	"github.com/pkg/errors"
)

type Client struct {
	config     *Configuration
	HttpClient *http.Client
}

var debug bool

// SetDebug enables additional tracing globally.
// The method is designed for used during testing.
func (c *Client) SetDebug(value bool) {
	debug = value
}

func (c *Client) getURL() string {
	if c.config.Live {
		return constants.LiveURL
	}
	return constants.SandboxURL
}

func New(config *Configuration) *Client {

	transport := &http.Transport{}

	if proxy, err := config.GetHttpProxy(); err != nil {
		log.Error(
			errors.Wrap(err, "Adyen - Client - Proxy configuration"),
		)
	} else {
		transport.Proxy = proxy
	}

	return &Client{
		config: config,
		HttpClient: &http.Client{
			Transport: transport,
		},
	}
}

// Call is used by Call to generate an http.Request. It handles encoding parameters and attaching the Authorization header.
func (c *Client) call(method string, path string, body interface{}, v interface{}) error {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	var b io.Reader
	path = c.getURL() + path

	if debug {
		log.Debugf("Call %v\n", path)
	}

	if strings.ToUpper(method) == "GET" {
		return errors.New("Method GET not allowed.")
	} else {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return errors.Wrap(err, "Marshal body")
		}
		b = bytes.NewBuffer(jsonBody)
	}

	req, err := http.NewRequest(method, path, b)
	if err != nil {
		return errors.Wrap(err, "Cannot create Adyen request")
	}

	req.SetBasicAuth(c.config.Username, c.config.Password)
	req.Header.Add("User-Agent", "Mcornut/v1 GoAdyenSdk/"+Version)
	if body != nil {
		req.Header.Add("Content-Type", "application/json")
	}

	return c.do(req, v)
}

// do is used by Call to execute an API request and parse the response. It uses
// the environment's HTTP client to execute the request and unmarshals the response
// into v. It also handles unmarshaling errors returned by the API.
func (c *Client) do(req *http.Request, v interface{}) error {
	log.Infof("Requesting %v %q\n", req.Method, req.URL.Path)
	start := time.Now()

	res, err := c.HttpClient.Do(req)

	if debug {
		log.Debugf("Completed in %v\n", time.Since(start))
	}

	if res.StatusCode >= 400 {
		if debug {
			log.Debugf("StatusCode %v with Status %v\n", res.StatusCode, res.Status)
		}

		adyenErr := &Error{
			Type:           ErrorType(res.Status),
			HTTPStatusCode: res.StatusCode,
			Message:        res.Status,
		}

		switch adyenErr.Type {
		case ErrorTypeAPIConnection:
			adyenErr.Err = &APIConnectionError{adyenErr: adyenErr}

		case ErrorTypeAuthentication:
			adyenErr.Err = &AuthenticationError{adyenErr: adyenErr}

		case ErrorTypeInvalidRequest:
			adyenErr.Err = &InvalidRequestError{adyenErr: adyenErr}

		case ErrorTypePermission:
			adyenErr.Err = &PermissionError{adyenErr: adyenErr}

		case ErrorTypeNotFound:
			adyenErr.Err = &NotFoundError{adyenErr: adyenErr}

		case ErrorTypeUnprocessable:
			adyenErr.Err = &UnprocessableError{adyenErr: adyenErr}

		default:
			adyenErr.Err = &APIError{adyenErr: adyenErr}
		}

		return adyenErr

	}

	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Error(
			errors.Wrap(err, "Cannot parse Adyen response"),
		)
		return err
	}

	if debug {
		log.Debugf("Adyen Response: %v\n", string(resBody))
	}

	if v != nil {
		return json.Unmarshal(resBody, v)
	}
	return nil
}
