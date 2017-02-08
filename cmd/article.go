package cmd

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
	"github.com/terasum/lettor/structs"
	"github.com/terasum/lettor/utils"
	"github.com/terasum/lettor/core"
)

func init() {
	RootCmd.AddCommand(genArticleCmd)
}

var genArticleCmd = &cobra.Command{
	Use:   "gena",
	Short: "generate a html article for lettor",
	Long:  `genarate a html from specific the markdown file`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 3 {
			fmt.Println("Please specific the source md and target html")
			fmt.Println("Example:")
			fmt.Println("lettor article template.tmpl source.md target.html")
			os.Exit(1)
		}
		templatePath := utils.Abs(args[0])
		sourcePath := utils.Abs(args[1])
		targetPath := utils.Abs(args[2])
		if !utils.Exist(templatePath) {
			fmt.Println("Please specific the template path")
			os.Exit(1)
		}
		if !utils.Exist(sourcePath) {
			fmt.Println("Please specific the source md path")
			os.Exit(1)
		}

		genArticle(templatePath, sourcePath, targetPath)
	},
}

func genArticle(templatePath, sourcePath, targetPath string) error {
	output, err := core.GenerateMarkdown(sourcePath)
	if err != nil {
		return err
	}

	// add content into the template
	b := bytes.NewBuffer(make([]byte, 0))

	t := template.New("article_template")

	tmp, err := ioutil.ReadFile(templatePath)

	check(err)

	t, _ = t.Parse(string(tmp))
	a := structs.Article{
		Title:   "Article",
		Content: template.HTML(string(output)),
		//TOC:toc,
	}

	t.Execute(b, a)

	err = ioutil.WriteFile(targetPath, b.Bytes(), 0655)
	check(err)
	return nil
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
