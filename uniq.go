package main

import (
	"uniq/m/actions"
	"uniq/m/settings"
)

func main() {
	opt, err := settings.InitOptions()
	if err != nil {
		panic(err)
	}

	input, err := actions.Input(*opt)
	if err != nil {
		panic(err)
	}
	var output []string
	if opt.Doubled {
		output = actions.DetectDuplicateStrings(*opt, input)
	} else if opt.Unique {
		output = actions.DetectUniqueStrings(*opt, input)
	} else if opt.Count {
		output = actions.CountStringsInInput(*opt, input)
	} else {
		output = actions.DefaultMode(*opt, input)
	}

	if err = actions.Output(*opt, output); err != nil {
		panic(err)
	}
}
