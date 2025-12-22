package functions

import (
	"fmt"

	"github.com/fogleman/gg"
)

func Paint(pi *[]*PaintInstruction) {
	dc := gg.NewContext(1000, 1000)

	dc.SetRGB(1, 1, 1)
	dc.Clear()

	dc.SetRGB(0, 0, 0)

	if err := dc.LoadFontFace("functions/fonts.ttf", 16); err != nil {
		panic(err)
	}

	for _, i := range *pi {
		switch i.Function {
		case "WriteText":
			fmt.Println(i.CursorX, i.CursorY)
			dc.DrawString(i.Content, float64(i.CursorX), float64(i.CursorY))
		}
	}

	dc.SavePNG("output.png")
}
