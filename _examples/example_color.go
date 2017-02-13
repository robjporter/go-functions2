package main

import (
    "fmt"
    "../colors"
)

func main() {
	fmt.Println(colors.Bold("BOLD"))
	fmt.Println(colors.Italic("ITALIC"))
	fmt.Println(colors.Underline("UNDERLINE"))
	fmt.Println(colors.StrikeThrough("STRIKETHROUGH"))
	fmt.Println(colors.Blink("BLINK"))
	fmt.Println(colors.Reversed("REVERSED"))
	fmt.Println(colors.Background("BACKGROUND", colors.RED))
	fmt.Println(colors.Color("COLOR", colors.YELLOW))
	fmt.Println(colors.Highlight("HIGHLIGHT", "GHLI", colors.BRIGHTYELLOW))
	fmt.Println(colors.Color(colors.Bold("BOLD YELLOW"), colors.YELLOW))
	fmt.Println(colors.Background(colors.Color(colors.Bold("BOLD YELLOW WITH BACKGROUND"), colors.YELLOW), colors.RED))
	fmt.Println(colors.Color("BLACK", colors.BLACK))
	fmt.Println(colors.Color("BRIGHTBLACK", colors.BRIGHTBLACK))
	fmt.Println(colors.Color("RED", colors.RED))
	fmt.Println(colors.Color("BRIGHTRED", colors.BRIGHTRED))
	fmt.Println(colors.Color("GREEN", colors.GREEN))
	fmt.Println(colors.Color("BRIGHTGREEN", colors.BRIGHTGREEN))
	fmt.Println(colors.Color("YELLOW", colors.YELLOW))
	fmt.Println(colors.Color("BRIGHTYELLOW", colors.BRIGHTYELLOW))
	fmt.Println(colors.Color("BLUE", colors.BLUE))
	fmt.Println(colors.Color("BRIGHTBLUE", colors.BRIGHTBLUE))
	fmt.Println(colors.Color("MAGENTA", colors.MAGENTA))
	fmt.Println(colors.Color("BRIGHTMAGENTA", colors.BRIGHTMAGENTA))
	fmt.Println(colors.Color("WHITE", colors.WHITE))
	fmt.Println(colors.Color("BRIGHTWHITE", colors.BRIGHTWHITE))
}