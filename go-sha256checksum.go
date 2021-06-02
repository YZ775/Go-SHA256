package main

import(
    "fmt"
	"flag"
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
	flag.Parse()
    args := flag.Args()
	if args ==nil{
		fmt.Println("file path required")
		os.Exit(1)
	}

	f, err := os.Open(args[0]) //第一引数のファイルを開く
	filelength :=0

    if err != nil{
        fmt.Println("file open error")
		os.Exit(1)
    }

	input_array := make([]string, 0, 0) //入力を保持するための配列

	fileScanner := bufio.NewScanner(f)
	for fileScanner.Scan() { //ファイルをすべて読み込み配列に入れておく
		line := fileScanner.Text()
		input_array=append(input_array, line)
		filelength+=1
	}

	ch := make([]chan string, filelength)//行数分のチャネルを作る
	for i,s := range input_array{
		ch[i] = make(chan string)
		go culculate_sha256(s, ch[i]) //各行に対してSHA256を計算するgo routineを走らせることで並列処理を行う
	}

	for i, _ :=range input_array{ //各行のチャネルを待ち順番にprintする
		x:= <-ch[i]
		fmt.Printf("%s\n",x)
	}

	f.Close()
}