package benchtrie

import (
	"fmt"
	"math/rand"
	"testing"
)

func BenchmarkLookupSlice(b *testing.B) {
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
		name := fmt.Sprintf("Lookup-slice-%d", bc.length)
		hosts := GenerateHosts(bc.length)
		key := hosts[rand.Intn(len(hosts))]

		b.Run(name, func(b *testing.B) {
			for _, s := range hosts {
				if s == key {
					break
				}
			}
		})
	}
}
