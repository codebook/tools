package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func CheckErr(err error) {
	if err != nil {
		fmt.Println("err:%s\n", err.Error())
		return
	}
}

func GetComment(files []string, tofile string) {
	var content []byte
	var commentfile, err = os.Create(tofile)
	CheckErr(err)
	fmt.Printf("Create file %s\n", tofile)

	for _, filename := range files {
		file, err := os.Open(filename)
		CheckErr(err)
		fmt.Printf("Open file %s\n", filename)
		defer file.Close()

		finfo, err := file.Stat()
		CheckErr(err)
		//不是文件夹（即是文件）才处理
		if !finfo.IsDir() {
			fileReader := bufio.NewReader(file)
			preIsComment := false
			for {
				line, _, err := fileReader.ReadLine()
				if err == io.EOF {
					break
				}
				strline := strings.Trim(string(line), " ")
				if strings.HasPrefix(strings.ToLower(strline), "comment on") {
					content = append(content, line...)
					content = append(content, []byte{'\r', '\n'}...)
					if !strings.HasSuffix(strline, `;`) { //如果Comment是多行，需要将下面的一行给加上来
						preIsComment = true
					}
				} else if strings.HasSuffix(strline, ";") && preIsComment {
					content = append(content, line...)
					content = append(content, []byte{'\r', '\n'}...)
					preIsComment = false
				} else {
					if preIsComment { //如果Comment是多行，需要将下面的一行给加上来
						content = append(content, line...)
						content = append(content, []byte{'\r', '\n'}...)
						preIsComment = true
					} else {
						preIsComment = false
					}
				}
			}
		}
	}
	commentfile.Write(content) //写入文件
	commentfile.Sync()         //刷入磁盘
}

func main() {
	finfos, err := ioutil.ReadDir("comment")
	CheckErr(err)
	var files []string
	for _, finfo := range finfos {
		fname := "comment/" + finfo.Name()
		files = append(files, fname)
	}
	tofile := "comment.sql"
	GetComment(files, tofile)
}
