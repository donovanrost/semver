package semver

import (
	"reflect"
	"testing"
)


func TestStr(t *testing.T) {
	
	type test struct {
		given Semver
		expected string
	}
	
	tests := []test {
		{given: Semver{Major: 1, Minor: 2, Patch: 3, Pre: []string{"dev"}, Meta: []string{"asd"}}, expected: "1.2.3-dev+asd"},
		{given: Semver{Major: 1, Minor: 2, Patch: 3, Pre: []string{"dev", "pod"}, Meta: []string{"asd", "qwe"}}, expected: "1.2.3-dev.pod+asd.qwe"},
		
	}
	
	for _, tc := range tests {
		got := tc.given.String()
		if !reflect.DeepEqual(tc.expected, got) {
			t.Fatalf("expected: %v, got: %v", tc.expected, got)
		}
	}
}

func TestParse(t *testing.T) {

	given := "1.2.3-dev+asd"

	got, _ := NewFromString(given)

	if got.Major != 1 {
		t.Errorf("expect major was 1, got %d", got.Major)
	}
	if got.Minor != 2 {
		t.Errorf("expect minor was 2, got %d", got.Minor)
	}
	if got.Patch != 3 {
		t.Errorf("expect patch was 3, got %d", got.Patch)
	}
	if got.Pre[0] != "dev" {
		t.Errorf("expect prerelease was \"dev\", got %s", got.Pre[0])
	}
	if got.Meta[0] != "asd" {
		t.Errorf("expect build metadata was \"asd\", got %s", got.Meta[0])
	}
}
