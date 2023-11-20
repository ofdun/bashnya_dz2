package actions

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"uniq/m/settings"
)

func Output(opt settings.Options, output []string) error {
	if opt.StdoutOutput {
		for _, line := range output {
			fmt.Println(line)
		}
	} else {
		f, err := os.Create(opt.OutputFile)
		if err != nil {
			return err
		}
		defer func() {
			if err = f.Close(); err != nil {
				panic(err)
			}
		}()
		for _, line := range output {
			_, err = f.WriteString(line + "\n")
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func Input(opt settings.Options) ([]string, error) {
	var output []string

	if !opt.StdinInput {
		reader, err := os.Open(opt.InputFile)

		if err != nil {
			return output, errors.New("Invalid input file: " + opt.InputFile)
		}

		defer func() {
			if err = reader.Close(); err != nil {
				panic(err)
			}
		}()
		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			output = append(output, scanner.Text())
		}

		if err = scanner.Err(); err != nil {
			return output, err
		}
	} else {
		reader := os.Stdin
		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			output = append(output, scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			return output, err
		}
	}
	return output, nil
}

// UniqueStringsIndexes returns indexes of strings in options u need to output
func UniqueStringsIndexes(opt settings.Options, input []string) []int {
	var c int
	indexes := make([]int, 0)
	hashset := make(map[string]struct{})
	f, charsToSkip := opt.SkipFields, opt.SkipChars
	for i, line := range input {
		splitLine := make([]string, 0)
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

func DefaultMode(opt settings.Options, input []string) []string {
	output := make([]string, 0)
	indexes := UniqueStringsIndexes(opt, input)
	for _, i := range indexes {
		output = append(output, input[i])
	}
	return output
}

// stringsIndexesCounter returns a hashmap of strings and their count
func stringsIndexesCounter(opt settings.Options, input []string) map[string]struct {
	count int
	index int
} {
	var c int
	counter := make(map[string]struct {
		count int
		index int
	})
	f, charsToSkip := opt.SkipFields, opt.SkipChars
	for i, line := range input {
		splitLine := make([]string, 0)
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
			counter[stringified] = struct {
				count int
				index int
			}{count: 1, index: i}
		} else {
			counter[stringified] = struct {
				count int
				index int
			}{count: 2, index: counter[stringified].index}
		}
	}
	return counter
}

func DetectDuplicateStrings(opt settings.Options, input []string) []string {
	indexes := make([]int, 0)
	output := make([]string, 0)

	counter := stringsIndexesCounter(opt, input)

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

func DetectUniqueStrings(opt settings.Options, input []string) []string {
	indexes := make([]int, 0)
	output := make([]string, 0)

	counter := stringsIndexesCounter(opt, input)

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

func CountStringsInInput(opt settings.Options, input []string) []string {
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
