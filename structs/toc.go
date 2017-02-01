package structs

type TOC struct {
	Items []TOCItem
}

type TOCItem struct {
	parentIdx int
	currentIdx int
	content string
	// current level h1-h4
	currentLevel int
}
//生成目录树需要广度优先遍历

//func ()

func (toc *TOC)SetH(i int ,s string){
}
