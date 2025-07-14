//atomic counters are used to safely increment, decrement, or read shared integer variables across multiple goroutines, without using mutexes
package main
import(
	"sync/atomic"
	"fmt"
	"sync"
)

func main(){
	var wg sync.WaitGroup //it will wait for all goroutines to finish the work
	var ops atomic.Int64 //We’ll use an atomic integer type to represent our (always-positive) counter

	// starting 50 goroutines, each loops over 1000 times
	for range 50 {
		wg.Add(1) //it means it is adding 1 goroutine
		go func(){
			for range 1000 {
				ops.Add(1) // to atomically increase counter
			}
			wg.Done()
		}()

	}

	wg.Wait() //to wait for all goroutines to finish the execution
	fmt.Println("ops:",ops.Load())
}