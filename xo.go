// An armor format for Crockford-base32 messages.

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var alphabet = []string{
	"ooooo", "oooox", "oooxo", "oooxx", "ooxoo",
	"ooxox", "ooxxo", "ooxxx", "oxooo", "oxoox",
	"oxoxo", "oxoxx", "oxxoo", "oxxox", "oxxxo",
	"oxxxx", "xoooo", "xooox", "xooxo", "xooxx",
	"xoxoo", "xoxox", "xoxxo", "xoxxx", "xxooo",
	"xxoox", "xxoxo", "xxoxx", "xxxoo", "xxxox",
	"xxxxo", "xxxxx",
}

func e(input string) string {
	var encoded strings.Builder
	for _, r := range input {
		if r == ' ' {
			fmt.Println("the input must not contain spaces")
			os.Exit(1)
		}
		index := strings.IndexRune("0123456789ABCDEFGHJKMNPQRSTVWXYZ", r)
		if index == -1 {
			fmt.Println("invalid character")
			os.Exit(1)
		}
		encoded.WriteString(alphabet[index])
	}
	return encoded.String()
}

func d(input string) string {
	var decoded strings.Builder
	groups := strings.Fields(input) // groups the input by spaces
	for _, group := range groups {
		index := -1
		for j, a := range alphabet {
			if a == group {
				index = j
				break
			}
		}
		if index == -1 {
			fmt.Println("invalid group")
			os.Exit(1)
		}
		char := rune("0123456789ABCDEFGHJKMNPQRSTVWXYZ"[index])
		decoded.WriteRune(char)
	}
	return decoded.String()
}

func formatGroups(encoded string) string {
	var formatted strings.Builder
	groups := strings.Split(encoded, "")

	for i, group := range groups {
		formatted.WriteString(group)
		if (i+1)%5 == 0 && (i+1)%30 != 0 {
			formatted.WriteString(" ")
		}
		if (i+1)%30 == 0 {
			formatted.WriteString("\n")
		}
	}

	return formatted.String()
}

func usage() {
    fmt.Println("usage: xo [-d]")
    fmt.Println("  -d: decodes the input from stdin")
    fmt.Println("  if no flag is provided, encodes the input from stdin")
}

func main() {
    flag.Usage = func() {
        usage()
        os.Exit(2)
    }

    decodeFlag := flag.Bool("d", false, "decodes the input from stdin")
    flag.Parse()

    if *decodeFlag {
        inputBytes, err := ioutil.ReadAll(os.Stdin)
        if err != nil {
            fmt.Println(err)
            os.Exit(1)
        }
        input := strings.TrimSpace(string(inputBytes))

        decoded := d(input)
        fmt.Println(decoded)
    } else {
        inputBytes, err := ioutil.ReadAll(os.Stdin)
        if err != nil {
            fmt.Println(err)
            os.Exit(1)
        }
        input := strings.TrimSpace(string(inputBytes))
        input = strings.ReplaceAll(input, "\n", "")

        encoded := e(input)
        formatted := formatGroups(encoded)
        fmt.Println(formatted)
    }
}
