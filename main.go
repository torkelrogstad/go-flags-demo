package main

import (
	"errors"
	"log"

	"github.com/miraiex/go-flags"
)

type baseConfig struct {
	Base string `long:"base"`
}

type firstConfig struct {
	First int `long:"first"`
}

// Execute implements flags.Commander
func (c *firstConfig) Execute(args []string) error {
	log.Printf("i'm executing the first subcommand: %+v", c)

	return nil
}

type secondConfig struct {
	Second bool `long:"second"`
}

// Execute implements flags.Commander
func (c *secondConfig) Execute(args []string) error {
	log.Printf("i'm executing the second subcommand: %+v", c)

	return nil
}

func main() {
	parser := flags.NewParser(&baseConfig{}, flags.Default)
	parser.AddCommand("first", "", "", new(firstConfig))
	parser.AddCommand("second", "", "", new(secondConfig))
	_, err := parser.Parse()
	if err == nil {
		return
	}

	if fErr := new(flags.Error); errors.As(err, &fErr) && fErr.Type == flags.ErrHelp {
		return
	}

	log.Fatal(err)
}
