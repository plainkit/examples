package views

import (
	"fmt"
	"strings"

	"modern_todo_plain/internal/store"

	. "github.com/plainkit/html"
	icons "github.com/plainkit/icons/lucide"
)

type PageData struct {
	Todos  []store.Todo
	Filter store.Filter
	Stats  store.Stats
}

func TodoPage(data PageData) Component {
	return Layout("Modern Todo", appShell(data))
}

func appShell(data PageData) Node {
	return Div(
		AId("todo-app"),
		AClass("flex w-full min-h-screen bg-background"),
		todoSidebar(data),
		Div(
			AClass("flex-1 flex flex-col"),
			AppHeader(),
			Main(
				AClass("flex-1 bg-background p-6"),
				TodoListSection(data),
			),
		),
		AddTodoDialog(data.Filter),
		EditTodoDialog(data.Filter),
	)
}

func todoSidebar(data PageData) Node {
	filters := []struct {
		key   store.Filter
		label string
		icon  Node
		count int
	}{
		{store.FilterAll, "All Tasks", icons.List(icons.Size("18")), data.Stats.Total},
		{store.FilterActive, "Active", icons.Circle(icons.Size("18")), data.Stats.Active},
		{store.FilterCompleted, "Completed", icons.CircleCheck(icons.Size("18")), data.Stats.Completed},
	}

	items := make([]ChildOpt, 0, len(filters))
	for _, f := range filters {
		isActive := f.key == data.Filter
		buttonClass := "filter-button"
		if isActive {
			buttonClass += " is-active"
		}

		items = append(items,
			Child(
				Button(
					AType("button"),
					AClass(buttonClass),
					AAria("pressed", fmt.Sprintf("%t", isActive)),
					ACustom("hx-get", fmt.Sprintf("/?filter=%s&partial=app", f.key)),
					ACustom("hx-target", "#todo-app"),
					ACustom("hx-swap", "outerHTML"),
					ACustom("hx-push-url", fmt.Sprintf("/?filter=%s", f.key)),
					Div(
						AClass("flex items-center gap-3"),
						Span(AClass("flex h-9 w-9 items-center justify-center rounded-lg bg-sidebar-muted"), Child(f.icon)),
						Span(AClass("font-medium"), T(f.label)),
					),
					Span(
						AId(fmt.Sprintf("count-%s", f.key)),
						AClass("filter-count"),
						T(fmt.Sprintf("%d", f.count)),
					),
				),
			),
		)
	}

	percent := completionPercent(data.Stats)

	buttonArgs := make([]DivArg, len(items)+1)
	buttonArgs[0] = AClass("space-y-2")
	for i, item := range items {
		buttonArgs[i+1] = item
	}

	return Aside(
		AClass("hidden md:block w-72 border-r border-border bg-sidebar text-sidebar-foreground"),
		Child(
			Div(
				AClass("p-6 space-y-6"),
				Div(buttonArgs...),
				Div(
					AClass("space-y-2 border-t border-sidebar-muted pt-4"),
					Div(
						AClass("flex items-center justify-between text-sm"),
						Span(AClass("text-sidebar-foreground"), T("Completion")),
						Span(
							AId("completion-label"),
							AClass("text-sidebar-accent font-semibold"),
							T(fmt.Sprintf("%d%%", percent)),
						),
					),
					Div(
						AClass("h-2 w-full rounded-full bg-sidebar-muted"),
						Div(
							AId("progress-bar"),
							AClass("h-2 rounded-full bg-sidebar-accent transition-all"),
							AStyle(fmt.Sprintf("wdth: %d%%", percent)),
						),
					),
				),
			),
		),
	)
}

func TodoListSection(data PageData) Node {
	listChildren := []ChildOpt{
		Child(Input(AType("hidden"), AId("todo-current-filter"), AName("filter"), AValue(string(data.Filter)))),
	}

	if len(data.Todos) == 0 {
		listChildren = append(listChildren,
			Child(
				Div(
					AClass("flex flex-col items-center justify-center rounded-2xl border border-dashed border-border py-16 text-center"),
					Div(
						AClass("mb-4 flex h-24 w-24 items-center justify-center rounded-full bg-muted"),
						icons.ShieldCheck(icons.Size("48"), AClass("text-muted-foreground")),
					),
					H3(AClass("text-lg font-semibold"), T("No tasks found")),
					P(AClass("max-w-sm text-sm text-muted-foreground"), T("You're all caught up! Add a new task to get started.")),
				),
			),
		)
	} else {
		for _, todo := range data.Todos {
			listChildren = append(listChildren, Child(todoCard(todo)))
		}
	}

	listArgs := make([]DivArg, len(listChildren)+2)
	listArgs[0] = AId("todo-list")
	listArgs[1] = AClass("space-y-3")
	for i, child := range listChildren {
		listArgs[i+2] = child
	}

	return Section(
		Div(listArgs...),
	)
}

