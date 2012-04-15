package trie

import (
	"fmt"
	. "github.com/sdegutis/go.bdd"
	"testing"
)

func ExampleTrie_Wildcard() {
	t := &Trie{}
	t.Put("Hello", "World")
	t.Put("Hilly", "World")
	t.Put("Hello, bob", "World")
	t.Wildcard("H*ll.") // []string{"Hello", "Hilly"}
	t.Wildcard("Hel")   // []string(nil)
}

func ExampleTrie_LongestPrefix() {
	t := &Trie{}
	t.Put("Hello", "World")
	t.Put("Hello, brother", "World")
	t.Put("Hello, bob", "World")
	t.LongestPrefix("Hello, brandon") // "Hello"
	t.LongestPrefix("Hel")            // ""
	t.LongestPrefix("Hello")          // "Hello"
}

func ExampleTrie_Get() {
	t := &Trie{}
	t.Put("hello", "world")
	fmt.Println(t.Get("hello").(string))
	fmt.Println(t.Get("non-existant"))
	// Output:
	// world
	// <nil>
}

func ExampleTrie_Put() {
	t := &Trie{}
	t.Put("hello", "world")
	t.Put("hello", "world") // does the same thing
	fmt.Println(t.Get("hello"))
	t.Put("1", 1)
	fmt.Println(t.Get("1").(int))
	// Output:
	// world
	// 1
}

func TestEverything(t *testing.T) {
	defer PrintSpecReport()

	Describe("Trie", func() {
		It("Returns nil when a key doesn't exist", func() {
			trie := &Trie{}
			Expect(trie.Get("foo"), ToEqual, nil)
		})

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

		trie := &Trie{}
		trie.Put("Hello", "Oshizushi")
		trie.Put("Hello", "Nigirizushi")
		trie.Put("Hilly", "Narezushi")
		trie.Put("Hello, brother", "Makizushi")
		trie.Put("Hello, bob", "Inarizushi")

		It("Matches wildcards", func() {
			var s []string
			Expect(trie.Wildcard("H*ll."), ToDeepEqual, []string{"Hilly", "Hello"})
			Expect(trie.Wildcard("Hel"), ToDeepEqual, s)
		})

		It("Returns the longest prefix", func() {
			Expect(trie.LongestPrefix("Hello, brandon"), ToEqual, "Hello")
			Expect(trie.LongestPrefix("Hel"), ToEqual, "")
			Expect(trie.LongestPrefix("Hello"), ToEqual, "Hello")
			Expect(trie.LongestPrefix("Hello, bob"), ToEqual, "Hello, bob")
		})

		It("Iterates with a channel", func() {
			keys := []string{}
			values := []string{}
			trie.Each(func(key string, value interface{}) bool {
				keys = append(keys, key)
				values = append(values, value.(string))
				return true
			})
			Expect(keys, ToDeepEqual, []string{"Hello", "Hello, bob", "Hello, brother", "Hilly"})
			Expect(values, ToDeepEqual, []string{"Nigirizushi", "Inarizushi", "Makizushi", "Narezushi"})
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
