package chinese_commercial_code

func ToChineseRuneFromCommercialCode(commercialCode int) rune {
	return CommercialCodeToChineseMap[commercialCode]
}

func ToChineseRunesFromCommercialCodeSlice(commercialCodeSlice []int) []rune {
	resultSlice := make([]rune, len(commercialCodeSlice))
	for index, commercialCode := range commercialCodeSlice {
		resultSlice[index] = CommercialCodeToChineseMap[commercialCode]
	}
	return resultSlice
}

func ToChineseStringFromCommercialCodeSlice(commercialCodeSlice []int) string {
	return string(ToChineseRunesFromCommercialCodeSlice(commercialCodeSlice))
}
