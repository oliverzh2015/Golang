// 编写一个小程序：
// 给定一个字符串数组
// [“I”,“am”,“stupid”,“and”,“weak”]
// 用 for 循环遍历该数组并修改为
// [“I”,“am”,“smart”,“and”,“strong”]

package main

import "fmt"

func main() {
	myArray := [5]string{"I", "am", "stupid", "and", "weak"}
	for x, y := range myArray {
		fmt.Println(x, y)
	}

	myArray[2] = "smart"
	myArray[4] = "strong"

	for i := 0; i < len(myArray); i++ {
		fmt.Println(myArray[i])
	}
}
