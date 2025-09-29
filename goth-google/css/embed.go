package css

import _ "embed"

// TailwindCSS contains the compiled Tailwind stylesheet for the demo application.
//
//go:embed output.css
var TailwindCSS string
