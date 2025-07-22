// Copyright 2025 PT Espay Debit Indonesia Koe
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package config

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

const (
	ENV_SANDBOX    = "SANDBOX"
	ENV_PRODUCTION = "PRODUCTION"
)

// APIKey struct to store all Dana API keys
type APIKey struct {
	ENV              string
	ORIGIN           string
	X_PARTNER_ID     string
	CHANNEL_ID       string
	PRIVATE_KEY      string
	PRIVATE_KEY_PATH string
	// OPEN_API specific fields
	CLIENT_SECRET string
	CLIENT_ID     string
	ACCESS_TOKEN  string
}

// ServerVariable stores the information about a server variable
type ServerVariable struct {
	Description  string
	DefaultValue string
	EnumValues   []string
}

// ServerConfiguration stores the information about a server
type ServerConfiguration struct {
	URL         string
	Description string
	Variables   map[string]ServerVariable
}

// ServerConfigurations stores multiple ServerConfiguration items
type ServerConfigurations []ServerConfiguration

// Configuration stores the configuration of the API client
type Configuration struct {
	Host             string
	Scheme           string
	DefaultHeader    map[string]string
	UserAgent        string
	Debug            bool
	Servers          ServerConfigurations
	OperationServers map[string]ServerConfigurations
	HTTPClient       *http.Client
	APIKey           *APIKey
}

// NewConfiguration returns a new Configuration object
func NewConfiguration() *Configuration {
	cfg := &Configuration{
		DefaultHeader: make(map[string]string),
		UserAgent:     "DANA-API-Client/1.0.0/go",
		Debug:         false,
		Servers: ServerConfigurations{
			{
				URL:         "https://api.sandbox.dana.id",
				Description: "Dana API Gateway Sandbox",
			},
			{
				URL:         "https://api.saas.dana.id",
				Description: "Dana API Gateway",
			},
		},
		OperationServers: map[string]ServerConfigurations{},
		APIKey:           &APIKey{},
	}
	return cfg
}

// AddDefaultHeader adds a new HTTP header to the default header in the request
func (c *Configuration) AddDefaultHeader(key string, value string) {
	c.DefaultHeader[key] = value
}

// URL formats template on a index using given variables
func (sc ServerConfigurations) URL(index int, variables map[string]string) (string, error) {
	if index < 0 || len(sc) <= index {
		return "", fmt.Errorf("index %v out of range %v", index, len(sc)-1)
	}
	server := sc[index]
	url := server.URL

	// go through variables and replace placeholders
	for name, variable := range server.Variables {
		var value string
		if val, ok := variables[name]; ok {
			value = val
			found := bool(len(variable.EnumValues) == 0)
			for _, enumValue := range variable.EnumValues {
				if value == enumValue {
					found = true
				}
			}
			if !found {
				return "", fmt.Errorf("the variable %s in the server URL has invalid value %v. Must be %v", name, value, variable.EnumValues)
			}
		} else {
			// use default value
			value = variable.DefaultValue
		}

		url = strings.Replace(url, fmt.Sprintf("{%s}", name), value, -1)
	}
	return url, nil
}

// ServerURL returns URL based on server settings
func (c *Configuration) ServerURL(index int, variables map[string]string) (string, error) {
	return c.Servers.URL(index, variables)
}

// ServerURLWithContext returns a new server URL given an endpoint
func (c *Configuration) ServerURLWithContext(ctx context.Context, endpoint string) (string, error) {
	// Check if there's a specific server configuration for this endpoint

	sc := c.Servers

	// For simplicity, just use the first server and no variables
	if len(sc) == 0 {
		return "", fmt.Errorf("no server configuration available for endpoint: %s", endpoint)
	}

	if c.APIKey.ENV == ENV_PRODUCTION {
		return sc.URL(1, nil)
	}
	return sc.URL(0, nil)
}
