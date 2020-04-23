package main

var BooklistData []string
//var BookInfos []os.FileInfo
var BookInfos []string

var BookListNum int = 0
var curBookName string = ""
var curBookIndex int = 0

var Scene int = MenuScene

const (
	MenuScene int = iota
	BookScene
	InfoScene
	BossScene
)

var bookmd5 string = ""
var bookc int = 0
