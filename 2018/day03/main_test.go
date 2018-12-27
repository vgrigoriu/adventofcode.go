package main

import "testing"

func TestNewClaim(t *testing.T) {
	input := "#123 @ 3,2: 5x4"
	claim := new(input)
	if claim.id != 123 {
		t.Errorf("id is %d, expecting 123", claim.id)
	}
	if claim.left != 3 {
		t.Errorf("left is %d, expecting 3", claim.left)
	}
	if claim.top != 2 {
		t.Errorf("top is %d, expecting 2", claim.top)
	}
	if claim.w != 5 {
		t.Errorf("w is %d, expecting 5", claim.w)
	}
	if claim.h != 4 {
		t.Errorf("h is %d, expecting 4", claim.h)
	}
}
