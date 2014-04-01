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
	if len(pad)%pageSize != 0 {
		return nil, fmt.Errorf("pad size must be divisible by page size")
	}

	o := OTP{
		pad:      pad,
		pageSize: pageSize,
		offset:   offset,
	}

	return &o
}
