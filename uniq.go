package main

import (
	"uniq/m/actions"
	"uniq/m/io"
	"uniq/m/settings"
)

func main() {
	opt, ioOpt, err := settings.InitOptions()
	if err != nil {
		panic(err)
	}

	input, err := io.Input(*ioOpt)
	if err != nil {
		panic(err)
	}

	var output []string
	if opt.Duplicated {
		output = actions.DetectDuplicateStrings(*opt, input)
	} else if opt.Unique {
		output = actions.DetectUniqueStrings(*opt, input)
	} else if opt.Count {
		output = actions.CountStringsInInput(*opt, input)
	} else {
		output = actions.DefaultMode(*opt, input)
	}

	if err = io.Output(*ioOpt, output); err != nil {
		panic(err)
	}
}
