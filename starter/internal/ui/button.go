package ui

import (
	x "github.com/plainkit/html"
)

// Base button classes (shadcn/ui parity)
func buttonBase() x.ButtonArg {
	return x.AClass("inline-flex items-center justify-center gap-2 whitespace-nowrap rounded-md text-sm font-medium transition-colors focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring disabled:pointer-events-none disabled:opacity-50 [&_svg]:pointer-events-none [&_svg]:size-4 [&_svg]:shrink-0")
}

// Internal wrappers embed x.Global so they are valid x.ButtonArg and keep their class string
type buttonVariantArg struct {
	x.Global
	cls string
}

type buttonSizeArg struct {
	x.Global
	cls string
}

// Variants

func ButtonDefault() x.ButtonArg {
	s := "bg-primary text-primary-foreground shadow hover:bg-primary/90"
	return buttonVariantArg{Global: x.AClass(s), cls: s}
}

func ButtonDestructive() x.ButtonArg {
	s := "bg-destructive text-destructive-foreground shadow-sm hover:bg-destructive/90"
	return buttonVariantArg{Global: x.AClass(s), cls: s}
}

func ButtonOutline() x.ButtonArg {
	s := "border border-input bg-background shadow-sm hover:bg-accent hover:text-accent-foreground"
	return buttonVariantArg{Global: x.AClass(s), cls: s}
}

func ButtonOutlineBlue() x.ButtonArg {
	s := "border-2 border-blue-500 text-blue-600 bg-background shadow-sm hover:bg-blue-50 dark:border-blue-400 dark:text-blue-400 dark:hover:bg-blue-950"
	return buttonVariantArg{Global: x.AClass(s), cls: s}
}

func ButtonOutlineYellow() x.ButtonArg {
	s := "border-2 border-yellow-500 text-yellow-600 bg-background shadow-sm hover:bg-yellow-50 dark:border-yellow-400 dark:text-yellow-400 dark:hover:bg-yellow-950"
	return buttonVariantArg{Global: x.AClass(s), cls: s}
}

func ButtonOutlineRed() x.ButtonArg {
	s := "border-2 border-red-500 text-red-600 bg-background shadow-sm hover:bg-red-50 dark:border-red-400 dark:text-red-400 dark:hover:bg-red-950"
	return buttonVariantArg{Global: x.AClass(s), cls: s}
}

func ButtonOutlineMuted() x.ButtonArg {
	s := "border-2 border-current text-current bg-background shadow-sm hover:bg-current/10"
	return buttonVariantArg{Global: x.AClass(s), cls: s}
}

func ButtonSecondary() x.ButtonArg {
	s := "bg-secondary text-secondary-foreground shadow-sm hover:bg-secondary/80"
	return buttonVariantArg{Global: x.AClass(s), cls: s}
}

func ButtonGhost() x.ButtonArg {
	s := "hover:bg-accent hover:text-accent-foreground"
	return buttonVariantArg{Global: x.AClass(s), cls: s}
}

func ButtonLink() x.ButtonArg {
	s := "text-primary underline-offset-4 hover:underline"
	return buttonVariantArg{Global: x.AClass(s), cls: s}
}

// Sizes (prefixed)

func ButtonDefaultSize() x.ButtonArg {
	s := "h-9 px-4 py-2"
	return buttonSizeArg{Global: x.AClass(s), cls: s}
}

func ButtonSm() x.ButtonArg {
	s := "h-8 rounded-md px-3 text-xs"
	return buttonSizeArg{Global: x.AClass(s), cls: s}
}

func ButtonLg() x.ButtonArg {
	s := "h-10 rounded-md px-8"
	return buttonSizeArg{Global: x.AClass(s), cls: s}
}

func ButtonIcon() x.ButtonArg {
	s := "h-6 w-6"
	return buttonSizeArg{Global: x.AClass(s), cls: s}
}

// Button creates a UI button with styling. Strictly accepts x.ButtonArg values.
// Adds base classes and applies default variant/size if not provided.
func Button(args ...x.ButtonArg) x.Component {
	buttonArgs := make([]x.ButtonArg, 0, len(args)+3)
	buttonArgs = append(buttonArgs, buttonBase())

	var hasVariant, hasSize bool
	for _, a := range args {
		switch a.(type) {
		case buttonVariantArg:
			hasVariant = true
		case buttonSizeArg:
			hasSize = true
		}
		buttonArgs = append(buttonArgs, a)
	}

	if !hasVariant {
		buttonArgs = append(buttonArgs, ButtonDefault())
	}
	if !hasSize {
		buttonArgs = append(buttonArgs, ButtonDefaultSize())
	}

	return x.Button(buttonArgs...)
}

// ButtonClass returns a single x.AClass with base + variant + size classes; useful for asChild-like usage
func ButtonClass(args ...x.ButtonArg) x.Global {
	base := "inline-flex items-center justify-center gap-2 whitespace-nowrap rounded-md text-sm font-medium transition-colors focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring disabled:pointer-events-none disabled:opacity-50 [&_svg]:pointer-events-none [&_svg]:size-4 [&_svg]:shrink-0"
	var variantCls, sizeCls string
	for _, a := range args {
		switch v := a.(type) {
		case buttonVariantArg:
			if variantCls == "" {
				variantCls = v.cls
			}
		case buttonSizeArg:
			if sizeCls == "" {
				sizeCls = v.cls
			}
		}
	}

	if variantCls == "" {
		variantCls = ButtonDefault().(buttonVariantArg).cls
	}

	if sizeCls == "" {
		sizeCls = ButtonDefaultSize().(buttonSizeArg).cls
	}

	full := base
	if variantCls != "" {
		full += " " + variantCls
	}

	if sizeCls != "" {
		full += " " + sizeCls
	}

	return x.AClass(full)
}
