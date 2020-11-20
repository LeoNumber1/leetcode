package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	method1()
	//method2()
	//method3()
	//method4()
}

func method1() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
	}
	fmt.Printf("bufio.NewScanner:%q\r\n", scanner.Text())
}

func method2() {
	inputBytes := make([]byte, 512)
	_, err := os.Stdin.Read(inputBytes)
	if err != nil {
		fmt.Println("read error:", err)
	}
	textBytes := bytes.TrimRight(inputBytes, "\x00")
	fmt.Printf("os.Stdin.Read: %q\r\n", strings.TrimSpace(string(textBytes)))
}

func method3() {
	inputText := ""
	fmt.Scanf("%s", &inputText) //注意此方法在win下会因为\r\n读取两次
	fmt.Printf("fmt.Scanf: %q\r\n", inputText)
}

func method4() {
	inputText2 := ""
	stdin := bufio.NewReader(os.Stdin)
	fmt.Fscan(stdin, &inputText2)
	fmt.Printf("fmt.Fscan: %q\r\n", inputText2)
}
