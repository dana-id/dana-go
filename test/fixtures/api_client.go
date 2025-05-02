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

package fixtures

import (
	"context"
	"os"

	dana "github.com/dana-id/dana-go-api-client"
	"github.com/dana-id/dana-go-api-client/config"
)

// ApiClient fixture structure
type ApiClientFixture struct {
	Client *dana.APIClient
	Ctx    context.Context
}

// GetApiClient returns a fixture for APIClient
func GetApiClient() ApiClientFixture {
	cfg := config.NewConfiguration()

	// Set API keys
	cfg.APIKey = &config.APIKey{
		ENV:          config.ENV_SANDBOX,
		X_PARTNER_ID: os.Getenv("X_PARTNER_ID"),
		CHANNEL_ID:   os.Getenv("CHANNEL_ID"),
		PRIVATE_KEY:  os.Getenv("PRIVATE_KEY"),
		ORIGIN:       os.Getenv("ORIGIN"),
		// PRIVATE_KEY_PATH: os.Getenv("PRIVATE_KEY_PATH"),
	}

	// Create API client with config
	apiClient := dana.NewAPIClient(cfg)
	ctx := context.Background()

	return ApiClientFixture{
		Client: apiClient,
		Ctx:    ctx,
	}
}
