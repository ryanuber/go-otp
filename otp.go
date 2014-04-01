package otp

import (
	"fmt"
)

type Pad struct {
	pages       [][]byte
	currentPage int
}

func NewPad(material []byte, pageSize int, startPage int) (*OTP, error) {
	if startPage < 1 {
		return nil, fmt.Errorf("minimum start page is 1")
	}

	if len(material)%pageSize != 0 {
		return nil, fmt.Errorf("pad size must be divisible by page size")
	}

	var pages [][]byte
	for i := 0; i < len(material); i += pageSize {
		pages = append(pages, material[i:i+pageSize])
	}

	p := Pad{
		pages:       pages,
		currentPage: startPage - 1,
	}

	return &p, nil
}

func (p *Pad) TotalPages() int {
	return len(o.pages)
}

func (p *Pad) RemainingPages() int {
	return len(o.pages) - o.currentPage
}

func (p *Pad) UsedPages() int {
	return o.currentPage + 1
}

func (p *Pad) Previous() ([]byte, error) {
	if o.currentPage == 0 {
		return nil, fmt.Errorf("no previous pages")
	}
	return o.pages[o.currentPage-1], nil
}

func (p *Pad) Current() []byte {
	return o.pages[o.currentPage]
}

func (p *Pad) Next() ([]byte, error) {
	if o.RemainingPages() == 0 {
		return nil, fmt.Errorf("pad depleted")
	}
	o.currentPage++
	return o.Current(), nil
}

func (p *Pad) PeekNext() ([]byte, error) {
	if o.RemainingPages() == 0 {
		return nil, fmt.Errorf("pad depleted")
	}
	return o.pages[o.currentPage+1], nil
}
