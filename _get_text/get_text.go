package main

import (
    "fmt"
    // "github.com/atotto/clipboard"

    "syscall"
    "unsafe"
)

func main() {
    // fmt.Println("a")
    // content, err := clipboard.ReadAll()
    // if err != nil {
    //     panic(err)
    // }
    if true {
        fmt.Printf("haha ha")
//         var h syscall.Handle
//         b := make([]uint16, 200)
//         _, err := GetWindowText(h, &b[0], int32(len(b)))
//         if err != nil {
//             fmt.Printf("%v\n", b)
//         }
    }
    // fmt.Println(content)
}

var (
    user32             = syscall.MustLoadDLL("user32.dll")
    procEnumWindows    = user32.MustFindProc("EnumWindows")
    procGetWindowTextW = user32.MustFindProc("GetWindowTextW")
)

func GetWindowText(hwnd syscall.Handle, str *uint16, maxCount int32) (len int32, err error) {
    r0, _, e1 := syscall.Syscall(procGetWindowTextW.Addr(), 3, uintptr(hwnd), uintptr(unsafe.Pointer(str)), uintptr(maxCount))
    len = int32(r0)
    if len == 0 {
        if e1 != 0 {
            err = error(e1)
        } else {
            err = syscall.EINVAL
        }
    }
    return
}
