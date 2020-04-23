package main

import (
	"fmt"
	"strconv"
)

type Books struct {
	Text  			string
	BorderLabel		string
}

var Book = Books{}

var title string = ""

var wen []rune
var wenl int = 0
var context string = ""

var curCharStart int = 0
var curCharEnd int = 0
//var curCharPrev int = 0

func initBookScene(_title string, _context string) {
	curCharStart = 0
	curCharEnd = getBookC(bookmd5)
	fmt.Printf("%s\n",curCharEnd)
	title = _title
	//Book.BorderLabel = title
	context = _context
	wen = []rune(context)
	wenl = len(wen)

	NextPage()
	//updateBookContext()
}

func updateBookContext() {
	var b int = int(float32(curCharStart) / float32(wenl) * 100)

	Book.Text = string(wen[curCharStart:curCharEnd])
	//var b float32 = float32(curCharStart) / float32(wenl)
	bookc = curCharStart

	Book.BorderLabel = "【" + title + "】 [" + strconv.Itoa(b) + "%]"
	updateConfig()
}

func NextPage() {
	if curCharEnd < len(wen) {
		curCharStart = curCharEnd
		var linenum = 20 //可用行数
		var charnum = 60  //每行可用字符数
		var cchar int = charnum

		var i int = 0
		for i = 0; i < (len(wen) - curCharStart); i++ {
			if string(wen[curCharStart+i]) == "\n" {
				linenum--
				cchar = charnum
			} else if string(wen[curCharStart+i]) == "\t" {
				cchar -= 4
			} else {
				cchar -= 2
			}

			if cchar <= 0 {
				linenum--
				cchar += charnum
			}
			curCharEnd = i + curCharStart
			if linenum == 0 {
				break
			}
		}
		updateBookContext()
	}
}

func PrevPage() {
	if curCharStart > 0 {
		var linenum = 20 //可用行数
		var charnum = 60  //每行可用字符数
		var cchar int = charnum

		var i int = 0
		for i = 0; i <= curCharStart; i++ {
			if string(wen[curCharStart-i]) == "\n" {
				linenum--
				cchar = charnum
			} else if string(wen[curCharStart-i]) == "\t" {
				cchar -= 4
			} else {
				cchar -= 2
			}

			if cchar <= 0 {
				linenum--
				cchar += charnum
			}
			curCharEnd = curCharStart - i
			if linenum == 0 {
				break
			}
		}

		v := curCharEnd
		curCharEnd = curCharStart
		curCharStart = v

		updateBookContext()
	}
}

func NextBooks()  {
	if curBookIndex > BookListNum{
		return
	}
	BooklistData[curBookIndex] = "[" + strconv.Itoa(curBookIndex+1) +"] " + BookInfos[curBookIndex]
	curBookIndex ++
	BooklistData[curBookIndex] = Red(BooklistData[curBookIndex])

}

func PrevBooks()  {
	if curBookIndex <= 0 {
		return
	}
	BooklistData[curBookIndex] = "[" + strconv.Itoa(curBookIndex+1) +"] " + BookInfos[curBookIndex]
	curBookIndex --
	BooklistData[curBookIndex] = Red(BooklistData[curBookIndex])
}
