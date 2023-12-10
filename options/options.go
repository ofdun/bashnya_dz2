package options

import (
	"errors"
	"flag"
)

type Options struct {
	Count      bool
	Duplicated bool
	Unique     bool
	SkipFields int
	SkipChars  int
	IgnoreCase bool
}

type IoOptions struct {
	StdinInput   bool
	StdoutOutput bool
	InputFile    string
	OutputFile   string
}

func InitOptions() (*Options, *IoOptions, error) {
	countPtr := flag.Bool("c", false, "Count number of each string and output them")
	doubledPtr := flag.Bool("d", false, "Output only duplicated lines")
	uniquePtr := flag.Bool("u", false, "Output only unique lines")
	skipfPtr := flag.Int("f", 0, "Doesn't count first n fields in a string")
	skipCPtr := flag.Int("s", 0, "Doesn't count first n chars in a string")
	ignorePtr := flag.Bool("i", false, "Ignore case")
	stdin := true
	stdout := true
	var inputFilePath, outputFilePath string

	flag.Parse()
	if flag.NArg() == 1 {
		inputFilePath = flag.Arg(0)
		stdin = false
	} else if flag.NArg() == 2 {
		inputFilePath = flag.Arg(0)
		outputFilePath = flag.Arg(1)
		stdin = false
		stdout = false
	}

	c := *countPtr
	d := *doubledPtr
	u := *uniquePtr
	if (c && (d || u)) || (d && (c || u)) || (u && (c || d)) {
		return nil, nil, errors.New("Invalid input! -с, -d and -u are interchangeable! ")
	}

	opt := Options{c, d, u, *skipfPtr, *skipCPtr, *ignorePtr}
	ioOpt := IoOptions{stdin, stdout, inputFilePath, outputFilePath}

	return &opt, &ioOpt, nil
}
