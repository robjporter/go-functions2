package inline

import "html/template"

var (
	Bold = func(content string) (template.HTML, error) {
		return template.HTML("<b>" + content + "</b>"), nil
	}
	Italic = func(content string) (template.HTML, error) {
		return template.HTML("<i>" + content + "</i>"), nil
	}
	Underline = func(content string) (template.HTML, error) {
		return template.HTML("<u>" + content + "</u>"), nil
	}
)
