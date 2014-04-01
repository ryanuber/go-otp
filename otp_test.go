package otp

import (
	"bytes"
	"crypto/rand"
	"testing"
)

func TestNewPad(t *testing.T) {
	m := make([]byte, 32)
	_, err := rand.Read(m)
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	// Make sure we can properly create a new pad
	if _, err := NewPad(m, 8, 1); err != nil {
		t.Fatalf("bad: %s", err)
	}

	// Make sure an error is thrown if the pad size and the page size are not
	// cleanly divisible
	if _, err := NewPad(m, 7, 1); err == nil {
		t.Fatalf("Expected page size error")
	}

	// Make sure an error is thrown if the startPage is out of bounds
	if _, err := NewPad(m, 8, 5); err == nil {
		t.Fatalf("Expected out of bounds error")
	}
}

func TestTotalPages(t *testing.T) {
	m := make([]byte, 32)
	_, err := rand.Read(m)
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	p, err := NewPad(m, 8, 1)
	if err != nil {
		t.Fatalf("bad: %s", err)
	}

	totalPages := p.TotalPages()
	if totalPages != 4 {
		t.Fatalf("Expected 4 pages, got %d", totalPages)
	}
}

func TestPages(t *testing.T) {
	m := make([]byte, 32)
	_, err := rand.Read(m)
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	p, err := NewPad(m, 8, 1)
	if err != nil {
		t.Fatalf("bad: %s", err)
	}

	// Error is thrown if we attempt to look at pages < 1
	if _, err := p.PreviousPage(); err == nil {
		t.Fatalf("Expected out of bounds error")
	}

	// Current page is properly returned
	if !bytes.Equal(p.CurrentPage(), m[0:8]) {
		t.Fatalf("First page does not match expected bytes")
	}

	// Advancing the page works properly
	nextPage, err := p.NextPage()
	if err != nil {
		t.Fatalf("bad: %s", err)
	}

	if !bytes.Equal(nextPage, m[8:16]) {
		t.Fatalf("Page does not match expected bytes")
	}

	// Page pointer is updated after page advance
	if !bytes.Equal(p.CurrentPage(), m[8:16]) {
		t.Fatalf("Page pointer not properly advanced")
	}

	// Previous page is properly returned
	previousPage, err := p.PreviousPage()
	if err != nil {
		t.Fatalf("bad: %s", err)
	}

	if !bytes.Equal(previousPage, m[0:8]) {
		t.Fatalf("Previous page did not match expected bytes")
	}

	// By advancing the page twice more, we should be at the last page, where
	// requests for yet another page should fail.
	p.NextPage()
	p.NextPage()
	if _, err := p.NextPage(); err == nil {
		t.Fatalf("Expected out of bounds error")
	}
}
