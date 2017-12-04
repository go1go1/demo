package main

import (
	"net/http"
	"math/rand"
	"fmt"
	"time"
)
var countNum int = 0;

func lazyHandler(w http.ResponseWriter, req *http.Request) {
	ranNum := rand.Intn(2)
	nNum := func() int {
		countNum ++
		return countNum
	}()
	if ranNum == 0 {
		time.Sleep(6 * time.Second)
		
		fmt.Fprintf(w, "slow response, %d \n", nNum)
		fmt.Printf("slow response, %d\n", nNum)
		return
	}
	fmt.Fprintf(w, "quick response, %d\n", nNum)
	fmt.Printf("quick response, %d\n", nNum)
	return
}

func goHandler(w http.ResponseWriter, req *http.Request) {
	go lazyHandler(w, req)
	return 
}

func main(){
	http.HandleFunc("/", goHandler)
	http.ListenAndServe(":8000", nil)
}

