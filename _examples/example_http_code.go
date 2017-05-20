package main

import (
	"fmt"

	"../web/http"
)

func main() {
	fmt.Println(http.GetCodeShortInfo(100))
	fmt.Println(http.GetCodeLongInfo(100))
}
