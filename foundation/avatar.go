package foundation

import (
	"github.com/gobuffalo/packr"
	"github.com/lyupengpublish/golibs/myutils/calc"
	"image/color"
)
var colors = []uint32{
	0xff6200, 0x42c58e, 0x5a8de1, 0x785fe0,
}
//生成头像的方式
func GenerateAvatar(width int,height int,content string)( []byte,error) {
	box := packr.NewBox("../static")
	s, err := box.Find("SourceHanSansSC-Medium.ttf")
	if err!=nil {
		return nil,err
	}
	ab := calc.NewAvatarBuilder(s, &calc.SourceHansSansSCMedium{})
	ab.SetBackgroundColorHex(colors[3])
	ab.SetFrontgroundColor(color.White)
	ab.SetFontSize(80)
	ab.SetAvatarSize(width, height)
	data,err:=ab.GenerateImage(content)
	if err!=nil {
		return nil,err
	}
	return  data,err

}
