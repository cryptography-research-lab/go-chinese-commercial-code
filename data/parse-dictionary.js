/**
 *
 * 用于从这个页面上解析出字典：
 *
 * https://zh.wiktionary.org/wiki/Appendix:%E4%B8%AD%E6%96%87%E7%94%B5%E7%A0%81/%E4%B8%AD%E5%9B%BD%E5%A4%A7%E9%99%861983#04##
 *
 * @type {NodeListOf<Element>}
 */
const items = document.querySelectorAll(".wikitable td");
const itemArray = [];
items.forEach(x => {
    // 数据样例： '0100\n他'
    const s = x.innerText;
    const item = s.replace("\n", " ");
    itemArray.push(item);
});
console.log(itemArray.join("\n"));
