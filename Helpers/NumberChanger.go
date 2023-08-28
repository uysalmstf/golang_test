package Helpers

import (
	"fmt"
	"strconv"
)

func StrToInt32(str string) (int32, error) {
	fmt.Println("str: ", str)
	num, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}

	return int32(num), nil
}
