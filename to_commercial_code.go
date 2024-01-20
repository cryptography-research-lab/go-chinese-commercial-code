package chinese_commercial_code

// ToCommercialCodeFromChineseRune 为单个的中文字符获取对应的电报编码
func ToCommercialCodeFromChineseRune(chineseRune rune) int {
	return ChineseToCommercialCodeMap[chineseRune]
}

// ToCommercialCodeFromChineseRunes 为一组中文字符获取对应的电报编码
func ToCommercialCodeFromChineseRunes(chineseRunes []rune) []int {
	resultSlice := make([]int, len(chineseRunes))
	for index, chineseRune := range chineseRunes {
		resultSlice[index] = ChineseToCommercialCodeMap[chineseRune]
	}
	return resultSlice
}

// ToCommercialCodeSliceFromChineseString 为一句中文编码获取对应的
func ToCommercialCodeSliceFromChineseString(chineseString string) []int {
	return ToCommercialCodeFromChineseRunes([]rune(chineseString))
}
