package chinese_commercial_code

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToChineseRuneFromCommercialCode(t *testing.T) {
	chineseRune := ToChineseRuneFromCommercialCode(7115)
	assert.Equal(t, '陈', chineseRune)
}

func TestToChineseRunesFromCommercialCodeSlice(t *testing.T) {
	chineseRunes := ToChineseRunesFromCommercialCodeSlice([]int{7115, 59})
	assert.Equal(t, []rune{'陈', '二'}, chineseRunes)
}

func TestToChineseStringFromCommercialCodeSlice(t *testing.T) {
	chineseString := ToChineseStringFromCommercialCodeSlice([]int{7115, 59})
	assert.Equal(t, "陈二", chineseString)
}
