package xmlselect

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"strings"
)

type Element struct {
	Type     string
	Attrs    []xml.Attr
	Children []*Element
}

func Select(xmlDoc io.Reader, args map[string]map[string]string) (*Element, error) {
	dec := xml.NewDecoder(xmlDoc)
	xmlData := make(map[string][]xml.Attr)
	var stack []*Element
	var root *Element

	for {
		t, err := dec.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("can't parse xml %v\n", err)
		}

		switch item := t.(type) {
		case xml.StartElement:
			xmlData[item.Name.Local] = item.Attr
			el := Element{
				Type:  item.Name.Local,
				Attrs: item.Attr,
			}
			stack = append(stack, &el)
		case xml.EndElement:
			delete(xmlData, item.Name.Local)
			stack, root = updTree(stack)
		case xml.CharData:
			text := bytes.TrimSpace(item)
			if containsAll(xmlData, args) {
				fmt.Printf("%s: %s\n", strings.Join(elements(args), " "), text)
			}
		}
	}
	return root, nil
}

func updTree(stack []*Element) ([]*Element, *Element) {
	if len(stack) > 1 {
		parent := stack[len(stack)-2]
		parent.Children = append(parent.Children, stack[len(stack)-1])
		return stack[:len(stack)-1], nil
	}
	return nil, stack[0]
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

func elements(xmlData map[string]map[string]string) []string {
	nodes := make([]string, len(xmlData))
	for k := range xmlData {
		nodes = append(nodes, k)
	}
	return nodes
}
