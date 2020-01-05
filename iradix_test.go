package benchtrie

import (
	"fmt"
	"math/rand"
	"testing"

	iradix "github.com/hashicorp/go-immutable-radix"
)

func BenchmarkImmutableRadixInsert(b *testing.B) {
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
		name := fmt.Sprintf("Insert (iradix) within %d hosts", bc.length)
		hosts := GenerateHosts(bc.length)
		r := iradix.New()
		b.Run(name, func(b *testing.B) {
			for _, host := range hosts {
				r, _, _ = r.Insert([]byte(host), nil)
			}
		})
	}
}

func BenchmarkImmutableRadixLookup(b *testing.B) {
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
		name := fmt.Sprintf("Lookup (iradix) within %d hosts", bc.length)

		hosts := GenerateHosts(bc.length)
		r := iradix.New()
		for _, host := range hosts {
			r, _, _ = r.Insert([]byte(host), nil)
		}
		key := []byte(hosts[rand.Intn(len(hosts))])
		b.Run(name, func(b *testing.B) {
			_, ok := r.Get(key)
			if !ok {
				b.Errorf("expected positive lookup for key: %s", key)
			}
		})
	}
}
