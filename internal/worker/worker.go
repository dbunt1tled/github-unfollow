package worker

import "sync"

type Worker struct {
	tasks chan func()
	wg    sync.WaitGroup
	num   int
}

func NewWorker(numWorkers int, queueSize int) *Worker {
	return &Worker{
		tasks: make(chan func(), queueSize),
		wg:    sync.WaitGroup{},
		num:   numWorkers,
	}
}

func (w *Worker) Start() {
	for i := 0; i < w.num; i++ {
		go func(workerID int) {
			for task := range w.tasks {
				task()
			}
		}(i)
	}
}

func (w *Worker) AddTask(task func()) {
	w.wg.Add(1)
	w.tasks <- func() {
		defer w.wg.Done()
		task()
	}
}

func (w *Worker) Wait() {
	w.wg.Wait()
}

func (w *Worker) Stop() {
	close(w.tasks)
}
