package patterns

import (
	"reflect"
	"testing"
)

func TestRightTriangle(t *testing.T) {
	got, err := RightTriangle(3)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := []string{
		"*",
		"**",
		"***",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("RightTriangle(3) = %#v, want %#v", got, want)
	}
}

func TestCenteredPyramid(t *testing.T) {
	got, err := CenteredPyramid(3)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := []string{
		"  *",
		" ***",
		"*****",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("CenteredPyramid(3) = %#v, want %#v", got, want)
	}
}

func TestDiamond(t *testing.T) {
	got, err := Diamond(3)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := []string{
		"  *",
		" ***",
		"*****",
		" ***",
		"  *",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Diamond(3) = %#v, want %#v", got, want)
	}
}

func TestHollowSquare(t *testing.T) {
	got, err := HollowSquare(4)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := []string{
		"****",
		"*  *",
		"*  *",
		"****",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("HollowSquare(4) = %#v, want %#v", got, want)
	}
}

func TestRightPascal(t *testing.T) {
	got, err := RightPascal(3)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := []string{
		"*",
		"**",
		"***",
		"**",
		"*",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("RightPascal(3) = %#v, want %#v", got, want)
	}
}

func TestInvalidSizes(t *testing.T) {
	cases := []struct {
		name string
		fn   func(int) ([]string, error))
	}{
		{"RightTriangle", RightTriangle},
		{"CenteredPyramid", CenteredPyramid},
		{"Diamond", Diamond},
		{"HollowSquare", HollowSquare},
		{"RightPascal", RightPascal},
	}
	for _, tc := range cases {
		_, err := tc.fn(0)
		if err == nil {
			t.Errorf("%s(0) expected error, got nil", tc.name)
		}
	}
}
