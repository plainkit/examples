package ui

import x "github.com/plainkit/html"

func Label(args ...x.LabelArg) x.Node {
	classes := "flex items-center gap-2 text-sm leading-none font-medium select-none group-data-[disabled=true]:pointer-events-none group-data-[disabled=true]:opacity-50 peer-disabled:cursor-not-allowed peer-disabled:opacity-50 cursor-pointer"
	labelArgs := []x.LabelArg{x.AClass(classes)}
	labelArgs = append(labelArgs, args...)

	return x.Label(labelArgs...)
}
