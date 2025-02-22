package main

func main() {
	var a = 1
	switch a {
	case 1:
		println("a is 1") // go 保留了 break 语句，不过不需要显式写出来，会自动 break
		// fallthrough // 如果需要继续执行下一个 case，可以使用 fallthrough 语句
	case 2:
		println("a is 2")
	default:
		println("a is not 1 or 2")
	}

	// switch 支持初始化语句
	switch b := 2; b {
	case 1:
		println("b is 1")
	case 2, 3, 4: // case 支持多个值
		println("b is 2 or 3 or 4")
	default:
		println("b is not 1 or 2")
	}
	//b = 3 // 无法访问 b，因为 b 的作用域仅限于 switch 语句块内

	score := 90
	// switch 后面可以不跟表达式，相当于 if-else 的用法
	switch {
	case score >= 90:
		println("优秀")
	case score >= 80:
		println("良好")
	case score >= 60:
		println("及格")
	default:
		println("不及格")
	}
}
