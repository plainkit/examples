package ui

import (
	x "github.com/plainkit/html"
)

// Card creates a UI card with shadcn styling. Strictly accepts x.DivArg.
func Card(args ...x.DivArg) x.Node {
	classes := "bg-card text-card-foreground flex flex-col gap-6 rounded-xl border p-6 shadow-sm"
	cardArgs := []x.DivArg{x.AClass(classes), x.AData("slot", "card")}
	cardArgs = append(cardArgs, args...)

	return x.Div(cardArgs...)
}

// CardHeader creates a UI card header with shadcn styling. Strictly accepts x.DivArg.
func CardHeader(args ...x.DivArg) x.Node {
	classes := "@container/card-header grid auto-rows-min grid-rows-[auto_auto] items-start gap-1.5 has-data-[slot=card-action]:grid-cols-[1fr_auto] [.border-b]:pb-6"
	headerArgs := []x.DivArg{x.AClass(classes), x.AData("slot", "card-header")}
	headerArgs = append(headerArgs, args...)

	return x.Div(headerArgs...)
}

// CardTitle creates a UI card title with shadcn styling. Strictly accepts x.DivArg.
// For text, pass x.Text/x.T; for children, pass x.Child/x.C.
func CardTitle(args ...x.DivArg) x.Node {
	classes := "leading-none font-semibold"
	titleArgs := []x.DivArg{x.AClass(classes), x.AData("slot", "card-title")}
	titleArgs = append(titleArgs, args...)

	return x.Div(titleArgs...)
}

// CardDescription creates a UI card description with shadcn styling. Strictly accepts x.DivArg.
// For text, pass x.Text/x.T; for children, pass x.Child/x.C.
func CardDescription(args ...x.DivArg) x.Node {
	classes := "text-muted-foreground text-sm"
	descArgs := []x.DivArg{x.AClass(classes), x.AData("slot", "card-description")}
	descArgs = append(descArgs, args...)

	return x.Div(descArgs...)
}

// CardAction creates a UI card action area with shadcn styling. Strictly accepts x.DivArg.
func CardAction(args ...x.DivArg) x.Node {
	classes := "col-start-2 row-span-2 row-start-1 self-start justify-self-end"
	actionArgs := []x.DivArg{x.AClass(classes), x.AData("slot", "card-action")}
	actionArgs = append(actionArgs, args...)

	return x.Div(actionArgs...)
}

// CardContent creates a UI card content with shadcn styling. Strictly accepts x.DivArg.
func CardContent(args ...x.DivArg) x.Node {
	classes := ""
	contentArgs := []x.DivArg{x.AClass(classes), x.AData("slot", "card-content")}
	contentArgs = append(contentArgs, args...)

	return x.Div(contentArgs...)
}

// CardFooter creates a UI card footer with shadcn styling. Strictly accepts x.DivArg.
func CardFooter(args ...x.DivArg) x.Node {
	classes := "flex items-center px-6 [.border-t]:pt-6"
	footerArgs := []x.DivArg{x.AClass(classes), x.AData("slot", "card-footer")}
	footerArgs = append(footerArgs, args...)

	return x.Div(footerArgs...)
}
