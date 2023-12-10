package io

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"uniq/m/settings"
)

func Input(opt settings.IoOptions) ([]string, error) {
	var output []string
	var reader io.Reader
	if opt.StdinInput {
		reader = os.Stdin
	} else {
		file, err := os.Open(opt.InputFile)

		reader = file
		if err != nil {
			return output, errors.New("Invalid input file: " + opt.InputFile)
		}
		defer func() {
			if err = file.Close(); err != nil {
				panic(err)
			}
		}()
	}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		output = append(output, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return output, err
	}

	return output, nil
}

func Output(opt settings.IoOptions, output []string) error {
	var outputStream io.Writer
	if opt.StdoutOutput {
		outputStream = os.Stdout
	} else {
		file, err := os.Create(opt.OutputFile)

		if err != nil {
			return err
		}

		outputStream = file
		defer func() {
			if err = file.Close(); err != nil {
				panic(err)
			}
		}()
	}
	for _, line := range output {
		if _, err := fmt.Fprintln(outputStream, line); err != nil {
			return err
		}
	}
	return nil
}
