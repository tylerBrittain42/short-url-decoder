package decoder

import (
	"slices"
	"testing"
)

func TestGetUrl(t *testing.T) {
	type testCase struct {
		description    string
		input          string
		expectedOutput string
	}
	tests := []testCase{
		{
			description:    "bitly",
			input:          "http://bit.ly/4cXanz8",
			expectedOutput: "https://www.reddit.com/r/C_Programming/comments/1obr86e/beejs_guide_to_c_programming/",
		},
		{
			description:    "tinyurl",
			input:          "https://tinyurl.com/2npz4jk5",
			expectedOutput: "https://github.com/tylerBrittain42/EZ-Blog/blob/main/article_handler.go",
		},
	}
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			actualOutput, err := Decode(tt.input)
			if err != nil {
				t.Errorf("Encountered error %v", err)
			}
			if actualOutput != tt.expectedOutput {
				t.Errorf("Expected %v, got %v", tt.expectedOutput, actualOutput)
			}

		})

	}
}
