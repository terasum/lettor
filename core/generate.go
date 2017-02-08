package core

import (
	"fmt"
	"io/ioutil"

	"github.com/russross/blackfriday"
)

func GenerateMarkdown(sourcePath string) ([]byte, error) {
	input, err := ioutil.ReadFile(sourcePath)
	if err != nil {
		fmt.Printf("occur a err %v \n", err)
		return nil, err
	}
	output := blackfriday.MarkdownCommon(input)
	return output, nil
}

func fillTemplate() {

}
