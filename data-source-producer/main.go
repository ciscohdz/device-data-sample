package main

import (
	DataPublisher "ciscohdz/samples/devicestreams/data_publisher"
	"context"
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"runtime/trace"
	"time"
)

var ctx context.Context = nil
var task *trace.Task = nil

func main() {

	// f, err := os.Create("trace.out")
	// if err != nil {
	// 	log.Fatal("Error creating a file", err)
	// }
	// defer func() {
	// 	if err := f.Close(); err != nil {
	// 		log.Fatalf("Error closing trace file")
	// 	}
	// }()
	// if err := trace.Start(f); err != nil {
	// 	log.Fatalf("failed to start trace: %v", err)
	// }
	// defer trace.Stop()

	//ctx, task := trace.NewTask(context.Background(), "main start")

	// defer task.End()
	// defer ctx.Done()

	fmt.Printf("Logical Cpus: %d\n", runtime.NumCPU())
	fmt.Printf("Starting up. Cores %d\n", runtime.GOMAXPROCS(4))

	mainChan := make(chan DataPublisher.DeviceData, 100)
	//go processMessage(mainChan)
	go runner(mainChan)
	go DataPublisher.Publish(mainChan)

	select {
	case <-time.After(time.Minute * 1):
		fmt.Println("Exiting from main...")
	}

}

func processMessage(mainChan chan DataPublisher.DeviceData) {

	// if ctx == nil {
	// 	log.Fatal("trace is nil")
	// }
	// r := trace.StartRegion(ctx, "processing channels")
	// defer r.End()

	counter := 0
	start := time.Now()
	for true {
		select {
		case <-time.After(time.Second * 5):
			duration := time.Now().Sub(start)
			fmt.Println("")
			fmt.Printf("Performed %d in %v", counter, duration)
			os.Exit(1)
		case <-mainChan:
			counter += 1
			if counter%1_000_000 == 0 {
				duration := time.Now().Sub(start)
				perSec := float64(counter) / float64(duration.Milliseconds())
				fmt.Printf(" --> %s %d %s\n%v ops/ms\n", Yellow, counter, Reset, perSec)
			}
		}

	}

}

func runner(mainChan chan DataPublisher.DeviceData) {

	// r := trace.StartRegion(ctx, "runner")
	// defer r.End()

	for {
		for i := 0; i < 50; i++ {
			go pushRandomData(mainChan)
		}
		time.Sleep(time.Millisecond * 1000)
	}

}

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Purple = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

func pushRandomData(mainChan chan<- DataPublisher.DeviceData) {
	// cpu bound which means no syscalls and harder for go to timeslice
	for i := 0; i < 10_000; i++ {
		dd := proucdeRandomDeviceData()
		mainChan <- *dd
	}
}

func proucdeRandomDeviceData() *DataPublisher.DeviceData {

	//s1 := rand.NewSource(time.Now().UnixNano())
	//r1 := rand.New(s1)

	randomValue := rand.Float32() * 10_000
	randomId := rand.Intn(255)

	dd := DataPublisher.DeviceData{
		DeviceId: fmt.Sprintf("device%d", randomId),
		Value:    randomValue,
		Ts:       time.Now(),
	}

	return &dd

}
