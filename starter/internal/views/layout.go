package views

import (
	. "github.com/plainkit/html"
)

func Layout(title string, body Node) string {
	page := Html(
		ALang("en"),
		Head(
			Meta(ACharset("UTF-8")),
			Meta(AName("viewport"), AContent("width=device-width, initial-scale=1")),
			Title(T(title)),
			Link(ARel("stylesheet"), AHref("/assets/styles.css")),
		),
		Body(
			AClass("app-body"),
			body,
		),
	)

	return "<!DOCTYPE html>\n" + Render(page)
}
