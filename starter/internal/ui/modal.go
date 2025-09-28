package ui

import x "github.com/plainkit/html"

const modalCSS = `
.modal:target {
    opacity: 1 !important;
    pointer-events: auto !important;
}

.modal:target .modal-content {
    transform: scale(1) translateY(0) !important;
    opacity: 1 !important;
}

.modal:target .modal-content a:focus,
.modal:target .modal-content button:focus {
    outline: 2px solid #3b82f6;
    outline-offset: 2px;
}`

const modalJS = `
(function() {
    let currentModal = null;

    // Track when modal opens
    function handleHashChange() {
        const hash = window.location.hash;
        if (hash && hash !== '#') {
            const modal = document.querySelector(hash);
            if (modal && modal.classList.contains('fixed') && modal.getAttribute('role') === 'dialog') {
                currentModal = modal;
            }
        } else {
            currentModal = null;
        }
    }

    // Handle ESC key
    function handleEscKey(e) {
        if (e.key === 'Escape' && currentModal) {
            window.location.hash = '#';
            currentModal = null;
        }
    }

    // Set up event listeners
    window.addEventListener('hashchange', handleHashChange);
    document.addEventListener('keydown', handleEscKey);

    // Initialize on page load
    handleHashChange();
})();`

// Modal creates a modal container with shadcn/ui styling and accessibility features.
func Modal(args ...x.DivArg) x.Node {
	modalClasses := "modal fixed inset-0 z-50 bg-black/50 opacity-0 pointer-events-none transition-opacity duration-300 flex items-center justify-center"

	modalArgs := append([]x.DivArg{
		x.AClass(modalClasses),
		x.AAria("modal", "true"),
		x.AAria("labelledby", "modal-title"),
	}, args...)

	// Add backdrop link for closing modal (click outside)
	modalArgs = append(modalArgs, x.A(
		x.AHref("#"),
		x.AClass("absolute inset-0 z-[-1]"),
		x.AAria("label", "Close dialog"),
	))

	// Add accessible close link that can be reached with keyboard navigation
	modalArgs = append(modalArgs, x.A(
		x.AHref("#"),
		x.AClass("modal-esc-link sr-only focus:not-sr-only focus:absolute focus:top-4 focus:right-4 focus:bg-background focus:border focus:px-2 focus:py-1 focus:rounded focus:text-sm focus:z-[1002]"),
		x.Text("Press Tab then Enter to close"),
		x.AAria("label", "Close modal with keyboard"),
	))

	return x.Div(modalArgs...).WithAssets(modalCSS, modalJS, "modal")
}

// ModalTrigger creates a trigger link for opening the modal. Pass x.AArg like x.AHref("#id"), x.Text/x.T, classes, etc.
func ModalTrigger(args ...x.AArg) x.Node {
	return x.A(args...)
}

// ModalContent creates the modal content container with shadcn/ui styling
func ModalContent(args ...x.DivArg) x.Node {

	// Modal content styling - with entrance animation and focus management
	contentClasses := "modal-content relative bg-background border shadow-lg p-6 w-full max-w-lg grid gap-4 rounded-lg transform scale-90 translate-y-[-20px] opacity-0 transition-all duration-200"

	contentArgs := append([]x.DivArg{x.AClass(contentClasses)}, args...)

	// Add close button (×) in top-right
	contentArgs = append(contentArgs, x.A(
		x.AHref("#"),
		x.AClass("absolute right-5 top-5 rounded-sm opacity-70 hover:opacity-100 transition-opacity text-2xl leading-none w-4 h-4 flex items-center justify-center"),
		x.AAria("label", "Close modal"),
		x.Text("×"),
	))

	return x.Div(contentArgs...)
}

// ModalHeader creates a modal header with shadcn/ui styling
func ModalHeader(args ...x.DivArg) x.Node {
	headerClasses := "flex flex-col gap-2 text-center sm:text-left"
	headerArgs := append([]x.DivArg{x.AClass(headerClasses)}, args...)
	return x.Div(headerArgs...)
}

// ModalTitle creates a modal title with shadcn/ui styling. Pass x.H2Arg (x.Text/x.T, x.Child, etc.)
func ModalTitle(args ...x.H2Arg) x.Node {
	titleClasses := "text-lg leading-none font-semibold"
	titleArgs := append([]x.H2Arg{x.AClass(titleClasses), x.AId("modal-title")}, args...)
	return x.H2(titleArgs...)
}

// ModalDescription creates a modal description with shadcn/ui styling
func ModalDescription(args ...x.PArg) x.Node {
	descClasses := "text-muted-foreground text-sm"
	descArgs := append([]x.PArg{x.AClass(descClasses)}, args...)
	return x.P(descArgs...)
}

// ModalFooter creates a modal footer with shadcn/ui styling
func ModalFooter(args ...x.DivArg) x.Node {
	footerClasses := "flex flex-col-reverse gap-2 sm:flex-row sm:justify-end"
	footerArgs := append([]x.DivArg{x.AClass(footerClasses)}, args...)
	return x.Div(footerArgs...)
}
