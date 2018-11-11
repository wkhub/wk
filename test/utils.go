package test

import "testing"

// HandlTestError properly handle errors returned during tests
func CheckError(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}
