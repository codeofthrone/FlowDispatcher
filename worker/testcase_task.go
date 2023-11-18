// task.go
package worker

import "log"

type Task struct {
	Args string
}

func (t Task) RunTask(request interface{}) {

	log.Println("Worker run : ", t.Args)
}
