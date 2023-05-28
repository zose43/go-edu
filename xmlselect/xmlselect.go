package xmlselect

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"strings"
)

func Select(xmlDoc io.Reader, args map[string]map[string]string) error {

	dec := xml.NewDecoder(xmlDoc)
	xmlData := make(map[string][]xml.Attr)

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
			xmlData[item.Name.Local] = item.Attr
		case xml.EndElement:
			delete(xmlData, item.Name.Local)
		case xml.CharData:
			if containsAll(xmlData, args) {
				fmt.Printf("%s: %s\n", strings.Join(createStack(args), " "), bytes.TrimSpace(item))
			}
		}
	}
	return nil
}

func containsAll(xmlData map[string][]xml.Attr, args map[string]map[string]string) bool {
	var i int
	for s, attrs := range xmlData {
		sel, ok := args[s]
		if ok {
			i++
		}
		if !compareAttr(attrs, sel) {
			return false
		}
	}
	return len(args) == i
}

func compareAttr(find []xml.Attr, source map[string]string) bool {
	if source == nil {
		return true
	}
	for _, attr := range find {
		if sel, ok := source[attr.Name.Local]; ok && sel == attr.Value {
			return true
		}
	}
	return false
}

func createStack(xmlData map[string]map[string]string) []string {
	stack := make([]string, len(xmlData))
	for k := range xmlData {
		stack = append(stack, k)
	}
	return stack
}
