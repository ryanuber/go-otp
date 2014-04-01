package otp

import (
	"fmt"
)

type Pad struct {
	pages       [][]byte
	currentPage int
}

func NewPad(material []byte, pageSize int, startPage int) (*Pad, error) {
	if len(material)%pageSize != 0 {
		return nil, fmt.Errorf("pad size must be divisible by page size")
	}

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

func (p *Pad) TotalPages() int {
	return len(p.pages)
}

func (p *Pad) RemainingPages() int {
	return len(p.pages) - p.currentPage
}

func (p *Pad) UsedPages() int {
	return p.currentPage + 1
}

func (p *Pad) PreviousPage() ([]byte, error) {
	if p.currentPage == 0 {
		return nil, fmt.Errorf("no previous pages")
	}
	return p.pages[p.currentPage-1], nil
}

func (p *Pad) CurrentPage() []byte {
	return p.pages[p.currentPage]
}

func (p *Pad) NextPage() ([]byte, error) {
	if p.RemainingPages() == 0 {
		return nil, fmt.Errorf("pad depleted")
	}
	p.currentPage++
	return p.CurrentPage(), nil
}

func (p *Pad) PeekNextPage() ([]byte, error) {
	if p.RemainingPages() == 0 {
		return nil, fmt.Errorf("pad depleted")
	}
	return p.pages[p.currentPage+1], nil
}
