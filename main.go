// This program implements a prototype Forth system
// Its entended as an experimental platform
package main

import (
	"bufio"
	"fmt"
	"github.com/fh-wedel/DBTI/bufferinterface"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Data Base Theory and Implementation Interaction\n")
	fmt.Println()

	terminal := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("DBTI> ")
		cmd, _ := terminal.ReadString('\n')
		cmd = strings.TrimSpace(cmd)

		if strings.HasPrefix(cmd, "help") {
			fmt.Printf("Data Base Theory and Implementation Interaction\n")
			fmt.Println()
			fmt.Printf("help\t\tThis help\n")
			fmt.Printf("quit\t\tExit interaction\n")
			fmt.Printf("page <n>\tShow page n\n")
			fmt.Printf("erase <n>\tFill page n with zero byte\n")
			fmt.Println()
			continue
		}

		if strings.HasPrefix(cmd, "quit") {
			bufferinterface.Flush()
			break
		}

		if strings.HasPrefix(cmd, "erase") {
			pageNo, err := strconv.Atoi(strings.Split(cmd, " ")[1])
			if err != nil {
				fmt.Println(err.Error)
				continue
			}
			p, err := bufferinterface.Request(pageNo)
			if err != nil {
				fmt.Printf(err.Error())
			} else {
				for i := 5; i < bufferinterface.PageSize; i++ {
					p[i] = 0
				}
				bufferinterface.Update(pageNo)
			}
			continue
		}

		if strings.HasPrefix(cmd, "page") {
			pageNo, err := strconv.Atoi(strings.Split(cmd, " ")[1])
			if err != nil {
				fmt.Println(err.Error)
				continue
			}
			p, err := bufferinterface.Request(pageNo)
			if err != nil {
				fmt.Printf(err.Error())
			} else {
				for i := 0; i < bufferinterface.PageSize; i++ {
					fmt.Printf("%02X ", p[i])
					if (i+1)%32 == 0 {
						fmt.Println()
					}
				}
			}
			continue
		}

		fmt.Printf("Unknown command '%s'. Try help\n", cmd)
	}
}
