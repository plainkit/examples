package views

import (
	. "github.com/plainkit/html"
)

// HomePage renders the home page.
func HomePage() string {
	content := Div(
		AClass("space-y-12 max-w-3xl mx-auto px-6 py-12"),
		Div(
			H2(
				AClass("text-5xl font-extrabold tracking-tight bg-gradient-to-r from-primary to-purple-500 bg-clip-text text-transparent mb-6"),
				T("Welcome to PlainKit Starter"),
			),
			P(
				AClass("text-lg text-muted-foreground leading-relaxed"),
				T("This is a minimal starter template for building web applications with Go and PlainKit."),
			),
		),
		Div(
			AClass("space-y-5"),
			H3(AClass("text-2xl font-semibold text-foreground border-b border-border pb-2"), T("Features")),
			Ul(
				AClass("list-disc list-inside space-y-2 text-base text-muted-foreground ml-4"),
				Li(T("Clean architecture with separated concerns")),
				Li(T("In-memory data store (easy to swap for a database)")),
				Li(T("Type-safe HTML generation with PlainKit")),
				Li(T("Tailwind CSS for styling")),
				Li(T("Integration tests with Testify")),
			),
		),
		Div(
			AClass("space-y-5 pt-4"),
			H3(AClass("text-2xl font-semibold text-foreground border-b border-border pb-2"), T("Quick Links")),
			Div(
				AClass("flex gap-4"),
				A(
					AHref("/todos"),
					AClass("inline-block px-6 py-3 rounded-lg bg-primary text-primary-foreground font-semibold shadow hover:shadow-lg hover:-translate-y-0.5 transition"),
					T("View Todos"),
				),
				A(
					AHref("/health"),
					AClass("inline-block px-6 py-3 rounded-lg border border-border text-foreground font-medium hover:bg-muted transition"),
					T("Health Check"),
				),
			),
		),
	)

	return Layout("Home", content)
}
