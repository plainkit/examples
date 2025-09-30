package avatar

import (
	"fmt"

	"github.com/plainkit/html"
)

type Size string
type GroupSpacing string

type Props struct {
	ID      string
	Class   string
	Attrs   []html.Global
	Size    Size
	InGroup bool
}

type ImageProps struct {
	ID    string
	Class string
	Attrs []html.Global
	Alt   string
	Src   string
}

type FallbackProps struct {
	ID    string
	Class string
	Attrs []html.Global
}

type GroupProps struct {
	ID      string
	Class   string
	Attrs   []html.Global
	Spacing GroupSpacing
}

const (
	SizeSm Size = "sm"
	SizeMd Size = "md"
	SizeLg Size = "lg"
)

const (
	GroupSpacingSm GroupSpacing = "sm"
	GroupSpacingMd GroupSpacing = "md"
	GroupSpacingLg GroupSpacing = "lg"
)

func divArgsFromProps(baseClass string, extra ...string) func(p Props) []html.DivArg {
	return func(p Props) []html.DivArg {
		className := html.ClassMerge(
			append([]string{baseClass},
				append(extra, p.Class)...)...)

		args := []html.DivArg{html.AClass(className)}

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
	args := divArgsFromProps(
		"relative inline-flex shrink-0 items-center justify-center overflow-hidden",
		"size-12 text-base",
		"data-[pui-avatar-size=sm]:size-8 data-[pui-avatar-size=sm]:text-xs",
		"data-[pui-avatar-size=md]:size-12 data-[pui-avatar-size=md]:text-base",
		"data-[pui-avatar-size=lg]:size-16 data-[pui-avatar-size=lg]:text-xl",
		"rounded-full bg-muted",
		"data-[pui-avatar-in-group=true]:ring-2 data-[pui-avatar-in-group=true]:ring-background",
	)(p)

	args = append([]html.DivArg{html.AData("pui-avatar", "")}, args...)

	if p.Size != "" {
		args = append(args, html.AData("pui-avatar-size", string(p.Size)))
	}

	if p.InGroup {
		args = append(args, html.AData("pui-avatar-in-group", "true"))
	}

	for _, a := range args {
		a.ApplyDiv(attrs, children)
	}
}

// Avatar renders an avatar using the composable pattern.
// Accepts variadic html.DivArg arguments, with Props as an optional first argument.
func Avatar(args ...html.DivArg) html.Node {
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

	return html.Div(append([]html.DivArg{props}, rest...)...)
}

func imgArgsFromProps(baseClass string, extra ...string) func(p ImageProps) []html.ImgArg {
	return func(p ImageProps) []html.ImgArg {
		className := html.ClassMerge(
			append([]string{baseClass},
				append(extra, p.Class)...)...)

		args := []html.ImgArg{html.AClass(className)}

		if p.ID != "" {
			args = append(args, html.AId(p.ID))
		}

		for _, a := range p.Attrs {
			args = append(args, a)
		}

		return args
	}
}

// ApplyImg implements the html.ImgArg interface for ImageProps
func (p ImageProps) ApplyImg(attrs *html.ImgAttrs, children *[]html.Component) {
	args := imgArgsFromProps(
		"absolute inset-0 w-full h-full",
		"object-cover",
		"z-10",
	)(p)

	args = append([]html.ImgArg{html.AData("pui-avatar-image", "")}, args...)

	if p.Src != "" {
		args = append(args, html.ASrc(p.Src))
	}

	args = append(args, html.AAlt(p.Alt))

	for _, a := range args {
		a.ApplyImg(attrs, children)
	}
}

// Image renders an avatar image using the composable pattern.
// Accepts variadic html.ImgArg arguments, with ImageProps as an optional first argument.
func Image(args ...html.ImgArg) html.Node {
	var (
		props ImageProps
		rest  []html.ImgArg
	)

	// Separate Props from other arguments
	for _, a := range args {
		if v, ok := a.(ImageProps); ok {
			props = v
		} else {
			rest = append(rest, a)
		}
	}

	return html.Img(append([]html.ImgArg{props}, rest...)...)
}

func spanArgsFromProps(baseClass string, extra ...string) func(p FallbackProps) []html.SpanArg {
	return func(p FallbackProps) []html.SpanArg {
		className := html.ClassMerge(
			append([]string{baseClass},
				append(extra, p.Class)...)...)

		args := []html.SpanArg{html.AClass(className)}

		if p.ID != "" {
			args = append(args, html.AId(p.ID))
		}

		for _, a := range p.Attrs {
			args = append(args, a)
		}

		return args
	}
}

// ApplySpan implements the html.SpanArg interface for FallbackProps
func (p FallbackProps) ApplySpan(attrs *html.SpanAttrs, children *[]html.Component) {
	args := spanArgsFromProps(
		"font-medium text-muted-foreground",
	)(p)

	args = append([]html.SpanArg{html.AData("pui-avatar-fallback", "")}, args...)

	for _, a := range args {
		a.ApplySpan(attrs, children)
	}
}

// Fallback renders an avatar fallback using the composable pattern.
// Accepts variadic html.SpanArg arguments, with FallbackProps as an optional first argument.
func Fallback(args ...html.SpanArg) html.Node {
	var (
		props FallbackProps
		rest  []html.SpanArg
	)

	// Separate Props from other arguments
	for _, a := range args {
		if v, ok := a.(FallbackProps); ok {
			props = v
		} else {
			rest = append(rest, a)
		}
	}

	return html.Span(append([]html.SpanArg{props}, rest...)...)
}

func groupDivArgsFromProps(baseClass string, extra ...string) func(p GroupProps) []html.DivArg {
	return func(p GroupProps) []html.DivArg {
		className := html.ClassMerge(
			append([]string{baseClass},
				append(extra, p.Class)...)...)

		args := []html.DivArg{html.AClass(className)}

		if p.ID != "" {
			args = append(args, html.AId(p.ID))
		}

		for _, a := range p.Attrs {
			args = append(args, a)
		}

		return args
	}
}

// ApplyDiv implements the html.DivArg interface for GroupProps
func (p GroupProps) ApplyDiv(attrs *html.DivAttrs, children *[]html.Component) {
	spacing := p.Spacing
	if spacing == "" {
		spacing = GroupSpacingMd
	}

	args := groupDivArgsFromProps(
		"flex items-center",
		"data-[pui-avatar-spacing=sm]:-space-x-1",
		"data-[pui-avatar-spacing=md]:-space-x-2",
		"data-[pui-avatar-spacing=lg]:-space-x-4",
	)(p)

	args = append([]html.DivArg{html.AData("pui-avatar-spacing", string(spacing))}, args...)

	for _, a := range args {
		a.ApplyDiv(attrs, children)
	}
}

// Group renders an avatar group using the composable pattern.
// Accepts variadic html.DivArg arguments, with GroupProps as an optional first argument.
func Group(args ...html.DivArg) html.Node {
	var (
		props GroupProps
		rest  []html.DivArg
	)

	// Separate Props from other arguments
	for _, a := range args {
		if v, ok := a.(GroupProps); ok {
			props = v
		} else {
			rest = append(rest, a)
		}
	}

	return html.Div(append([]html.DivArg{props}, rest...)...)
}

// GroupOverflow renders an overflow indicator for avatar groups that have more avatars than displayed.
// This function maintains its existing signature for backward compatibility.
func GroupOverflow(count int, props Props, args ...html.DivArg) html.Node {
	className := html.ClassMerge(
		"inline-flex items-center justify-center",
		"w-12 h-12 text-base",
		"data-[pui-avatar-size=sm]:w-8 data-[pui-avatar-size=sm]:h-8 data-[pui-avatar-size=sm]:text-xs",
		"data-[pui-avatar-size=md]:w-12 data-[pui-avatar-size=md]:h-12 data-[pui-avatar-size=md]:text-base",
		"data-[pui-avatar-size=lg]:w-16 data-[pui-avatar-size=lg]:h-16 data-[pui-avatar-size=lg]:text-xl",
		"rounded-full bg-muted ring-2 ring-background",
		props.Class,
	)

	dArgs := []html.DivArg{
		html.AClass(className),
	}
	if props.ID != "" {
		dArgs = append(dArgs, html.AId(props.ID))
	}

	if props.Size != "" {
		dArgs = append(dArgs, html.AData("pui-avatar-size", string(props.Size)))
	}

	for _, attr := range props.Attrs {
		dArgs = append(dArgs, attr)
	}

	dArgs = append(dArgs, args...)

	return html.Div(append(dArgs, html.Span(
		html.AClass("text-xs font-medium"),
		html.Text("+"+fmt.Sprint(count)),
	))...)
}
