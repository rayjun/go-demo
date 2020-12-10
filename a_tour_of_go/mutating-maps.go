package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42                        // 插入或者更新一个值
	fmt.Println("The value: ", m["Answer"]) // 取出一个值

	m["Answer"] = 48
	fmt.Println("The value: ", m["Answer"])

	delete(m, "Answer") // 删除一个元素
	fmt.Println("The value: ", m["Answer"])

	v, ok := m["Answer"] // 判断 map 中某个元素是否存在
	fmt.Println("The value: ", v, "Present?", ok)
}
