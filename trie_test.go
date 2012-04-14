package trie

import (
	. "github.com/sdegutis/go.bdd"
	"testing"
)

func TestEverything(t *testing.T) {
	defer PrintSpecReport()

	Describe("Trie", func() {
		It("Stores anything by its key", func() {
			trie := &Trie{}

			trie.Put("foo", 42)
			Expect(trie.Get("foo").(int), ToEqual, 42)

			trie.Put("bar", 31)
			Expect(trie.Get("bar").(int), ToEqual, 31)

			trie.Put("foobar", 21)
			Expect(trie.Get("foobar").(int), ToEqual, 21)
			Expect(trie.Get("foo").(int), ToEqual, 42)
		})
	})
}

func BenchmarkTrieLookup(b *testing.B) {
	trie := &Trie{}
	trie.Put("foo", 42)
	trie.Put("bar", 31)
	trie.Put("foobar", 21)
	trie.Put("foobarraboof", 100)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		trie.Get("zoo")
	}
}

func BenchmarkMapLookup(b *testing.B) {
	hash := map[string]int{
		"foo": 42, "bar": 31, "foobar": 21, "foobarraboof": 100,
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = hash["zoo"]
	}
}
