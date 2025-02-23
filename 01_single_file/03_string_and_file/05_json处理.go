package main

import (
	"encoding/json"
	"fmt"
)

// 成员变量首字母大写，表示可以被外部包访问
type It struct {
	Company  string   `json:"company"` // struct tag, 通过反射机制获取tag
	Subjects []string `json:"-"`       // - 表示不输出
	Isok     bool     `json:",string"` // 转换为字符串输出
	Price    float64  `json:",string"`
}

func main() {
	// 定义一个It
	var it1 = It{"Google", []string{"Go", "Python", "Java"}, true, 666.666}
	var jsonStr, err = json.Marshal(it1) // 通过结构体生成json字符串
	if err != nil {
		fmt.Println("json marshal error")
	} else {
		fmt.Println(string(jsonStr))
	}

	// 格式化输出
	data, err := json.MarshalIndent(it1, "", "    ")
	if err != nil {
		fmt.Println("json marshal error")
	} else {
		fmt.Println(string(data))
	}

	// 通过map生成json字符串
	m := make(map[string]interface{})
	m["company"] = "Google"
	m["subjects"] = []string{"Go", "Python", "Java"}
	m["isok"] = true
	m["price"] = 666.666
	data, err = json.MarshalIndent(m, "", "    ")
	if err != nil {
		fmt.Println("json marshal error")
	} else {
		fmt.Println(string(data))
	}

	// json字符串转换为结构体
	str := `{"company":"Google","subjects":["Go","Python","Java"],"isok":"true","price":"666.666"}` // json字符串
	var it2 It
	err = json.Unmarshal([]byte(str), &it2)
	if err != nil {
		fmt.Println("json unmarshal error")
	} else {
		fmt.Println(it2)
	}
	// 只想获取subjects
	type It2 struct {
		Subjects []string
	}
	var it3 It2
	err = json.Unmarshal([]byte(str), &it3)
	if err != nil {
		fmt.Println("json unmarshal error")
	} else {
		fmt.Println(it3)
	}

	// json字符串转换为map
	var m2 map[string]interface{}
	err = json.Unmarshal([]byte(str), &m2)
	if err != nil {
		fmt.Println("json unmarshal error")
	} else {
		fmt.Println(m2)
	}

	var s1 string
	s1 = m2["company"].(string) // 类型断言
	fmt.Println(s1)
}
