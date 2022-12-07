package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const MAX_LETTER_COUNT = 9

var ErrInvalidString = errors.New("Invalid string")

var CountRegexp = regexp.MustCompile(`(?P<letter>[\p{L}])(?P<count>[0-9]*)?`)

func unpack(s string) (string, error) {
	if len(s) == 0 {
		return "", nil
	}

	if !CountRegexp.MatchString(s) {
		return "", ErrInvalidString
	}

	matches := CountRegexp.FindAllStringSubmatch(s, -1)
	var builder strings.Builder

	for _, match := range matches {
		letter, countRaw := match[1], match[2]

		countInt, err := strconv.Atoi(countRaw)
		if err != nil && len(countRaw) == 0 {
			countInt = 1
		}

		if countInt == 0 || countInt > MAX_LETTER_COUNT {
			return "", ErrInvalidString
		}

		letterRepeat := strings.Repeat(letter, countInt)
		builder.WriteString(letterRepeat)

	}

	return builder.String(), nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter string to unpack:")

	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Print(err)
	}

	result, err := unpack(text)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Print(result)
}
