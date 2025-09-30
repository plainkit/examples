package views

import (
	"github.com/plainkit/examples/google-auth/internal/domain"
	"github.com/plainkit/examples/google-auth/internal/ui/avatar"
	"github.com/plainkit/examples/google-auth/internal/ui/button"
	"github.com/plainkit/examples/google-auth/internal/ui/card"
	"github.com/plainkit/examples/google-auth/internal/ui/separator"
	. "github.com/plainkit/html"
)

// DashboardPage renders the authenticated user dashboard.
func DashboardPage(user *domain.User) string {
	var avatarContent Node
	if user.AvatarURL != "" {
		avatarContent = card.Content(
			Div(
				AClass("flex justify-center mb-6"),
				avatar.Avatar(
					avatar.Props{Size: avatar.SizeLg},
					avatar.Image(avatar.ImageProps{Src: user.AvatarURL, Alt: "Profile picture"}),
					avatar.Fallback(T(string([]rune(user.Name)[0:1]))),
				),
			),
		)
	}

	content := card.Card(
		card.Header(
			card.Title(T("Welcome back, "+user.Name)),
			card.Description(T("Here's your profile information from Google.")),
		),
		avatarContent,
		card.Content(
			Div(
				AClass("space-y-4"),
				Div(
					AClass("flex justify-between items-center py-2"),
					Span(AClass("text-sm font-medium text-muted-foreground"), T("Full Name")),
					Span(AClass("text-sm font-semibold"), T(user.Name)),
				),
				separator.Separator(),
				Div(
					AClass("flex justify-between items-center py-2"),
					Span(AClass("text-sm font-medium text-muted-foreground"), T("Email Address")),
					Span(AClass("text-sm font-semibold"), T(user.Email)),
				),
			),
		),
		card.Footer(
			button.Button(
				button.Props{
					Variant:   button.VariantOutline,
					Size:      button.SizeDefault,
					FullWidth: true,
					Href:      "/logout",
				},
				T("Sign Out"),
			),
		),
	)

	return Layout("Dashboard", content)
}
