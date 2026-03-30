package patterns

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRightTriangle(t *testing.T) {
	tests := []struct {
		size int
		want []string
	}{
		{1, []string{"*"}},
		{3, []string{"*", "**", "***"}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("size_%d", tt.size), func(t *testing.T) {
			got, err := RightTriangle(tt.size)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RightTriangle(%d) = %#v, want %#v", tt.size, got, tt.want)
			}
		})
	}
}

func TestCenteredPyramid(t *testing.T) {
	tests := []struct {
		size int
		want []string
	}{
		{1, []string{"*"}},
		{3, []string{"  *", " ***", "*****"}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("size_%d", tt.size), func(t *testing.T) {
			got, err := CenteredPyramid(tt.size)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CenteredPyramid(%d) = %#v, want %#v", tt.size, got, tt.want)
			}
		})
	}
}

func TestDiamond(t *testing.T) {
	tests := []struct {
		size int
		want []string
	}{
		{1, []string{"*"}},
		{3, []string{"  *", " ***", "*****", " ***", "  *"}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("size_%d", tt.size), func(t *testing.T) {
			got, err := Diamond(tt.size)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Diamond(%d) = %#v, want %#v", tt.size, got, tt.want)
			}
		})
	}
}

func TestHollowSquare(t *testing.T) {
	tests := []struct {
		size int
		want []string
	}{
		{1, []string{"*"}},
		{4, []string{"****", "*  *", "*  *", "****"}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("size_%d", tt.size), func(t *testing.T) {
			got, err := HollowSquare(tt.size)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HollowSquare(%d) = %#v, want %#v", tt.size, got, tt.want)
			}
		})
	}
}

func TestRightPascal(t *testing.T) {
	tests := []struct {
		size int
		want []string
	}{
		{1, []string{"*"}},
		{3, []string{"*", "**", "***", "**", "*"}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("size_%d", tt.size), func(t *testing.T) {
			got, err := RightPascal(tt.size)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RightPascal(%d) = %#v, want %#v", tt.size, got, tt.want)
			}
		})
	}
}

func TestInvalidSizes(t *testing.T) {
	type gen func(int) ([]string, error)
	cases := []struct {
		name string
		fn   gen
	}{
		{"RightTriangle", RightTriangle},
		{"CenteredPyramid", CenteredPyramid},
		{"Diamond", Diamond},
		{"HollowSquare", HollowSquare},
		{"RightPascal", RightPascal},
	}
	for _, tc := range cases {
		for _, bad := range []int{0, -2} {
			t.Run(fmt.Sprintf("%s_bad_%d", tc.name, bad), func(t *testing.T) {
				if _, err := tc.fn(bad); err == nil {
					t.Fatalf("%s(%d) expected error, got nil", tc.name, bad)
				}
			})
		}
	}
}
