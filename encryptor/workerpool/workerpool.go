package workerpool

import (
	"github.com/k1dan/string-encryptor/encryptor/encrypt"
	"sync"
)

func PooledWork(strings []string, workersNumber int) (result []string) {
	var wg sync.WaitGroup

	r := make([]string, len(strings))

	dataCh := make(chan string, workersNumber)
	resultCh := make(chan string, workersNumber)

	for i := 0; i < workersNumber; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for string := range dataCh {
				res := encrypt.EncryptString(string)
				resultCh <- res
			}
		}()
	}

	for i, _ := range strings {
		dataCh <- strings[i]
	}

	for i, _ := range r {
		r[i] = <- resultCh
	}

	return r
}