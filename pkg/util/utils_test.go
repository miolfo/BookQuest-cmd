package util

import "testing"

func TestStripTitle(t *testing.T) {
	testStripping("The Dark Forest (Remembrance of Earth’s Past, #2)", false, true, "the dark forest", t)
	testStripping("The Dark Forest (Remembrance of Earth’s Past, #2)", true, true, "thedarkforest", t)
	testStripping("The Dark   Forest (Remembrance of Earth’s Past, #2) test", false, true, "the dark forest test", t)
	testStripping("The book: part to remove is here", false, false, "the book part to remove is here", t)
	testStripping("The book: part to remove is here", false, true, "the book", t)
	testStripping("The War of Vengence: The Great Betrayal (Time of Legends: War of Vengence)", false, true, "the war of vengence", t)
	testStripping("The War of Vengence: The Great Betrayal (Time of Legends: War of Vengence)", false, false, "the war of vengence the great betrayal", t)
}

func testStripping(title string, removeWhiteSpace bool, removePartsAfterColon bool, expected string, t *testing.T) {
	s := StripTitle(title, removeWhiteSpace, removePartsAfterColon)
	if s != expected {
		t.Fatalf("Expected %s, got %s", expected, s)
	}
}
