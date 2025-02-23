package main

import (
	"fmt"
	"regexp"
)

func main() {
	buf := `
	<div>
		<span>hello</span>
		<span>world</span>
	</div>
	`
	// 正则表达式提取
	re := regexp.MustCompile(`<span>(?s:(.*?))</span>`)
	if re == nil {
		fmt.Println("regexp error")
		return
	}
	// 提取关键信息
	alls := re.FindAllStringSubmatch(buf, -1)
	fmt.Println("alls = ", alls)

	// 过滤<><>标签
	for _, text := range alls {
		fmt.Println("text[0] = ", text[0]) // text[0] =  <span>hello</span>
		fmt.Println("text[1] = ", text[1]) // text[1] =  hello
	}
}
