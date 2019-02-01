package wordwrap_test

import (
	"fmt"

	"github.com/apatters/go-wordwrap"
)


func ExampleWrap() {

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
	
	// Output:
	// Now is the
	// time for all
	// good men to
	// come to the aid
	// of their
	// country.
	//
	// Have you ever
	// watched a crab
	// on the shore
	// crawling
	// backward in
	// search of the
	// Atlantic Ocean,
	// and missing?
	// That's the way
	// the mind of man
	// operates -- H.
	// L. Mencken
	//
	// Four score and
	// seven years ago
	// our fathers
	// brought forth
	// on this
	// continent, a
	// new nation,
	// conceived in
	// Liberty, and
	// dedicated to
	// the proposition
	// that all men
	// are created
	// equal.
	//
	// Now we are
	// engaged in a
	// great civil
	// war, testing
	// whether that
	// nation, or any
	// nation so
	// conceived and
	// so dedicated,
	// can long
	// endure. We are
	// met on a great
	// battle-field of
	// that war. We
	// have come to
	// dedicate a
	// portion of that
	// field, as a
	// final resting
	// place for those
	// who here gave
	// their lives
	// that that
	// nation might
	// live. It is
	// altogether
	// fitting and
	// proper that we
	// should do this.
	//
	// The quick
	// brown fox
	// jumped over the
	// lazy dog.
	//
	//     This text is not wrapped at all even though it goes past 15 columns.
	//
	// Now is the
	// time for all
	// good men to
	// come to the aid
	// of their
	// country.
}
