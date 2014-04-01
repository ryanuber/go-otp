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

	currentPage := o.PeekCurrent()
	out = append(out, fmt.Sprintf("Current page | %s", string(currentPage)))

	nextPage, err := o.PeekNext()
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}
	out = append(out, fmt.Sprintf("Next page | %s", string(nextPage)))

	previousPage, err := o.PeekPrevious()
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}
	out = append(out, fmt.Sprintf("Previous page | %s", string(previousPage)))

	output, _ := columnize.SimpleFormat(out)
	fmt.Println(output)
}
