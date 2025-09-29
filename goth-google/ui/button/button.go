package button

import (
	"github.com/plainkit/html"
)

type Variant string
type Size string
type Type string

type Props struct {
	ID        string
	Class     string
	Attrs     []html.Global
	Variant   Variant
	Size      Size
	FullWidth bool
	Href      string
	Target    string
	Disabled  bool
	Type      Type
	Form      string
}

const (
	VariantDefault     Variant = "default"
	VariantDestructive Variant = "destructive"
	VariantOutline     Variant = "outline"
	VariantSecondary   Variant = "secondary"
	VariantGhost       Variant = "ghost"
	VariantLink        Variant = "link"
)

const (
	SizeDefault Size = "default"
	SizeSm      Size = "sm"
	SizeLg      Size = "lg"
	SizeIcon    Size = "icon"
)

const (
	TypeButton Type = "button"
	TypeReset  Type = "reset"
	TypeSubmit Type = "submit"
)

func buttonArgsFromProps(baseClass string, extra ...string) func(p Props) []html.ButtonArg {
	return func(p Props) []html.ButtonArg {
		className := html.ClassMerge(
			append([]string{baseClass},
				append(extra,
					variantClass(p.Variant),
					sizeClass(p.Size),
					modifierClass(p.FullWidth),
					p.Class,
				)...)...)

		args := []html.ButtonArg{html.AClass(className)}

		if p.ID != "" {
			args = append(args, html.AId(p.ID))
		}

		if p.Type != "" {
			args = append(args, html.AType(string(p.Type)))
		} else {
			args = append(args, html.AType(string(TypeButton)))
		}

		if p.Form != "" {
			args = append(args, html.AForm(p.Form))
		}

		if p.Disabled {
			args = append(args, html.ADisabled())
		}

		for _, a := range p.Attrs {
			args = append(args, a)
		}

		return args
	}
}

func aArgsFromProps(baseClass string, extra ...string) func(p Props) []html.AArg {
	return func(p Props) []html.AArg {
		className := html.ClassMerge(
			append([]string{baseClass},
				append(extra,
					variantClass(p.Variant),
					sizeClass(p.Size),
					modifierClass(p.FullWidth),
					p.Class,
				)...)...)

		args := []html.AArg{
			html.AHref(p.Href),
			html.AClass(className),
		}

		if p.ID != "" {
			args = append(args, html.AId(p.ID))
		}

		if p.Target != "" {
			args = append(args, html.ATarget(p.Target))
		}

		for _, a := range p.Attrs {
			args = append(args, a)
		}

		return args
	}
}

// ApplyButton implements the html.ButtonArg interface for Props
func (p Props) ApplyButton(attrs *html.ButtonAttrs, children *[]html.Component) {
	args := buttonArgsFromProps(
		"inline-flex items-center justify-center gap-2 whitespace-nowrap rounded-md text-sm font-medium transition-all",
		"disabled:pointer-events-none disabled:opacity-50 [&_svg]:pointer-events-none [&_svg:not([class*='size-'])]:size-4 shrink-0 [&_svg]:shrink-0",
		"outline-none focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px]",
		"aria-invalid:ring-destructive/20 dark:aria-invalid:ring-destructive/40 aria-invalid:border-destructive",
		"cursor-pointer",
	)(p)

	for _, a := range args {
		a.ApplyButton(attrs, children)
	}
}

// ApplyA implements the html.AArg interface for Props
func (p Props) ApplyA(attrs *html.AAttrs, children *[]html.Component) {
	if p.Href == "" || p.Disabled {
		// If no href or disabled, don't apply anchor attributes
		return
	}

	args := aArgsFromProps(
		"inline-flex items-center justify-center gap-2 whitespace-nowrap rounded-md text-sm font-medium transition-all",
		"disabled:pointer-events-none disabled:opacity-50 [&_svg]:pointer-events-none [&_svg:not([class*='size-'])]:size-4 shrink-0 [&_svg]:shrink-0",
		"outline-none focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px]",
		"aria-invalid:ring-destructive/20 dark:aria-invalid:ring-destructive/40 aria-invalid:border-destructive",
		"cursor-pointer",
	)(p)

	for _, a := range args {
		a.ApplyA(attrs, children)
	}
}

// Button renders a button (or link if Href is provided) using the composable pattern.
// Accepts variadic html.ButtonArg arguments, with Props as an optional first argument.
func Button(args ...html.ButtonArg) html.Node {
	var (
		props Props
		rest  []html.ButtonArg
	)

	// Separate Props from other arguments

	for _, a := range args {
		if v, ok := a.(Props); ok {
			props = v
		} else {
			rest = append(rest, a)
		}
	}

	// Render as anchor if href is provided and not disabled
	if props.Href != "" && !props.Disabled {
		// Convert button args to anchor args for link rendering
		var anchorArgs []html.AArg

		anchorArgs = append(anchorArgs, props) // Props implements html.AArg

		for _, arg := range rest {
			if aArg, ok := arg.(html.AArg); ok {
				anchorArgs = append(anchorArgs, aArg)
			}
		}

		return html.A(anchorArgs...)
	}

	// Otherwise render as button
	return html.Button(append([]html.ButtonArg{props}, rest...)...)
}

func variantClass(v Variant) string {
	switch v {
	case VariantDestructive:
		return "bg-destructive text-destructive-foreground shadow-xs hover:bg-destructive/90 focus-visible:ring-destructive/20 dark:focus-visible:ring-destructive/40 dark:bg-destructive/60"
	case VariantOutline:
		return "border bg-background shadow-xs hover:bg-accent hover:text-accent-foreground dark:bg-input/30 dark:border-input dark:hover:bg-input/50"
	case VariantSecondary:
		return "bg-secondary text-secondary-foreground shadow-xs hover:bg-secondary/80"
	case VariantGhost:
		return "hover:bg-accent hover:text-accent-foreground dark:hover:bg-accent/50"
	case VariantLink:
		return "text-primary underline-offset-4 hover:underline"
	default:
		return "bg-primary text-primary-foreground shadow-xs hover:bg-primary/90"
	}
}

func sizeClass(s Size) string {
	switch s {
	case SizeSm:
		return "h-8 rounded-md gap-1.5 px-3 has-[>svg]:px-2.5"
	case SizeLg:
		return "h-10 rounded-md px-6 has-[>svg]:px-4"
	case SizeIcon:
		return "size-9"
	default:
		return "h-9 px-4 py-2 has-[>svg]:px-3"
	}
}

func modifierClass(fullWidth bool) string {
	if fullWidth {
		return "w-full"
	}

	return ""
}
