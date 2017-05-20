package main

import (
	"fmt"

	"../cli/banner"
)

func main() {
	banner.PrintNewFigure("TEST", "3x5", true)
	fmt.Println(banner.GetNewFigure("TEST", "rounded", true))
	banner.BannerPrintLineS("=", 40)
}
