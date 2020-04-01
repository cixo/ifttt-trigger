package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	method     = "POST"
	urlBase    = "https://maker.ifttt.com/trigger/%s/with/key/%s"
	timeoutSec = 10
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
		args      = os.Args
		err       error
		key       string
		event     string
		jsonBytes []byte
		req       *http.Request
		resp      *http.Response
		respBytes []byte
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

	if jsonBytes, err = json.Marshal(data); err != nil {
		throwError(fmt.Sprintf("fail to marshal data: %v", err), 4)
	}
	if req, err = http.NewRequest(method, fmt.Sprintf(urlBase, event, key), bytes.NewBuffer(jsonBytes)); err != nil {
		throwError(fmt.Sprintf("fail to create request: %v", err), 5)
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{
		Timeout: timeoutSec * time.Second,
	}

	if resp, err = client.Do(req); err != nil {
		throwError(fmt.Sprintf("fail to send request: %v", err), 6)
	} else {
		defer resp.Body.Close()
	}

	if respBytes, err = ioutil.ReadAll(resp.Body); err != nil {
		throwError(fmt.Sprintf("fail to read response: %v", err), 7)
	}

	fmt.Println("IFTTT says:", string(respBytes))
}
