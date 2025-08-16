package helper

import (
	"bufio"
	"os"
	"strings"
)

func GetInput(input string) string {
	reader := bufio.NewReader(os.Stdin)
	input, _ = reader.ReadString('\n')
	return strings.TrimSpace(input)
}
