// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package constants

// server database drivers
const (
	// DriverPostgres defines the driver type when integrating with a PostgreSQL database.
	DriverPostgres = "postgres"

	// DriverSqlite defines the driver type when integrating with a SQLite database.
	DriverSqlite = "sqlite3"
)

// agent executor drivers
const (
	// DriverLinux defines the driver type when integrating with a linux distribution.
	DriverLinux = "linux"

	// DriverWindows defines the driver type when integrating with a windows distribution.
	DriverWindows = "windows"

	// DriverDarwin defines the driver type when integrating with a darwin distribution.
	DriverDarwin = "darwin"
)

// server and agent queue drivers
const (

	// DriverKafka defines the driver type when integrating with a Kafka queue.
	DriverKafka = "kafka"

	// DriverRedis defines the driver type when integrating with a Redis queue.
	DriverRedis = "redis"
)

// agent runtime drivers
const (
	// DriverDocker defines the driver type when integrating with a Docker runtime.
	DriverDocker = "docker"

	// DriverKubernetes defines the driver type when integrating with a Kubernetes runtime.
	DriverKubernetes = "kubernetes"
)

// server and agent secret drivers
const (
	// DriverNative defines the driver type when integrating with a Vela secret service.
	DriverNative = "native"

	// DriverVault defines the driver type when integrating with a Vault secret service.
	DriverVault = "vault"
)

// server source drivers
const (
	// DriverGitHub defines the driver type when integrating with a Github source code system.
	DriverGithub = "github"

	// DriverGitLab defines the driver type when integrating with a Gitlab source code system.
	DriverGitlab = "gitlab"
)
