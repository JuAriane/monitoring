package infra

import (
	"fmt"
	"os"
	"time"
)

func RegisterLog(website string, status bool) {
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + website + " - online: " + fmt.Sprint(status) + "\n")

	file.Close()
}

func readLog() {
	file, err := os.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
		return
	}

	fmt.Println(string(file))
}
