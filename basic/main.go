package main

import (
	"fmt"

	. "github.com/plainkit/html"
)

func main() {
	page := Html(
		ALang("en"),
		Head(
			Title(T("My Page")),
			Meta(ACharset("UTF-8")),
			Style(T(".intro { color: blue; }")),
		),
		Body(
			H1(T("Hello, World!")),
			P(T("Built with Plain"), AClass("intro")),
		),
	)

	fmt.Println("<!DOCTYPE html>")
	fmt.Println(Render(page))
}
