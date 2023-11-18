package worker

var (
	MaxPoolSize   = 200
	MaxWorkerSize = 100
)

type Job interface {
	RunTask(request interface{})
}

// Job channel
type JobChan chan Job
