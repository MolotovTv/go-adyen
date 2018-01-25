// Copyright 2018 Matthieu CORNUT-CHAUVINC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package adyen

// Configuration represents the MySQL configuration
type Configuration struct {
	Password string `toml:"password" json:"password"`
	Username string `toml:"username" json:"username"`
	Live     bool   `toml:"live" json:"live"`
}

// NewConfiguration generates a Configuration
func NewConfiguration() Configuration {
	return Configuration{}
}
