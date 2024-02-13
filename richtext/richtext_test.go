package richtext

import (
	"image"
	"testing"
	"time"

	"github.com/utopiagio/gio/font/gofont"
	"github.com/utopiagio/gio/io/router"
	"github.com/utopiagio/gio/layout"
	"github.com/utopiagio/gio/op"
	"github.com/utopiagio/gio/text"
	"github.com/utopiagio/gio/unit"
	"github.com/utopiagio/gio/widget/material"
)

// TestNilInteractiveText ensures that it is safe to lay out
// richtext with a nil state when none of the spans are
// interactive.
func TestNilInteractiveText(t *testing.T) {
	th := material.NewTheme()
	th.Shaper = text.NewShaper(text.WithCollection(gofont.Collection()))
	spans := []SpanStyle{
		{
			Size:    12,
			Content: "Hello",
		},
		{
			Size:    12,
			Content: "world",
		},
	}
	var ops op.Ops
	gtx := layout.Context{
		Constraints: layout.Exact(image.Pt(100, 100)),
		Metric: unit.Metric{
			PxPerDp: 1,
			PxPerSp: 1,
		},
		Queue: &router.Router{},
		Now:   time.Now(),
		Ops:   &ops,
	}

	Text(nil, th.Shaper, spans...).Layout(gtx)
}
