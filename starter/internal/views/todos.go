package views

import (
	"fmt"

	"starter/internal/domain"

	. "github.com/plainkit/html"
)

// TodosPage renders the todo list page.
func TodosPage(todos []*domain.Todo) string {
	return Layout("Todos",
		Div(AClass("space-y-10 max-w-3xl mx-auto px-6 py-12"),
			H2(AClass("text-4xl font-extrabold tracking-tight mb-4"),
				T("Todo List"),
			),

			// Create todo form
			Form(AMethod("POST"), AAction("/todos"),
				Div(AClass("flex gap-3"),
					Input(
						AType("text"),
						AName("title"),
						APlaceholder("What needs to be done?"),
						AClass("flex-1 px-5 py-3 text-lg rounded-lg border border-input bg-background placeholder:text-muted-foreground focus:border-primary focus:ring-2 focus:ring-primary/50 outline-none transition"),
						ARequired(),
					),
					Button(AType("submit"),
						AClass("px-8 py-3 bg-primary text-primary-foreground rounded-lg font-medium shadow hover:shadow-md hover:-translate-y-0.5 transition"),
						T("Add"),
					),
				),
			),

			// Todo list container
			Div(AClass("border-t pt-6 space-y-3"),
				todoListContent(todos),
			),

			// Back link
			Div(AClass("pt-6"),
				A(AHref("/"), AClass("inline-flex items-center gap-2 text-primary hover:underline text-lg font-medium"),
					T("← Back to Home"),
				),
			),
		),
	)
}

// todoListContent returns either the empty state or the list of todos.
func todoListContent(todos []*domain.Todo) DivArg {
	if len(todos) == 0 {
		return P(AClass("text-muted-foreground text-lg italic py-12 text-center"),
			T("No todos yet. Create one above!"),
		)
	}

	return Div(AClass("space-y-3"),
		F(mapTodos(todos, todoItem)...),
	)
}

// todoItem renders a single todo item with complete and delete actions.
func todoItem(todo *domain.Todo) Component {
	return Div(AClass("py-4 px-5 border border-input rounded-lg flex justify-between items-center hover:shadow transition"),
		Span(AClass(todoTitleClass(todo)),
			T(todo.Title),
		),
		Div(AClass("flex gap-2"),
			completeButton(todo),
			deleteButton(todo),
		),
	)
}

// completeButton renders the complete button if the todo is not completed.
func completeButton(todo *domain.Todo) DivArg {
	if todo.Completed {
		return F() // Empty fragment
	}

	return Form(AMethod("POST"), AAction(fmt.Sprintf("/todos/%s/complete", todo.ID)), AClass("inline"),
		Button(AType("submit"),
			AClass("px-4 py-2 text-sm bg-primary text-primary-foreground rounded-md font-medium hover:bg-primary/90 transition"),
			T("Complete"),
		),
	)
}

// deleteButton renders the delete button.
func deleteButton(todo *domain.Todo) DivArg {
	return Form(AMethod("POST"), AAction(fmt.Sprintf("/todos/%s/delete", todo.ID)), AClass("inline"),
		Button(AType("submit"),
			AClass("px-4 py-2 text-sm border border-input rounded-md font-medium hover:bg-destructive/10 hover:text-destructive transition"),
			T("Delete"),
		),
	)
}

// todoTitleClass returns the CSS class for the todo title.
func todoTitleClass(todo *domain.Todo) string {
	if todo.Completed {
		return "text-lg line-through text-muted-foreground"
	}

	return "text-lg font-medium"
}

// mapTodos maps todos to components using the provided function.
func mapTodos(todos []*domain.Todo, fn func(*domain.Todo) Component) []Component {
	result := make([]Component, len(todos))
	for i, todo := range todos {
		result[i] = fn(todo)
	}

	return result
}
