package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}
var wg sync.WaitGroup // these are used as pointers
var mut sync.Mutex

func main() {
	//    go greeter("Hello")
	//    greeter("World")

	websiteList := []string{
		"https://google.com",
		"https://instagram.com",
		"https://github.com",

	}
	for _, web := range websiteList {
		go getStatusCode(web)
		wg.Add(1)
	}
    wg.Wait()
	fmt.Println(signals)
}



func getStatusCode(endpoint string) {
	defer wg.Done()
	
	resp, err := http.Get(endpoint)
	if err!=nil{
		fmt.Println("Oops in endpoint")
	}else{
		mut.Lock()
	    signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d Status Code for %s\n", resp.StatusCode, endpoint)
     
	}
}

// func greeter(s string) {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(10*time.Millisecond)
// 		fmt.Println(s)
// 	}
// }