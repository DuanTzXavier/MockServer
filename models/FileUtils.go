package models

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/axgle/mahonia"
)

func GetUIDList(dirPth string) (dirs []string, err error) {
	dirs = make([]string, 0, 10)
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}
	for _, fi := range dir {
		if fi.IsDir() { // 忽略目录
			print(fi.Name() + "\n")
			dirs = append(dirs, fi.Name())
		}
	}
	return dirs, nil
}

//获取指定目录下的所有文件，不进入下一级目录搜索，可以匹配后缀过滤。
func ListDir(dirPth string, suffix string) (files []string, err error) {
	files = make([]string, 0, 10)
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}
	PthSep := string(os.PathSeparator)
	suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写
	for _, fi := range dir {
		if fi.IsDir() { // 忽略目录
			continue
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) { //匹配文件
			files = append(files, dirPth+PthSep+fi.Name())
		}
	}
	return files, nil
}

func ReadAll(filePth string) ([]byte, error) {
	f, err := os.Open(filePth)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}

func ReadFile(filePath string) (string, error) {
	print(filePath)
	inputFile, inputError := os.Open(filePath)
	if inputError != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got acces to it?\n")
		return "", inputError
	}
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	context := ""
	for {
		inputString, readerError := inputReader.ReadString('\n')
		context = context + inputString
		fmt.Printf("The input was: %s", inputString)
		if readerError == io.EOF {
			break
		}
	}

	buf := make([]byte, 1024)
	//文件ex7.txt的编码是gb18030
	decoder := mahonia.NewDecoder("gbk")
	if decoder == nil {
		return "", nil
	}
	var str string = ""
	for {
		n, _ := inputFile.Read(buf)
		if 0 == n {
			break
		}
		//解码为UTF-8
		str += decoder.ConvertString(string(buf[:n]))
	}

	return str, nil
}

// func BytesToString(bs []byte) string {
// 	l := len(bs)
// 	buf := make([]string, 0, l)
// 	for i := 0; i < l; i++ {
// 		buf = appendString(buf, bs[i])
// 	}
// 	return strings.Join(buf, dot)
// }

// func appendString(bs []string, b byte) []string {
// 	var a byte
// 	var s int
// 	for i := 0; i < 8; i++ {
// 		a = b
// 		b <<= 1
// 		b >>= 1
// 		switch a {
// 		case b:
// 			s += 0
// 		default:
// 			temp := 1
// 			for j := 0; j < 7-i; j++ {
// 				temp = temp * 2
// 			}
// 			s += temp
// 		}

// 		b <<= 1
// 	}

// 	return append(bs, strconv.Itoa(s))
// }
