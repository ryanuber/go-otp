package main

import (
	"../../go-otp"
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func b64(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

func main() {
	// Initialize 16k of random OTP material
	material := make([]byte, 16384)
	_, err := rand.Read(material)
	if err != nil {
		fmt.Printf("err: %s\n", err)
		return
	}

	// Create the pad using the material with 16 byte pad size
	pad, err := otp.NewPad(material, 16, 1)

	fmt.Printf("Total pages: %d\n", pad.TotalPages())
	fmt.Printf("Remaining pages: %d\n", pad.RemainingPages())
	fmt.Printf("Current page: %s\n", b64(pad.CurrentPage()))
	nextPage, _ := pad.PeekNextPage()
	fmt.Printf("Next page: %s\n", b64(nextPage))
	fmt.Printf("\nRotating to next page...\n\n")
	currentPage, _ := pad.NextPage()
	fmt.Printf("Total pages: %d\n", pad.TotalPages())
	fmt.Printf("Remaining pages: %d\n", pad.RemainingPages())
	previousPage, _ := pad.PreviousPage()
	fmt.Printf("Previous page: %s\n", b64(previousPage))
	fmt.Printf("Current page: %s\n", b64(currentPage))
	nextPage, _ = pad.PeekNextPage()
	fmt.Printf("Next page: %s\n", b64(nextPage))
}
