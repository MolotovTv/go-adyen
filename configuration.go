// Copyright 2018 Matthieu CORNUT-CHAUVINC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package adyen

import (
	"fmt"
	"net/http"
	"net/url"
)

// Configuration represents the MySQL configuration
type Configuration struct {
	Password  string `toml:"password" json:"password"`
	Username  string `toml:"username" json:"username"`
	Live      bool   `toml:"live" json:"live"`
	ProxyHost string `toml:"proxy_host" json:"proxy_host"`
	ProxyPort int    `toml:"proxy_port" json:"proxy_port"`
}

// NewConfiguration generates a Configuration
func NewConfiguration() Configuration {
	return Configuration{}
}

func (c *Configuration) GetHttpProxy() (func(*http.Request) (*url.URL, error), error) {
	proxyURL, err := url.Parse(fmt.Sprintf("%s:%d", c.ProxyHost, c.ProxyPort))

	if err != nil {
		return nil, err
	}
	return http.ProxyURL(proxyURL), nil
}
