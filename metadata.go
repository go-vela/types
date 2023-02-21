// Copyright (c) 2023 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package types

import "time"

type (
	// Database is the extra set of database data passed to the compiler.
	Database struct {
		Driver string `json:"driver"`
		Host   string `json:"host"`
	}

	// Queue is the extra set of queue data passed to the compiler.
	Queue struct {
		Channel string `json:"channel"`
		Driver  string `json:"driver"`
		Host    string `json:"host"`
	}

	// Source is the extra set of source data passed to the compiler.
	Source struct {
		Driver string `json:"driver"`
		Host   string `json:"host"`
	}

	// TokenManager is the config data for the v2 token manager.
	TokenManager struct {
		RegTokenDuration         time.Duration `json:"reg_token_duration"`
		WorkerAuthTokenDuration  time.Duration `json:"worker_authtoken_duration"`
		UserAccessTokenDuration  time.Duration `json:"user_accesstoken_duration"`
		UserRefreshTokenDuration time.Duration `json:"user_refreshtoken_duration"`
		TokenCleanupTicker       time.Duration `json:"token_cleanup_ticker"`
		KeyCleanupTicker         time.Duration `json:"key_cleanup_ticker"`
		InvalidTokenTTL          time.Duration `json:"invalid_token_ttl"`
		SigningKeyTTL            time.Duration `json:"signing_key_ttl"`
	}

	// Vela is the extra set of Vela data passed to the compiler.
	Vela struct {
		Address              string        `json:"address"`
		WebAddress           string        `json:"web_address"`
		WebOauthCallbackPath string        `json:"web_oauth_callback_path"`
		AccessTokenDuration  time.Duration `json:"access_token_duration"`
		RefreshTokenDuration time.Duration `json:"refresh_token_duration"`
	}

	// Metadata is the extra set of data passed to the compiler for
	// converting a yaml configuration to an executable pipeline.
	Metadata struct {
		Database     *Database     `json:"database"`
		Queue        *Queue        `json:"queue"`
		Source       *Source       `json:"source"`
		Vela         *Vela         `json:"vela"`
		TokenManager *TokenManager `json:"token_manager"`
	}
)
