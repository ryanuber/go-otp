package main

import (
	"fmt"
	"github.com/ryanuber/columnize"
	"github.com/ryanuber/go-otp"
)

func main() {
	var out []string

	//                        |--||--||--||--||--||--|
	o, err := otp.New([]byte("lsjkdfjsdlfsyehdu2ue82kd"), 4, 2)
	if err != nil {
		fmt.Printf("failed: %s", err)
		return
	}

	totalPages := o.TotalPages()
	out = append(out, fmt.Sprintf("Total pages | %d", totalPages))

	remainingPages := o.RemainingPages()
	out = append(out, fmt.Sprintf("Remaining pages | %d", remainingPages))

	usedPages := o.UsedPages()
	out = append(out, fmt.Sprintf("Used pages | %d", usedPages))

	currentPage := o.Current()
	out = append(out, fmt.Sprintf("Current page | %s", string(currentPage)))

	nextPage, err := o.PeekNext()
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}
	out = append(out, fmt.Sprintf("Next page (peek) | %s", string(nextPage)))

	previousPage, err := o.Previous()
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}
	out = append(out, fmt.Sprintf("Previous page | %s", string(previousPage)))

	// -------
	// move to next page
	// -------
	out = append(out, "-------------- Moving to next page ------------------")

	currentPage, err = o.Next()
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}
	out = append(out, fmt.Sprintf("Current page | %s", currentPage))

	previousPage, err = o.Previous()
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}
	out = append(out, fmt.Sprintf("Previous page | %s", string(previousPage)))

	nextPage, err = o.PeekNext()
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}
	out = append(out, fmt.Sprintf("Next page (peek) | %s", string(nextPage)))

	output, _ := columnize.SimpleFormat(out)
	fmt.Println(output)
}
