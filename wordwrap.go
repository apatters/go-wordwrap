package wordwrap

import (
	"regexp"
	"strings"
)

var (
	blankLineRegExpr         = regexp.MustCompile(`(^\s+$)|(^$)`)
	leadingWhitespaceRegExpr = regexp.MustCompile(`(^\s+.*)|(^$)`)
)

// Wrap word-wraps text on the "limit" column boundary. Multiple
// whitespace is consolidated into a single space or newline. Leading
// blank lines and all trailing whitespace is stripped. Empty lines,
// or lines starting with whitespace are untouched, so you can include
// "codeblocks" and/or break text into paragraphs. Text is returned
// unchanged if limit <= 0.
//
// Example usage:
//
//  Wrap(10, "This string would be split onto several new lines")
//
func Wrap(limit int, input string) string {
	if limit < 1 {
		return indentText("", false, input)
	}

	return wrapText(limit, "", false, input)
}

// Indent indents each line in input with a string. A line is a set of
// runes followed by a newline. The first line is indented with
// indent. Subsequent lines either use the same indent or use a
// "hanging indent" (an indent using spaces with length equal to
// indent) depending on the setting of hanging. All leading and
// trailing whitespace is stripped.
func Indent(indent string, hanging bool, input string) string {
	return indentText(indent, hanging, input)
}

// IndentWithWrap wraps text on the limit column boundary but adds an
// indent to each line.
func IndentWithWrap(limit int, indent string, hanging bool, input string) string {
	if limit-len(indent) < 1 {
		return indentText(indent, hanging, input)
	}

	return wrapText(limit, indent, hanging, input)
}

func wrapText(limit int, indent string, hanging bool, input string) string {
	var output string
	var paragraph string

	hangingIndent := strings.Repeat(" ", len(indent))
	lines := strings.Split(strings.TrimRight(input, " \t\n"), "\n")
	textStarted := false
	for _, line := range lines {
		switch {
		case blankLineRegExpr.MatchString(line) && !textStarted:
			continue
		case leadingWhitespaceRegExpr.MatchString(line):
			textStarted = true
			if len(paragraph) > 0 {
				output += wrapParagraph(limit, indent, hanging, paragraph) + "\n"
				if hanging {
					indent = hangingIndent
				}
			}
			if !blankLineRegExpr.MatchString(line) {
				output += indentText(indent, hanging, line)
			}
			output += "\n"
			if hanging {
				indent = hangingIndent
			}
			paragraph = ""
		default:
			textStarted = true
			if len(paragraph) != 0 {
				paragraph += " "
			}
			paragraph += line
		}
	}
	if len(paragraph) > 0 {
		output += wrapParagraph(limit, indent, hanging, paragraph)
	}

	return output
}

func wrapParagraph(limit int, indent string, hanging bool, input string) string {
	limit -= len(indent)
	output := indent
	var hangingIndent string
	if hanging {
		hangingIndent = strings.Repeat(" ", len(indent))
	} else {
		hangingIndent = indent
	}
	// Split string into array of words
	words := strings.Fields(input)
	if len(words) <= 1 {
		return indent + input
	}

	remaining := limit
	for _, word := range words {
		if len(word)+1 > remaining {
			if len(output) > 0 {
				output += "\n" + hangingIndent
			}

			output += word
			remaining = limit - len(word)
		} else {
			if len(output) > len(indent) {
				output += " "
			}

			output += word
			remaining -= (len(word) + 1)
		}
	}

	return output
}

func indentText(indent string, hanging bool, input string) string {
	var output string

	if len(input) == 0 {
		return indent
	}

	var hangingIndent string
	if hanging {
		hangingIndent = strings.Repeat(" ", len(indent))
	} else {
		hangingIndent = indent
	}

	lines := strings.Split(strings.TrimRight(input, " \t\n"), "\n")
	textStarted := false
	for _, line := range lines {
		switch {
		case blankLineRegExpr.MatchString(line) && !textStarted:
			continue
		case !textStarted:
			output += indent + line
			textStarted = true
		default:
			output += hangingIndent + line
		}
		output += "\n"

	}
	output = strings.TrimRight(output, "\n")

	return output
}
