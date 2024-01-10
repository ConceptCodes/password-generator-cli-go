# Password Generator CLI

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)


A simple command-line tool for generating random passwords.

## Prerequisites

- [![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)](https://golang.org/doc/install) 

## Installation

1. Clone this repository:

   ```sh
   git clone https://github.com/conceptcodes/password-generator-cli-go.git
   cd password-generator-cli-go
   ```

2. Navigate to the CLI directory and build the application:

   ```sh
    go build -o password-generator cmd/main.go
   ```

3. Run the CLI:

   ```sh
    ./password-generator --help
   ```

If you want to install the CLI globally, you can run the following command:

(Mac and Linux)
```sh
sudo mv password-generator /usr/local/bin
```

## Usage
To run the CLI, use the following command:

```sh
password-generator --length 12 --uppercase --lowercase --numbers --symbols

NAME:
   password-generator - Generate a random password

USAGE:
   password-generator [global options] command [command options] 

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --length value  Length of the password, Max 100 (default: 10)
   --uppercase     Include uppercase letters (default: false)
   --lowercase     Include lowercase letters (default: false)
   --numbers       Include numbers (default: false)
   --symbols       Include special characters (default: false)
   --help, -h      show help
 
```