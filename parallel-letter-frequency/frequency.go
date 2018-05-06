package letter

type FreqMap map[rune]int

func (m FreqMap) Merge(other FreqMap) {
	for r, c := range other {
		m[r] += c
	}
}

func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(words []string) (m FreqMap) {
	ch := make(chan FreqMap, len(words))
	for _, word := range words {
		word := word
		go func() {
			ch <- Frequency(word)
		}()
	}

	m = FreqMap{}
	for range words {
		m.Merge(<-ch)
	}
	return
}
