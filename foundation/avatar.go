package foundation

import (
	"github.com/lyupengpublish/golibs/myutils/calc"
	"image/color"
)
var colors = []uint32{
	0xff6200, 0x42c58e, 0x5a8de1, 0x785fe0,
}
//生成头像的方式
func GeneraterAvater(width int,height int,content string) []byte {
	ab := calc.NewAvatarBuilder("../static/SourceHanSansSC-Medium.ttf", &calc.SourceHansSansSCMedium{})
	ab.SetBackgroundColorHex(colors[3])
	ab.SetFrontgroundColor(color.White)
	ab.SetFontSize(80)
	ab.SetAvatarSize(width, height)
	if data,err:=ab.GenerateImage(content);err==nil {
		return  data
	}
	return nil
}
