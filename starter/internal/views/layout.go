// Package views handles HTML rendering using PlainKit HTML.
// Views are responsible for converting data into HTML.
package views

import (
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
		),
		Body(
			AClass("min-h-screen flex flex-col bg-background text-foreground antialiased"),
			Header(
				AClass("border-b bg-card/50 backdrop-blur supports-[backdrop-filter]:bg-card/70 sticky top-0 z-40"),
				Div(
					AClass("px-6 py-4"),
					H1(AClass("text-2xl font-bold tracking-tight"), T("PlainKit Starter")),
				),
			),
			Main(
				AClass("flex-1 max-w-4xl w-full mx-auto px-6 py-12"),
				body,
			),
			Footer(
				AClass("border-t mt-auto bg-card/50 backdrop-blur supports-[backdrop-filter]:bg-card/70"),
				Div(
					AClass("max-w-4xl mx-auto px-6 py-6 text-center text-sm text-muted-foreground"),
					T("Built with PlainKit"),
				),
			),
		),
	)

	return "<!DOCTYPE html>\n" + Render(page)
}
