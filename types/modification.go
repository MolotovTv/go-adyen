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

// ModificationTechnicalCancelParams struct
type ModificationTechnicalCancelParams struct {
	MerchantAccount           string `json:"merchantAccount"`
	OriginalMerchantReference string `json:"originalMerchantReference"`
	Reference                 string `json:"reference"`
}

/********************************************/
/*            RESPONSE FROM ADYEN           */
/********************************************/

// ModificationCapture is the resource representing a Adyen payment.
type ModificationCapture struct {
	PspReference string `json:"pspReference"`
	Response     string `json:"response"` // "response": "[capture-received]"
}

// ModificationCancel is the resource representing a Adyen payment.
type ModificationCancel struct {
	PspReference string `json:"pspReference"`
	Response     string `json:"response"` // "[cancel-received]"
}

// ModificationRefund struct
type ModificationRefund struct {
	PspReference string `json:"pspReference"`
	Response     string `json:"response"` //  "[refund-received]"
}

// ModificationCancelOrRefund struct
type ModificationCancelOrRefund struct {
	PspReference string `json:"pspReference"`
	Response     string `json:"response"` // "[cancelOrRefund-received]"
}

// ModificationAdjustAuthorise struct
type ModificationAdjustAuthorise struct {
	PspReference string `json:"pspReference"`
	Response     string `json:"response"` //  "[adjustAuthorisation-received]"
}

// ModificationTechnicalCancel struct
type ModificationTechnicalCancel struct {
	PspReference string `json:"pspReference"`
	Response     string `json:"response"` // "[technical-cancel-received]"
}
