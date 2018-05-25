// Copyright 2018 Matthieu CORNUT-CHAUVINC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package adyen

import (
	"fmt"
	"net/http"

	"github.com/mcornut/go-adyen/constants"

	"github.com/mcornut/go-adyen/types"
)

// ClientRecurring interface
type ClientRecurring interface {
	ListRecurringDetails(params *types.RecurringDetailsParams) (*types.RecurringDetails, error)
	Disable(params *types.RecurringDisableParams) (*types.RecurringDisable, error)
}

// ListRecurringDetails func
func (c Client) ListRecurringDetails(params *types.RecurringDetailsParams) (*types.RecurringDetails, error) {
	adyenRecurringDetails := &types.RecurringDetails{}
	err := c.call(http.MethodPost, fmt.Sprintf("/pal/servlet/Recurring/%v/listRecurringDetails", constants.APIRecurringV25), params, adyenRecurringDetails)

	return adyenRecurringDetails, err
}

// Disable func
func (c Client) Disable(params *types.RecurringDisableParams) (*types.RecurringDisable, error) {
	adyenRecurringDisable := &types.RecurringDisable{}
	err := c.call("POST", fmt.Sprintf("/pal/servlet/Recurring/%v/disable", constants.APIRecurringV25), params, adyenRecurringDisable)
	return adyenRecurringDisable, err
}
