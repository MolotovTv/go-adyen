// Copyright 2018 Matthieu CORNUT-CHAUVINC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package adyen

import (
	"fmt"

	"github.com/molotovtv/go-adyen/constants"

	"github.com/molotovtv/go-adyen/types"
)

// ClientModification interface
type ClientModification interface {
	Capture(params *types.ModificationCaptureParams) (*types.ModificationCapture, error)
	Cancel(params *types.ModificationCancelParams) (*types.ModificationCancel, error)
	Refund(params *types.ModificationRefundParams) (*types.ModificationRefund, error)
	CancelOrRefund(params *types.ModificationCancelOrRefundParams) (*types.ModificationCancelOrRefund, error)
	AdjustAuthorise(params *types.ModificationAdjustAuthoriseParams) (*types.ModificationAdjustAuthorise, error)
	TechnicalCancel(params *types.ModificationTechnicalCancelParams) (*types.ModificationTechnicalCancelParams, error)
}

// Capture func
func (c Client) Capture(params *types.ModificationCaptureParams) (*types.ModificationCapture, error) {
	adyenCapture := &types.ModificationCapture{}
	err := c.call("POST", fmt.Sprintf("/pal/servlet/Payment/%v/capture", constants.APIPaymentV30), params, adyenCapture)

	return adyenCapture, err
}

// Cancel func
func (c Client) Cancel(params *types.ModificationCancelParams) (*types.ModificationCancel, error) {
	adyenCancel := &types.ModificationCancel{}
	err := c.call("POST", fmt.Sprintf("/pal/servlet/Payment/%v/cancel", constants.APIPaymentV30), params, adyenCancel)

	return adyenCancel, err
}

// Refund func
func (c Client) Refund(params *types.ModificationRefundParams) (*types.ModificationRefund, error) {
	adyenRefund := &types.ModificationRefund{}
	err := c.call("POST", fmt.Sprintf("/pal/servlet/Payment/%v/refund", constants.APIPaymentV30), params, adyenRefund)

	return adyenRefund, err
}

// CancelOrRefund func
func (c Client) CancelOrRefund(params *types.ModificationCancelOrRefundParams) (*types.ModificationCancelOrRefund, error) {
	adyenCancelOrRefund := &types.ModificationCancelOrRefund{}
	err := c.call("POST", fmt.Sprintf("/pal/servlet/Payment/%v/cancelOrRefund", constants.APIPaymentV30), params, adyenCancelOrRefund)

	return adyenCancelOrRefund, err
}

// AdjustAuthorise func
func (c Client) AdjustAuthorise(params *types.ModificationAdjustAuthoriseParams) (*types.ModificationAdjustAuthorise, error) {
	adyenAdjustAuthorise := &types.ModificationAdjustAuthorise{}
	err := c.call("POST", fmt.Sprintf("/pal/servlet/Payment/%v/adjustAuthorisation", constants.APIPaymentV30), params, adyenAdjustAuthorise)

	return adyenAdjustAuthorise, err
}

// TechnicalCancel func
func (c Client) TechnicalCancel(params *types.ModificationTechnicalCancelParams) (*types.ModificationTechnicalCancel, error) {
	adyenTechnicalCancel := &types.ModificationTechnicalCancel{}
	err := c.call("POST", fmt.Sprintf("/pal/servlet/Payment/%v/technicalCancel", constants.APIPaymentV30), params, adyenTechnicalCancel)

	return adyenTechnicalCancel, err
}
