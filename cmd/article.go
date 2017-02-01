package cmd

import (
	"io/ioutil"
	"fmt"
	"github.com/russross/blackfriday"
	"bytes"
	"html/template"
	"github.com/terasum/lettor/structs"
	"github.com/spf13/cobra"
	"os"
	"github.com/terasum/lettor/utils"
)
func init() {
	RootCmd.AddCommand(genArticleCmd)
}

var genArticleCmd = &cobra.Command{
	Use:   "gena",
	Short: "generate a html article for lettor",
	Long:  `genarate a html from specific the markdown file`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <3{
			fmt.Println("Please specific the source md and target html")
			fmt.Println("Example:")
			fmt.Println("lettor article template.tmpl source.md target.html")
			os.Exit(1)
		}
		templatePath := utils.Abs(args[0])
		sourcePath := utils.Abs(args[1])
		targetPath := utils.Abs(args[2])
		if !utils.Exist(templatePath){
			fmt.Println("Please specific the template path")
			os.Exit(1)
		}
		if !utils.Exist(sourcePath){
			fmt.Println("Please specific the source md path")
			os.Exit(1)
		}

		genArticle(templatePath,sourcePath,targetPath)
	},
}

func genArticle(templatePath,sourcePath,targetPath string) {
	input,err := ioutil.ReadFile(sourcePath)
	if err  != nil {
		fmt.Printf("occur a err %v \n",err)
	}
	output := blackfriday.MarkdownCommon(input)
	//parse the content and get the toc

	//toc := generateTOC(output)

	// add content into the template
	b := bytes.NewBuffer(make([]byte, 0))

	t := template.New("article_template")

	tmp,err  := ioutil.ReadFile(templatePath)

	check(err)


	t, _ = t.Parse(string(tmp))
	a := structs.Article{
		Title: "Article",
		Content:template.HTML(string(output)),
		//TOC:toc,
	}

	t.Execute(b, a)

	err = ioutil.WriteFile(targetPath,b.Bytes(),0655)
	check(err)

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}