go-wordwrap
===========

Wordwrap is a golang package used to wrap text on a column boundery,
e.g., lines are broken up along whitespace boundaries so that the
length of a line never goes past the column boundary.

Lines starting with whitespace are not touched even if their length
exceeds the column boundary.

Wordwrap can also be be used to indent one or more lines of text with
a given leader.

[![Build Status](https://travis-ci.org/apatters/go-wordwrap.svg)](https://travis-ci.org/apatters/go-wordwrap) [![GoDoc](https://godoc.org/github.com/apatters/go-wordwrap?status.svg)](https://godoc.org/github.com/apatters/go-wordwrap)


Features
-------

* Wraps text past a column width.
* Empty lines or lines with leading whitespace are untouched.
* Text can be indented with a leader.

Documentation
-------------

Documentation can be found at [GoDoc](https://godoc.org/github.com/apatters/go-wordwrap)


Installation
------------

Install wordwrap using the "go get" command:

```bash
$ go get github.com/apatters/go-wordwrap
```

The Go distribution is wordwrap's only dependency.


Examples
--------

### Wrap

``` golang
package main

import (
	"fmt"

	"github.com/apatters/go-wordwrap"
)


func main() {

	// Wrap a single line at column 15.
	fmt.Println(wordwrap.Wrap(15, "Now is the time for all good men to come to the aid of their country."))
	fmt.Println()

	// Wrap a paragraph at colume 15.
	paragraph := `
Have you ever watched a crab on the shore crawling backward in search of the
Atlantic Ocean, and missing? That's the way the mind of man operates -- H. L. Mencken`
	fmt.Println(wordwrap.Wrap(15, paragraph))
	fmt.Println()

	// Wrap multiple paragraphs at column 15. Paragraphs are
	// separated by one of more blank lines.
	multiParagraph := `
Four score and seven years ago our fathers brought forth on this
continent, a new nation, conceived in Liberty, and dedicated to the
proposition that all men are created equal.

Now we are engaged in a great civil war, testing whether that nation,
or any nation so conceived and so dedicated, can long endure. We are
met on a great battle-field of that war. We have come to dedicate a
portion of that field, as a final resting place for those who here
gave their lives that that nation might live. It is altogether fitting
and proper that we should do this.`
	fmt.Println(wordwrap.Wrap(15, multiParagraph))
	fmt.Println()

	// Multiparagraphs with indented text wrapped at column 15.
	indented := `
The quick brown fox jumped over the lazy dog.

    This text is not wrapped at all even though it goes past 15 columns.

Now is the time for all good men to come to the aid of their country.`
	fmt.Println(wordwrap.Wrap(15, indented))
}
```

The above example program outputs:

```
Now is the
time for all
good men to
come to the aid
of their
country.

Have you ever
watched a crab
on the shore
crawling
backward in
search of the
Atlantic Ocean,
and missing?
That's the way
the mind of man
operates -- H.
L. Mencken

Four score and
seven years ago
our fathers
brought forth
on this
continent, a
new nation,
conceived in
Liberty, and
dedicated to
the proposition
that all men
are created
equal.

Now we are
engaged in a
great civil
war, testing
whether that
nation, or any
nation so
conceived and
so dedicated,
can long
endure. We are
met on a great
battle-field of
that war. We
have come to
dedicate a
portion of that
field, as a
final resting
place for those
who here gave
their lives
that that
nation might
live. It is
altogether
fitting and
proper that we
should do this.

The quick
brown fox
jumped over the
lazy dog.

    This text is not wrapped at all even though it goes past 15 columns.

Now is the
time for all
good men to
come to the aid
of their
country.
```

### Indent

``` golang
package main

import (
	"fmt"

	"github.com/apatters/go-wordwrap"
)

func main() {

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

}
```

The above example program outputs:

```
leader: A single line, indented with the leader.

leader: The first line has a leader.
leader: The second line also has a leader.

leader: The first line has a leader.
        The second line has a hanging indent.
```

License
-------

The go-wordwrap package is available under the [MITLicense](https://mit-license.org/).


Thanks
------

Thanks to [Secure64](https://secure64.com/company/) for
contributing this code.




