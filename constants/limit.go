// SPDX-License-Identifier: Apache-2.0

package constants

// Limits and constraints.
const (
	// BuildLimitMin defines the minimum value for repo concurrent build limit.
	BuildLimitMin = 1

	// BuildLimitMax defines the maximum value for repo concurrent build limit.
	BuildLimitMax = 30

	// BuildLimitDefault defines the default value for repo concurrent build limit.
	BuildLimitDefault = 10

	// BuildTimeoutMin defines the minimum value in minutes for repo build timeout.
	BuildTimeoutMin = 1

	// BuildTimeoutMax defines the maximum value in minutes for repo build timeout.
	BuildTimeoutMax = 90

	// BuildTimeoutDefault defines the default value in minutes for repo build timeout.
	BuildTimeoutDefault = 30

	// FavoritesMaxSize defines the maximum size in characters for user favorites.
	FavoritesMaxSize = 5000

	// DashboardAdminMaxSize defines the maximum size in characters for dashboard admins.
	DashboardAdminMaxSize = 5000

	// RunningBuildIDsMaxSize defines the maximum size in characters for worker RunningBuildIDs.
	RunningBuildIDsMaxSize = 500

	// TopicsMaxSize defines the maximum size in characters for repo topics. Ex: GitHub has a 20-topic, 50-char limit.
	TopicsMaxSize = 1020

	// DeployBuildsMaxSize defines the maximum size in characters for deployment builds.
	DeployBuildsMaxSize = 500

	// ReportStepStatusLimit defines the maximum number of steps in a pipeline that may report their status to the SCM.
	ReportStepStatusLimit = 10
)
