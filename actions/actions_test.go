package actions

import (
	"github.com/stretchr/testify/require"
	"testing"
	"uniq/m/settings"
)

func TestDefaultMode(t *testing.T) {
	tests := map[string]struct {
		options   settings.Options
		inputData []string
		output    []string
	}{
		"No Flags": {
			options: settings.Options{},
			inputData: []string{"I love music.", "I love music.", "I love music.",
				"", "I love music of Kartik.", "I love music of Kartik.", "Thanks."},
			output: []string{"I love music.", "", "I love music of Kartik.", "Thanks."},
		},
		"With -f flag": {
			options: settings.Options{SkipFields: 1},
			inputData: []string{"We love music.", "I love music.", "They love music.",
				"", "I love music of Kartik.", "We love music of Kartik.", "Thanks."},
			output: []string{"We love music.", "", "I love music of Kartik.", "Thanks."},
		},
		"With -s flag": {
			options: settings.Options{SkipChars: 1},
			inputData: []string{"I love music.", "A love music.", "C love music.",
				"", "I love music of Kartik.", "We love music of Kartik.", "Thanks."},
			output: []string{"I love music.", "", "I love music of Kartik.", "We love music of Kartik.", "Thanks."},
		},
		"With -i flag": {
			options: settings.Options{IgnoreCase: true},
			inputData: []string{"AAAA", "aaaa", "BbBbBb",
				"", "BBBB.", "BbBbBb", ""},
			output: []string{"AAAA", "BbBbBb", "", "BBBB."},
		},
		"With -f and -s": {
			options: settings.Options{SkipChars: 1, SkipFields: 1},
			inputData: []string{"I love music.", "A Bove music.", "C Move music.",
				"", "I love music of Kartik.", "Weasdasd Gove music of Kartik.", "Thanks."},
			output: []string{"I love music.", "", "I love music of Kartik.", "Thanks."},
		},
		"With -f and -s and -i": {
			options: settings.Options{SkipFields: 1, SkipChars: 1, IgnoreCase: true},
			inputData: []string{"I love music.", "A BOve music.", "C Move music.",
				"", "I love music of Kartik.", "Weasdasd GoVe music of KaRtik.", "Thanks."},
			output: []string{"I love music.", "", "I love music of Kartik.", "Thanks."},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			result := DefaultMode(tc.options, tc.inputData)
			require.Equal(t, tc.output, result)
		})
	}
}

