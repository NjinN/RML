package view

import (

	. "github.com/NjinN/RML/go/core"

	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

type RForm struct {
	*vcl.TForm
	Wrap 	*View
}


func NewRForm(attr *Token, child *Token) *Token {
	if attr.Tp != OBJECT || child.Tp != BLOCK {
		return &Token{ERR, "Type Mismatch"}
	}

	var result Token
	result.Tp = VIEW

	var f *RForm
	
	vcl.Application.CreateForm(&f)

	var v View
	v.Tp = "Form"
	v.Attr = attr.Ctx().Clone()
	v.Raw = f
	v.Child = NewTks(4)

	f.Wrap = &v

	result.Val = &v
	return &result
}

func (f RForm) Update() {
	vcl.ThreadSync(func(){
		title := f.Wrap.Attr.GetNow("title")
		if title.Tp == STRING {
			f.SetCaption(title.Str())
		}

		size := f.Wrap.Attr.GetNow("size")
		if size.Tp == PAIR {
			f.SetClientWidth(int32(size.Tks()[0].Int()))
			f.SetClientHeight(int32(size.Tks()[1].Int()))
		}

		autoSize := f.Wrap.Attr.GetNow("auto-size")
		if autoSize.Tp == LOGIC {
			f.SetAutoSize(autoSize.ToBool())
		}

		position := f.Wrap.Attr.GetNow("position")
		if position.Tp == INTEGER {
			f.SetPosition(types.TPosition(position.Int()))
		}

		offset := f.Wrap.Attr.GetNow("offset")
		if offset.Tp == PAIR {
			f.SetTop(int32(offset.Tks()[0].Int()))
			f.SetLeft(int32(offset.Tks()[1].Int()))
		}

		align := f.Wrap.Attr.GetNow("align")
		if align.Tp == STRING {
			if align.Str() == "top" {
				f.SetAlign(types.AlTop)
			}else if align.Str() == "bottom" {
				f.SetAlign(types.AlBottom)
			}else if align.Str() == "left" {
				f.SetAlign(types.AlLeft)
			}else if align.Str() == "right" {
				f.SetAlign(types.AlRight)
			}else if align.Str() == "client" {
				f.SetAlign(types.AlClient)
			}else if align.Str() == "custom" {
				f.SetAlign(types.AlCustom)
			}
		}

		buttons := f.Wrap.Attr.GetNow("buttons") 
		if buttons.Tp == INTEGER {
			f.SetBorderIcons(types.TSet(buttons.Int()))
		}
	})
}
