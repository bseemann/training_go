package main

import "testing"

func TestBigger(t *testing.T) {
	g := bigger(12)
	e := 24.0
	// Case even number greater than 10
	if g != e {
		t.Errorf("Exp: %v, Got: %v", e, g)
	}

	g = bigger(13)
	e = 26.0
	// Case odd number greater than 10
	if g != e {
		t.Errorf("Exp: %v, Got: %v", e, g)
	}

	g = bigger(8)
	e = 4.0
	// Case even number lesser than 10
	if g != e {
		t.Errorf("Exp: %v, Got: %v", e, g)
	}

	g = bigger(7)
	e = 3.5
	// Case odd number lesser than 10
	if g != e {
		t.Errorf("Exp: %v, Got: %v", e, g)
	}

	g = bigger(10)
	e = 10.0
	// Case number is 10
	if g != e {
		t.Errorf("Exp: %v, Got: %v", e, g)
	}

}
