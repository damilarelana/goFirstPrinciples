package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// initialize the new instance of sync.Mutex
	newSyncMutexInstance := new(sync.Mutex)

	// start separate 5 goroutines with each having an infinite loop and mutual exclusive locks
	for counter := int64(0); counter < 5; counter++ {
		// create and immediately run an anonymous function
		go func(counter int64) {
			newSyncMutexInstance.Lock()
			fmt.Println("start of goroutine:", counter)
			time.Sleep(time.Second) // this is when the lock control gets changed between each go routine i.e. when they are sleeping
			fmt.Println("end of goroutine:", counter)
			fmt.Println(" ")
			fmt.Println("==== ==== ====")
			newSyncMutexInstance.Unlock()
		}(counter)

	}
	// create code termination hack using an expected input  that is not used for anything
	var unusedInput string
	fmt.Scanln(&unusedInput)
}
