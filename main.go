package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type EventData struct {
	Value1 string `json:"value1,omitempty"`
	Value2 string `json:"value2,omitempty"`
	Value3 string `json:"value3,omitempty"`
}

func throwError(msg string, code int) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(code)
}

func printJsonWithIndent(data interface{}) {
	if body, err := json.MarshalIndent(data, "", "  "); err != nil {
		fmt.Println("null")
	} else {
		fmt.Println(string(body))
	}
}

func main() {
	var (
		args  = os.Args
		key   string
		event string
	)

	if len(args) < 3 {
		throwError("missing arguments: ./trigger key event [value1] [value2] [value3]", 1)
	}
	if key = strings.TrimSpace(args[1]); len(key) == 0 {
		throwError("got blank ifttt key", 2)
	}
	if event = strings.TrimSpace(args[2]); len(event) == 0 {
		throwError("got blank event name", 3)
	}

	data := EventData{}
ValueLoop:
	for i, a := range args[3:] {
		s := strings.TrimSpace(a)
		switch i + 1 {
		case 1:
			data.Value1 = s
		case 2:
			data.Value2 = s
		case 3:
			data.Value3 = s
		default:
			break ValueLoop
		}
	}

	fmt.Printf("k: [%s]\n", key)
	fmt.Printf("e: [%s]\n", event)
	fmt.Printf("v: %v\n", data)
	printJsonWithIndent(data)

	//pwd, _ := os.Getwd()
	//host, _ := os.Hostname()
	//fmt.Printf("🌋: Hello World!\n💻: %s\n📂: %s\n⏰: %s\n", host, pwd, time.Now().Format("2006-01-02T15:04:05-0700"))
	//
	//fmt.Printf("\ngot arguments: %d\n", len(args))
	//for i, a := range args {
	//	fmt.Printf("#%d arg: [%v]\n", i, a)
	//}
}
