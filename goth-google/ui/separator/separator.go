package separator

import (
	"github.com/plainkit/html"
)

type Orientation string
type Decoration string

type Props struct {
	ID          string
	Class       string
	Attrs       []html.Global
	Orientation Orientation
	Decoration  Decoration
}

const (
	OrientationHorizontal Orientation = "horizontal"
	OrientationVertical   Orientation = "vertical"
)

const (
	DecorationDashed Decoration = "dashed"
	DecorationDotted Decoration = "dotted"
)

func divArgsFromProps(baseClass string, extra ...string) func(p Props) []html.DivArg {
	return func(p Props) []html.DivArg {
		args := []html.DivArg{html.AClass(html.ClassMerge(append([]string{baseClass}, append(extra, p.Class)...)...))}
		if p.ID != "" {
			args = append(args, html.AId(p.ID))
		}

		for _, a := range p.Attrs {
			args = append(args, a)
		}

		return args
	}
}

// ApplyDiv implements the html.DivArg interface for Props
func (p Props) ApplyDiv(attrs *html.DivAttrs, children *[]html.Component) {
	orientation := p.Orientation
	if orientation == "" {
		orientation = OrientationHorizontal
	}

	var args []html.DivArg
	if orientation == OrientationVertical {
		args = divArgsFromProps("shrink-0 h-full")(p)
		args = append([]html.DivArg{
			html.ACustom("role", "separator"),
			html.AAria("orientation", "vertical"),
		}, args...)
	} else {
		args = divArgsFromProps("shrink-0 w-full")(p)
		args = append([]html.DivArg{
			html.ACustom("role", "separator"),
			html.AAria("orientation", "horizontal"),
		}, args...)
	}

	for _, a := range args {
		a.ApplyDiv(attrs, children)
	}
}

// Separator renders a decorative divider using the composable pattern.
// Accepts variadic html.DivArg arguments, with Props as an optional first argument.
func Separator(args ...html.DivArg) html.Node {
	var (
		props Props
		rest  []html.DivArg
	)

	// Separate Props from other arguments
	for _, a := range args {
		if v, ok := a.(Props); ok {
			props = v
		} else {
			rest = append(rest, a)
		}
	}

	if props.Orientation == "" {
		props.Orientation = OrientationHorizontal
	}

	if props.Orientation == OrientationVertical {
		return verticalSeparator(props, rest)
	}

	return horizontalSeparator(props, rest)
}

func horizontalSeparator(props Props, divArgs []html.DivArg) html.Node {
	outerArgs := append([]html.DivArg{props}, divArgs...)

	inner := html.Div(
		html.AClass("relative flex items-center w-full"),
		html.Span(
			html.AClass(html.ClassMerge("absolute w-full border-t h-[1px]", decorationClass(props.Decoration))),
			html.ACustom("aria-hidden", "true"),
		),
	)

	outerArgs = append(outerArgs, inner)

	return html.Div(outerArgs...)
}

func verticalSeparator(props Props, divArgs []html.DivArg) html.Node {
	outerArgs := append([]html.DivArg{props}, divArgs...)

	inner := html.Div(
		html.AClass("relative flex flex-col items-center h-full"),
		html.Span(
			html.AClass(html.ClassMerge("absolute h-full border-l w-[1px]", decorationClass(props.Decoration))),
			html.ACustom("aria-hidden", "true"),
		),
	)

	outerArgs = append(outerArgs, inner)

	return html.Div(outerArgs...)
}

func decorationClass(dec Decoration) string {
	switch dec {
	case DecorationDashed:
		return "border-dashed"
	case DecorationDotted:
		return "border-dotted"
	default:
		return ""
	}
}
