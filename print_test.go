package testCodecov

import (
	"reflect"
	"testing"
)

func TestAnalyzeText(t *testing.T) {
	tests := []struct {
		input    string
		expected []WordFrequency
	}{
		{
			input: "Hello, world! Hello, Go.",
			expected: []WordFrequency{
				{Word: "hello", Count: 2},
				{Word: "world", Count: 1},
				{Word: "go", Count: 1},
			},
		},
		{
			input: "Go is great. Go is fun. Go is powerful.",
			expected: []WordFrequency{
				{Word: "go", Count: 3},
				{Word: "is", Count: 3},
				{Word: "great", Count: 1},
				{Word: "fun", Count: 1},
				{Word: "powerful", Count: 1},
			},
		},
		{
			input: "Test, test, test. Testing, one, two, three.",
			expected: []WordFrequency{
				{Word: "test", Count: 3},
				{Word: "testing", Count: 1},
				{Word: "one", Count: 1},
				{Word: "two", Count: 1},
				{Word: "three", Count: 1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := AnalyzeText(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}