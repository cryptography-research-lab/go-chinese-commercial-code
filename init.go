package chinese_commercial_code

import (
	"github.com/cryptography-research-lab/go-chinese-commercial-code/data"
	"strconv"
	"strings"
	"unicode/utf8"
)

var (

	// ChineseToCommercialCodeMap 中文到电码的映射
	ChineseToCommercialCodeMap = make(map[rune]int)

	// CommercialCodeToChineseMap 电码到中文的映射
	CommercialCodeToChineseMap = make(map[int]rune)
)

func init() {

	lines := strings.Split(data.Dictionary, "\n")
	for _, line := range lines {
		line = strings.ReplaceAll(line, "\r", "")
		split := strings.Split(line, " ")

		if len(split) != 2 {
			continue
		}

		// 电报码
		commercialCodeString := split[0]
		commercialCode, err := strconv.ParseInt(commercialCodeString, 10, 0)
		if err != nil {
			panic(err)
		}

		// 对应的中文
		chinese := split[1]
		// 只要一个字长的
		if utf8.RuneCountInString(chinese) != 1 {
			continue
		}
		chineseWord := ([]rune)(chinese)[0]

		ChineseToCommercialCodeMap[chineseWord] = int(commercialCode)
		CommercialCodeToChineseMap[int(commercialCode)] = chineseWord
	}

}
