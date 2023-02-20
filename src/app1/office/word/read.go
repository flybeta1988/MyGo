package main

import (
	"baliance.com/gooxml/document"
	"fmt"
	"log"
)

func main() {
	fmt.Println("read office word test...")

	doc, err := document.Open("./paper.docx")
	if err != nil {
		log.Fatalf("error opening document: %s", err)
	}

	for i, para := range doc.Paragraphs() {
		fmt.Println("-------第", i, "段---------")
		for j, run := range para.Runs() {
			fmt.Println("\t-------第", j, "格式片段---------")
			fmt.Println(run.Text())
		}
		fmt.Println()
	}

	for k, img := range doc.Images {
		fmt.Println("image:", k, img.Format(), img.Path(), img.Size())
	}
}
