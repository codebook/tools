package main

import (
  "os"
  "io/ioutil"
  "regexp"
  "fmt"
)

func DelFiles(dir,reg string) {
  
  files,err := ioutil.ReadDir(dir)
  if err != nil {
    fmt.Println("error:",err.Error())
  }
  
  for _,file := range files {
    fileName := file.Name()
    if isDir := file.IsDir(); isDir {
	  DelFiles(dir+"/"+fileName,reg)
	} else {
	  if matched,_ := regexp.MatchString(reg,fileName); !matched {
	    os.Remove(dir+"/"+fileName)
	  }
	}
  }
  // 如果文件夹中的文件都已经被删除，则将该空文件夹删除
  files,err = ioutil.ReadDir(dir)
  fileCount := len(files)
  if fileCount == 0 {
    os.RemoveAll(dir)
  }
}

func main() {
  
  if len(os.Args) != 3 {
    fmt.Println("Usage: command basedir reg-name\nbasedir:where to find the files\nreg-name:the string to mathch the filename")
	return
  } else {
    _,err := os.Open(os.Args[1])
	if err != nil {
	  if exist := os.IsExist(err); !exist {
	    fmt.Println(os.Args[1]+" is not dir")
	    return
	  }
	}
  }
  baseDir := os.Args[1]
  regName := os.Args[2]+".*"
  DelFiles(baseDir,regName)  
}