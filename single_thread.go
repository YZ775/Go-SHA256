package main

import(
    "fmt"
    "os"
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"time"
)

func culculate_sha256(text string,) string {
	time.Sleep(2 * time.Second)
	checksum := sha256.Sum256([]byte(text))
	result:= hex.EncodeToString(checksum[:])
	return result
}

func main(){
	f, err := os.Open("test.txt")
	filelength :=0

    if err != nil{
        fmt.Println("error")
    }

	fileScanner := bufio.NewScanner(f)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		result:= culculate_sha256(line)
		fmt.Printf("%d:%s\n",filelength,result)
		filelength+=1
	}

	f.Close()
}