package chinese_commercial_code

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToCommercialCodeFromChineseRune(t *testing.T) {
	commercialCode := ToCommercialCodeFromChineseRune('陈')
	assert.Equal(t, 7115, commercialCode)
}

func TestToCommercialCodeFromChineseRunes(t *testing.T) {
	commercialCodeSlice := ToCommercialCodeFromChineseRunes([]rune{'陈', '二'})
	assert.Equal(t, []int{7115, 59}, commercialCodeSlice)
}

func TestToCommercialCodeSliceFromChineseString(t *testing.T) {
	commercialCodeSlice := ToCommercialCodeSliceFromChineseString("陈二")
	assert.Equal(t, []int{7115, 59}, commercialCodeSlice)
}
