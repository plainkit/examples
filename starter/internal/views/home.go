package views

import (
	. "github.com/plainkit/html"
	icons "github.com/plainkit/icons/lucide"
	"github.com/plainkit/starter/internal/ui"
)

func HomePage() Node {
	return Div(
		AClass("grid gap-8"),

		// Hero Section
		Div(
			AClass("text-center py-12"),
			Div(
				AClass("flex items-center justify-center w-16 h-16 bg-gradient-to-br from-primary to-chart-4 rounded-xl mx-auto mb-6"),
				icons.Diamond(icons.Size("32"), AClass("text-primary-foreground")),
			),
			H1(
				AClass("text-4xl font-bold text-foreground mb-4"),
				T("Welcome to Plain"),
			),
			P(
				AClass("text-xl text-muted-foreground max-w-2xl mx-auto mb-8"),
				T("A modern, type-safe HTML component library for Go. This demo showcases beautiful interfaces built with compile-time guarantees."),
			),
			Div(
				AClass("flex items-center justify-center gap-4"),
				A(
					AHref("/users"),
					ui.ButtonClass(),
					AClass("flex items-center gap-2"),
					icons.Users(icons.Size("16")),
					T("View Demo"),
				),
				A(
					AHref("https://github.com/plainkit/html"),
					ATarget("_blank"),
					ui.ButtonClass(ui.ButtonSecondary()),
					AClass("flex items-center gap-2"),
					icons.Github(icons.Size("16")),
					T("GitHub"),
				),
			),
		),

		// Stats Section
		Div(
			AClass("text-center mb-8"),
			H2(AClass("text-2xl font-bold mb-6"), T("Key Features")),
		),
		Div(
			AClass("grid grid-cols-1 md:grid-cols-4 gap-4"),
			ui.Card(
				AClass("p-6 text-center"),
				Div(
					AClass("flex items-center justify-center w-12 h-12 bg-primary/10 rounded-lg mx-auto mb-3"),
					icons.Code(icons.Size("24"), AClass("text-primary")),
				),
				H3(AClass("text-2xl font-bold"), T("100%")),
				P(AClass("text-muted-foreground text-sm"), T("Type Safe")),
			),
			ui.Card(
				AClass("p-6 text-center"),
				Div(
					AClass("flex items-center justify-center w-12 h-12 bg-chart-1/10 rounded-lg mx-auto mb-3"),
					icons.Zap(icons.Size("24"), AClass("text-chart-1")),
				),
				H3(AClass("text-2xl font-bold"), T("0ms")),
				P(AClass("text-muted-foreground text-sm"), T("Runtime Overhead")),
			),
			ui.Card(
				AClass("p-6 text-center"),
				Div(
					AClass("flex items-center justify-center w-12 h-12 bg-chart-3/10 rounded-lg mx-auto mb-3"),
					icons.Blocks(icons.Size("24"), AClass("text-chart-3")),
				),
				H3(AClass("text-2xl font-bold"), T("50+")),
				P(AClass("text-muted-foreground text-sm"), T("Components")),
			),
			ui.Card(
				AClass("p-6 text-center"),
				Div(
					AClass("flex items-center justify-center w-12 h-12 bg-chart-5/10 rounded-lg mx-auto mb-3"),
					icons.Heart(icons.Size("24"), AClass("text-chart-5")),
				),
				H3(AClass("text-2xl font-bold"), T("1000+")),
				P(AClass("text-muted-foreground text-sm"), T("Beautiful Icons")),
			),
		),

		// Features Grid
		Div(
			AClass("text-center mb-8"),
			H2(AClass("text-2xl font-bold mb-6"), T("Why Choose Plain")),
		),
		Div(
			AClass("grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6"),

			// Feature 1: Type Safety
			ui.Card(
				AClass("hover:shadow-lg transition-shadow"),
				ui.CardHeader(
					Div(
						AClass("flex items-center gap-4"),
						Div(
							AClass("flex items-center justify-center w-12 h-12 bg-primary/10 rounded-lg"),
							icons.Shield(icons.Size("24"), AClass("text-primary")),
						),
						ui.CardTitle(T("Type Safe")),
					),
				),
				ui.CardContent(
					ui.CardDescription(
						T("Compile-time validation ensures your HTML is always correct. Invalid combinations fail at build time."),
					),
				),
			),

			// Feature 2: Performance
			ui.Card(
				AClass("hover:shadow-lg transition-shadow"),
				ui.CardHeader(
					Div(
						AClass("flex items-center gap-4"),
						Div(
							AClass("flex items-center justify-center w-12 h-12 bg-chart-2/10 rounded-lg"),
							icons.Zap(icons.Size("24"), AClass("text-chart-2")),
						),
						ui.CardTitle(T("Lightning Fast")),
					),
				),
				ui.CardContent(
					ui.CardDescription(
						T("Zero runtime overhead. Pure function calls generate HTML strings at compile time."),
					),
				),
			),

			// Feature 3: Beautiful Design
			ui.Card(
				AClass("hover:shadow-lg transition-shadow"),
				ui.CardHeader(
					Div(
						AClass("flex items-center gap-4"),
						Div(
							AClass("flex items-center justify-center w-12 h-12 bg-chart-4/10 rounded-lg"),
							icons.Palette(icons.Size("24"), AClass("text-chart-4")),
						),
						ui.CardTitle(T("Beautiful Design")),
					),
				),
				ui.CardContent(
					ui.CardDescription(
						T("Modern UI components with shadcn/ui styling and 1000+ Lucide icons included."),
					),
				),
			),
		),

		// Code Example
		ui.Card(
			AClass("bg-gradient-to-br from-muted/50 to-accent/50"),
			ui.CardHeader(
				ui.CardTitle(
					Div(
						AClass("flex items-center gap-2"),
						icons.Code(icons.Size("24")),
						T("Clean, Readable Code"),
					),
				),
				ui.CardDescription(T("This entire page is built with type-safe Go functions. No templates needed!")),
			),
			ui.CardContent(
				Pre(
					AClass("bg-muted p-4 rounded-lg border text-sm overflow-auto font-mono"),
					T(`import (
	. "github.com/plainkit/html"
	icons "github.com/plainkit/icons/lucide"
	"github.com/plainkit/starter/internal/ui"
)

func HomePage() Node {
    return ui.Card(
        ui.CardHeader(
            ui.CardTitle(
                Div(
                    AClass("flex items-center gap-2"),
                    icons.Zap(icons.Size("20")),
                    T("Plain Demo"),
                ),
            ),
        ),
        ui.CardContent(
            P(T("Type-safe HTML in Go!")),
        ),
    )
}`),
				),
			),
		),
	)
}
