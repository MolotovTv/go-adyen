// Copyright 2018 Matthieu CORNUT-CHAUVINC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExpiryTimeUnmarshalJSON(t *testing.T) {
	e := &ExpiryTime{}

	err := e.UnmarshalJSON([]byte(`"08\/2018"`))
	assert.NoError(t, err)
}

func TestExpiryTimeUnmarshalJSONTwo(t *testing.T) {
	e := &ExpiryTime{}

	err := e.UnmarshalJSON([]byte(`"8\/2018"`))
	assert.NoError(t, err)
}

func TestExpiryTimeUnmarshalTextTwo(t *testing.T) {
	e := &ExpiryTime{}

	err := e.UnmarshalText([]byte("8/2018"))
	assert.NoError(t, err)
}
func TestExpiryTimeUnmarshalText(t *testing.T) {
	e := &ExpiryTime{}

	err := e.UnmarshalText([]byte("08/2018"))
	assert.NoError(t, err)
}
