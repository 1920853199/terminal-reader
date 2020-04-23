package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

var   (
	String string
	Number int
	Input string
)

func init() {
	clear()
}

var flag = ""

func main() {
	ReadConfig()
	getFiles()

	//BooklistData[curBookIndex] = "[" + BooklistData[curBookIndex] + "](fg-white,bg-green)"
	BooklistData[curBookIndex] = Red(BooklistData[curBookIndex])

	flag = "list"

	Menu()

	/*for _,v := range BooklistData{
		fmt.Printf("%s\n",v)
	}*/

	//openBook()

	/*fmt.Printf("%s\n",Book.BorderLabel)
	fmt.Printf("%s\n",Book.Text)*/

	f := bufio.NewReader(os.Stdin) //读取输入的内容
	for {
		fmt.Print("请输入操作命名[输入h:查看帮助]>")

		//定义一行输入的内容分隔符。
		Input,_ = f.ReadString('\n')

		//如果用户输入的是一个空行就让用户继续输入。
		if len(Input) == 1 {
			continue
		}

		//将Input
		fmt.Sscan(Input,&String,&Number)

		if String == "上一页" && flag == "detail" {
			PrevPage()
			Read()
		}

		if String == "下一页" && flag == "detail" {
			NextPage()
			Read()
		}


		if String == "q" {
			os.Exit(1)
		}

		if String == "返回" && flag == "detail" {
			Menu()
			flag = "list"

		}

		if String == "上一本" && flag == "list" {
			PrevBooks()
			Menu()
		}

		if String == "下一本"  && flag == "list" {
			NextBooks()
			Menu()
		}

		if String == "阅读" && flag == "list"{
			openBook()
			flag = "detail"
			Read()
		}

		if String == "h" {
			clear()
			fmt.Println("命令详情：")
			fmt.Println("q：退出")
			fmt.Println("h：查看帮助")
			fmt.Println("上一本：在书籍列表页输入`上一本`选择上一本书籍")
			fmt.Println("下一本：在书籍列表页输入`下一本`选择下一本书籍")
			fmt.Println("阅读：在书籍列表页输入`阅读`进行阅读当前选择的书籍")
			fmt.Println("返回：在书籍列表页输入`返回`返回到书籍列表")
			fmt.Println("上一页：在书籍阅读页输入`上一页`进行翻回上一页")
			fmt.Println("下一页：在书籍列表页输入`下一本`选择翻回下一页")
		}

	}
}

func clear()  {
	cmd := exec.Command("clear") //Linux example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func Menu()  {
	clear()
	for _,v := range BooklistData{
		fmt.Printf("%s\n",v)
	}
}

func Read()  {
	clear()
	fmt.Printf("%s\n",Book.BorderLabel)
	fmt.Printf("%s\n",Book.Text)
}
