// Copyright (c) 2019 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package library

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/go-vela/types/pipeline"
)

func TestLibrary_Executor_Getters(t *testing.T) {
	// setup types
	num := 1
	num64 := int64(num)
	str := "foo"
	booL := true
	b := &Build{
		ID:           &num64,
		RepoID:       &num64,
		Number:       &num,
		Parent:       &num,
		Event:        &str,
		Status:       &str,
		Error:        &str,
		Enqueued:     &num64,
		Created:      &num64,
		Started:      &num64,
		Finished:     &num64,
		Deploy:       &str,
		Clone:        &str,
		Source:       &str,
		Title:        &str,
		Message:      &str,
		Commit:       &str,
		Sender:       &str,
		Author:       &str,
		Branch:       &str,
		Ref:          &str,
		BaseRef:      &str,
		Host:         &str,
		Runtime:      &str,
		Distribution: &str,
	}
	r := &Repo{
		ID:          &num64,
		UserID:      &num64,
		Org:         &str,
		Name:        &str,
		FullName:    &str,
		Link:        &str,
		Clone:       &str,
		Branch:      &str,
		Timeout:     &num64,
		Visibility:  &str,
		Private:     &booL,
		Trusted:     &booL,
		Active:      &booL,
		AllowPull:   &booL,
		AllowPush:   &booL,
		AllowDeploy: &booL,
		AllowTag:    &booL,
	}
	p := &pipeline.Build{
		Services: pipeline.ContainerSlice{
			&pipeline.Container{
				Image:  "postgres:latest",
				Name:   "postgres",
				Number: 1,
			}},
		Worker: pipeline.Worker{
			Flavor:   "16cpu8gb",
			Platform: "gcp",
		},
		Stages: pipeline.StageSlice{
			&pipeline.Stage{
				Name:  "install",
				Needs: []string{"clone"},
				Steps: pipeline.ContainerSlice{
					&pipeline.Container{
						Commands: []string{"./gradlew downloadDependencies"},
						Image:    "openjdk:latest",
						Name:     "install",
						Number:   1,
						Pull:     true,
					},
				},
			},
			&pipeline.Stage{
				Name:  "test",
				Needs: []string{"install"},
				Steps: pipeline.ContainerSlice{
					&pipeline.Container{
						Commands: []string{"./gradlew check"},
						Image:    "openjdk:latest",
						Name:     "test",
						Number:   2,
						Pull:     true,
						Ruleset: pipeline.Ruleset{
							If: pipeline.Rules{
								Event: []string{"push"},
							},
							Operator: "and",
						},
					},
				},
			},
		},
	}
	e := &Executor{
		ID:           &num64,
		Host:         &str,
		Runtime:      &str,
		Distribution: &str,
		Build:        b,
		Repo:         r,
		Pipeline:     p,
	}
	wantID := num64
	wantHost := str
	wantRuntime := str
	wantDistribution := str
	wantBuild := *e.Build
	wantRepo := *e.Repo
	wantPipeline := *e.Pipeline

	// run test
	gotID := e.GetID()
	gotHost := e.GetHost()
	gotRuntime := e.GetRuntime()
	gotDistribution := e.GetDistribution()
	gotBuild := e.GetBuild()
	gotRepo := e.GetRepo()
	gotPipeline := e.GetPipeline()

	if gotID != wantID {
		t.Errorf("GetID is %v, want %v", gotID, wantID)
	}
	if gotHost != wantHost {
		t.Errorf("GetHost is %v, want %v", gotHost, wantHost)
	}
	if gotRuntime != wantRuntime {
		t.Errorf("GetRuntime is %v, want %v", gotRuntime, wantRuntime)
	}
	if gotDistribution != wantDistribution {
		t.Errorf("GetDistribution is %v, want %v", gotDistribution, wantDistribution)
	}
	if !reflect.DeepEqual(gotBuild, wantBuild) {
		t.Errorf("GetBuild is %v, want %v", gotBuild, wantBuild)
	}
	if !reflect.DeepEqual(gotRepo, wantRepo) {
		t.Errorf("GetRepo is %v, want %v", gotRepo, wantRepo)
	}
	if !reflect.DeepEqual(gotPipeline, wantPipeline) {
		t.Errorf("GetPipeline is %v, want %v", gotPipeline, wantPipeline)
	}
}

