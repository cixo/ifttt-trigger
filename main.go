package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	args := os.Args
	pwd, _ := os.Getwd()
	host, _ := os.Hostname()
	fmt.Printf("🌋: Hello World!\n💻: %s\n📂: %s\n⏰: %s\n", host, pwd, time.Now().Format("2006-01-02T15:04:05-0700"))

	fmt.Printf("\ngot arguments: %d\n", len(args))
	for i, a := range args {
		fmt.Printf("#%d arg: [%v]\n", i, a)
	}
}
