package datas

type Iterator struct {
	C    <-chan Item
	stop chan struct{}
}

func (i *Iterator) Stop() {
	defer func() {
		recover()
	}()

	close(i.stop)

	for range i.C {
	}
}

func newIterator() (*Iterator, chan<- Item, <-chan struct{}) {
	itemChan := make(chan Item)
	stopChan := make(chan struct{})
	return &Iterator{
		C:    itemChan,
		stop: stopChan,
	}, itemChan, stopChan
}
