package ui

import x "github.com/plainkit/html"

// Input renders a styled input. Pass standard input attributes via x.InputArg
func Input(args ...x.InputArg) x.Node {
	classes := "flex h-9 w-full rounded-md border border-muted-foreground/50 bg-background dark:bg-input px-3 py-1 text-base shadow-inner transition-colors file:border-0 file:bg-transparent file:text-sm file:font-medium file:text-foreground placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-blue-500 dark:focus-visible:ring-blue-400 disabled:cursor-not-allowed disabled:opacity-50 md:text-sm"
	inputArgs := []x.InputArg{x.AClass(classes)}
	inputArgs = append(inputArgs, args...)

	return x.Input(inputArgs...)
}
