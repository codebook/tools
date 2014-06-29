package main

import (
	//"bufio"
	"fmt"
	"os"
	"strings"
)

func instead(fpath string) {
	
	jsold_str := `<script src="http://www.google.com/jsapi" type="text/javascript"></script>`
	jsnew_str := `<script src="./assets/js/jsapi.js" type="text/javascript"></script>`
	cssold_str := `<link rel="stylesheet"
href="http://fonts.googleapis.com/css?family=Roboto:regular,medium,thin,italic,mediumitalic,bold" title="roboto">`
	cssnew_str := `<link rel="stylesheet"
href="./assets/css.fonts.css" title="roboto">`
	
	

	file,e := os.OpenFile(fpath,os.O_RDWR|os.O_CREATE,0666)
	defer file.Close()
	if e != nil {
		fmt.Println("os.Open("+fpath+") error")
		return
	}

	//reader := bufio.NewReader(file)
	//writer := bufio.NewWriter(file)

	/*
	var content string
	k := 1
	cont, err := reader.ReadString('\n')
	for  err != nil {
		fmt.Println(k)
		k++
		content += cont
		cont, err = reader.ReadString('\n')
	}

	fmt.Println("1 content")
	fmt.Println("2 before content: \n" + content)
	fmt.Println("3 ")
	*/

	
	readbytes := make([]byte,1024*1014*10) // 10M
	//n, err := reader.Read(readbytes)
	n, err := file.Read(readbytes)
	if err != nil {
		fmt.Println("io.ReadFull() is error")
	}
	if n > 0 {
		fmt.Println("n =", n)
		//buf := bytes.NewBuffer(readbytes)
		content := string(readbytes[:n])
		fmt.Println("len(content) =", len(content))
		/*
		fmt.Println(len(content))
		fmt.Println("1 content")
		fmt.Println("2 before content: \n" + content)
		fmt.Println("3 ")
		fmt.Println(content)
		fmt.Println("5 ")
		
		//content := string(bytes[:1024])
		//content_aft := bytes[1024:]
		*/	
		//fmt.Println(string(readbytes[:1024]))
		content = strings.Replace(content, jsold_str, jsnew_str, -1)
		content = strings.Replace(content, cssold_str, cssnew_str, -1)
		
		//file.Truncate(int64(0))
		readbytes = []byte(content) 
		//fmt.Println(string(readbytes[:1024]))
		file.Seek(0,0)  //移动文件的开始位置
		m, write_err := file.Write(readbytes)
		fmt.Println("m =", m)
		fmt.Println("len(content) =", len(content))
		if write_err != nil {
			fmt.Println("file.Write() error", write_err.Error())
		}
		flush_err := file.Sync()
		if flush_err != nil {
			fmt.Println("file.Flush() error", flush_err.Error())
		}
		
	}
	

}

func main() {
	fpath := "path/to/doc/file.html"
	instead(fpath)

}