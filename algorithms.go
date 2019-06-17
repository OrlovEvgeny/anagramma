package anagramma

import (
	"sort"
	"strings"
	"sync"
)

//HashMap
type HashMap struct {
	mu      *sync.RWMutex
	safeMap map[string][]string
}

//New
func NewHmap() *HashMap {
	return &HashMap{
		mu:      &sync.RWMutex{},
		safeMap: make(map[string][]string),
	}
}

//Store - return load size
func (hm *HashMap) Store(str ...string) int {
	if len(str) < 1 {
		return 0
	}
	hm.mu.Lock()
	hm.safeMap = make(map[string][]string)
	for _, v := range str {
		sorted := sortAbc(v)
		hm.safeMap[sorted] = append(hm.safeMap[sorted], v)
	}
	tsize := len(hm.safeMap)
	hm.mu.Unlock()
	return tsize
}

//Load - return anagramm words slice equal arg word
func (hm *HashMap) Load(str string) []string {
	sorted := sortAbc(str)

	hm.mu.RLock()
	defer hm.mu.RUnlock()
	if v, ok := hm.safeMap[sorted]; ok {
		return v
	}
	return []string{}
}

//sortAbc - return abc sorted string
func sortAbc(str string) string {
	char := []rune(strings.ToLower(str))
	sort.Slice(char, func(i, j int) bool {
		return char[i] < char[j]
	})

	return string(char)
}
