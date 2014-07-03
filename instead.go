package main

import (
	"fmt"
	"os"
	"strings"
	"path/filepath"
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
		fmt.Println("os.Open("+fpath+") error:",e.Error())
		return
	}
	
	readbytes := make([]byte,1024*1014*10) // 10M
	n, err := file.Read(readbytes)
	if err != nil {
		fmt.Println("file.Read() error:",err.Error())
	}
	if n > 0 {	
		content := string(readbytes[:n])		
		content = strings.Replace(content, jsold_str, jsnew_str, -1)
		content = strings.Replace(content, cssold_str, cssnew_str, -1)
		
		file.Truncate(int64(0)) //删除原先文件中的内容
		file.Seek(0,0)  //移动文件的开始位置
		readbytes = []byte(content) 
		_, write_err := file.Write(readbytes)
		if write_err != nil {
			fmt.Println("file.Write() error", write_err.Error())
		}
		flush_err := file.Sync()
		if flush_err != nil {
			fmt.Println("file.Flush() error", flush_err.Error())
		}
		
	}
	

}

func batchDeal(fpath string) {
	abspath,e := filepath.Abs(fpath)
	if e != nil {
		fmt.Println("filepath.Abs("+abspath+") error:",e.Error())
		return
	}
	file,e := os.Open(abspath)
	defer file.Close()
	if e != nil {
		fmt.Println("os.Open("+abspath+") error:",e.Error())
		return
	}
	fi, e := file.Stat()
	if e != nil {
		fmt.Println("file.Stat() error:", e.Error())
		return
	}
	isDir := fi.IsDir()
	if isDir {
		fnames, e := file.Readdirnames(-1)
		if e != nil {
			fmt.Println("file.Readdirnames(-1) error:", e.Error())
			return 
		}
		for _,fname := range fnames {
			batchDeal(abspath+"\\"+fname)
		}
	} else {
		if name := strings.ToLower(file.Name());  strings.HasSuffix(name,".html") {
			instead(name)
		}
	}

}

func main() {
	fpath := `D:\Android\adt-bundle-windows-x86_64-20140321\sdk\docs`
	//instead(fpath)
	//fpath := "docs"
	batchDeal(fpath)
}