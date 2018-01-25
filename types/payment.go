// Copyright 2018 Matthieu CORNUT-CHAUVINC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package types

// PaymentAuthoriseParams struct
type PaymentAuthoriseParams struct {
	AdditionalData   *AdditionalData `json:"additionalData"` // 3D Secure
	Amount           *Amount         `json:"amount"`
	Reference        string          `json:"reference"`
	ShopperEmail     string          `json:"shopperEmail"`
	ShopperReference string          `json:"shopperReference"`
	ShopperStatement string          `json:"shopperStatement"`
	MerchantAccount  string          `json:"merchantAccount"`
}

// PaymentAuthorise3dParams struct
type PaymentAuthorise3dParams struct {
	Md              string `json:"md"`
	PaResponse      string `json:"paResponse"`
	ShopperIP       string `json:"shopperIP"`
	MerchantAccount string `json:"merchantAccount"`
}

// PaymentAuthoriseForRecurringParams is the set of parameters that can be used when process an payment.
type PaymentAuthoriseForRecurringParams struct {
	AdditionalData   *AdditionalData    `json:"additionalData"` // 3D Secure
	Amount           *Amount            `json:"amount"`
	Reference        string             `json:"reference"`
	ShopperEmail     string             `json:"shopperEmail"`
	ShopperIP        string             `json:"shopperIP"`
	ShopperReference string             `json:"shopperReference"`
	ShopperStatement string             `json:"shopperStatement"`
	Recurring        *RecurringContract `json:"recurring"`
	MerchantAccount  string             `json:"merchantAccount"`
}

// PaymentAuthoriseOneClickParams is the set of parameters that can be used when process an payment.
type PaymentAuthoriseOneClickParams struct {
	AdditionalData                   *AdditionalData    `json:"additionalData"` // 3D Secure
	Amount                           *Amount            `json:"amount"`
	Card                             *Card              `json:"card"`
	Reference                        string             `json:"reference"`
	ShopperEmail                     string             `json:"shopperEmail"`
	ShopperIP                        string             `json:"shopperIP"`
	ShopperReference                 string             `json:"shopperReference"`
	ShopperStatement                 string             `json:"shopperStatement"`
	SelectedRecurringDetailReference string             `json:"selectedRecurringDetailReference"`
	Recurring                        *RecurringContract `json:"recurring"`
	ShopperInteraction               string             `json:"shopperInteraction"` //Ecommerce
	MerchantAccount                  string             `json:"merchantAccount"`
}

// PaymentAuthoriseRecurringParams is the set of parameters that can be used when process an payment.
type PaymentAuthoriseRecurringParams struct {
	AdditionalData                   *AdditionalData    `json:"additionalData"` // 3D Secure
	Amount                           *Amount            `json:"amount"`
	Reference                        string             `json:"reference"`
	ShopperEmail                     string             `json:"shopperEmail"`
	ShopperIP                        string             `json:"shopperIP"`
	ShopperReference                 string             `json:"shopperReference"`
	ShopperStatement                 string             `json:"shopperStatement"`
	SelectedRecurringDetailReference string             `json:"selectedRecurringDetailReference"`
	Recurring                        *RecurringContract `json:"recurring"`
	ShopperInteraction               string             `json:"shopperInteraction"` // ContAuth
	MerchantAccount                  string             `json:"merchantAccount"`
}

/********************************************/
/*            RESPONSE FROM ADYEN           */
/********************************************/
// PaymentAuthorise is the resource representing a Adyen payment respnse.
type PaymentAuthorise struct {
	AdditionalData *AdditionalData `json:"additionalData"`
	PspReference   string          `json:"pspReference"`
	ResultCode     string          `json:"resultCode"`
	AuthCode       string          `json:"authCode"`
}
