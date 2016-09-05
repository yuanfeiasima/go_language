package main

import (
	"errors"
	"fmt"
)

//错误类型
func main() {

	err := errors.New("this is error")
	if err != nil {
		fmt.Print(err)
	}



}
