go-otp
======

OTP ("One-Time Pad") Implementation in Go

This simple Go package implements an [OTP](http://en.wikipedia.org/wiki/One-time_pad)
container to ease the process of creating and utilizing single-use tokens.
One-time pads are useful in crytology to create
[perfect forward secrecy](http://en.wikipedia.org/wiki/Forward_secrecy#Perfect_Forward_Secrecy)
systems. While this implementation does not deter you from re-using keys, it
makes switching keys or "pages" out simple and easy, and lends itself well to
the concept in general.

Internals
=========

A likely use case in using OTP is generating enough "pages" of cryptographically
secure random bytes to permit encrypting/decrypting many messages. In `go-otp`,
you generate some random bytes, and then use them to create a `otp.Pad`. Within
this pad, an internal pointer keeps track of used pads. This pointer may be
passed in during creation of an `otp.Pad`, allowing the same pad data to be used
until it has been exhausted. This makes it easy to pre-share many pages of OTP
material in advance, and utilize the pages over time.

Usage
=====

```go
func NewPad(material []byte, pageSize int, startPage int) (*Pad, error)
type Pad
    func (p *Pad) TotalPages() int
    func (p *Pad) RemainingPages() int
    func (p *Pad) UsedPages() int
    func (p *Pad) PreviousPage() ([]byte, error)
    func (p *Pad) CurrentPage() []byte
    func (p *Pad) NextPage() ([]byte, error)
```

Example
=======

Following is a basic usage example which creates enough page material to
facilitate 2048 unique one-time use byte slices, each 16 bytes in size:

```go
package main

import (
	"crypto/rand"
	"fmt"
	"github.com/ryanuber/go-otp"
)

func main() {
	m := make([]byte, 16*2048)
	rand.Read(m)

	// Create a new pad with 16-byte pages. Set pointer to page 1.
	pad, err := otp.NewPad(m, 16, 1)
	if err != nil {
		fmt.Printf("%s", err)
		return
	}

	// Display the page data at page 1
	currentData := pad.CurrentPage()

	// Display the page data at page 2 (advances pointer)
	nextPage := pad.NextPage()
}
```
