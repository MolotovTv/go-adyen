// Copyright 2018 Matthieu CORNUT-CHAUVINC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package types

// Notification struct
type Notification struct {
	Live              string              `json:"live"`
	NotificationItems []*NotificationItem `json:"notificationItems"`
}

// NotificationItem struct
type NotificationItem struct {
	NotificationRequestItem *NotificationRequestItem `json:"NotificationRequestItem"`
}

// NotificationRequestItem struct
type NotificationRequestItem struct {
	AdditionalData      *NotificationAdditionalData `json:"additionalData"`
	Amount              *Amount                     `json:"amount"`
	EventCode           string                      `json:"eventCode"`
	EventDate           *DateTime                   `json:"eventDate"`
	MerchantAccountCode string                      `json:"merchantAccountCode"`
	MerchantReference   string                      `json:"merchantReference"`
	Operations          []string                    `json:"operations"`
	OriginalReference   string                      `json:"originalReference"`
	PaymentMethod       string                      `json:"paymentMethod"`
	PspReference        string                      `json:"pspReference"`
	Reason              string                      `json:"reason"`
	Success             string                      `json:"success"`
}

// NotificationAdditionalData struct
type NotificationAdditionalData struct {
	AuthCode                 string      `json:"authCode"`
	CardHolderName           string      `json:"cardHolderName"`
	CardSummary              string      `json:"cardSummary"`
	ShopperEmail             string      `json:"shopperEmail"`
	AuthorisedAmountValue    string      `json:"authorisedAmountValue"`
	ExpiryDate               *ExpiryTime `json:"expiryDate"`
	AuthorisedAmountCurrency string      `json:"authorisedAmountCurrency"`
	CardBin                  string      `json:"cardBin"`
	AliasType                string      `json:"aliasType"`
	Alias                    string      `json:"alias"`
	CardPaymentMethod        string      `json:"cardPaymentMethod"`
	PaymentMethodVariant     string      `json:"paymentMethodVariant"`
	CardIssuingCountry       string      `json:"cardIssuingCountry"`
	ShopperReference         string      `json:"shopperReference"`
	IssuerCountry            string      `json:"issuerCountry"`
}
