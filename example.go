package main

import (
	"fmt"
	"github.com/ryanuber/columnize"
	"github.com/ryanuber/go-otp"
)

func main() {
	var out []string

	p, err := otp.NewPad([]byte("lsjkdfjsdlfsyehdu2ue82kd"), 4, 2)
	if err != nil {
		fmt.Printf("failed: %s", err)
		return
	}

	totalPages := p.TotalPages()
	out = append(out, fmt.Sprintf("Total pages | %d", totalPages))

	remainingPages := p.RemainingPages()
	out = append(out, fmt.Sprintf("Remaining pages | %d", remainingPages))

	usedPages := p.UsedPages()
	out = append(out, fmt.Sprintf("Used pages | %d", usedPages))

	currentPage := p.CurrentPage()
	out = append(out, fmt.Sprintf("Current page | %s", string(currentPage)))

	nextPage, err := p.PeekNextPage()
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}
	out = append(out, fmt.Sprintf("Next page (peek) | %s", string(nextPage)))

	previousPage, err := p.PreviousPage()
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}
	out = append(out, fmt.Sprintf("Previous page | %s", string(previousPage)))

	// -------
	// move to next page
	// -------
	out = append(out, "-------------- Moving to next page ------------------")

	currentPage, err = p.NextPage()
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}
	out = append(out, fmt.Sprintf("Current page | %s", currentPage))

	previousPage, err = p.PreviousPage()
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}
	out = append(out, fmt.Sprintf("Previous page | %s", string(previousPage)))

	nextPage, err = p.PeekNextPage()
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}
	out = append(out, fmt.Sprintf("Next page (peek) | %s", string(nextPage)))

	output, _ := columnize.SimpleFormat(out)
	fmt.Println(output)
}
