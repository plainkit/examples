package views

import (
	"fmt"

	. "github.com/plainkit/html"

	"github.com/plainkit/bloxui/examples/starter/internal/domain"
)

func HomePage(todos []*domain.Todo) string {
	var items []Node

	for _, todo := range todos {
		var checkboxChildren []Component
		if todo.Completed {
			checkboxChildren = append(checkboxChildren, Span(AClass("checkbox__indicator")))
		}

		statusClass := "todo-card"
		if todo.Completed {
			statusClass += " todo-card--done"
		}

		items = append(items, Article(
			AClass(statusClass),
			Div(
				AClass("todo-card__main"),
				Form(
					AMethod("post"),
					AAction(fmt.Sprintf("/todos/%s/toggle", todo.ID)),
					Button(
						AClass("todo-card__toggle"),
						AType("submit"),
						Span(
							AClass("checkbox"),
							Fragment(checkboxChildren...),
						),
						Span(
							AClass("visually-hidden"),
							T(fmt.Sprintf("Toggle completion for %s", todo.Title)),
						),
					),
				),
				Div(
					AClass("todo-card__content"),
					H3(AClass("todo-card__title"), T(todo.Title)),
				),
			),
			Form(
				AClass("todo-card__actions"),
				AMethod("post"),
				AAction(fmt.Sprintf("/todos/%s/delete", todo.ID)),
				Button(
					AClass("ghost-button"),
					AType("submit"),
					T("Remove"),
				),
			),
		))
	}

	var listBody Node
	if len(todos) == 0 {
		listBody = Div(
			AClass("empty"),
			P(T("You have a clean slate. Add your first task.")),
		)
	} else {
		listBody = Div(
			AClass("todo-list"),
			Fragment(itemsToComponents(items)...),
		)
	}

	layout := Layout(
		"Tasks",
		Main(
			AClass("shell"),
			Header(
				AClass("shell__header"),
				Div(
					AClass("hero"),
					Span(AClass("hero__tag"), T("plainkit")),
					H1(AClass("hero__title"), T("Tasks")),
					P(AClass("hero__subtitle"), T("A sleek todo list focused on clarity.")),
				),
				Form(
					AClass("new-todo"),
					AMethod("post"),
					AAction("/todos"),
					Input(
						AClass("new-todo__input"),
						AType("text"),
						AName("title"),
						APlaceholder("Type your next win"),
						AAutocomplete("off"),
						ARequired(),
					),
					Button(
						AClass("primary-button"),
						AType("submit"),
						T("Add"),
					),
				),
			),
			Section(
				AClass("card"),
				listBody,
			),
		),
	)

	return layout
}

func itemsToComponents(items []Node) []Component {
	components := make([]Component, 0, len(items))
	for _, item := range items {
		components = append(components, item)
	}

	return components
}
