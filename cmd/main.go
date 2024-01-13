// Description: Main file for the password generator CLI

package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

const (
	LOWER_CASE_LETTERS = "abcdefghijklmnopqrstuvwxyz"
	UPPER_CASE_LETTERS = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	NUMBERS            = "0123456789"
	SPECIAL_CHARACTERS = "!@#$%^&*()_+{}[]"
	MAX_LENGTH         = 100
	DEFAULT_LENGTH     = 10
)

type PasswordGenerationRequest struct {
	length    int
	uppercase bool
	lowercase bool
	numbers   bool
	symbols   bool
}

func generatePassword(input PasswordGenerationRequest) (string, error) {
	length := input.length
	str := ""

	if length > MAX_LENGTH {
		return "", errors.New(color.RedString("Length cannot be greater than %d\n", MAX_LENGTH))
	}

	if length == 0 {
		length = DEFAULT_LENGTH
	}

	if input.lowercase {
		str += LOWER_CASE_LETTERS
	}

	if input.uppercase {
		str += UPPER_CASE_LETTERS
	}

	if input.numbers {
		str += NUMBERS
	}

	if input.symbols {
		str += SPECIAL_CHARACTERS
	}

	chars := strings.Split(str, "")

	password := ""
	for i := 0; i < length; i++ {
		randomIndex := rand.Intn(len(chars))
		password += chars[randomIndex]
	}

	if password == "" {
		return "", errors.New(color.YellowString("Please select at least one option.\n"))
	}

	return password, nil
}

func main() {

	app := &cli.App{
		Name:  "password-generator",
		Usage: "Generate a random password",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    "length",
				Aliases: []string{"l"},
				Value:   DEFAULT_LENGTH,
				Usage:   fmt.Sprintf("Length of the password, Max %s", color.YellowString(strconv.Itoa(MAX_LENGTH))),
			},
			&cli.BoolFlag{
				Name:    "uppercase",
				Aliases: []string{"uc"},
				Value:   false,
				Usage:   "Include uppercase letters",
			},
			&cli.BoolFlag{
				Name:    "lowercase",
				Aliases: []string{"lc"},
				Value:   true,
				Usage:   "Include lowercase letters",
			},
			&cli.BoolFlag{
				Name:    "numbers",
				Aliases: []string{"n"},
				Value:   false,
				Usage:   "Include numbers",
			},
			&cli.BoolFlag{
				Name:    "symbols",
				Aliases: []string{"s"},
				Value:   false,
				Usage:   "Include special characters",
			},
		},
		Action: func(c *cli.Context) error {

			input := PasswordGenerationRequest{
				length:    c.Int("length"),
				uppercase: c.Bool("uppercase"),
				lowercase: c.Bool("lowercase"),
				symbols:   c.Bool("symbols"),
				numbers:   c.Bool("numbers"),
			}

			password, err := generatePassword(input)

			if err != nil {
				fmt.Println(err)
			}

			if password != "" {
				fmt.Println("Your password is:", color.CyanString(password))
			}

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
