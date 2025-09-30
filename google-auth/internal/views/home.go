package views

import (
	"github.com/plainkit/examples/google-auth/internal/icons"
	"github.com/plainkit/examples/google-auth/internal/ui/button"
	"github.com/plainkit/examples/google-auth/internal/ui/card"
	. "github.com/plainkit/html"
)

// HomePage renders the home page with Google login button.
func HomePage() string {
	content := card.Card(
		card.Header(
			card.Title(T("Welcome to PlainKit + Goth")),
			card.Description(T("Authenticate with Google to access your personalized dashboard.")),
		),
		card.Content(
			button.Button(
				button.Props{
					Variant:   button.VariantDefault,
					Size:      button.SizeLg,
					FullWidth: true,
					Href:      "/auth/google",
				},
				icons.Google(AClass("w-5 h-5")),
				T("Continue with Google"),
			),
		),
	)

	return Layout("Home", content)
}
