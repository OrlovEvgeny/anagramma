package anagramma

import (
	"log"
	"testing"
)

var hmap *HashMap

//init
func init() {
	hmap = NewHmap()
}

//TestSortAbc
func TestSortAbc(t *testing.T) {
	en := sortAbc("zkHntmslAgEro")
	if en != "aeghklmnorstz" {
		t.Errorf("sortAbc eng fail result %s\n", en)
	}
	log.Printf("sortAbc eng %s : OK\n", t.Name())
	ru := sortAbc("флЯпрзДкнгцаЬ")
	if ru != "агдзклнпрфцья" {
		t.Errorf("sortAbc rus fail result %s\n", ru)
	}
	log.Printf("sortAbc rus %s : OK\n", t.Name())

}

//TestStore
func TestStore(t *testing.T) {
	length := hmap.Store("kod", "dok", "sok", "rock", "kos", "cokr", "salt", "feed", "rolk", "olkr")
	if length != 6 {
		t.Errorf("hmap store fail result size %d\n", length)
	}
	log.Printf("%s : OK\n", t.Name())
}

func TestLoad(t *testing.T) {
	dok := hmap.Load("dok")
	if len(dok) != 2 {
		t.Errorf("hmap load fail result size %d\n", len(dok))
	}

	log.Printf("%s : OK\n", t.Name())
}
