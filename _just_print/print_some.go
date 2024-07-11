package main

import (
	"fmt"
	"github.com/lxn/win"
	"os"
	"strings"
	"syscall"
	"time"
)

func main() {
	fmt.Print("See...\n")
	if len(os.Args) > 1 {
		fmt.Printf("args:%v\n", os.Args)
	} else {
		fmt.Println("None.")
	}
	fmt.Println()
}

func getSelectionByClipboard() (string, error) {
	hwnd := win.GetForegroundWindow()
	winText := make([]uint16, 1000)
	GetWindowText(hwnd, &winText[0], uint32(len(winText)-1))
	windowName := syscall.UTF16ToString(winText)

	// save old clipboard content
	oldClipboardText, _ := GetClipboardText()
	// current sequentNumber
	oldSeqNumber, _ := GetClipboardSequenceNumber()

	// send ctrl+c copy event
	if strings.Contains(windowName, "Internet Explore") || strings.Contains(windowName, "Microsoft Edge") {
		// Internet Explore ignore ctrl+c click event
		win.PostMessage(hwnd, win.WM_COMMAND, 0x0001000f, 0)
	} else {
		// send selected to clipboard
		SendCopy()
	}

	clipboardChange := <-waitForClipboardChange(oldSeqNumber)
	if clipboardChange {

		text, _ := GetClipboardText()

		// restore old clipborad state
		UpdateClipboardText(oldClipboardText)

		return text, nil
	}

	//Failed :(
	return "", fmt.Errorf("get selected failed")
}

func waitForClipboardChange(oldSeqNumber uint32) chan bool {
	waitChangeCh := make(chan bool)

	timeout := 200
	runTime := 0
	changeCheckTimer := time.NewTimer(10 * time.Microsecond)

	go func() {

		for {
			<-changeCheckTimer.C

			seqNumber, _ := GetClipboardSequenceNumber()
			if seqNumber != oldSeqNumber {
				waitChangeCh <- true
				return
			}

			runTime += 10
			if runTime > timeout {
				waitChangeCh <- false
				return
			}
			changeCheckTimer.Reset(10 * time.Microsecond)
		}
	}()

	return waitChangeCh
}
