package xmlselect

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

func Select(args []string) error {
	xmlDoc, err := os.Open("xmlselect/random.xml")
	if err != nil {
		return fmt.Errorf("can't open file %s", err)
	}
	dec := xml.NewDecoder(xmlDoc)
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
				fmt.Printf("%s: %s\n", strings.Join(args, " "), item)
			}
		}
	}
	return nil
}

func containsAll(args, stack []string) bool {
	for len(args) <= len(stack) {
		if len(args) == 0 {
			return true
		}
		if args[0] == stack[0] {
			args = args[1:]
		}
		stack = stack[1:]
	}
	return false
}
