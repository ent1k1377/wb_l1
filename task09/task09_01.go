package main

func main() {
	writeCh := make(chan int, 1)
	readCh := make(chan<- int, 1)

	list := []int{1, 2, 3, 4, 5, 6, 7, 8}
	for _, v := range list {
		readCh <- v
	}

	go func() {
		for ch := range readCh {

		}
		for ch := range readCh {
			writeCh <- 1
		}
	}()
}
