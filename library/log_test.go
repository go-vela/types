// SPDX-License-Identifier: Apache-2.0

package library

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLibrary_Log_AppendData(t *testing.T) {
	// setup types
	data := []byte("bar")

	want := testLog()
	want.SetData([]byte("foobar"))

	// setup tests
	tests := []struct {
		log  *Log
		want *Log
	}{
		{
			log:  testLog(),
			want: want,
		},
		{
			log:  new(Log),
			want: &Log{Data: &data},
		},
	}

	// run tests
	for _, test := range tests {
		test.log.AppendData(data)

		if !reflect.DeepEqual(test.log, test.want) {
			t.Errorf("AppendData is %v, want %v", test.log, test.want)
		}
	}
}

func TestLibrary_Log_MaskData(t *testing.T) {
	// set up test secrets
	sVals := []string{"secret", "((%.YY245***pP.><@@}}", "littlesecret", "extrasecret"}

	// set up test logs
	s1 := "$ echo $NO_SECRET\nnosecret\n"
	s2 := "((%.YY245***pP.><@@}}"
	s2Masked := "***"
	s3 := "$ echo $SECRET1\n((%.YY245***pP.><@@}}\n$ echo $SECRET2\nlittlesecret\n"
	s3Masked := "$ echo $SECRET1\n***\n$ echo $SECRET2\n***\n"
	s4 := "SOME_SECRET=((%.YY245***pP.><@@}}"
	s4Masked := "SOME_SECRET=***"
	s5 := "www.example.com?username=secret&password=extrasecret"
	s5Masked := "www.example.com?username=***&password=***"
	s6 := "[token: extrasecret]"
	s6Masked := "[token: ***]"

	tests := []struct {
		want    []byte
		log     []byte
		secrets []string
	}{
		{ // no secrets in log
			want:    []byte(s1),
			log:     []byte(s1),
			secrets: sVals,
		},
		{ // one secret in log
			want:    []byte(s2Masked),
			log:     []byte(s2),
			secrets: sVals,
		},
		{ // multiple secrets in log
			want:    []byte(s3Masked),
			log:     []byte(s3),
			secrets: sVals,
		},
		{ // secret with leading =
			want:    []byte(s4Masked),
			log:     []byte(s4),
			secrets: sVals,
		},
		{ // secret baked in URL query params
			want:    []byte(s5Masked),
			log:     []byte(s5),
			secrets: sVals,
		},
		{ // secret in verbose brackets
			want:    []byte(s6Masked),
			log:     []byte(s6),
			secrets: sVals,
		},
		{ // empty secrets slice
			want:    []byte(s3),
			log:     []byte(s3),
			secrets: []string{},
		},
	}
	// run tests
	l := testLog()
	for _, test := range tests {
		l.SetData(test.log)
		l.MaskData(test.secrets)

		got := l.GetData()
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("MaskData is %v, want %v", string(got), string(test.want))
		}
	}
}

func TestLibrary_Log_Getters(t *testing.T) {
	// setup tests
	tests := []struct {
		log  *Log
		want *Log
	}{
		{
			log:  testLog(),
			want: testLog(),
		},
		{
			log:  new(Log),
			want: new(Log),
		},
	}

	// run tests
	for _, test := range tests {
		if test.log.GetID() != test.want.GetID() {
			t.Errorf("GetID is %v, want %v", test.log.GetID(), test.want.GetID())
		}

		if test.log.GetServiceID() != test.want.GetServiceID() {
			t.Errorf("GetServiceID is %v, want %v", test.log.GetServiceID(), test.want.GetServiceID())
		}

		if test.log.GetStepID() != test.want.GetStepID() {
			t.Errorf("GetStepID is %v, want %v", test.log.GetStepID(), test.want.GetStepID())
		}

		if test.log.GetBuildID() != test.want.GetBuildID() {
			t.Errorf("GetBuildID is %v, want %v", test.log.GetBuildID(), test.want.GetBuildID())
		}

		if test.log.GetRepoID() != test.want.GetRepoID() {
			t.Errorf("GetRepoID is %v, want %v", test.log.GetRepoID(), test.want.GetRepoID())
		}

		if !reflect.DeepEqual(test.log.GetData(), test.want.GetData()) {
			t.Errorf("GetData is %v, want %v", test.log.GetData(), test.want.GetData())
		}
	}
}

func TestLibrary_Log_Setters(t *testing.T) {
	// setup types
	var l *Log

	// setup tests
	tests := []struct {
		log  *Log
		want *Log
	}{
		{
			log:  testLog(),
			want: testLog(),
		},
		{
			log:  l,
			want: new(Log),
		},
	}

	// run tests
	for _, test := range tests {
		test.log.SetID(test.want.GetID())
		test.log.SetServiceID(test.want.GetServiceID())
		test.log.SetStepID(test.want.GetStepID())
		test.log.SetBuildID(test.want.GetBuildID())
		test.log.SetRepoID(test.want.GetRepoID())
		test.log.SetData(test.want.GetData())

		if test.log.GetID() != test.want.GetID() {
			t.Errorf("SetID is %v, want %v", test.log.GetID(), test.want.GetID())
		}

		if test.log.GetServiceID() != test.want.GetServiceID() {
			t.Errorf("SetServiceID is %v, want %v", test.log.GetServiceID(), test.want.GetServiceID())
		}

		if test.log.GetStepID() != test.want.GetStepID() {
			t.Errorf("SetStepID is %v, want %v", test.log.GetStepID(), test.want.GetStepID())
		}

		if test.log.GetBuildID() != test.want.GetBuildID() {
			t.Errorf("SetBuildID is %v, want %v", test.log.GetBuildID(), test.want.GetBuildID())
		}

		if test.log.GetRepoID() != test.want.GetRepoID() {
			t.Errorf("SetRepoID is %v, want %v", test.log.GetRepoID(), test.want.GetRepoID())
		}

		if !reflect.DeepEqual(test.log.GetData(), test.want.GetData()) {
			t.Errorf("SetData is %v, want %v", test.log.GetData(), test.want.GetData())
		}
	}
}

func TestLibrary_Log_String(t *testing.T) {
	// setup types
	l := testLog()

	want := fmt.Sprintf(`{
  BuildID: %d,
  Data: %s,
  ID: %d,
  RepoID: %d,
  ServiceID: %d,
  StepID: %d,
}`,
		l.GetBuildID(),
		l.GetData(),
		l.GetID(),
		l.GetRepoID(),
		l.GetServiceID(),
		l.GetStepID(),
	)

	// run test
	got := l.String()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("String is %v, want %v", got, want)
	}
}

// testLog is a test helper function to create a Log
// type with all fields set to a fake value.
func testLog() *Log {
	l := new(Log)

	l.SetID(1)
	l.SetServiceID(1)
	l.SetStepID(1)
	l.SetBuildID(1)
	l.SetRepoID(1)
	l.SetData([]byte("foo"))

	return l
}