func todoCard(todo store.Todo) Node {
	cardClass := "group rounded-xl border border-border bg-card/80 p-5 shadow-sm transition hover:shadow-md"
	if todo.Completed {
		cardClass += " opacity-80"
	}

	metaContent := []ChildOpt{
		Child(
			Div(
				AClass("inline-flex items-center gap-1"),
				icons.Calendar(icons.Size("14"), AClass("text-muted-foreground")),
				T(todo.CreatedAt.Format("Jan 02")),
			),
		),
		Child(
			Span(
				AClass(priorityBadgeAClass(todo.Priority)),
				T(capitalize(string(todo.Priority))),
			),
		),
	}

	textContent := []ChildOpt{
		Child(H3(AClass(titleClasses(todo.Completed)), T(todo.Title))),
	}
	if todo.Description != "" {
		textContent = append(textContent,
			Child(P(AClass(descriptionClasses(todo.Completed)), T(todo.Description))),
		)
	}
	metaArgs := make([]DivArg, len(metaContent)+1)
	metaArgs[0] = AClass("flex items-center gap-3 text-xs text-muted-foreground")
	for i, child := range metaContent {
		metaArgs[i+1] = child
	}

	textContent = append(textContent,
		Child(Div(metaArgs...)),
	)

	textArgs := make([]DivArg, len(textContent)+1)
	textArgs[0] = AClass("space-y-2")
	for i, child := range textContent {
		textArgs[i+1] = child
	}

	return Article(
		AClass(cardClass),
		Div(
			AClass("flex items-start gap-4"),
			Button(
				AType("button"),
				AClass(toggleButtonClasses(todo.Completed)),
				ACustom("hx-post", fmt.Sprintf("/todos/toggle?id=%s", todo.ID)),
				ACustom("hx-target", "#todo-app"),
				ACustom("hx-swap", "outerHTML"),
				ACustom("hx-include", "#todo-current-filter"),
				icons.Check(icons.Size("16")),
			),
			Div(
				AClass("flex-1 space-y-3"),
				Div(
					AClass("flex items-start justify-between gap-3"),
					Div(textArgs...),
					Div(
						AClass("flex items-center gap-1"),
						Button(
							AType("button"),
							AClass("inline-flex h-8 w-8 items-center justify-center rounded-lg text-muted-foreground hover:bg-muted"),
							AData("dialog-target", "edit-dialog"),
							AData("edit-id", todo.ID),
							AData("edit-title", todo.Title),
							AData("edit-description", todo.Description),
							AData("edit-priority", string(todo.Priority)),
							icons.PencilLine(icons.Size("16")),
						),
						Button(
							AType("button"),
							AClass("inline-flex h-8 w-8 items-center justify-center rounded-lg text-muted-foreground hover:bg-destructive/10 hover:text-destructive"),
							ACustom("hx-post", fmt.Sprintf("/todos/delete?id=%s", todo.ID)),
							ACustom("hx-target", "#todo-app"),
							ACustom("hx-swap", "outerHTML"),
							ACustom("hx-confirm", "Delete this task?"),
							ACustom("hx-include", "#todo-current-filter"),
							icons.Trash2(icons.Size("16")),
						),
					),
				),
			),
		),
	)
}

func toggleButtonClasses(completed bool) string {
	base := "mt-1 flex h-5 w-5 items-center justify-center rounded border"
	if completed {
		return base + " border-primary bg-primary text-primary-foreground"
	}
	return base + " border-muted-foreground/40 text-muted-foreground"
}

func titleClasses(completed bool) string {
	base := "text-base font-semibold text-card-foreground"
	if completed {
		return base + " line-through text-muted-foreground"
	}
	return base
}

func descriptionClasses(completed bool) string {
	base := "text-sm text-muted-foreground"
	if completed {
		return base + " line-through"
	}
	return base
}

func priorityBadgeAClass(priority store.Priority) string {
	switch priority {
	case store.PriorityLow:
		return "priority-low inline-flex items-center rounded-full px-3 py-1 text-xs font-medium"
	case store.PriorityHigh:
		return "priority-high inline-flex items-center rounded-full px-3 py-1 text-xs font-medium"
	default:
		return "priority-medium inline-flex items-center rounded-full px-3 py-1 text-xs font-medium"
	}
}

func completionPercent(stats store.Stats) int {
	if stats.Total == 0 {
		return 0
	}
	return int((float64(stats.Completed)/float64(stats.Total))*100 + 0.5)
}

