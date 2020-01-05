package benchtrie

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
)

func BenchmarkTreeInsert(b *testing.B) {
	cases := []struct {
		length int
	}{
		{1},
		{10},
		{100},
		{1000},
		{10000},
		{100000},
		{1000000},
	}

	for _, bc := range cases {
		name := fmt.Sprintf("Insert (tree) within %d hosts", bc.length)
		hosts := GenerateHosts(bc.length)
		tree := NewTree()
		b.Run(name, func(b *testing.B) {
			for _, s := range hosts {
				tree.Insert(strings.Split(s, "."))
			}
		})
	}
}

func BenchmarkTreeLookup(b *testing.B) {
	cases := []struct {
		length int
	}{
		{1},
		{10},
		{100},
		{1000},
		{10000},
		{100000},
		{1000000},
	}

	for _, bc := range cases {
		name := fmt.Sprintf("Lookup (tree) within %d hosts", bc.length)
		hosts := GenerateHosts(bc.length)
		tree := NewTree()
		for _, s := range hosts {
			tree.Insert(strings.Split(s, "."))
		}
		key := hosts[rand.Intn(len(hosts))]
		b.Run(name, func(b *testing.B) {
			ok := tree.LookupHost(key)
			if !ok {
				b.Errorf("expected positive lookup for key: %s", key)
			}
		})
	}
}
