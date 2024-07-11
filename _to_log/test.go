package main

import (
    "fmt"
    "github.com/go-vgo/robotgo"
    "golang.org/x/text/encoding/simplifiedchinese"
    // "os/exec"
    // "bytes"
    "crypto/md5"
    "encoding/hex"
    "os"
    "time"
)

func main() {
    tellInfo()
}

func tellInfo() {

    x, y := robotgo.GetMousePos()
    fmt.Printf("pos-x:%v [%T]\n", x, x)
    fmt.Printf("pos-y:%v [%T]\n", y, y)

    title := robotgo.GetTitle()
    titleGbk, _ := simplifiedchinese.GBK.NewDecoder().String(title)
    fmt.Printf("titleGbk:%v [%T]\n", titleGbk, titleGbk)

    color := robotgo.GetPixelColor(x, y)
    fmt.Printf("color:%v [%T]\n", color, color)

    currPid := robotgo.GetPID()
    fmt.Printf("currPid:%v [%T]\n", currPid, currPid)
    currTitle, _ := simplifiedchinese.GBK.NewDecoder().String(robotgo.GetTitle(currPid))
    fmt.Printf("titleGbk:%v\n", titleGbk)
    fmt.Printf("currTitle:%v\n", currTitle)

    // activePid := robotgo.ActivePID()
    // fmt.Println(robotgo.GetActive())
    // fmt.Println(robotgo.GetHandle())
    // fmt.Println(robotgo.Process())

    // pids, _ := robotgo.Pids()
    // fmt.Printf("pids:%v [%T]\n", pids, pids)
    // fmt.Printf("robotgo.KeySleep:%v\n", robotgo.KeySleep)

    // var output1 bytes.Buffer
    // cmd1 := exec.Command("cmd", "/c", "dir", "/b")
    // cmd1.Stdout = &output1
    // // cmd1.Start()
    // cmd1.Run()
    // fmt.Printf("\n[begin...]\n%v[end...]\n", output1.String())

    // cmd2 := exec.Command("cmd", "/c", "pause")
    // var output2 bytes.Buffer
    // cmd2.Stdout = &output2
    // // cmd2.Start()
    // cmd2.Run()
    // gbkOutput, _ := simplifiedchinese.GBK.NewDecoder().String(output2.String())
    // fmt.Printf("%v\n", gbkOutput)

    fileName := "log_a.txt"
    file, opErr := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND, 0755)
    if opErr != nil {
        fmt.Printf("ERROR! opErr: %v\n", opErr)
        return
    }
    defer file.Close()

    timeFormatter := "2006-01-02 15:04:05"
    timeStr := time.Now().Format(timeFormatter)

    var argList = []interface{}{
    // argList := []interface{}{
        x,
        y,
        currPid,
        color,
        titleGbk,
    }

    // var argList2 []interface{}
    // {
    // }
    // fmt.Printf("argList2:%T\n", argList2)

    // if false {
    //     println(argList[0])
    //     return
    // }
    // argList = append(argList, x)
    // argList = append(argList, y)
    // argList = append(argList, currPid)
    // argList = append(argList, color)
    // argList = append(argList, titleGbk)

    allArg2Str := fmt.Sprint(argList...)
    byteStr := []byte(allArg2Str)
    hash := md5.Sum(byteStr)
    md5Str := hex.EncodeToString(hash[:])

    preSlice := []interface{}{
        md5Str,
        timeStr,
    }

    argList = append(preSlice, argList...)
    s :=  fmt.Sprintf("%s[%s]%d_%d_%d_%s_%s\t\n", argList...)
    cnt, _ := file.WriteString(s)
    fmt.Printf("cnt:[%d]\n", cnt)
    file.Sync()
}
