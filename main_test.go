package funplugin

import (
	"testing"
)

func TestSchema_ID2id(t *testing.T) {
	tests := []struct {
		Value    interface{}
		Expected uint
	}{
		{"nil", 0},
		{"", 0},
		{nil, 0},
		{1, 0},
		{"a", 0},
		{"-a", 0},
		{"1", 1},
		{"-1", 1},
		{"a-1", 1},
		{"1-a", 0},
		{"123456", 123456},
		{"-123456", 123456},
		{"abc-123456", 123456},
		{"123456-", 0},
	}

	for _, test := range tests {
		val, err := ID2id(test.Value)
		if test.Expected == 0 && err == nil {
			t.Fatalf("wrong result")
		}

		if val != test.Expected {
			t.Fatalf("failed ID2id(%s), expected: %v, got %v", test.Value, test.Expected, val)
		}
	}
}
