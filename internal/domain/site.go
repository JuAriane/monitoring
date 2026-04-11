package domain

import (
	"bufio"
	"fmt"
	"io"
	"monitoring/internal/infra"
	"net/http"
	"os"
	"strings"
	"time"
)

const monitoring = 2
const delay = 5

func initializeMonitoring() {
	fmt.Println("Monitorando...")
	website := readFileWebsite()

	for i := 0; i < monitoring; i++ {
		for i, webwebsite := range website {
			fmt.Printf("Testando site %d: %s\n", i, webwebsite)
			testWebsite(webwebsite)
		}

		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

	fmt.Println("")
}

func testWebsite(website string) {
	resp, err := http.Get(website)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Printf("Site: %s foi carregado com sucesso!\n", website)
		infra.RegisterLog(website, true)
	} else {
		fmt.Printf("Site: %s está com problemas. Status Code: %d\n", website, resp.StatusCode)
		infra.RegisterLog(website, false)
	}

	// Teste commit
}

func readFileWebsite() []string {
	var websites []string
	file, err := os.Open("websites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		websites = append(websites, line)

		if err == io.EOF {
			break
		}
	}

	file.Close()

	return websites
}