func TestLibrary_Executor_Getters_Empty(t *testing.T) {
	// setup types
	e := new(Executor)

	// run test
	gotID := e.GetID()
	gotHost := e.GetHost()
	gotRuntime := e.GetRuntime()
	gotDistribution := e.GetDistribution()
	gotBuild := e.GetBuild()
	gotRepo := e.GetRepo()
	gotPipeline := e.GetPipeline()

	if gotID != 0 {
		t.Errorf("GetID is %v, want 0", gotID)
	}
	if gotHost != "" {
		t.Errorf("GetHost is %v, want \"\"", gotHost)
	}
	if gotRuntime != "" {
		t.Errorf("GetRuntime is %v, want \"\"", gotRuntime)
	}
	if gotDistribution != "" {
		t.Errorf("GetDistribution is %v, want \"\"", gotDistribution)
	}
	if !reflect.DeepEqual(gotBuild, Build{}) {
		t.Errorf("GetBuild is %v, want \"\"", gotBuild)
	}
	if !reflect.DeepEqual(gotRepo, Repo{}) {
		t.Errorf("GetRepo is %v, want \"\"", gotRepo)
	}
	if !reflect.DeepEqual(gotPipeline, pipeline.Build{}) {
		t.Errorf("GetPipeline is %v, want \"\"", gotPipeline)
	}
}

func TestLibrary_Executor_Setters(t *testing.T) {
	// setup types
	num := 1
	num64 := int64(num)
	str := "foo"
	booL := true
	b := Build{
		ID:           &num64,
		RepoID:       &num64,
		Number:       &num,
		Parent:       &num,
		Event:        &str,
		Status:       &str,
		Error:        &str,
		Enqueued:     &num64,
		Created:      &num64,
		Started:      &num64,
		Finished:     &num64,
		Deploy:       &str,
		Clone:        &str,
		Source:       &str,
		Title:        &str,
		Message:      &str,
		Commit:       &str,
		Sender:       &str,
		Author:       &str,
		Branch:       &str,
		Ref:          &str,
		BaseRef:      &str,
		Host:         &str,
		Runtime:      &str,
		Distribution: &str,
	}
	r := Repo{
		ID:          &num64,
		UserID:      &num64,
		Org:         &str,
		Name:        &str,
		FullName:    &str,
		Link:        &str,
		Clone:       &str,
		Branch:      &str,
		Timeout:     &num64,
		Visibility:  &str,
		Private:     &booL,
		Trusted:     &booL,
		Active:      &booL,
		AllowPull:   &booL,
		AllowPush:   &booL,
		AllowDeploy: &booL,
		AllowTag:    &booL,
	}
	p := pipeline.Build{
		Services: pipeline.ContainerSlice{
			&pipeline.Container{
				Image:  "postgres:latest",
				Name:   "postgres",
				Number: 1,
			}},
		Worker: pipeline.Worker{
			Flavor:   "16cpu8gb",
			Platform: "gcp",
		},
		Stages: pipeline.StageSlice{
			&pipeline.Stage{
				Name:  "install",
				Needs: []string{"clone"},
				Steps: pipeline.ContainerSlice{
					&pipeline.Container{
						Commands: []string{"./gradlew downloadDependencies"},
						Image:    "openjdk:latest",
						Name:     "install",
						Number:   1,
						Pull:     true,
					},
				},
			},
			&pipeline.Stage{
				Name:  "test",
				Needs: []string{"install"},
				Steps: pipeline.ContainerSlice{
					&pipeline.Container{
						Commands: []string{"./gradlew check"},
						Image:    "openjdk:latest",
						Name:     "test",
						Number:   2,
						Pull:     true,
						Ruleset: pipeline.Ruleset{
							If: pipeline.Rules{
								Event: []string{"push"},
							},
							Operator: "and",
						},
					},
				},
			},
		},
	}
	e := new(Executor)

	wantID := num64
	wantHost := str
	wantRuntime := str
	wantDistribution := str
	wantBuild := b
	wantRepo := r
	wantPipeline := p

	// Run tests
	e.SetID(wantID)
	e.SetHost(wantHost)
	e.SetRuntime(wantRuntime)
	e.SetDistribution(wantDistribution)
	e.SetBuild(wantBuild)
	e.SetRepo(wantRepo)
	e.SetPipeline(wantPipeline)

	if e.GetID() != wantID {
		t.Errorf("SetID is %v, want %v", e.GetID(), wantID)
	}
	if e.GetHost() != wantHost {
		t.Errorf("SetHost is %v, want %v", e.GetHost(), wantHost)
	}
	if e.GetRuntime() != wantRuntime {
		t.Errorf("SetRuntime is %v, want %v", e.GetRuntime(), wantRuntime)
	}
	if e.GetDistribution() != wantDistribution {
		t.Errorf("SetDistribution is %v, want %v", e.GetDistribution(), wantDistribution)
	}
	if !reflect.DeepEqual(e.GetBuild(), wantBuild) {
		t.Errorf("SetBuild is %v, want %v", e.GetBuild(), wantBuild)
	}
	if !reflect.DeepEqual(e.GetRepo(), wantRepo) {
		t.Errorf("SetRepo is %v, want %v", e.GetRepo(), wantRepo)
	}
	if !reflect.DeepEqual(e.GetPipeline(), wantPipeline) {
		t.Errorf("SetPipeline is %v, want %v", e.GetPipeline(), wantPipeline)
	}
}

