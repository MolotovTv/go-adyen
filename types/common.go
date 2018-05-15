// Copyright 2018 Matthieu CORNUT-CHAUVINC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package types

import (
	"encoding/json"
	"time"
)

// ContractType const
type ContractType string

// Contract types
const (
	ContractRecurring         ContractType = "RECURRING"
	ContractOneClick          ContractType = "ONECLICK"
	ContractRecurringOneClick ContractType = "RECURRING,ONECLICK"
)

// InteractionType const
type InteractionType string

// Interaction types
const (
	InteractionEcommerce InteractionType = "Ecommerce"
	InteractionContAuth  InteractionType = "ContAuth"
)

const expiryTimeFormat = "01/2006"

// Amount struct
type Amount struct {
	Value    int    `json:"value"`
	Currency string `json:"currency"`
}

// Card struct
type Card struct {
	Cvc string `json:"cvc"`
}

// RecurringContract struct
type RecurringContract struct {
	Contract ContractType `json:"contract"`
}

// AdditionalData struct
type AdditionalData struct {
	AuthorisedAmountCurrency string      `json:"authorisedAmountCurrency"`
	AuthorisedAmountValue    string      `json:"authorisedAmountValue"`
	CardBin                  string      `json:"cardBin"`
	CardEncryptedJSON        string      `json:"card.encrypted.json"`
	CardIssuingCountry       string      `json:"cardIssuingCountry"`
	CardHolderName           string      `json:"cardHolderName"`
	CardPaymentMethod        string      `json:"cardPaymentMethod"`
	CardSummary              string      `json:"cardSummary"`
	ExecuteThreeD            string      `json:"executeThreeD"`
	ExpiryDate               *ExpiryTime `json:"expiryDate"`
	IndustryUsage            string      `json:"industryUsage"`
	PaymentMethod            string      `json:"paymentMethod"`
	PaymentMethodVariant     string      `json:"paymentMethodVariant"`
}

// DateTime struct
type DateTime struct {
	Time time.Time
}

// UnmarshalJSON implements json.Unmarshaler.
// It supports string, object and null input.
func (t *DateTime) UnmarshalJSON(b []byte) error {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	return t.UnmarshalText(b)
}

// UnmarshalText allows RFC3339 to implement the TextUnmarshaler interface
func (t *DateTime) UnmarshalText(b []byte) error {
	str := string(b)
	var err error

	if str == "" || str == "null" {
		t.Time = time.Time{}

		return nil
	}
	t.Time, err = time.Parse(time.RFC3339, str)
	return err
}

// ExpiryTime struct
type ExpiryTime struct {
	Time time.Time
}

// UnmarshalJSON implements json.Unmarshaler.
// It supports string, object and null input.
func (t *ExpiryTime) UnmarshalJSON(b []byte) error {
	var d string
	err := json.Unmarshal(b, &d)
	if err != nil {
		return err
	}
	return t.UnmarshalText([]byte(d))
}

// UnmarshalText allows to implement the TextUnmarshaler interface
func (t *ExpiryTime) UnmarshalText(b []byte) error {
	str := string(b)
	var err error
	if len(str) == 6 {
		str = "0" + str
	}
	if str == "" || str == "null" {
		t.Time = time.Time{}

		return nil
	}
	t.Time, err = time.Parse(expiryTimeFormat, str)
	return err
}
