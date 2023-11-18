package main

import (
	"FlowDispatcher/device"
	"FlowDispatcher/worker"
	"log"
	"math/rand"
	"strconv"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomStatus() string {
	rand.Seed(time.Now().UnixNano())
	if rand.Intn(2) != 0 {
		log.Println("wait")
		return "wait"
	} else {
		log.Println("dispatch")
		return "dispatch"
	}
}

func main() {
	worker.InitWorker()
	go func() {
		for i := 0; i < 10; i++ {
			rendStatus := RandomStatus()
			device.SetDeviceStatus("test_"+strconv.Itoa(i), rendStatus)
			//device.SetDeviceStatus("test_"+strconv.Itoa(i), "wait")
		}
	}()
	for {
		for i := 0; i < 10; i++ {
			for flow := 0; flow < 10; flow++ {
				for k := 0; k < 10; k++ {
					log.Println("test_" + strconv.Itoa(i) + "_" + strconv.Itoa(flow) + "_" + strconv.Itoa(k))
					dStatus, exist := device.GetDeviceStatus("test_" + strconv.Itoa(k))
					log.Println("DEVICE STATUS   ", dStatus, exist)
					if exist {
						if dStatus.Status == "wait" && dStatus != nil {
							Si := strconv.Itoa(i)
							Sj := strconv.Itoa(flow)
							Sk := strconv.Itoa(k)
							task := &worker.Task{
								Args: "test_" + Si + "_" + Sj + "_" + Sk,
							}
							taskPool := worker.GetWorkerPool(20, 20)
							taskPool.JobQueue <- task
						} else {
							break
						}
					}
				}
			}
		}
	}

}