func TestLibrary_Executor_Setters_Empty(t *testing.T) {
	// setup types
	var e *Executor

	// Run tests
	e.SetID(0)
	e.SetHost("")
	e.SetRuntime("")
	e.SetDistribution("")
	e.SetBuild(Build{})
	e.SetRepo(Repo{})
	e.SetPipeline(pipeline.Build{})

	if e.GetID() != 0 {
		t.Errorf("SetID is %v, want 0", e.GetID())
	}
	if e.GetHost() != "" {
		t.Errorf("SetHost is %v, want \"\"", e.GetHost())
	}
	if e.GetRuntime() != "" {
		t.Errorf("SetRuntime is %v, want \"\"", e.GetRuntime())
	}
	if e.GetDistribution() != "" {
		t.Errorf("SetDistribution is %v, want \"\"", e.GetDistribution())
	}
	if !reflect.DeepEqual(e.GetBuild(), Build{}) {
		t.Errorf("GetBuild is %v, want \"\"", e.GetBuild())
	}
	if !reflect.DeepEqual(e.GetRepo(), Repo{}) {
		t.Errorf("GetRepo is %v, want \"\"", e.GetRepo())
	}
	if !reflect.DeepEqual(e.GetPipeline(), pipeline.Build{}) {
		t.Errorf("GetPipeline is %v, want \"\"", e.GetPipeline())
	}
}

func TestLibrary_Executor_String(t *testing.T) {
	// setup types
	num := 1
	num64 := int64(num)
	str := "foo"
	booL := true
	b := &Build{
		ID:           &num64,
		RepoID:       &num64,
		Number:       &num,
		Parent:       &num,
		Event:        &str,
		Status:       &str,
		Error:        &str,
		Enqueued:     &num64,
		Created:      &num64,
		Started:      &num64,
		Finished:     &num64,
		Deploy:       &str,
		Clone:        &str,
		Source:       &str,
		Title:        &str,
		Message:      &str,
		Commit:       &str,
		Sender:       &str,
		Author:       &str,
		Branch:       &str,
		Ref:          &str,
		BaseRef:      &str,
		Host:         &str,
		Runtime:      &str,
		Distribution: &str,
	}
	r := &Repo{
		ID:          &num64,
		UserID:      &num64,
		Org:         &str,
		Name:        &str,
		FullName:    &str,
		Link:        &str,
		Clone:       &str,
		Branch:      &str,
		Timeout:     &num64,
		Visibility:  &str,
		Private:     &booL,
		Trusted:     &booL,
		Active:      &booL,
		AllowPull:   &booL,
		AllowPush:   &booL,
		AllowDeploy: &booL,
		AllowTag:    &booL,
	}
	p := &pipeline.Build{
		Services: pipeline.ContainerSlice{
			&pipeline.Container{
				Image:  "postgres:latest",
				Name:   "postgres",
				Number: 1,
			}},
		Worker: pipeline.Worker{
			Flavor:   "16cpu8gb",
			Platform: "gcp",
		},
		Stages: pipeline.StageSlice{
			&pipeline.Stage{
				Name:  "install",
				Needs: []string{"clone"},
				Steps: pipeline.ContainerSlice{
					&pipeline.Container{
						Commands: []string{"./gradlew downloadDependencies"},
						Image:    "openjdk:latest",
						Name:     "install",
						Number:   1,
						Pull:     true,
					},
				},
			},
			&pipeline.Stage{
				Name:  "test",
				Needs: []string{"install"},
				Steps: pipeline.ContainerSlice{
					&pipeline.Container{
						Commands: []string{"./gradlew check"},
						Image:    "openjdk:latest",
						Name:     "test",
						Number:   2,
						Pull:     true,
						Ruleset: pipeline.Ruleset{
							If: pipeline.Rules{
								Event: []string{"push"},
							},
							Operator: "and",
						},
					},
				},
			},
		},
	}
	e := &Executor{
		ID:           &num64,
		Host:         &str,
		Runtime:      &str,
		Distribution: &str,
		Build:        b,
		Repo:         r,
		Pipeline:     p,
	}
	want := fmt.Sprintf("%+v", *e)

	// run test
	got := e.String()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("String is %v, want %v", got, want)
	}
}
