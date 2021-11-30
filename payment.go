// Copyright 2018 Matthieu CORNUT-CHAUVINC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package adyen

import (
	"fmt"

	"github.com/molotovtv/go-adyen/types"
)

// ClientPayment interface
type ClientPayment interface {
	Authorise(params *types.PaymentAuthoriseParams, version string) (*types.PaymentAuthorise, error)
	Authorise3d(params *types.PaymentAuthorise3dParams, version string) (*types.PaymentAuthorise, error)
	AuthoriseForRecurring(params *types.PaymentAuthoriseForRecurringParams, version string) (*types.PaymentAuthorise, error)
	AuthoriseRecurring(params *types.PaymentAuthoriseRecurringParams, version string) (*types.PaymentAuthorise, error)
	AuthoriseOneClick(params *types.PaymentAuthoriseOneClickParams, version string) (*types.PaymentAuthorise, error)
}

// Authorise func
func (c Client) Authorise(params *types.PaymentAuthoriseParams, version string) (*types.PaymentAuthorise, error) {
	adyenPayment := &types.PaymentAuthorise{}
	err := c.call("POST", fmt.Sprintf("/pal/servlet/Payment/%v/authorise", version), params, adyenPayment)

	return adyenPayment, err
}

// Authorise3d func
func (c Client) Authorise3d(params *types.PaymentAuthorise3dParams, version string) (*types.PaymentAuthorise, error) {
	adyenPayment3d := &types.PaymentAuthorise{}
	err := c.call("POST", fmt.Sprintf("/pal/servlet/Payment/%v/authorise3d", version), params, adyenPayment3d)

	return adyenPayment3d, err
}

// AuthoriseForRecurring func
func (c Client) AuthoriseForRecurring(params *types.PaymentAuthoriseForRecurringParams, version string) (*types.PaymentAuthorise, error) {
	adyenPayment := &types.PaymentAuthorise{}
	err := c.call("POST", fmt.Sprintf("/pal/servlet/Payment/%v/authorise", version), params, adyenPayment)

	return adyenPayment, err
}

// AuthoriseRecurring func
func (c Client) AuthoriseRecurring(params *types.PaymentAuthoriseRecurringParams, version string) (*types.PaymentAuthorise, error) {
	adyenPayment := &types.PaymentAuthorise{}
	err := c.call("POST", fmt.Sprintf("/pal/servlet/Payment/%v/authorise", version), params, adyenPayment)

	return adyenPayment, err
}

// AuthoriseOneClick func
func (c Client) AuthoriseOneClick(params *types.PaymentAuthoriseOneClickParams, version string) (*types.PaymentAuthorise, error) {
	adyenPayment := &types.PaymentAuthorise{}
	err := c.call("POST", fmt.Sprintf("/pal/servlet/Payment/%v/authorise", version), params, adyenPayment)

	return adyenPayment, err
}
