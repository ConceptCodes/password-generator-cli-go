package main

import (
	"fmt"
	"math/rand"
	"strings"

	"os"

	"github.com/urfave/cli/v2"
)

const (
	LOWER_CASE_LETTERS = "abcdefghijklmnopqrstuvwxyz"
	NUMBERS            = "0123456789"
	SPECIAL_CHARACTERS = "!@#$%^&*()_+{}[]"
	MAX_LENGTH         = 100
	DEFAULT_LENGTH     = 10
)

var UPPER_CASE_LETTERS = strings.ToUpper(LOWER_CASE_LETTERS)

func logPassword(password string) {
	fmt.Println("Your password is: ", password)
}

func main() {

	app := &cli.App{
		Name:  "password-generator",
		Usage: "Generate a random password",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "length",
				Aliases: []string{"l"},
				Value: DEFAULT_LENGTH,
				Usage: fmt.Sprintf("Length of the password, Max %d", MAX_LENGTH),
			},
			&cli.BoolFlag{
				Name:  "uppercase",
				Aliases: []string{"uc"},
				Value: false,
				Usage: "Include uppercase letters",
			},
			&cli.BoolFlag{
				Name:  "lowercase",
				Aliases: []string{"lc"},
				Value: true,
				Usage: "Include lowercase letters",
			},
			&cli.BoolFlag{
				Name:  "numbers",
				Aliases: []string{"n"},
				Value: false,
				Usage: "Include numbers",
			},
			&cli.BoolFlag{
				Name:  "symbols",
				Aliases: []string{"s"},
				Value: false,
				Usage: "Include special characters",
			},
		},
		Action: func(c *cli.Context) error {
			if !c.IsSet("lowercase") && !c.IsSet("uppercase") {
				fmt.Println("No arguments were provided. Please provide arguments.")
				cli.ShowAppHelp(c)
				return nil
			}

			str := ""

			if c.Bool("lowercase") {
				str += LOWER_CASE_LETTERS
			}

			if c.Bool("uppercase") {
				str += UPPER_CASE_LETTERS
			}

			if c.Bool("numbers") {
				str += NUMBERS
			}

			if c.Bool("symbols") {
				str += SPECIAL_CHARACTERS
			}

			chars := strings.Split(str, "")
			length := c.Int("length")

			if length == 0 || length > MAX_LENGTH {
				length = DEFAULT_LENGTH
			}

			password := ""
			for i := 0; i < length; i++ {
				randomIndex := rand.Intn(len(chars))
				password += chars[randomIndex]
			}

			if password == "" {
				fmt.Println("Please select at least one option")
				return nil
			}

			logPassword(password)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
