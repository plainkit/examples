package views

import (
	. "github.com/plainkit/html"
	icons "github.com/plainkit/icons/lucide"
)

// Layout wraps content with head, tailwind, and collected component assets.
func Layout(title string, content Node) Component {
	assets := NewAssets()
	assets.Collect(content)
	return baseHTML(title, content, assets)
}

// LayoutWithAssets collects assets from content and additional components
// that may not be reachable by the collector (e.g., nested components).
func LayoutWithAssets(title string, content Node, extras ...Component) Component {
	assets := NewAssets()
	assets.Collect(content)
	if len(extras) > 0 {
		assets.Collect(extras...)
	}
	return baseHTML(title, content, assets)
}

// LayoutWithAssetsProvided renders using a pre-collected assets bundle.
// If assets is nil, it falls back to collecting from content.
func LayoutWithAssetsProvided(title string, content Node, assets *Assets) Component {
	if assets == nil {
		assets = NewAssets()
		assets.Collect(content)
	}
	return baseHTML(title, content, assets)
}

func baseHTML(title string, content Node, assets *Assets) Component {
	return Html(
		ALang("en"),
		Head(
			Meta(ACharset("utf-8")),
			Meta(AName("viewport"), AContent("width=device-width, initial-scale=1")),
			Meta(AName("description"), AContent("Plain - A modern, type-safe HTML component library for Go with beautiful interfaces and compile-time guarantees.")),
			ATitle(title),
			Link(ARel("preload"), AHref("/assets/styles.css"), AType("text/css")),
			Link(ARel("stylesheet"), AHref("/assets/styles.css")),
		),
		Body(
			AClass("bg-background text-foreground antialiased min-h-screen font-sans"),
			siteHeader(title),
			Main(AClass("container mx-auto p-6"), content),
		),
	)
}

func siteHeader(title string) Node {
	return Header(
		AClass("border-b border-border bg-background/95 backdrop-blur supports-[backdrop-filter]:bg-background/80 shadow-sm"),
		Div(
			AClass("container mx-auto px-6 h-16 flex items-center justify-between"),
			Div(
				AClass("font-bold text-lg tracking-tight inline-flex items-center gap-3"),
				Div(
					AClass("flex items-center justify-center w-8 h-8 bg-primary rounded-lg"),
					icons.Diamond(icons.Size("18"), AClass("text-primary-foreground")),
				),
				Span(AClass("text-foreground"), T("Plain")),
				Span(AClass("text-muted-foreground text-sm font-normal"), T("/ "+title)),
			),
			Nav(
				AClass("flex items-center gap-1 text-sm"),
				A(
					AHref("/"),
					AClass("inline-flex items-center gap-2 px-3 py-2 rounded-md hover:bg-muted transition-colors"),
					icons.House(icons.Size("16"), AClass("text-muted-foreground")),
					T("Home"),
				),
				A(
					AHref("/users"),
					AClass("inline-flex items-center gap-2 px-3 py-2 rounded-md hover:bg-muted transition-colors bg-muted text-foreground"),
					icons.Users(icons.Size("16")),
					T("Users"),
				),
			),
		),
	)
}
