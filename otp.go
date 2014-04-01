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
		return nil, fmt.Errorf("start page (%d) out of bounds", startPage)
	}

	p := Pad{
		pages:       pages,
		currentPage: startPage - 1,
	}

	return &p, nil
}

// TotalPages returns the number of pages in the pad
func (p *Pad) TotalPages() int {
	return len(p.pages)
}

// UnusedPages returns the number of unused pages in the pad
func (p *Pad) RemainingPages() int {
	return len(p.pages) - (p.currentPage + 1)
}

// UsedPages returns the number of pages that have been used
func (p *Pad) UsedPages() int {
	return p.currentPage + 1
}

// PreviousPage returns the payload of the last used page
func (p *Pad) PreviousPage() ([]byte, error) {
	if p.currentPage == 0 {
		return nil, fmt.Errorf("no previous pages")
	}
	return p.pages[p.currentPage-1], nil
}

// CurrentPage returns the payload of the current page
func (p *Pad) CurrentPage() []byte {
	return p.pages[p.currentPage]
}

// NextPage will advance the page pointer, and return the payload of the
// new current key.
func (p *Pad) NextPage() ([]byte, error) {
	if p.RemainingPages() == 0 {
		return nil, fmt.Errorf("pad depleted")
	}
	p.currentPage++
	return p.CurrentPage(), nil
}

// Encode will take a byte slice and use modular addition to encrypt the
// payload using the current page.
func (p *Pad) Encode(payload []byte) ([]byte, error) {
	var result []byte
	page := p.CurrentPage()

	// Page must be at least as long as plain text
	if len(page) < len(payload) {
		return nil, fmt.Errorf("insufficient page size")
	}

	for i := 0; i < len(payload); i++ {
		bdec := int64(payload[i])
		kdec := int64(page[i])
		encoded := uint64(bdec+kdec) % ((1 << 64) - 1)
		result = append(result, byte(encoded))
	}
	return result, nil
}

// Decode will accept a byte slice and reverse the process taken by Encode to
// translate encrypted text back into raw bytes. It is required that the page
// pointer be set to the same position as it was during Encode().
func (p *Pad) Decode(payload []byte) ([]byte, error) {
	var result []byte
	page := p.CurrentPage()

	// Page must be at least as long as plain text
	if len(page) < len(payload) {
		return nil, fmt.Errorf("insufficient page size")
	}

	for i := 0; i < len(payload); i++ {
		bdec := int64(payload[i])
		kdec := int64(page[i])
		decoded := uint64(bdec-kdec) % ((1 << 64) - 1)
		if decoded < 0 {
			decoded += 26
		}
		result = append(result, byte(decoded))
	}
	return result, nil
}
