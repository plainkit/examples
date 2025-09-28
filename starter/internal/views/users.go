package views

import (
	. "github.com/plainkit/html"
	icons "github.com/plainkit/icons/lucide"
	"github.com/plainkit/starter/internal/domain"
	"github.com/plainkit/starter/internal/ui"
)

// UsersPage renders the users list and includes the modal markup.
func UsersPage(users []domain.User) Node {
	// Enhanced table with better styling and action buttons
	list := Div(
		AClass("overflow-auto border border-border rounded-lg"),
		Table(
			AClass("w-full text-sm"),
			Thead(
				Tr(AClass("border-b border-border"),
					Th(AClass("text-left p-4 font-medium text-muted-foreground border-r border-border"), T("ID")),
					Th(AClass("text-left p-4 font-medium text-muted-foreground border-r border-border"), T("User")),
					Th(AClass("text-left p-4 font-medium text-muted-foreground border-r border-border"), T("Email")),
					Th(AClass("text-left p-4 font-medium text-muted-foreground border-r border-border"), T("Status")),
					Th(AClass("text-right p-4 font-medium text-muted-foreground"), T("Actions")),
				),
			),
			Tbody(rows(users)...),
		),
	)

	// Enhanced modal with better UX and icons
	modal := ui.Modal(
		AId("add-user"),
		ui.ModalContent(
			ui.ModalHeader(
				ui.ModalTitle(
					Div(
						AClass("flex items-center gap-2"),
						Div(
							AClass("flex items-center justify-center w-10 h-10 bg-primary/10 rounded-lg"),
							icons.UserPlus(icons.Size("20"), AClass("text-primary")),
						),
						T("Add New Team Member"),
					),
				),
				ui.ModalDescription(T("Fill in the details below to invite a new member to your team.")),
			),
			Form(
				AMethod("post"),
				AAction("/users"),
				AClass("grid gap-6"),
				Div(
					AClass("grid gap-2"),
					ui.Label(
						AFor("name"),
						T("Full Name"),
					),
					ui.Input(
						AId("name"),
						AName("name"),
						APlaceholder("Enter full name"),
						ARequired(),
					),
				),
				Div(
					AClass("grid gap-2 relative"),
					ui.Label(
						AFor("email"),
						T("Email Address"),
					),
					ui.Input(
						AId("email"),
						AName("email"),
						AType("email"),
						APlaceholder("name@company.com"),
						ARequired(),
					),
				),
				ui.ModalFooter(
					AClass("flex items-center gap-4"),
					A(
						AHref("#"),
						ui.ButtonClass(ui.ButtonSecondary()),
						T("Cancel"),
					),
					Button(
						AType("submit"),
						ui.ButtonClass(),
						icons.Plus(icons.Size("16")),
						T("Create User"),
					),
				),
			),
		),
	)

	// Overview tab content
	overviewContent := Div(
		AClass("grid gap-6"),
		// Stats cards
		Div(
			AClass("grid grid-cols-1 md:grid-cols-3 gap-4"),
			ui.Card(
				AClass("p-6"),
				Div(
					AClass("flex items-center gap-4"),
					Div(
						AClass("flex items-center justify-center w-12 h-12 bg-primary/10 rounded-lg"),
						icons.Users(icons.Size("24"), AClass("text-primary")),
					),
					Div(
						Div(AClass("text-2xl font-bold"), T(itoa(len(users)))),
						P(AClass("text-muted-foreground text-sm"), T("Total Users")),
					),
				),
			),
			ui.Card(
				AClass("p-6"),
				Div(
					AClass("flex items-center gap-4"),
					Div(
						AClass("flex items-center justify-center w-12 h-12 bg-chart-2/10 rounded-lg"),
						icons.Check(icons.Size("24"), AClass("text-chart-2")),
					),
					Div(
						Div(AClass("text-2xl font-bold"), T(itoa(len(users)))),
						P(AClass("text-muted-foreground text-sm"), T("Active Users")),
					),
				),
			),
			ui.Card(
				AClass("p-6"),
				Div(
					AClass("flex items-center gap-4"),
					Div(
						AClass("flex items-center justify-center w-12 h-12 bg-chart-4/10 rounded-lg"),
						icons.Star(icons.Size("24"), AClass("text-chart-4")),
					),
					Div(
						Div(AClass("text-2xl font-bold"), T("4.8")),
						P(AClass("text-muted-foreground text-sm"), T("Avg Rating")),
					),
				),
			),
		),
		// Activity overview
		ui.Card(
			ui.CardHeader(
				ui.CardTitle(
					Div(
						AClass("flex items-center gap-2"),
						icons.Clock(icons.Size("20")),
						T("Recent Activity"),
					),
				),
				ui.CardDescription(T("Latest user activity and system events.")),
			),
			ui.CardContent(
				Div(
					AClass("space-y-4"),
					Div(
						AClass("flex items-center gap-3 p-3 bg-muted/30 rounded-lg"),
						Div(
							AClass("flex items-center justify-center w-8 h-8 bg-chart-2/10 rounded-full"),
							icons.UserPlus(icons.Size("16"), AClass("text-chart-2")),
						),
						Div(
							P(AClass("text-sm font-medium"), T("New user registered")),
							P(AClass("text-xs text-muted-foreground"), T("2 minutes ago")),
						),
					),
					Div(
						AClass("flex items-center gap-3 p-3 bg-muted/30 rounded-lg"),
						Div(
							AClass("flex items-center justify-center w-8 h-8 bg-chart-1/10 rounded-full"),
							icons.Settings(icons.Size("16"), AClass("text-chart-1")),
						),
						Div(
							P(AClass("text-sm font-medium"), T("System configuration updated")),
							P(AClass("text-xs text-muted-foreground"), T("1 hour ago")),
						),
					),
					Div(
						AClass("flex items-center gap-3 p-3 bg-muted/30 rounded-lg"),
						Div(
							AClass("flex items-center justify-center w-8 h-8 bg-chart-4/10 rounded-full"),
							icons.Shield(icons.Size("16"), AClass("text-chart-4")),
						),
						Div(
							P(AClass("text-sm font-medium"), T("Security scan completed")),
							P(AClass("text-xs text-muted-foreground"), T("3 hours ago")),
						),
					),
				),
			),
		),
	)

	// Users tab content
	usersContent := ui.Card(
		ui.CardHeader(
			Div(
				AClass("flex items-center justify-between"),
				Div(
					ui.CardTitle(
						Div(
							AClass("flex items-center gap-2"),
							icons.Users(icons.Size("20")),
							T("Team Members"),
						),
					),
					ui.CardDescription(T("Manage your team members and their permissions.")),
				),
				ui.ModalTrigger(
					AHref("#add-user"),
					ui.ButtonClass(),
					icons.UserPlus(icons.Size("16")),
					T("Add User"),
				),
			),
		),
		ui.CardContent(list),
	)

	// Settings tab content
	settingsContent := Div(
		AClass("grid gap-6"),
		ui.Card(
			ui.CardHeader(
				ui.CardTitle(
					Div(
						AClass("flex items-center gap-2"),
						icons.Settings(icons.Size("20")),
						T("User Management Settings"),
					),
				),
				ui.CardDescription(T("Configure user permissions and system preferences.")),
			),
			ui.CardContent(
				Div(
					AClass("space-y-6"),
					Div(
						AClass("flex items-center justify-between p-4 border border-border rounded-lg"),
						Div(
							H3(AClass("text-sm font-medium"), T("Auto-approve new users")),
							P(AClass("text-xs text-muted-foreground"), T("Automatically approve user registrations")),
						),
						ui.Checkbox(AId("auto-approve"), AName("auto-approve")),
					),
					Div(
						AClass("flex items-center justify-between p-4 border border-border rounded-lg"),
						Div(
							H3(AClass("text-sm font-medium"), T("Email notifications")),
							P(AClass("text-xs text-muted-foreground"), T("Send notifications for user activities")),
						),
						ui.Checkbox(AId("email-notifications"), AName("email-notifications"), AChecked()),
					),
					Div(
						AClass("flex items-center justify-between p-4 border border-border rounded-lg"),
						Div(
							H3(AClass("text-sm font-medium"), T("Two-factor authentication")),
							P(AClass("text-xs text-muted-foreground"), T("Require 2FA for all users")),
						),
						ui.Checkbox(AId("two-factor"), AName("two-factor")),
					),
				),
			),
		),
		ui.Card(
			ui.CardHeader(
				ui.CardTitle(
					Div(
						AClass("flex items-center gap-2"),
						icons.Shield(icons.Size("20")),
						T("Security Settings"),
					),
				),
				ui.CardDescription(T("Configure security policies and access controls.")),
			),
			ui.CardContent(
				Div(
					AClass("space-y-4"),
					Div(
						AClass("grid gap-2"),
						ui.Label(AFor("session-timeout"), T("Session Timeout (minutes)")),
						ui.Input(
							AId("session-timeout"),
							AName("session-timeout"),
							AType("number"),
							AValue("30"),
							AClass("w-32"),
						),
					),
					Div(
						AClass("grid gap-2"),
						ui.Label(AFor("password-policy"), T("Password Policy")),
						Div(
							AClass("flex items-center gap-2"),
							ui.Checkbox(AId("require-uppercase"), AName("require-uppercase"), AChecked()),
							ui.Label(AFor("require-uppercase"), T("Require uppercase letters")),
						),
					),
				),
			),
		),
	)

	return Div(
		AClass("grid gap-6"),
		// Page header
		Div(
			AClass("flex items-center justify-between"),
			Div(
				H1(AClass("text-3xl font-bold"), T("User Management")),
				P(AClass("text-muted-foreground"), T("Manage team members, permissions, and settings.")),
			),
		),
		// Tabs container using new component structure
		ui.Tabs(
			AClass("w-full"),
			ui.TabsList(
				ui.TabsTrigger(
					AData("value", "overview"),
					AData("state", "active"),
					icons.TrendingUp(icons.Size("16")),
					T("Overview"),
				),
				ui.TabsTrigger(
					AData("value", "users"),
					icons.Users(icons.Size("16")),
					T("Users"),
				),
				ui.TabsTrigger(
					AData("value", "settings"),
					icons.Settings(icons.Size("16")),
					T("Settings"),
				),
			),
			ui.TabsContent(
				AData("value", "overview"),
				overviewContent,
			),
			ui.TabsContent(
				AData("value", "users"),
				usersContent,
			),
			ui.TabsContent(
				AData("value", "settings"),
				settingsContent,
			),
		),
		modal,
	)
}

