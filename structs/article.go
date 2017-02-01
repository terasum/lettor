package structs

import "html/template"

type Article struct{
	//标题
	Title string
	//别名
	Alias string
	//作者
	Author string
	//发布日期
	Date string
	//发布时间
	Time string
	//分类
	Categories []string
	//标签
	Tags []string
	//内容
	Content template.HTML
	//摘要
	Abstract template.HTML
	//目录
	TOC TOC
}
