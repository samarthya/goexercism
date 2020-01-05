package letter

import "sync"

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

type SafeMap struct {
	val  FreqMap
	lock sync.Mutex
}

func (sm *SafeMap) Increment(e rune) {
	sm.lock.Lock()
	sm.val[e]++
	sm.lock.Unlock()
}

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ParallelFreq(s string, sm *SafeMap, wg *sync.WaitGroup) {
	for _, r := range s {
		sm.Increment(r)
	}

	if wg != nil {
		wg.Done()
	}
}

func ConcurrentFrequency(elements []string) FreqMap {
	var sm = SafeMap{FreqMap{}, sync.Mutex{}}
	var wg sync.WaitGroup

	wg.Add(len(elements))
	for _, r := range elements {
		go ParallelFreq(r, &sm, &wg)
	}
	wg.Wait()

	return sm.val
}
