package actions

import (
	"sort"
	"strconv"
	"strings"
	"uniq/m/options"
)

// UniqueStringsIndexes returns indexes of strings in options u need to output
func UniqueStringsIndexes(opt options.Options, input []string) []int {
	var c int
	indexes := make([]int, 0)
	hashset := make(map[string]struct{})
	var splitLine []string
	f, charsToSkip := opt.SkipFields, opt.SkipChars
	for i, line := range input {
		if opt.IgnoreCase {
			splitLine = strings.Split(strings.ToLower(line), " ")
		} else {
			splitLine = strings.Split(line, " ")
		}
		if f < len(splitLine) {
			splitLine = splitLine[f:]
		}
		stringified := ""
		c = charsToSkip
		for _, field := range splitLine {
			for _, elem := range field {
				if c == 0 {
					stringified += string(elem)
				} else {
					c--
				}
			}
			stringified += " "
		}
		if _, exists := hashset[stringified]; !exists {
			indexes = append(indexes, i)
			hashset[stringified] = struct{}{}
		}
	}
	return indexes
}

func DefaultMode(opt options.Options, input []string) []string {
	output := make([]string, 0)
	indexes := UniqueStringsIndexes(opt, input)
	for _, i := range indexes {
		output = append(output, input[i])
	}
	return output
}

type Counter struct {
	count int
	index int
}

// inputIndexCounter returns a hashmap of strings and their count
func inputIndexCounter(opt options.Options, input []string) map[string]Counter {
	var c int
	var splitLine []string
	counter := make(map[string]Counter)
	f, charsToSkip := opt.SkipFields, opt.SkipChars
	for i, line := range input {
		if opt.IgnoreCase {
			splitLine = strings.Split(strings.ToLower(line), " ")
		} else {
			splitLine = strings.Split(line, " ")
		}
		if f < len(splitLine) {
			splitLine = splitLine[f:]
		}
		stringified := ""
		c = charsToSkip
		for _, field := range splitLine {
			for _, elem := range field {
				if c == 0 {
					stringified += string(elem)
				} else {
					c--
				}
			}
			stringified += " "
		}
		if _, exists := counter[stringified]; !exists {
			counter[stringified] = Counter{count: 1, index: i}
		} else {
			counter[stringified] = Counter{count: 2, index: counter[stringified].index}
		}
	}
	return counter
}

func DetectDuplicateStrings(opt options.Options, input []string) []string {
	indexes := make([]int, 0)
	output := make([]string, 0)

	counter := inputIndexCounter(opt, input)

	for _, s := range counter {
		if s.count > 1 {
			indexes = append(indexes, s.index)
		}
	}
	sort.Ints(indexes)
	for _, index := range indexes {
		output = append(output, input[index])
	}
	return output
}

func DetectUniqueStrings(opt options.Options, input []string) []string {
	indexes := make([]int, 0)
	output := make([]string, 0)

	counter := inputIndexCounter(opt, input)

	for _, s := range counter {
		if s.count == 1 {
			indexes = append(indexes, s.index)
		}
	}
	sort.Ints(indexes)
	for _, index := range indexes {
		output = append(output, input[index])
	}
	return output
}

func CountStringsInInput(opt options.Options, input []string) []string {
	output := make([]string, 0)
	count := 0
	indexesOfUnique := UniqueStringsIndexes(opt, input)
	for i, index := range indexesOfUnique {
		if i == len(indexesOfUnique)-1 {
			count = len(input) - index
		} else {
			count = indexesOfUnique[i+1] - indexesOfUnique[i]
		}
		output = append(output, strconv.Itoa(count)+" "+input[index])
	}
	return output
}
