package main
import (
	"os"
	"fmt"
	"log"
	"bufio"
	"net/http"
	"crypto/tls"
)
func main() {
	var terrible string
	colorReset := "\033[0m"
	
	colorRed := "\033[31m"
	colorBlue := "\033[34m"
	url := os.Args[1]
	filePath := os.Args[2]
	readFile, err :=os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()
	
	scanner := bufio.NewScanner(readFile)
	
	for scanner.Scan() {
		line := scanner.Text()
	
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		terrible = url + line
		resp,  err := http.Get(terrible)
			if err != nil{
			log.Fatal(err)
			}
		if err != nil {
			log.Fatal(err)
		}
		
	if err := scanner.Err(); err != nil{
		log.Fatal(err)
		}
	
	
	fmt.Println(string(colorRed), "StatusCode:", string(colorReset), resp.StatusCode,string(colorRed), "Content-Length:", string(colorReset), resp.Header["Content-Length"],string(colorRed), "Server:", string(colorReset), resp.Header["Server"],string(colorBlue), "DIR:", string(colorReset), line)

	}
}
