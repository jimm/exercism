package letter

type FreqMap map[rune]int

func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(xs []string) FreqMap {
	c := make(chan FreqMap)
	for _, s := range xs {
		go func(s string) { c <- Frequency(s) }(s)
	}

	m := FreqMap{}
	for range xs {
		fmap := <-c
		for k, v := range fmap {
			m[k] += v
		}
	}
	return m
}
