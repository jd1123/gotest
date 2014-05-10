package main

import (
	"fmt"
	"os"
	"io/ioutil"
	)

func readfile(){
	bs, err := ioutil.ReadFile("test.txt")
	if err != nil {
		return
	}
	str := string(bs)
	fmt.Println(str)
	
}
	
func longreadfile(){
	file, err := os.Open("test.txt")
	if err != nil{
		fmt.Println("File Error")
		return
	}
	defer file.Close()
	
	stat, err := file.Stat()
	if err != nil{
		return
	}
	
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil{
		return
	}
	
	str:=string(bs)
	fmt.Println(str)
}

func writefile(){
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	
	file.WriteString("This is a test")
}

func readdir(){
	dir, err := os.Open(".")
	if err != nil {
		return
	}
	
	defer dir.Close()
	
	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return
	}
	
	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}
	
}
		
func main(){
	readdir()
}