func capitalize(s string) string {
	if s == "" {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

func AddTodoDialog(filter store.Filter) Node {
	return Dialog(
		AId("add-dialog"),
		AClass("modal"),
		Child(
			Form(
				AId("add-form"),
				AClass("space-y-4 p-6"),
				ACustom("hx-post", "/todos/create"),
				ACustom("hx-target", "#todo-app"),
				ACustom("hx-swap", "outerHTML"),
				ACustom("hx-include", "#todo-current-filter"),
				ACustom("hx-on::afterRequest", "if(event.detail.successful){ todoDialogs.closeDialog('add-dialog'); this.reset(); }"),
				H2(AClass("text-xl font-semibold"), T("Add New Task")),
				Label(
					AClass("block space-y-2"),
					AFor("add-title"),
					Span(AClass("text-sm font-medium"), T("Title")),
					Input(
						AId("add-title"),
						AName("title"),
						ARequired(),
						APlaceholder("What needs to be done?"),
						AClass("w-full rounded-lg border bg-background px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-primary"),
					),
				),
				Label(
					AClass("block space-y-2"),
					AFor("add-description"),
					Span(AClass("text-sm font-medium"), T("Description")),
					Textarea(
						AId("add-description"),
						AName("description"),
						ARows("3"),
						APlaceholder("Add more details... (optional)"),
						AClass("w-full rounded-lg border bg-background px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-primary"),
					),
				),
				Label(
					AClass("block space-y-2"),
					AFor("add-priority"),
					Span(AClass("text-sm font-medium"), T("Priority")),
					Select(
						AId("add-priority"),
						ACustom("name", "priority"),
						Child(Option(ACustom("value", string(store.PriorityLow)), T("Low"))),
						Child(Option(ACustom("value", string(store.PriorityMedium)), ASelected(), T("Medium"))),
						Child(Option(ACustom("value", string(store.PriorityHigh)), T("High"))),
						AClass("w-full rounded-lg border bg-background px-3 py-2 text-sm"),
					),
				),
				Input(AType("hidden"), AName("filter"), AValue(string(filter))),
				Div(
					AClass("flex gap-2 pt-2"),
					Button(
						AType("button"),
						AClass("flex-1 rounded-lg border border-border px-4 py-2 text-sm font-medium hover:bg-muted"),
						AData("close-dialog", "add-dialog"),
						T("Cancel"),
					),
					Button(
						AType("submit"),
						AClass("flex-1 rounded-lg bg-primary px-4 py-2 text-sm font-medium text-primary-foreground hover:bg-primary/90"),
						T("Add Task"),
					),
				),
			),
		),
	)
}

func EditTodoDialog(filter store.Filter) Node {
	return Dialog(
		AId("edit-dialog"),
		AClass("modal"),
		Child(
			Form(
				AId("edit-form"),
				AClass("space-y-4 p-6"),
				ACustom("hx-post", "/todos/update"),
				ACustom("hx-target", "#todo-app"),
				ACustom("hx-swap", "outerHTML"),
				ACustom("hx-include", "#todo-current-filter"),
				ACustom("hx-on::afterRequest", "if(event.detail.successful){ todoDialogs.closeDialog('edit-dialog'); }"),
				Input(AType("hidden"), AId("edit-id"), AName("id")),
				Input(AType("hidden"), AId("edit-filter"), AName("filter"), AValue(string(filter))),
				H2(AClass("text-xl font-semibold"), T("Edit Task")),
				Label(
					AClass("block space-y-2"),
					AFor("edit-title"),
					Span(AClass("text-sm font-medium"), T("Title")),
					Input(
						AId("edit-title"),
						AName("title"),
						ARequired(),
						APlaceholder("What needs to be done?"),
						AClass("w-full rounded-lg border bg-background px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-primary"),
					),
				),
				Label(
					AClass("block space-y-2"),
					AFor("edit-description"),
					Span(AClass("text-sm font-medium"), T("Description")),
					Textarea(
						AId("edit-description"),
						AName("description"),
						ARows("3"),
						APlaceholder("Add more details... (optional)"),
						AClass("w-full rounded-lg border bg-background px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-primary"),
					),
				),
				Label(
					AClass("block space-y-2"),
					AFor("edit-priority"),
					Span(AClass("text-sm font-medium"), T("Priority")),
					Select(
						AId("edit-priority"),
						ACustom("name", "priority"),
						Child(Option(ACustom("value", string(store.PriorityLow)), T("Low"))),
						Child(Option(ACustom("value", string(store.PriorityMedium)), T("Medium"))),
						Child(Option(ACustom("value", string(store.PriorityHigh)), T("High"))),
						AClass("w-full rounded-lg border bg-background px-3 py-2 text-sm"),
					),
				),
				Div(
					AClass("flex gap-2 pt-2"),
					Button(
						AType("button"),
						AClass("flex-1 rounded-lg border border-border px-4 py-2 text-sm font-medium hover:bg-muted"),
						AData("close-dialog", "edit-dialog"),
						T("Cancel"),
					),
					Button(
						AType("submit"),
						AClass("flex-1 rounded-lg bg-primary px-4 py-2 text-sm font-medium text-primary-foreground hover:bg-primary/90"),
						T("Save Changes"),
					),
				),
			),
		),
	)
}

func RenderFullPage(data PageData) string {
	return Render(TodoPage(data))
}

func RenderAppShell(data PageData) string {
	return Render(appShell(data))
}
