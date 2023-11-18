// dispatcher.go
package worker

import (
	"sync"
)

var workerPool *WorkerPool
var once sync.Once

// thread pool
type WorkerPool struct {
	Size        int
	JobQueue    JobChan
	WorkerQueue chan *Worker
}

// singletone模式創建WorkerPool
func GetWorkerPool(poolSize, jobQueueLen int) *WorkerPool {
	once.Do(func() {
		workerPool = NewWorkerPool(poolSize, jobQueueLen)
	})
	return workerPool
}

func NewWorkerPool(poolSize, jobQueueLen int) *WorkerPool {
	return &WorkerPool{
		poolSize,
		make(JobChan, jobQueueLen),
		make(chan *Worker, poolSize),
	}
}

func (wp *WorkerPool) Start() {

	// 啟動所有worker
	for i := 0; i < wp.Size; i++ {
		worker := NewWorker()
		worker.Start(wp)
	}

	// 監聽JobQueue，如果接收到request，隨機取一個Worker，並把Job發到該worker的job queue
	// 為了保證不阻塞，因此需要開一個新thread
	go func() {
		for {
			select {
			case job := <-wp.JobQueue:
				worker := <-wp.WorkerQueue
				worker.JobQueue <- job
			}
		}
	}()

}
