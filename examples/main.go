package main

import (
	"fmt"
	chinese_commercial_code "github.com/cryptography-research-lab/go-chinese-commercial-code"
)

func main() {

	// 对中文进行电报码编码
	commercialCodeSlice := chinese_commercial_code.ToCommercialCodeSliceFromChineseString("加密算法很有趣")
	fmt.Println(commercialCodeSlice) // Output: [502 1378 4615 3127 1771 2589 6393]

	// 将电报码编码转换为中文
	chineseString := chinese_commercial_code.ToChineseStringFromCommercialCodeSlice(commercialCodeSlice)
	fmt.Println(chineseString) // Output: 加密算法很有趣

}
