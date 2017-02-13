package main

import (
    "fmt"
    "../banner"
)

func main() {
	banner.PrintNewFigure("TEST", "3x5", true)
	fmt.Println(banner.GetNewFigure("TEST", "rounded", true))
	fmt.Println(banner.BannerPrintLineS("=", 40))
}