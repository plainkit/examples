package card

import (
	"github.com/plainkit/html"
)

type Props struct {
	ID    string
	Class string
	Attrs []html.Global
}

type HeaderProps Props

type TitleProps Props

type DescriptionProps Props

type ContentProps Props

type FooterProps Props

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
	args := divArgsFromProps("w-full rounded-lg border bg-card text-card-foreground shadow-xs")(p)

	for _, a := range args {
		a.ApplyDiv(attrs, children)
	}
}

// Card renders a card using the composable pattern.
// Accepts variadic html.DivArg arguments, with Props as an optional first argument.
func Card(args ...html.DivArg) html.Node {
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

func headerDivArgsFromProps(baseClass string, extra ...string) func(p HeaderProps) []html.DivArg {
	return func(p HeaderProps) []html.DivArg {
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

// ApplyDiv implements the html.DivArg interface for HeaderProps
func (p HeaderProps) ApplyDiv(attrs *html.DivAttrs, children *[]html.Component) {
	args := headerDivArgsFromProps("flex flex-col space-y-1.5 p-6 pb-0")(p)

	for _, a := range args {
		a.ApplyDiv(attrs, children)
	}
}

// Header renders a card header using the composable pattern.
// Accepts variadic html.DivArg arguments, with HeaderProps as an optional first argument.
func Header(args ...html.DivArg) html.Node {
	var (
		props HeaderProps
		rest  []html.DivArg
	)

	// Separate Props from other arguments
	for _, a := range args {
		if v, ok := a.(HeaderProps); ok {
			props = v
		} else {
			rest = append(rest, a)
		}
	}

	return html.Div(append([]html.DivArg{props}, rest...)...)
}

func h3ArgsFromProps(baseClass string, extra ...string) func(p TitleProps) []html.H3Arg {
	return func(p TitleProps) []html.H3Arg {
		args := []html.H3Arg{html.AClass(html.ClassMerge(append([]string{baseClass}, append(extra, p.Class)...)...))}
		if p.ID != "" {
			args = append(args, html.AId(p.ID))
		}

		for _, a := range p.Attrs {
			args = append(args, a)
		}

		return args
	}
}

// ApplyH3 implements the html.H3Arg interface for TitleProps
func (p TitleProps) ApplyH3(attrs *html.H3Attrs, children *[]html.Component) {
	args := h3ArgsFromProps("text-lg font-semibold leading-none tracking-tight")(p)

	for _, a := range args {
		a.ApplyH3(attrs, children)
	}
}

// Title renders a card title using the composable pattern.
// Accepts variadic html.H3Arg arguments, with TitleProps as an optional first argument.
func Title(args ...html.H3Arg) html.Node {
	var (
		props TitleProps
		rest  []html.H3Arg
	)

	// Separate Props from other arguments
	for _, a := range args {
		if v, ok := a.(TitleProps); ok {
			props = v
		} else {
			rest = append(rest, a)
		}
	}

	return html.H3(append([]html.H3Arg{props}, rest...)...)
}

func pArgsFromProps(baseClass string, extra ...string) func(p DescriptionProps) []html.PArg {
	return func(p DescriptionProps) []html.PArg {
		args := []html.PArg{html.AClass(html.ClassMerge(append([]string{baseClass}, append(extra, p.Class)...)...))}
		if p.ID != "" {
			args = append(args, html.AId(p.ID))
		}

		for _, a := range p.Attrs {
			args = append(args, a)
		}

		return args
	}
}

// ApplyP implements the html.PArg interface for DescriptionProps
func (p DescriptionProps) ApplyP(attrs *html.PAttrs, children *[]html.Component) {
	args := pArgsFromProps("text-sm text-muted-foreground")(p)

	for _, a := range args {
		a.ApplyP(attrs, children)
	}
}

// Description renders a card description using the composable pattern.
// Accepts variadic html.PArg arguments, with DescriptionProps as an optional first argument.
func Description(args ...html.PArg) html.Node {
	var (
		props DescriptionProps
		rest  []html.PArg
	)

	// Separate Props from other arguments
	for _, a := range args {
		if v, ok := a.(DescriptionProps); ok {
			props = v
		} else {
			rest = append(rest, a)
		}
	}

	return html.P(append([]html.PArg{props}, rest...)...)
}

func contentDivArgsFromProps(baseClass string, extra ...string) func(p ContentProps) []html.DivArg {
	return func(p ContentProps) []html.DivArg {
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

// ApplyDiv implements the html.DivArg interface for ContentProps
func (p ContentProps) ApplyDiv(attrs *html.DivAttrs, children *[]html.Component) {
	args := contentDivArgsFromProps("p-6")(p)

	for _, a := range args {
		a.ApplyDiv(attrs, children)
	}
}

// Content renders card content using the composable pattern.
// Accepts variadic html.DivArg arguments, with ContentProps as an optional first argument.
func Content(args ...html.DivArg) html.Node {
	var (
		props ContentProps
		rest  []html.DivArg
	)

	// Separate Props from other arguments
	for _, a := range args {
		if v, ok := a.(ContentProps); ok {
			props = v
		} else {
			rest = append(rest, a)
		}
	}

	return html.Div(append([]html.DivArg{props}, rest...)...)
}

func footerDivArgsFromProps(baseClass string, extra ...string) func(p FooterProps) []html.DivArg {
	return func(p FooterProps) []html.DivArg {
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

// ApplyDiv implements the html.DivArg interface for FooterProps
func (p FooterProps) ApplyDiv(attrs *html.DivAttrs, children *[]html.Component) {
	args := footerDivArgsFromProps("flex items-center p-6 pt-0")(p)

	for _, a := range args {
		a.ApplyDiv(attrs, children)
	}
}

// Footer renders a card footer using the composable pattern.
// Accepts variadic html.DivArg arguments, with FooterProps as an optional first argument.
func Footer(args ...html.DivArg) html.Node {
	var (
		props FooterProps
		rest  []html.DivArg
	)

	// Separate Props from other arguments
	for _, a := range args {
		if v, ok := a.(FooterProps); ok {
			props = v
		} else {
			rest = append(rest, a)
		}
	}

	return html.Div(append([]html.DivArg{props}, rest...)...)
}
