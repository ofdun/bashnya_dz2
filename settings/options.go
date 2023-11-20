package settings

import (
	"errors"
	"flag"
)

type Options struct {
	Count        bool
	Doubled      bool
	Unique       bool
	SkipFields   int
	SkipChars    int
	IgnoreCase   bool
	StdinInput   bool
	StdoutOutput bool
	InputFile    string
	OutputFile   string
}

func NewOptions(c, d, u, i, stdin, stdout bool,
	skipf, skipc int,
	inputfile, outputfile string) Options {
	return Options{c, d, u, skipf, skipc,
		i, stdin, stdout, inputfile, outputfile}
}

func InitOptions() (*Options, error) {
	countPtr := flag.Bool("c", false, "Count number of each string and output them")
	doubledPtr := flag.Bool("d", false, "Output only duplicated lines")
	uniquePtr := flag.Bool("u", false, "Output only unique lines")
	skipfPtr := flag.Int("f", 0, "Doesn't count first n fields in a string")
	skipCPtr := flag.Int("s", 0, "Doesn't count first n chars in a string")
	ignorePtr := flag.Bool("i", false, "Ignore case")
	stdin := true
	inputFilePath := "stdin"
	stdout := true
	outputFilePath := "stdout"

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
		return nil, errors.New("Invalid input! -с, -d and -u are interchangeable! ")
	}

	opt := NewOptions(c, d, u, *ignorePtr, stdin, stdout,
		*skipfPtr, *skipCPtr, inputFilePath, outputFilePath)

	return &opt, nil
	//return initOptions(c, d, u, *skipfPtr, *skipCPtr, *ignorePtr), nil
}
