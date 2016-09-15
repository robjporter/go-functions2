package functions

import (
	"html/template"
)

var (
	Bold = func(content string) (template.HTML, error) {
		return template.HTML("<b>" + content + "</b>"),nil
	}
	Italic = func(content string) (template.HTML, error) {
		return template.HTML("<i>" + content + "</i>"),nil
	}
	Underline = func(content string) (template.HTML, error) {
		return template.HTML("<u>" + content + "</u>"),nil
	}
)

func RemoveDuplicates(a []string) []string { 
	result := []string{} 
	seen := map[string]string{} 
	for _, val := range a { 
		if _, ok := seen[val]; !ok { 
			result = append(result, val) 
			seen[val] = val 
		} 
	} 
	return result 
} 
