package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
	// "sync"
)

//txtfilelegacy
//精简构建-1：go build -ldflags "-s -w" -o rsw.exe routine.go
//精简构建-2：go build -ldflags "-s -w -H=windowsgui" -o rswH.exe routine.go

type RoutineTool struct {
}

const (
	DateTimeFormatter = "2006-01-02 15:04:05"
	DateFormatter     = "2006-01-02"
	TimeFormatter     = "15:04:05"
)

func main() {
	fmt.Println("======== The Start ========")
	defer fmt.Println("======== The End ========")

	rt := new(RoutineTool)
	portList := []int{
		18001, 18002, 18003,
	}
	ch := make(chan int, len(portList)+1)
	for _, v := range portList {
		go rt.catchContent(v, ch)
	}
	whoIsNo1 := <-ch
	fmt.Printf("No.1 pt:[%d]\n", whoIsNo1)
	//rt.catchContent(18003)
}

func (rt *RoutineTool) callRun() {
	var baseChan = make(chan int, 10)
	var newChan = make(chan int, 10)
	go rt.sendItems(baseChan)
	go rt.addSomething(newChan, baseChan)
	rt.receiveItems(newChan)
}

func (rt *RoutineTool) catchContent(port int, ch chan<- int) {
	log.Printf("come here %d", port)
	timeStr := time.Now().Format(TimeFormatter)
	url := fmt.Sprintf("http://localhost:%d?time=%s", port, timeStr)
	values := map[string]string{
		"youPort": fmt.Sprintf("port-%d", port),
	}
	jsonData, _ := json.Marshal(values)
	req, errGet := http.NewRequest(http.MethodGet, url, bytes.NewBuffer(jsonData))
	if errGet != nil {
		fmt.Printf("get url-get:%s\n", errGet.Error())
		return
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Errored when sending request to the server", err.Error())
		return
	}

	defer resp.Body.Close()
	responseBody, errRead := ioutil.ReadAll(resp.Body)
	if errRead != nil {
		log.Fatal(errRead)
	}

	fmt.Println(resp.Status)
	fmt.Println(string(responseBody))
	fmt.Printf("[port:%d] OK\n", port)
	ch <- port
}

func (rt *RoutineTool) replaceAndPrint(beginStr string) {
	var str string
	str = "abc nba ban"
	newStr := strings.ReplaceAll(str, "a", "A")
	fmt.Printf("%s begin to test:%s\n", beginStr, newStr)
}

func (rt *RoutineTool) sendItems(out chan int) {
	fmt.Printf("Go now:%v\n", out)
	list := []int{1, 2, 3, 4}
	for k, v := range list {
		fmt.Printf("Put [%d] %d to channel \n", k, v)
		out <- v
	}
	//如果遍历过了，则后续需要重新发送数据到管道中
	// for itemInCh := range out {
	// 	fmt.Printf("Get %d from channel \n", itemInCh)
	// 	// break
	// }
	close(out)
}

func (rt *RoutineTool) addSomething(out chan<- int, in <-chan int) {
	for v := range in {
		fmt.Printf("before add %d \n", v)
		out <- v * 100
	}
	close(out)
}

func (rt *RoutineTool) receiveItems(in chan int) {
	//取管道内容，range时只有值(value)，没有键(key)
	for v := range in {
		fmt.Printf("See InOut %d \n", v)
	}
}
