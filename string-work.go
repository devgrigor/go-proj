package main

import (
	"fmt"
	"strings"
	"bytes"
	"os"
	"io/ioutil"
)

func createFile(fileName string, context string) {
	file, err := os.Create(fileName);

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close();

	file.WriteString(context);
}

func openDir(dirName string) {
	dir, err := os.Open("sub")
	
	if err != nil {
		fmt.Println(err)
		return
	}

	defer dir.Close()

	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}

}

func main() {
	createFile("some.txt", "something else here")
	openDir("sub")
	var buf bytes.Buffer
	buf.Write([]byte("somehting"))

	bs1, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	str1 := string(bs1)

	fmt.Println(str1)

	// Working with files
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
		return 
	}
	
	// This will happen in the end of function
	defer file.Close()

	// get file size
	stat, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("File size", stat.Size())

	// Read file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return
	}

	str := string(bs)
	fmt.Println(str)

	input := string([]byte{'t','e','s'})
	arr := []byte("test")
	fmt.Println(arr)

	fmt.Scanln(&input)

	fmt.Println(
		strings.Contains(input, "es"),

		// 2
		strings.Count(input, "t"),

		strings.HasPrefix(input, "te"),

		strings.Index(input, "t"),

		strings.Split(input, "-"),
	)

	fmt.Scanln(&input)

	fmt.Println(input)
}