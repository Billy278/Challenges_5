package main

import (
	"fmt"
	"sync"
	"time"
)

type infData interface{}

func main() {

	dataTemp1 := []infData{}
	dataTemp2 := []infData{}

	var data1 infData = []string{"bisa1", "bisa2", "bisa3", "bisa4"}
	var data2 infData = []string{"coba1", "coba2", "coba3", "coba4"}
	fmt.Println("Without mutex")
	ch1 := generateChannelWithoutMutex([]infData{data1, data2})

	for i := 0; i < 8; i++ {
		dataTemp1 = append(dataTemp1, <-ch1)
	}
	for _, v := range dataTemp1 {
		fmt.Printf("%+v \n", v)
	}

	fmt.Println("With mutex")
	mutex := sync.Mutex{}
	ch2 := generateChannelMutex([]infData{data1, data2}, &mutex)
	for i := 0; i < 8; i++ {
		dataTemp2 = append(dataTemp2, <-ch2)
	}
	for _, v := range dataTemp2 {
		fmt.Printf("%v \n", v)
	}
}

func generateChannelWithoutMutex(data []infData) chan infData {
	// design pattern: generator
	chanData := make(chan infData, 8)

	func() {
		defer close(chanData)
		for i := 0; i < 4; i++ {
			go func() {
				chanData <- data[0]
			}()
			go func() {
				chanData <- data[1]
			}()

		}
		time.Sleep(time.Second)
	}()

	return chanData
}

func generateChannelMutex(data []infData, mx *sync.Mutex) chan infData {
	// design pattern: generator
	chanData := make(chan infData, 8)

	func() {
		defer close(chanData)
		for i := 0; i < 4; i++ {
			go func() {
				mx.Lock()
				chanData <- data[0]
				mx.Unlock()
			}()
			go func() {
				mx.Lock()
				chanData <- data[1]
				mx.Unlock()
			}()
		}
		time.Sleep(time.Second)
	}()

	return chanData
}
