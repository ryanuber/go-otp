package otp

import (
	"fmt"
)

type Pad struct {
	pages       [][]byte
	currentPage int
}

// NewPad creates a new "one-time pad"
func NewPad(material []byte, pageSize int, startPage int) (*Pad, error) {
	if len(material)%pageSize != 0 {
		return nil, fmt.Errorf("pad size must be divisible by page size")
	}

	// Do the page-splitting work up front
	var pages [][]byte
	for i := 0; i < len(material); i += pageSize {
		pages = append(pages, material[i:i+pageSize])
	}

	if startPage < 1 || startPage > len(pages) {
		return nil, fmt.Errorf("page %d out of bounds", startPage)
	}

	// Create the new OTP pad
	p := Pad{
		pages:       pages,
		currentPage: startPage,
	}

	return &p, nil
}

// TotalPages returns the number of pages in the pad
func (p *Pad) TotalPages() int {
	return len(p.pages)
}

// RemainingPages returns the number of unused pages in the pad
func (p *Pad) RemainingPages() int {
	return len(p.pages) - p.currentPage
}

// CurrentPage returns the current position of the page pointer
func (p *Pad) CurrentPage() int {
	return p.currentPage
}

// getPage returns the payload of the current page
func (p *Pad) getPage() []byte {
	return p.pages[p.currentPage-1]
}

// SetPage will set the page pointer
func (p *Pad) SetPage(page int) error {
	if page < 1 || page > p.TotalPages() {
		return fmt.Errorf("page %d out of bounds", page)
	}
	p.currentPage = page
	return nil
}

// NextPage will advance the page pointer
func (p *Pad) NextPage() error {
	if p.RemainingPages() == 0 {
		return fmt.Errorf("pad depleted")
	}
	p.currentPage++
	return nil
}

// Encrypt will take a byte slice and use modular addition to encrypt the
// payload using the current page.
func (p *Pad) Encrypt(payload []byte) ([]byte, error) {
	var result []byte
	page := p.getPage()

	// Page must be at least as long as plain text
	if len(page) < len(payload) {
		return nil, fmt.Errorf("insufficient page size")
	}

	for i := 0; i < len(payload); i++ {
		bdec := uint8(payload[i])
		kdec := uint8(page[i])
		encoded := uint8(bdec+kdec) % 255
		result = append(result, byte(encoded))
	}
	return result, nil
}

// Decrypt will accept a byte slice and reverse the process taken by Encode to
// translate encrypted text back into raw bytes. It is required that the page
// pointer be set to the same position as it was during Encode().
func (p *Pad) Decrypt(payload []byte) ([]byte, error) {
	var result []byte
	page := p.getPage()

	// Page must be at least as long as plain text
	if len(page) < len(payload) {
		return nil, fmt.Errorf("insufficient page size")
	}

	for i := 0; i < len(payload); i++ {
		bdec := uint8(payload[i])
		kdec := uint8(page[i])
		decoded := uint8(bdec-kdec) % 255
		result = append(result, byte(decoded))
	}
	return result, nil
}
