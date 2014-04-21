package otp

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"testing"
)

func TestNewPad(t *testing.T) {
	m := make([]byte, 35)
	_, err := rand.Read(m)
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	// New pads are properly created
	if _, err := NewPad(m, 8, 1); err != nil {
		t.Fatalf("bad: %s", err)
	}

	// An error is thrown on zero-length pages
	if _, err := NewPad(m, 0, 1); err == nil {
		t.Fatalf("Expected zero-length page error")
	}

	// An error is thrown if no pages can be created
	if _, err := NewPad(m, 37, 1); err == nil {
		t.Fatalf("Expected page size error")
	}

	// The maximum number of pages possible are created
	p, err := NewPad(m, 8, 1)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	totalPages := p.TotalPages()
	if totalPages != 4 {
		t.Fatalf("Expected 4 total pages, got %d", totalPages)
	}

	// An error is thrown if the startPage is out of bounds
	if _, err := NewPad(m, 8, 5); err == nil {
		t.Fatalf("Expected out of bounds error")
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

	// Make sure the total number of pages is correct
	totalPages := p.TotalPages()
	if totalPages != 4 {
		t.Fatalf("Expected 4 pages, got %d", totalPages)
	}

	// Current page is properly returned
	page := p.CurrentPage()
	if page != 1 {
		t.Fatalf("Expected page pointer to be 1, got %d", page)
	}

	// Advancing the page works properly
	if err := p.NextPage(); err != nil {
		t.Fatalf("bad: %s", err)
	}

	// Page pointer is updated after page advance
	page = p.CurrentPage()
	if page != 2 {
		t.Fatalf("Expected page pointer to be 2, got %d", page)
	}

	// Explicitly setting the page works properly
	if err := p.SetPage(4); err != nil {
		t.Fatalf("err: %s", err)
	}

	// Advancing the page past the end of the pad should throw an error
	if err := p.NextPage(); err == nil {
		t.Fatalf("Expected pad depleted error")
	}
}

func TestEncryption(t *testing.T) {
	m := []byte("123456789abcdefghijklmno")
	p, err := NewPad(m, 4, 1)
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	encrypted1, err := p.Encrypt([]byte("test"))
	encoded1 := base64.StdEncoding.EncodeToString(encrypted1)
	if encoded1 != "pZemqA==" {
		t.Fatalf("bad: %s", encoded1)
	}

	if err := p.NextPage(); err != nil {
		t.Fatalf("err: %s", err)
	}

	encrypted2, err := p.Encrypt([]byte("test"))
	encoded2 := base64.StdEncoding.EncodeToString(encrypted2)
	if encoded2 != "qZuqrA==" {
		t.Fatalf("bad: %s", encoded2)
	}

	if err := p.SetPage(6); err != nil {
		t.Fatalf("err: %s", err)
	}

	encrypted3, err := p.Encrypt([]byte("test"))
	encoded3 := base64.StdEncoding.EncodeToString(encrypted3)
	if encoded3 != "4NLh4w==" {
		t.Fatalf("bad: %s", encoded3)
	}
}

func TestEncryption_Failure(t *testing.T) {
	m := []byte("12345")
	p, err := NewPad(m, 5, 1)
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	if _, err := p.Encrypt([]byte("123456")); err == nil {
		t.Fatalf("expected page size error")
	}
}

func TestDecryption(t *testing.T) {
	message := []byte("hello world")

	// Use at least one byte with ASCII value > 127 so that operations can be
	// verified to reverse correctly when modular addition folds values.
	m := []byte("ÿ0123456789")
	p, err := NewPad(m, 11, 1)
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	encrypted, err := p.Encrypt(message)
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	decrypted, err := p.Decrypt(encrypted)
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	if !bytes.Equal(message, decrypted) {
		t.Fatalf("bad: %v", decrypted)
	}
}

func TestDecryption_Failure(t *testing.T) {
	message := []byte("hello world")

	m := []byte("12345")
	p, err := NewPad(m, 5, 1)
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	if _, err := p.Decrypt(message); err == nil {
		t.Fatalf("expected page size error")
	}
}
