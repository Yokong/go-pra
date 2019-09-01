package main

import (
	"fmt"
	"time"

	"github.com/hpcloud/tail"
)

func main() {
	filename := "/Users/yoko/logs/my.log"
	tails, err := tail.TailFile(filename, tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	})
	if err != nil {
		fmt.Println("tail file err:", err)
	}

	var msg *tail.Line
	var ok bool
	for {
		msg, ok = <-tails.Lines
		if !ok {
			fmt.Println("tail file close reopen, filenam:", tails, filename)
			time.Sleep(100 * time.Millisecond)
			continue
		}
		fmt.Println("msg:", msg.Text)
	}
}
