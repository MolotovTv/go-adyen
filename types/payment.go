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
	BrowserInfo      *BrowserInfo       `json:"browserInfo"`
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
	ShopperInteraction               InteractionType    `json:"shopperInteraction"` // Ecommerce
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
	ShopperInteraction               InteractionType    `json:"shopperInteraction"` // ContAuth
	MerchantAccount                  string             `json:"merchantAccount"`
}

// BrowserInfo is the shopper's browser information, for 3D Secure, the full object is required for web integrations.
type BrowserInfo struct {

	// The accept header value of the shopper's browser.
	AcceptHeader string `json:"acceptHeader"`

	// The color depth of the shopper's browser in bits per pixel. This should be obtained by using the browser's screen.colorDepth property. Accepted values: 1, 4, 8, 15, 16, 24, 30, 32 or 48 bit color depth.
	ColorDepth int `json:"colorDepth"`

	// Boolean value indicating if the shopper's browser is able to execute Java.
	JavaEnabled bool `json:"javaEnabled"`

	// Boolean value indicating if the shopper's browser is able to execute JavaScript. A default 'true' value is assumed if the field is not present.
	JavaScriptEnabled bool `json:"javaScriptEnabled"`

	// The navigator.language value of the shopper's browser (as defined in IETF BCP 47).
	Language string `json:"language"`

	// The total height of the shopper's device screen in pixels.
	ScreenHeight int `json:"screenHeight"`

	// The total width of the shopper's device screen in pixels.
	ScreenWidth int `json:"screenWidth"`

	// Time difference between UTC time and the shopper's browser local time, in minutes.
	TimeZoneOffset int `json:"timeZoneOffset"`

	// The user agent value of the shopper's browser.
	UserAgent string `json:"userAgent"`
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
	Md             string          `json:"md,omitempty"`
	PaRequest      string          `json:"paRequest,omitempty"`
	IssuerUrl      string          `json:"issuerUrl,omitempty"`
}
