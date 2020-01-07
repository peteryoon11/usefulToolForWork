package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("test")

	// 출력파일 생성
	fo, err := os.Create("./2.txt")
	if err != nil {
		panic(err)
	}
	defer fo.Close()

	file, err := os.Open("./1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// fmt.Println(scanner.Text())

		temp := "," + scanner.Text() + "\n"
		// fmt.Println(temp)
		fo.Write([]byte(temp))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
