package main

func main() {
	// 复数类型
	var c1 complex64
	c1 = 1 + 2i
	var c2 complex64
	c2 = 2 + 3i
	var c3 = c1 + c2
	println("c3 = ", c3)             // (3+5i)
	println("real(c3) = ", real(c3)) // 3
	println("imag(c3) = ", imag(c3)) // 5
}
