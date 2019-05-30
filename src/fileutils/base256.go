// base256
package fileutils

import (
	"bytes"

	"github.com/axgle/mahonia"
)

const BASE_CODE string = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz+0123456789亡门义斤爪反介叮电号劝双书幻玉刊示末未击打比互切瓦止少日中不太巧正扑扒功扔去甘内水见午牛手毛气升长子卫乏公仓月氏勿欠田由史只央兄叼叫另叨叹四生失禾丘付仗了力乃刀又三父北占业旧帅归且旦目叶甲申仁什片仆化仇币代仙们仪白仔他斥瓜乎丛令用甩冈贝从今凶分允予云扎艺木五支厅印川亿个勺久凡及夕丸么广历尤友匹车巨牙屯犬区一乙二十丁厂七卜人入八九几儿仍仅飞刃习叉马乡丰王井开夫天无元专世古节本术可丙左厉右石布龙平灭轧东卡风丹匀乌凤勾文六方火为斗忆订计户认心尺引丑巴孔队办以也女之尸弓己已于干亏士工土才寸下大丈与万上小口巾山千乞"

func Base256_encode_gbk(data string) string {
	bytes := []byte(mahonia.NewEncoder("gbk").ConvertString(data))
	return encode(bytes)
}

func Base256_decode_gbk(data string) string {
	decode_str := decode(data)
	return mahonia.NewDecoder("gbk").ConvertString(decode_str)
}

func encode(data_bytes []byte) string {

	var buffer bytes.Buffer
	for _, index := range data_bytes {
		buffer.WriteString(string([]rune(BASE_CODE)[index]))
	}
	return buffer.String()
}

func decode(data string) string {

	var buffer bytes.Buffer
	bytes := []rune(data)
	for _, item := range bytes {
		var index int
		for i, cell := range []rune(BASE_CODE) {
			if item == cell {
				index = i
				break
			}
		}
		buffer.WriteByte(byte(index))
	}

	return buffer.String()
}
