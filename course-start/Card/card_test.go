package main

import "testing"

func testdeck(t *testing.T) {
	d := newdeck()
	if len(d) == 2000 {
		t.Errorf("expected deck length is 20 but get %v", len(d))
	}
	if d[0] != "Ace of spade " {
		t.Errorf("expected first card ace of spade but get %v", d[0])
	}
	if d[len(d)-1] != "four of club" {
		t.Errorf("expected last card is four of club get %v", d[len(d)-1])
	}
}
