// SPDX-License-Identifier: Apache-2.0

package types

import "time"

type (
	// Database is the extra set of database data passed to the compiler.
	//
	// Deprecated: use Database from github.com/go-vela/server/internal instead.
	Database struct {
		Driver string `json:"driver"`
		Host   string `json:"host"`
	}

	// Queue is the extra set of queue data passed to the compiler.
	//
	// Deprecated: use Queue from github.com/go-vela/server/internal instead.
	Queue struct {
		Channel string `json:"channel"`
		Driver  string `json:"driver"`
		Host    string `json:"host"`
	}

	// Source is the extra set of source data passed to the compiler.
	//
	// Deprecated: use Source from github.com/go-vela/server/internal instead.
	Source struct {
		Driver string `json:"driver"`
		Host   string `json:"host"`
	}

	// Vela is the extra set of Vela data passed to the compiler.
	//
	// Deprecated: use Vela from github.com/go-vela/server/internal instead.
	Vela struct {
		Address              string        `json:"address"`
		WebAddress           string        `json:"web_address"`
		WebOauthCallbackPath string        `json:"web_oauth_callback_path"`
		AccessTokenDuration  time.Duration `json:"access_token_duration"`
		RefreshTokenDuration time.Duration `json:"refresh_token_duration"`
	}

	// Metadata is the extra set of data passed to the compiler for
	// converting a yaml configuration to an executable pipeline.
	//
	// Deprecated: use Metadata from github.com/go-vela/server/internal instead.
	Metadata struct {
		Database *Database `json:"database"`
		Queue    *Queue    `json:"queue"`
		Source   *Source   `json:"source"`
		Vela     *Vela     `json:"vela"`
	}
)
