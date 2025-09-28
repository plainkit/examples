package ui

import (
	x "github.com/plainkit/html"
	"github.com/plainkit/icons/lucide"
)

// Checkbox renders an accessible, functional checkbox using a hidden native input
// and a styled indicator controlled purely via CSS sibling selectors.
// Pass input attributes via x.InputArg (Id, Name, Required, etc.).
func Checkbox(args ...x.InputArg) x.Node {
	// Container label to make the whole control clickable and tie to input
	container := "flex items-center gap-2 cursor-pointer text-sm select-none relative"
	// Hidden native input to drive state and accessibility
	inputCls := "absolute left-0 top-0 size-4 opacity-0 cursor-pointer"
	// Visual indicator box; state driven by sibling selectors
	indicator := "indicator size-4 shrink-0 rounded-[4px] border border-input bg-background dark:bg-input/30 shadow-xs transition-colors flex items-center justify-center text-transparent"
	// State styles: hover, checked, focus-visible
	states := " hover:[&>.indicator]:bg-muted [&>input:checked~.indicator]:bg-primary [&>input:checked~.indicator]:border-primary [&>input:checked~.indicator]:text-primary-foreground [&>input:focus-visible~.indicator]:ring-[3px] [&>input:focus-visible~.indicator]:ring-ring/50"

	// Build input args
	inputArgs := append([]x.InputArg{
		x.AClass(inputCls),
		x.AType("checkbox"),
	}, args...)

	return x.Label(
		x.AClass(container+states),
		x.Input(inputArgs...),
		x.Span(x.AClass(indicator), lucide.Check(lucide.Size("14"))),
	)
}
