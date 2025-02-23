package main

import "errors"

func MyDiv(a, b int) (result int, err error) {
	if b == 0 {
		err = errors.New("除数不能为0") // 普通错误
	} else {
		result = a / b
	}
	return
}

func main() {
	result, err := MyDiv(10, 0)
	if err != nil {
		println(err.Error())
	} else {
		println(result)
	}
}
