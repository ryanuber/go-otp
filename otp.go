package otp

import (
	"fmt"
)

type OTP struct {
	pad      []byte
	pageSize int
	offset   int
}

func New(pad []byte, pageSize int, offset int) (*OTP, error) {
	if offset < 1 {
		return nil, fmt.Errorf("minimum offset is 1")
	}

	if len(pad)%pageSize != 0 {
		return nil, fmt.Errorf("pad size must be divisible by page size")
	}

	o := OTP{
		pad:      pad,
		pageSize: pageSize,
		offset:   (offset * pageSize) - pageSize,
	}

	return &o, nil
}

func (o *OTP) TotalPages() int {
	return len(o.pad) / o.pageSize
}

func (o *OTP) RemainingPages() int {
	return (len(o.pad) - (o.offset + o.pageSize)) / o.pageSize
}

func (o *OTP) UsedPages() int {
	return o.TotalPages() - o.RemainingPages()
}

func (o *OTP) PeekPrevious() ([]byte, error) {
	if o.UsedPages() == 1 {
		return nil, fmt.Errorf("no previous pages")
	}
	start := (o.offset - o.pageSize)
	end := o.offset
	return o.pad[start:end], nil
}

func (o *OTP) PeekCurrent() []byte {
	start := o.offset
	end := start + o.pageSize
	return o.pad[start:end]
}

func (o *OTP) PeekNext() ([]byte, error) {
	if o.RemainingPages() == 0 {
		return nil, fmt.Errorf("pad depleted")
	}
	start := o.offset + o.pageSize
	end := start + o.pageSize
	return o.pad[start:end], nil
}
