package smpkg

func GetQueueKernel() *Kernel {
	return &Kernel{}
}

type Kernel struct {
	list map[string]Worker
}

func (k *Kernel) Register(queueName string, consumer Worker) {
	if k.list == nil {
		k.list = make(map[string]Worker)
	}
	k.list[queueName] = consumer
}

func (k *Kernel) Start() {
	done := make(chan int, 1)
	for qn, c := range k.list {
		go c.Do(qn)
	}
	<-done
}

type Worker struct {
	Do func(queue string)
}
