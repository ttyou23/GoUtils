// thread_pool
package spider

import (
	"sync"
)

var wg sync.WaitGroup

// type ThreadManager struct {
// 	var aw sync.WaitGroup
// }

func Start_work(f func(map[string]string), param map[string]string, cout int) {

	wg.Add(cout)
	for i := 0; i < cout; i++ {
		go working(f, param)
	}
	wg.Wait()
}

func working(f func(map[string]string), param map[string]string) {
	f(param)
	wg.Done()
}