func TestDetectDuplicateStrings(t *testing.T) {
	tests := map[string]struct {
		options   settings.Options
		inputData []string
		output    []string
	}{
		"With -d flag": {
			options: settings.Options{Duplicated: true},
			inputData: []string{"I love music.", "I love music.", "I love music.",
				"", "I love music of Kartik.", "I love music of Kartik.", "Thanks."},
			output: []string{"I love music.", "I love music of Kartik."},
		},
		"With -d and -f flag": {
			options: settings.Options{Duplicated: true, SkipFields: 1},
			inputData: []string{"KKK I love music.", "AS I love music.", "BSD I love music.",
				"", "I love music of Kartik.", "I love music of Kartik.", "Thanks."},
			output: []string{"KKK I love music.", "I love music of Kartik."},
		},
		"With -d and -f and -s flag": {
			options: settings.Options{Duplicated: true, SkipFields: 1, SkipChars: 2},
			inputData: []string{"KKK Ia love music.", "AS Ib love music.", "BSD IZ love music.",
				"", "I love music of Kartik.", "I love music of Kartik.", "Thanks."},
			output: []string{"KKK Ia love music.", "I love music of Kartik."},
		},
		"With -d and -f and -s and -i flag": {
			options: settings.Options{Duplicated: true, SkipFields: 1, SkipChars: 2, IgnoreCase: true},
			inputData: []string{"KKK Ia love music.", "AS Ib Love mUsic.", "BSD IZ love music.",
				"", "I loVe music of KARTik.", "I love mUsic of Kartik.", "Thanks."},
			output: []string{"KKK Ia love music.", "I loVe music of KARTik."},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			result := DetectDuplicateStrings(tc.options, tc.inputData)
			require.Equal(t, tc.output, result)
		})
	}
}
func TestDetectUniqueStrings(t *testing.T) {
	tests := map[string]struct {
		options   settings.Options
		inputData []string
		output    []string
	}{
		"With -u flag": {
			options: settings.Options{Unique: true},
			inputData: []string{"I love music.", "I love music.", "I love music.",
				"", "I love music of Kartik.", "I love music of Kartik.", "Thanks."},
			output: []string{"", "Thanks."},
		},
		"With -u and -f flag": {
			options: settings.Options{Unique: true, SkipFields: 1},
			inputData: []string{"Iasd love music.", "I love music.", "I love music.",
				"", "I love music of Kartik.", "I love music of Kartik.", "Thanks."},
			output: []string{"", "Thanks."},
		},
		"With -u and -s flag": {
			options: settings.Options{Unique: true, SkipChars: 5},
			inputData: []string{"Iasd love music.", "zxcI love music.", "I love music.",
				"", "I love music of Kartik.", "I love music of Kartik.", "Thanks."},
			output: []string{"I love music.", "", "Thanks."},
		},
		"With -u and -f and -s flag": {
			options: settings.Options{Unique: true, SkipFields: 1, SkipChars: 1},
			inputData: []string{"Iasd alove music.", "bI nlove music.", "I ]love music.",
				"", "I love music of Kartik.", "I love music of Kartik.", "Thanks."},
			output: []string{"", "Thanks."},
		},
		"With -u and -f and -s and -i flag": {
			options: settings.Options{Unique: true, SkipFields: 1, SkipChars: 1, IgnoreCase: true},
			inputData: []string{"Iasd alove musiC.", "bI nlove mUsic.", "I ]Love music.",
				"", "I love music of Kartik.", "I love music of Kartik.", "Thanks."},
			output: []string{"", "Thanks."},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			result := DetectUniqueStrings(tc.options, tc.inputData)
			require.Equal(t, tc.output, result)
		})
	}
}
func TestCountSubstringsInInput(t *testing.T) {
	tests := map[string]struct {
		options   settings.Options
		inputData []string
		output    []string
	}{
		"With -c flag": {
			options: settings.Options{Count: true},
			inputData: []string{"I love music.", "I love music.", "I love music.",
				"", "I love music of Kartik.", "I love music of Kartik.", "Thanks."},
			output: []string{"3 I love music.", "1 ", "2 I love music of Kartik.", "1 Thanks."},
		},
		"With -c and -f flag": {
			options: settings.Options{Count: true, SkipFields: 1},
			inputData: []string{"I love music.", "asdI love music.", "Ixcv love music.",
				"", "I love music of Kartik.", "Iasdasd love music of Kartik.", "Thanks."},
			output: []string{"3 I love music.", "1 ", "2 I love music of Kartik.", "1 Thanks."},
		},
		"With -c and -s flag": {
			options: settings.Options{Count: true, SkipChars: 4},
			inputData: []string{"zxcI love music.", "asdI love music.", "Ixcv love music.",
				"", "Izxczxc love music of Kartik.", "Iasdzxc love music of Kartik.", "Thanks."},
			output: []string{"3 zxcI love music.", "1 ", "2 Izxczxc love music of Kartik.", "1 Thanks."},
		},
		"With -c and -f and -s flag": {
			options: settings.Options{Count: true, SkipChars: 4, SkipFields: 1},
			inputData: []string{"zxcI love music.", "asdI love music.", "Ixcv love music.",
				"", "Izxczxc love music of Kartik.", "Iasdzxc lve music of Kartik.", "Thanks."},
			output: []string{"3 zxcI love music.", "1 ", "1 Izxczxc love music of Kartik.",
				"1 Iasdzxc lve music of Kartik.", "1 Thanks."},
		},
		"With -c and -f and -s and -i flag": {
			options: settings.Options{Count: true, SkipChars: 4, SkipFields: 1, IgnoreCase: true},
			inputData: []string{"zxcI love mUsIc.", "asdI love music.", "Ixcv love music.",
				"", "Izxczxc love music of Kartik.", "Iasdzxc lve music of KARtik.", "Thanks."},
			output: []string{"3 zxcI love mUsIc.", "1 ", "1 Izxczxc love music of Kartik.",
				"1 Iasdzxc lve music of KARtik.", "1 Thanks."},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			result := CountStringsInInput(tc.options, tc.inputData)
			require.Equal(t, tc.output, result)
		})
	}
}
