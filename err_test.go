package sak

import (
	"runtime"
	"testing"
)

func TestErrf(t *testing.T) {
	_, heref, herel, _ := runtime.Caller(0)
	got := Errf("Steve %s!", "Holt").(*sErr)
	want := &sErr{filename: heref, line: herel + 1, payload: "Steve Holt!"}
	want.format()
	if diff := compare(got, want); diff != "" {
		t.Errorf("error not formatted properly: (-got,+want):\n%s", diff)
	}
}

func TestErr(t *testing.T) {
	_, heref, herel, _ := runtime.Caller(0)
	got := Err("Steve Holt!").(*sErr)
	want := &sErr{filename: heref, line: herel + 1, payload: "Steve Holt!"}
	want.format()
	if diff := compare(got, want); diff != "" {
		t.Errorf("error not formatted properly: (-got,+want):\n%s", diff)
	}
}

func TestNilErr(t *testing.T) {
	got := Err(nil)
	if got != nil {
		t.Error("returned error should have been nil")
	}
}
