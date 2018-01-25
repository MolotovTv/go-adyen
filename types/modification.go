// Copyright 2018 Matthieu CORNUT-CHAUVINC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package types

// ModificationCaptureParams struct
type ModificationCaptureParams struct {
	MerchantAccount    string  `json:"merchantAccount"`
	ModificationAmount *Amount `json:"modificationAmount"`
	OriginalReference  string  `json:"originalReference"`
	Reference          string  `json:"reference"`
}

// ModificationCancelParams struct
type ModificationCancelParams struct {
	MerchantAccount   string `json:"merchantAccount"`
	OriginalReference string `json:"originalReference"`
	Reference         string `json:"reference"`
}

// ModificationRefundParams struct
type ModificationRefundParams struct {
	MerchantAccount    string  `json:"merchantAccount"`
	ModificationAmount *Amount `json:"modificationAmount"`
	Reference          string  `json:"reference"`
	OriginalReference  string  `json:"originalReference"`
}

// ModificationCancelOrRefundParams struct
type ModificationCancelOrRefundParams struct {
	MerchantAccount   string `json:"merchantAccount"`
	OriginalReference string `json:"originalReference"`
	Reference         string `json:"reference"`
}

// ModificationAdjustAuthoriseParams struct
type ModificationAdjustAuthoriseParams struct {
	AdditionalData     *AdditionalData `json:"additionalData"`
	MerchantAccount    string          `json:"merchantAccount"`
	ModificationAmount *Amount         `json:"modificationAmount"`
	OriginalReference  string          `json:"originalReference"`
	Reference          string          `json:"reference"`
}

/********************************************/
/*            RESPONSE FROM ADYEN           */
/********************************************/

// ModificationCapture is the resource representing a Adyen payment.
type ModificationCapture struct {
	// FIXME with good response from adyen
	Todo string `json:"todo"`
}

// ModificationCancel is the resource representing a Adyen payment.
type ModificationCancel struct {
	// FIXME with good response from adyen
	Todo string `json:"todo"`
}

// ModificationRefund struct
type ModificationRefund struct {
	// FIXME with good response from adyen
	Todo string `json:"todo"`
}

// ModificationCancelOrRefund struct
type ModificationCancelOrRefund struct {
	// FIXME with good response from adyen
	Todo string `json:"todo"`
}

// ModificationAdjustAuthorise struct
type ModificationAdjustAuthorise struct {
	// FIXME with good response from adyen
	Todo string `json:"todo"`
}