func rows(users []domain.User) []TbodyArg {
	out := make([]TbodyArg, 0, len(users))
	for _, u := range users {
		tr := Tr(
			AClass("hover:bg-muted/50 transition-colors"),
			// ID Column
			Td(
				AClass("p-4 border-r border-border"),
				Span(AClass("font-mono text-muted-foreground text-xs"), T("#"+itoa(u.ID))),
			),
			// User Column with Avatar
			Td(
				AClass("p-4 border-r border-border"),
				Div(
					AClass("flex items-center gap-3"),
					Div(
						AClass("flex items-center justify-center w-8 h-8 bg-gradient-to-br from-primary/10 to-chart-4/10 rounded-full"),
						icons.User(icons.Size("16"), AClass("text-primary")),
					),
					Div(
						P(AClass("font-medium"), T(u.Name)),
						P(AClass("text-muted-foreground text-xs"), T("Team Member")),
					),
				),
			),
			// Email Column
			Td(
				AClass("p-4 border-r border-border"),
				Div(
					AClass("flex items-center gap-2"),
					icons.Mail(icons.Size("14"), AClass("text-muted-foreground")),
					T(u.Email),
				),
			),
			// Status Column
			Td(
				AClass("p-4 border-r border-border"),
				Div(
					AClass("inline-flex items-center gap-1 px-2 py-1 bg-chart-2/10 text-chart-2 text-xs rounded-full"),
					icons.Check(icons.Size("12")),
					T("Active"),
				),
			),
			// Actions Column
			Td(
				AClass("p-4"),
				Div(
					AClass("flex items-center justify-end gap-1"),
					Button(
						AType("button"),
						AClass("inline-flex items-center justify-center h-8 w-8 rounded-md hover:bg-muted transition-colors"),
						ATitle("View user"),
						icons.Eye(AClass("text-muted-foreground"), icons.Size("14")),
					),
					Button(
						AType("button"),
						AClass("inline-flex items-center justify-center h-8 w-8 rounded-md hover:bg-muted transition-colors"),
						ATitle("Edit user"),
						icons.Pen(icons.Size("14"), AClass("text-muted-foreground")),
					),
					Button(
						AType("button"),
						AClass("inline-flex items-center justify-center h-8 w-8 rounded-md hover:bg-destructive/10 hover:text-destructive transition-colors"),
						ATitle("Delete user"),
						icons.Trash2(icons.Size("14"), AClass("text-muted-foreground")),
					),
				),
			),
		)
		out = append(out, tr)
	}

	// Add empty state if no users
	if len(users) == 0 {
		emptyRow := Tr(
			Td(
				AClass("p-8 text-center"),
				AColspan("5"),
				Div(
					AClass("flex flex-col items-center gap-3 text-muted-foreground"),
					icons.Users(icons.Size("48"), AClass("opacity-50")),
					H2(AClass("text-lg font-medium"), T("No users found")),
					P(AClass("text-sm"), T("Get started by adding your first team member.")),
				),
			),
		)
		out = append(out, emptyRow)
	}

	return out
}

// helper: integer to string using blox internal pattern
func itoa(i int) string {
	// small helper to avoid importing strconv all over
	// mirrored from blox's internal itoa
	// but here we simply do a minimal import if needed
	// reimplement to avoid public dependency footprint
	// This is deliberately simple.
	// In real code, prefer strconv.Itoa.
	if i == 0 {
		return "0"
	}
	n := i
	b := [20]byte{}
	bp := len(b)
	neg := n < 0
	if neg {
		n = -n
	}
	for n > 0 {
		bp--
		b[bp] = byte('0' + n%10)
		n /= 10
	}
	if neg {
		bp--
		b[bp] = '-'
	}
	return string(b[bp:])
}
