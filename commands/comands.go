package commands

import (
	"bufio"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func GetInput() (string, error) {
	str, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return str, nil
}