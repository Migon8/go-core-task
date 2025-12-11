package main

// SemaphoreWaitGroup — упрощённая версия WaitGroup,
// построенная на семафоре (буферизированном канале).
type SemaphoreWaitGroup struct {
	ch    chan struct{} // сюда горутины будут "кидать" сигнал о завершении
	total int           // сколько всего горутин мы ждём
}

// NewSemaphoreWaitGroup создаёт новую "wait group" на N горутин.
func NewSemaphoreWaitGroup(total int) *SemaphoreWaitGroup {
	if total < 0 {
		total = 0
	}

	return &SemaphoreWaitGroup{
		ch:    make(chan struct{}, total),
		total: total,
	}
}

// Done нужно вызвать в конце каждой горутины, которую мы ждём.
func (w *SemaphoreWaitGroup) Done() {
	// Если total == 0, просто выходим, чтобы не пытаться писать
	// в канал нулевой ёмкости и не словить дедлок.
	if w.total == 0 {
		return
	}
	w.ch <- struct{}{}
}

// Wait блокируется, пока не получит "total" сигналов завершения.
func (w *SemaphoreWaitGroup) Wait() {
	for i := 0; i < w.total; i++ {
		<-w.ch
	}
}
