// worker.go
package worker

type Worker struct {
	// 此queue沒有緩衝
	JobQueue JobChan
	Quit     chan bool
}

func InitWorker() {
	workerPool := GetWorkerPool(20, 20)
	workerPool.Start()
}

func (w Worker) Start(workerPool *WorkerPool) {
	// 開新thread，避免阻塞
	go func() {
		for {
			// 註冊worker到pool內
			workerPool.WorkerQueue <- &w
			select {
			case job := <-w.JobQueue:
				job.RunTask(nil)
			case <-w.Quit:
				return
			}
		}
	}()
}

func (w Worker) Stop() {
	go func() {
		w.Quit <- true
	}()
}

func NewWorker() Worker {
	return Worker{
		make(JobChan),
		make(chan bool),
	}
}
