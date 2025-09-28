package ui

import x "github.com/plainkit/html"

type RadioArg interface {
	applyUIRadio(*radioState)
}

type radioState struct {
	label    string
	baseArgs []x.InputArg
}

type RadioLabelWrapper struct{ text string }

func (w RadioLabelWrapper) applyUIRadio(s *radioState) {
	s.label = w.text
}

func RadioLabel(text string) RadioLabelWrapper {
	return RadioLabelWrapper{text}
}

type RadioArgAdapter struct{ arg x.InputArg }

func (a RadioArgAdapter) applyUIRadio(s *radioState) {
	s.baseArgs = append(s.baseArgs, a.arg)
}

func adaptRadioArg(arg interface{}) RadioArg {
	if uiArg, ok := arg.(RadioArg); ok {
		return uiArg
	}
	if coreArg, ok := arg.(x.InputArg); ok {
		return RadioArgAdapter{coreArg}
	}
	return nil
}

func Radio(args ...interface{}) x.Component {
	state := &radioState{}

	for _, arg := range args {
		if adapted := adaptRadioArg(arg); adapted != nil {
			adapted.applyUIRadio(state)
		}
	}

	// W3Schools approach with shadcn/ui styling (same as checkbox but circular)
	containerClasses := "flex items-center gap-2 cursor-pointer text-sm select-none relative"
	inputClasses := "absolute opacity-0 cursor-pointer h-0 w-0"
	checkmarkClasses := "size-4 shrink-0 rounded-full border border-input bg-background dark:bg-input/30 shadow-xs transition-colors flex items-center justify-center after:content-[''] after:absolute after:top-[6px] after:left-[4px] after:w-[8px] after:h-[8px] after:rounded-full after:bg-white after:opacity-0 after:transition-opacity"

	// Add hover and checked states using peer selectors
	containerWithStates := containerClasses + " hover:[&>.checkmark]:bg-muted [&>input:checked~.checkmark]:bg-primary [&>input:checked~.checkmark]:border-primary [&>input:checked~.checkmark:after]:opacity-100 [&>input:focus-visible~.checkmark]:ring-[3px] [&>input:focus-visible~.checkmark]:ring-ring/50"

	radioArgs := append([]x.InputArg{
		x.AClass(inputClasses),
		x.AType("radio"),
	}, state.baseArgs...)

	return x.Label(
		x.AClass(containerWithStates),
		x.Child(x.Input(radioArgs...)),
		x.Child(x.Span(x.AClass(checkmarkClasses+" checkmark"))),
		x.Text(state.label),
	)
}

type RadioGroupArg interface {
	applyUIRadioGroup(*radioGroupState)
}

type radioGroupState struct {
	baseArgs []x.DivArg
}

type RadioGroupArgAdapter struct{ arg x.DivArg }

func (a RadioGroupArgAdapter) applyUIRadioGroup(s *radioGroupState) {
	s.baseArgs = append(s.baseArgs, a.arg)
}

func adaptRadioGroupArg(arg interface{}) RadioGroupArg {
	if uiArg, ok := arg.(RadioGroupArg); ok {
		return uiArg
	}
	if coreArg, ok := arg.(x.DivArg); ok {
		return RadioGroupArgAdapter{coreArg}
	}
	return nil
}

func RadioGroup(args ...interface{}) x.Component {
	state := &radioGroupState{}

	for _, arg := range args {
		if adapted := adaptRadioGroupArg(arg); adapted != nil {
			adapted.applyUIRadioGroup(state)
		}
	}

	groupArgs := append([]x.DivArg{x.AClass("space-y-2")}, state.baseArgs...)

	return x.Div(groupArgs...)
}
