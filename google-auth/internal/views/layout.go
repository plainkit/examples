package views

import (
	"github.com/plainkit/fonts/inter"
	. "github.com/plainkit/html"
)

// Layout creates the base HTML structure for all pages.
func Layout(title string, body Node) string {
	page := Html(
		ALang("en"),
		Head(
			Title(T(title)),
			Meta(ACharset("UTF-8")),
			Meta(AName("viewport"), AContent("width=device-width, initial-scale=1")),
			Link(ARel("stylesheet"), AHref("/assets/styles.css")),
			inter.HeadComponents("/assets/fonts"),
		),
		Body(
			AClass("min-h-screen bg-background flex items-center justify-center p-4"),
			Div(
				AClass("w-full max-w-md"),
				body,
			),
		),
	)

	return "<!DOCTYPE html>\n" + Render(page)
}
