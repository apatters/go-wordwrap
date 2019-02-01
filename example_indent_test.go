package wordwrap_test

import (
	"fmt"

	"github.com/apatters/go-wordwrap"
)

func ExampleIndent() {

	// A single, indented line.
	fmt.Println(wordwrap.Indent("leader: ", false, "A single line, indented with the leader."))
	fmt.Println()

	// Indent multiple lines without a hanging indent.
	lines := `
The first line has a leader.
The second line also has a leader.`
	fmt.Println(wordwrap.Indent("leader: ", false, lines))
	fmt.Println()

	// Indent multiple lines without hanging indent.
	lines = `
The first line has a leader.
The second line has a hanging indent.`
	fmt.Println(wordwrap.Indent("leader: ", true, lines))
	fmt.Println()

	// Output:
	// leader: A single line, indented with the leader.
	//
	// leader: The first line has a leader.
	// leader: The second line also has a leader.
	//
	// leader: The first line has a leader.
	//         The second line has a hanging indent.
}
