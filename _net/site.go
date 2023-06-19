package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	ServerHost = "localhost"
	BasePort   = 18000
)

type WebAnywhere struct {
	SleepTime *int
}

func main() {
	w := new(WebAnywhere)

	port := BasePort
	//如果有传入至少1个参数，则第1个参数视为端口号
	argLength := len(os.Args)
	if argLength >= 2 {
		addPort, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Printf("strconv err:%s \n", err.Error())
		} else {
			port = BasePort + addPort
		}
		//如果有传入至少2个参数，则第2个参数视为休眠时间
		if argLength >= 3 {
			tmpTime, _ := strconv.Atoi(os.Args[2])
			w.SleepTime = &tmpTime
		}
	}
	listenUrl := fmt.Sprintf("%s:%d", ServerHost, port)
	w.webV2(listenUrl)
}

//////////////////////////////////
//func webV1() {
//	listenUrl := fmt.Sprintf("%s:%s", ServerHost, BasePort)
//	listen, err := net.Listen("tcp", listenUrl)
//	if err != nil {
//		fmt.Printf("err:%s", err.Error())
//	}
//	println("OK~")
//	defer listen.Close()
//
//	for {
//		// 在没有接收到数据前这里是阻塞的
//		conn, err2 := listen.Accept()
//		if err2 != nil {
//			// 处理异常
//			fmt.Printf("err2:%s", err2.Error())
//			continue
//		}
//		// 开启协程
//		go handleConnectV1(conn)
//
//		// // 字符串写入到客户端
//		// fmt.Fprintln(conn, "See~")
//	}
//}
//func handleConnectV1(conn net.Conn) {
//	data := make([]byte, 4096)
//	n, err := conn.Read(data)
//	if err != nil {
//		// 处理异常
//	}
//	fmt.Println("n is :", n)
//	// fmt.Fprintln(conn, "Hello from TCP server")
//}
//
//////////////////////////////////

func (w *WebAnywhere) webV2(listenUrl string) {
	// indexForV2 为向 url发送请求时，调用的函数
	http.HandleFunc("/", w.indexForV2)
	fmt.Printf("This is %s \n", listenUrl)
	log.Fatal(http.ListenAndServe(listenUrl, nil))
}

func (w *WebAnywhere) indexForV2(rw http.ResponseWriter, r *http.Request) {
	//默认不解析，需手动调用 r.ParseForm 方法进行解析
	_ = r.ParseForm()
	body, _ := io.ReadAll(r.Body)
	//解析方案-a
	// bodyStr, _ := json.Marshal(string(body))
	// fmt.Printf("os\t bodyStr:%s\n", bodyStr)

	//解析方案-b
	//如果没有传第2个参数，则将休眠时间，定为随机
	if w.SleepTime == nil {
		fmt.Printf("w-sleepTime is nil\n")
		rand.Seed(time.Now().UnixNano())
		var maxSleepTime = 5
		tmpTime := rand.Intn(maxSleepTime)
		w.SleepTime = &tmpTime
	}
	fmt.Printf("final-sleep %d second(s)\n", *w.SleepTime)

	var bodyMap map[string]interface{}
	_ = json.Unmarshal(body, &bodyMap)
	fmt.Printf("os bodyMap:%v\n======================\n", bodyMap)
	for k, v := range bodyMap {
		fmt.Printf("dataMap-[%s]:[%v]\n", k, v)
	}
	time.Sleep(time.Second * time.Duration(*w.SleepTime))

	for k, v := range r.Form {
		//默认v的值是数组，需要Join转换一下
		vInString := strings.Join(v, "")
		//api输出
		_, _ = fmt.Fprintf(rw, "k:%s, v:%s \n", k, vInString)
		//控制台输出
		fmt.Printf("params k:%s v:%s\n", k, vInString)
	}
	writeNumber, _ := fmt.Fprintf(rw, "\n")
	timeFormatter := "2006-01-02 15:04:05"
	fmt.Printf("number:%d[%s]\n", writeNumber, time.Now().Format(timeFormatter))
}
