package main

import (
	"fmt"
	"net"
	// "sync"
	"time"
	// "bufio"
)


func worker(ports, results chan int){
	// take ports and perform scan and sent that result
	for port := range ports{
		address := fmt.Sprintf("scanme.nmap.org:%d",port)
		conn, err := net.Dial("tcp",address)
		if err!=nil{
			results <- -1
			continue
		}
		conn.Close()
		results <- port
	}
}


func main() {

	fmt.Println("Scanning started...")
	start := time.Now()
	procSize := 1024
	ports := make(chan int,procSize)
	results := make(chan int, procSize)
	for i:=0; i<100; i++{
		go worker(ports, results)
	}

	maxPortLimit := 1023
	
	go func() {
		for i:=0 ; i< maxPortLimit; i++{
			ports <- i
		}
	}()
	
	count := 0
	for {
		port := <-results
		if port!=-1{
			fmt.Println("Open port: ", <-results)
		}
		count++
		fmt.Println(count)
		if count>1024{
			fmt.Println("leakage warn")
		}
		fmt.Printf("\n Time elapsed: %v \n", time.Since(start))
	}
	fmt.Println("done")

	
}

