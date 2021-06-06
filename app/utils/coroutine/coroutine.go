package coroutine

import "sync"

func Concurrently(fn ...func()) (result []interface{}) {
	var wg sync.WaitGroup

	wg.Add(len(fn))

	for _, f := range fn {
		go func() {
			defer wg.Done()

			result = append(result, f)
		}()
	}

	wg.Wait()

	return result
}
