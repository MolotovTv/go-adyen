// Copyright 2018 Matthieu CORNUT-CHAUVINC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package types

// RecurringDetailsParams is the set of parameters that can be used when process an recurringDetails.
type RecurringDetailsParams struct {
	MerchantAccount  string                   `json:"merchantAccount"`
	ShopperReference string                   `json:"shopperReference"`
	Recurring        *RecurringContractParams `json:"recurring"`
}

// RecurringContractParams struct
type RecurringContractParams struct {
	Contract ContractType `json:"contract"`
}

// RecurringDisableParams is the set of parameters that can be used when process an recurringDisable.
type RecurringDisableParams struct {
	ShopperReference         string `json:"shopperReference"`
	RecurringDetailReference string `json:"recurringDetailReference"`
	MerchantAccount          string `json:"merchantAccount"`
}

/********************************************/
/*            RESPONSE FROM ADYEN           */
/********************************************/

// RecurringDetails is the resource representing a Adyen recurringDetails.
type RecurringDetails struct {
	CreationDate             string     `json:"creationDate,omitempty"`
	LastKnownShopperEmail    string     `json:"lastKnownShopperEmail,omitempty"`
	ShopperReference         string     `json:"shopperReference,omitempty"`
	Details                  []*Details `json:"details,omitempty"`
	InvalidOneclickContracts string     `json:"invalidOneclickContracts"`
}

// Details struct
type Details struct {
	RecurringDetail *RecurringDetail `json:"RecurringDetail,omitempty"`
}

// RecurringDetail struct
type RecurringDetail struct {
	Acquirer                 string                         `json:"acquierer,omitempty"`
	AcquirerAccount          string                         `json:"acquirerAccount,omitempty"`
	AdditionalData           *RecurringDetailAdditionalData `json:"additionalData,omitempty"`
	Alias                    string                         `json:"alias,omitempty"`
	AliasType                string                         `json:"aliasType,omitempty"`
	Card                     *RecurringDetailCard           `json:"card,omitempty"`
	ContractTypes            []string                       `json:"contractTypes,omitempty"`
	CreationDate             *DateTime                      `json:"creationDate,omitempty"`
	FirstPspReference        string                         `json:"firstPspReference,omitempty"`
	PaymentMethodVariant     string                         `json:"paymentMethodVariant,omitempty"`
	RecurringDetailReference string                         `json:"recurringDetailReference,omitempty"`
	Variant                  string                         `json:"variant,omitempty"`
}

// RecurringDetailAdditionalData struct
type RecurringDetailAdditionalData struct {
	CardBin string `json:"cardBin,omitempty"`
}

// RecurringDetailCard struct
type RecurringDetailCard struct {
	ExpiryMonth string `json:"expiryMonth,omitempty"`
	ExpiryYear  string `json:"expiryYear,omitempty"`
	HolderName  string `json:"holderName,omitempty"`
	Number      string `json:"number,omitempty"`
}

// RecurringDisable struct
type RecurringDisable struct {
	Response string     `json:"response"` //"[detail-successfully-disabled]"
	Details  *[]Details `json:"details"`
}
