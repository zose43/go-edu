package xmlselect

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

func Select(args []string) error {
	dec := xml.NewDecoder(os.Stdin)
	var stack []string
	for {
		t, err := dec.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("can't parse xml %v\n", err)
		}

		switch item := t.(type) {
		case xml.StartElement:
			stack = append(stack, item.Name.Local)
		case xml.EndElement:
			stack = stack[:len(stack)-1]
		case xml.CharData:
			if containsAll(args, stack) {
				2
				fmt.Printf("")
			}
		}
	}
	return nil
}

func containsAll(args, stack []string) bool {
	if len(args) == 0 {
		return true
	}
	for len(args) <= len(stack) {
		if args[0] == stack[0] {
			args = args[1:]
		}
		stack = stack[1:]
	}
	return false
}
