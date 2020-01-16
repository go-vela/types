// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package types

import (
	"fmt"
	"reflect"
	"testing"
)

func TestResp_String(t *testing.T) {
	// setup types
	str := "foo"
	e := &Error{
		Message: &str,
	}
	want := fmt.Sprintf("%+v", *e)

	// run test
	got := e.String()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("String is %v, want %v", got, want)
	}
}
