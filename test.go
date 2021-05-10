package main

import(
    "fmt"
    "os"
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

func culculate_sha256(text string, ch chan string) { //SHA256チェックサムをchannelで返す
	checksum := sha256.Sum256([]byte(text))
	result:= strings.ToUpper(hex.EncodeToString(checksum[:]))
	ch <- result
}

func main(){

	array := make([]string, 0, 0)

	f, err := os.Open("test.txt")
	filelength :=0

    if err != nil{
        fmt.Println("file open error")
    }

	fileScanner := bufio.NewScanner(f)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		array=append(array, line)
		filelength+=1
	}

	ch := make([]chan string, filelength)
	for i,s := range array{
		ch[i] = make(chan string)
		go culculate_sha256(s, ch[i])
	}
	for i, _ :=range array{
		x:= <-ch[i]
		fmt.Printf("%d:%s\n",i,x)
	}

	f.Close()
}