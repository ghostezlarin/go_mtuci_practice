package main

import (
	"fmt"
	"time"
)

func main() {
	name := "Ларин Павел Сергеевич"
	currentTime := time.Now().Format("2006-01-02 15:04:05")

	fmt.Printf("Привет, %s!\n", name)
	fmt.Println("Текущая дата и время:", currentTime)
}
