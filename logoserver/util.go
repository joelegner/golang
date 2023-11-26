package main

import "fmt"

func html(head string, content string) string {
	return tag("html", (head + content))
}

func head(content string) string {
	return tag("head", content)
}

func body(content string) string {
	return tag("body", content)
}

func tag(tag string, content string) string {
	return fmt.Sprintf("<%s>%s</%s>", tag, content, tag)
}