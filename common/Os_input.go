package common

import (
	"bufio"
	"os"
)

func ReadDateFromInput() string {
	reader := bufio.NewReader(os.Stdin)
	data, _, err := reader.ReadLine()
	CheckError(err)
	return string(data)
}
