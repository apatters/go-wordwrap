package wordwrap_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/apatters/go-wordwrap"
)

func TestWrap(t *testing.T) {
	tests := []struct {
		desc     string
		limit    int
		input    string
		expected string
	}{
		{
			"Empty input should result in empty output.",
			4,
			"",
			"",
		},
		{
			"Should wrap text so words fit on lines with the given limit.",
			10,
			"Test text Test text Test text Test text",
			"Test text\nTest text\nTest text\nTest text",
		},
		{
			"Should remove additional whitespace.",
			10,
			"Test  text  Test  text  Test  text  Test  text",
			"Test text\nTest text\nTest text\nTest text",
		},
		{
			"Should trim text.",
			10,
			"Test  text  Test  text  Test  text  Test  text    ",
			"Test text\nTest text\nTest text\nTest text",
		},
		{
			"Should not wrap if only one word.",
			4,
			"Oneword",
			"Oneword",
		},
		{
			"When given slightly more realistic data should be broken properly.",
			40,
			"This is a bunch of text that should be split at somewhere near 40 characters.",
			"This is a bunch of text that should be\nsplit at somewhere near 40 characters.",
		},
		{
			"0 limit should result in no word wrap.",
			0,
			"This text should not be wrapped",
			"This text should not be wrapped",
		},
		{
			"A leading and trailing whitespace should be stripped.",
			30,
			"\n   \nThe leading/trailing whitespace will be stripped from this text.  \n\n   ",
			"The leading/trailing\nwhitespace will be stripped\nfrom this text.",
		},
		{
			"Newlines within paragraphs should be replaced by spaces.",
			20,
			`
This paragraph
has multiple
newlines.

Which is followed by
another paragraph
with multiple
lines.
`,
			"This paragraph has\nmultiple newlines.\n\nWhich is followed\nby another paragraph\nwith multiple lines.",
		},
		{
			"Indented lines should not be wrapped.",
			20,
			`
This paragraph has multiple newlines.

  * An longish indented line that should not be wrapped.
  * A longish indented line with a 'hanging'
    indent.

Which is followed by another paragraph with multiple lines.
`,
			"This paragraph has\nmultiple newlines.\n\n  * An longish indented line that should not be wrapped.\n  * A longish indented line with a 'hanging'\n    indent.\n\nWhich is followed\nby another paragraph\nwith multiple lines.",
		},
	}

	for _, test := range tests {
		wrappedText := wordwrap.Wrap(test.limit, test.input)
		t.Logf("Description: %s", test.desc)
		t.Logf("limit: %d", test.limit)
		t.Logf("output:   %q", wrappedText)
		t.Logf("expected: %q", test.expected)
		assert.EqualValues(t, test.expected, wrappedText)
	}
}

func TestIndent(t *testing.T) {
	tests := []struct {
		desc          string
		input         string
		leader        string
		hangingIndent bool
		expected      string
	}{
		{
			"Empty input.",
			"",
			"indent: ",
			true,
			"indent: ",
		},
		{
			"A single line.",
			"A single line",
			"indent: ",
			true,
			"indent: A single line",
		},
		{
			"A single line with trailing newline.",
			"A single line",
			"indent: ",
			true,
			"indent: A single line",
		},
		{
			"Simple indent with no hanging indent.",
			"Three lines\nof indented\ntext",
			"indent: ",
			false,
			"indent: Three lines\nindent: of indented\nindent: text",
		},
		{
			"Simple indent with hanging indent.",
			"Three lines\nof indented\ntext",
			"indent: ",
			true,
			"indent: Three lines\n        of indented\n        text",
		},
		{
			"A leading blank line should be stripped.",
			"\nThe leading newline will be stripped from this text.",
			"indent: ",
			false,
			"indent: The leading newline will be stripped from this text.",
		},
	}

	for _, test := range tests {
		indentedText := wordwrap.Indent(test.leader, test.hangingIndent, test.input)
		t.Logf("Description: %s", test.desc)
		t.Logf("leader: %q", test.leader)
		t.Logf("hangingIndent: %t", test.hangingIndent)
		t.Logf("output:   %q", indentedText)
		t.Logf("expected: %q", test.expected)
		assert.EqualValues(t, test.expected, indentedText)
	}
}

func TestIndentWithWrap(t *testing.T) {
	tests := []struct {
		desc          string
		limit         int
		input         string
		leader        string
		hangingIndent bool
		expected      string
	}{
		{
			"Empty input should produce empty output.",
			4,
			"",
			"",
			false,
			"",
		},
		{
			"Hanging indent should only add non-blank indent to first line.",
			30,
			"Test text that fills over first line",
			"First line:",
			true,
			"First line:Test text that\n           fills over first\n           line",
		},
		// Should allow prefixes to simply be spaces
		{
			"Indent can be just whitespace.",
			20,
			"Test text that fills over first line",
			"  ",
			false,
			"  Test text that\n  fills over first\n  line",
		},
		{
			"Indent should apply to all lines when not using hanging indent.",
			30,
			"Test text that fills over first line should have a hanging indent",
			"indent: ",
			false,
			"indent: Test text that fills\nindent: over first line should\nindent: have a hanging indent",
		},
		// Should not break when only one word.
		{
			"No wrapping if only one word.",
			15,
			"TestText",
			"FirstLine:",
			false,
			"FirstLine:TestText",
		},
		// This text should not be wrapped.
		{
			"Don't wrap if limit == 0.",
			0,
			"This text should not be wrapped, but should have a leader",
			"indent:",
			true,
			"indent:This text should not be wrapped, but should have a leader",
		},
	}

	for _, test := range tests {
		wrappedText := wordwrap.IndentWithWrap(
			test.limit,
			test.leader,
			test.hangingIndent,
			test.input)
		t.Logf("Description: %s", test.desc)
		t.Logf("limit: %d", test.limit)
		t.Logf("leader: %q", test.leader)
		t.Logf("hangingIndent: %t", test.hangingIndent)
		t.Logf("expected: %q", test.expected)
		t.Logf("wrapped:  %q", wrappedText)
		assert.EqualValues(t, test.expected, wrappedText)
	}
}
