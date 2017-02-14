package main

import (
	"fmt"

	"../terminal"
)

func main() {
	terminal.ClearScreen()
	terminal.WaitD(4)
	terminal.ClearLine()
	terminal.Wait(4, 1)
	terminal.PrintXY(10, 10, "HERE\n")
	cx, cy := terminal.CursorXY()
	fmt.Println(cx, cy)
}
