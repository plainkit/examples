// Package css embeds the compiled Tailwind CSS stylesheet.
package css

import _ "embed"

// TailwindCSS contains the compiled Tailwind stylesheet.
// The file is embedded at compile time using go:embed.
//
//go:embed output.css
var TailwindCSS string
