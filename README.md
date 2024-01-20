# 中文电码 

# 一、中文电码是什么？应用场景？

中文电码，又称：中文商用电码（Chinese commercial code, CCC）、中文电报码（Chinese telegraph code, CTC）或中文电报明码（Chinese ordinary telegraph code, COTC），原本是用于电报之中传送中文信息的方法。它是第一个将汉字化作电子信号的编码表。

自摩尔斯电码在1835年发明后，一直只能用来传送英语或以拉丁字母拼写的文字。在1880年，清朝政府雇用丹麦人（也有说是法国人）设计了中文汉字电报。

中文电码表采用了四位阿拉伯数字作代号，简称“四码电报”，从0001到9999按四位数顺序排列，用四位数字表示最多一万个汉字、字母和符号。汉字先按部首，后按笔划排列。字母和符号放到电码表的最尾。后来由于一万个汉字不足以应付户籍管理的要求，又有第二字面汉字的出现。在香港，两个字面都采用同一编码，由输入员人手选择字面；在台湾，第二字面的汉字会在开首补上“1”字，变成5个数字的编码。

中国汉字多达6万字，常用的汉字只有一万个左右， 所以用10的4次方（10,000）来表示。中文电码表采用了四位阿拉伯数字作代号，简称“四码电报”，从0001到9999按四位数顺序排列，用四位数字表示最多一万个汉字、字母和符号。汉字先按部首，后按笔划排列。字母和符号放到电码表的最尾。后来由于一万个汉字不足以应付户籍管理的要求，又有第二字面汉字的出现。在香港，两个字面都采用同一编码，由输入员人手选择字面；在台湾，第二字面的汉字会在开首补上“1”字，变成5个数字的编码。

中文电码应用

中文电码可用作电脑里的中文输入法，但因中文电码是“无理码”，记忆困难，一般用户几乎无法熟练地掌握使用。

在香港，每个有中文姓名的市民的身份证上，均会在他的姓名下面印有中文电码，外国人取得的入港签证亦有印上。在很多政府或商业机构的表格中，都会要求填写者填写他的中文电码，以便输入电脑。

办理签证时，有些国家（美国、英国、澳大利亚等）需要填写中文电码。比如：申请美国签证的DS156表，在姓名一项，要求填写“Chinese Commercial Code Number”，即为姓名的中文电码。

# 二、安装依赖

```bash
go get -u github.com/cryptography-research-lab/go-chinese-commercial-code
```

# 三、API示例

```go
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
```


# 四、参考资料
- [附錄:中文电码/中国大陆1983](https://zh.wiktionary.org/wiki/Appendix:%E4%B8%AD%E6%96%87%E7%94%B5%E7%A0%81/%E4%B8%AD%E5%9B%BD%E5%A4%A7%E9%99%861983)
- [标准中文电码(Chinese Commercial Code)查询](https://apps.chasedream.com/chinese-commercial-code/)
- [姓名中汉字无对应中文电码的处理办法](https://www.chasedream.com/show.aspx?id=4488&cid=30)

