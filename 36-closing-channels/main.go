package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			job, more := <-jobs
			if more {
				fmt.Println("job received:", job)
			} else {
				fmt.Println("all jobs received")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("job sent:", j)
	}

	close(jobs)
	fmt.Println("all jobs sent")

	<-done

	_, ok := <-jobs
	fmt.Println("more jobs received:", ok)
}
