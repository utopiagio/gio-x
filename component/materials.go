package component

import (
	"image"
	"image/color"

	"github.com/utopiagio/gioui/gio/layout"
	"github.com/utopiagio/gioui/gio/op/clip"
	"github.com/utopiagio/gioui/gio/op/paint"
)

type (
	C = layout.Context
	D = layout.Dimensions
)

type Rect struct {
	Color color.NRGBA
	Size  image.Point
	Radii int
}

func (r Rect) Layout(gtx C) D {
	paint.FillShape(
		gtx.Ops,
		r.Color,
		clip.UniformRRect(
			image.Rectangle{
				Max: r.Size,
			},
			r.Radii,
		).Op(gtx.Ops))
	return layout.Dimensions{Size: r.Size}
}
