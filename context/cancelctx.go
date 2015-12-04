package main

import (
	"fmt"
	"time"

	"golang.org/x/net/context"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // automatically cancel when done

	result := make(chan int64, 1)
	go func() {
		result <- doWork(ctx)
	}()

	// wait for work to finish
	// change time.After() to control work time window
	// the longer the time window, the more work gets done
	// before cancel() is called.
	select {
	case <-time.After(5 * time.Millisecond):
		fmt.Println("Work too long... cancelling")
		cancel() // cancel downstream
		fmt.Println("Calculated ", <-result)

	case <-ctx.Done():
		fmt.Println("Work was cancelled...")
	case val := <-result:
		fmt.Println("Work completed with value ", val)
	}

}

func doWork(ctx context.Context) int64 {
	_, cancel := context.WithCancel(ctx)
	defer cancel() // cancel everthing downstream (if necessary)

	number := int64(0)
	for i := 0; i < 100000; i++ {
		number = number + (func(i int) int64 {
			cnt := int64(0)
			for k := 0; k <= i; k++ {
				cnt = cnt + int64(k)
			}
			return cnt
		}(i))

		// should we contnue?
		// if not, return to cancel  upstream
		if number > 10000000000 {
			return number
		}

		// was the work cancelled upstream?
		select {
		// yes
		case <-ctx.Done():
			fmt.Println("Work cancelled...")
			return number

		// No, continue
		default:
		}
	}
	return number
}
