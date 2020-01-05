# benchtrie

Benchmarking prefix lookups, e.g. given a number of hostnames like

```shell
ads.0.xvlbz.com
ads.1.gbaic.com
ads.2.mrajw.com
ads.3.whthc.com
ads.4.tcuax.com
ads.5.hxkqf.com
ads.6.dafpl.com
ads.7.sjfbc.com
ads.8.xoeff.com
ads.9.rswxp.com
```

How long does a name lookup take with different data structures?

* slice
* tree (https://play.golang.org/p/63j5BvHIS1z)
* [iradix](https://github.com/hashicorp/go-immutable-radix)

Summary:

The *tree* has fastest insert time, *iradix* seems slightly faster on lookup,
slices are ok for up to 100 items, maybe.

## Details

```shell
$ go test -bench .
goos: linux
goarch: amd64
pkg: github.com/miku/benchtrie
BenchmarkTreeInsert/Insert_(tree)_within_1_hosts-8              1000000000               0.000009 ns/op
BenchmarkTreeInsert/Insert_(tree)_within_10_hosts-8             1000000000               0.000013 ns/op
BenchmarkTreeInsert/Insert_(tree)_within_100_hosts-8            1000000000               0.000080 ns/op
BenchmarkTreeInsert/Insert_(tree)_within_1000_hosts-8           1000000000               0.000767 ns/op
BenchmarkTreeInsert/Insert_(tree)_within_10000_hosts-8          1000000000               0.00396 ns/op
BenchmarkTreeInsert/Insert_(tree)_within_100000_hosts-8         1000000000               0.0427 ns/op
BenchmarkTreeInsert/Insert_(tree)_within_1000000_hosts-8               1        2370229227 ns/op
BenchmarkTreeLookup/Lookup_(tree)_within_1_hosts-8              1000000000               0.000003 ns/op
BenchmarkTreeLookup/Lookup_(tree)_within_10_hosts-8             1000000000               0.000003 ns/op
BenchmarkTreeLookup/Lookup_(tree)_within_100_hosts-8            1000000000               0.000002 ns/op
BenchmarkTreeLookup/Lookup_(tree)_within_1000_hosts-8           1000000000               0.000003 ns/op
BenchmarkTreeLookup/Lookup_(tree)_within_10000_hosts-8          1000000000               0.000005 ns/op
BenchmarkTreeLookup/Lookup_(tree)_within_100000_hosts-8         1000000000               0.000007 ns/op
BenchmarkTreeLookup/Lookup_(tree)_within_1000000_hosts-8        1000000000               0.000008 ns/op
BenchmarkImmutableRadixInsert/Insert_(iradix)_within_1_hosts-8          1000000000               0.000023 ns/op
BenchmarkImmutableRadixInsert/Insert_(iradix)_within_10_hosts-8         1000000000               0.000068 ns/op
BenchmarkImmutableRadixInsert/Insert_(iradix)_within_100_hosts-8        1000000000               0.000614 ns/op
BenchmarkImmutableRadixInsert/Insert_(iradix)_within_1000_hosts-8       1000000000               0.00707 ns/op
BenchmarkImmutableRadixInsert/Insert_(iradix)_within_10000_hosts-8      1000000000               0.132 ns/op
BenchmarkImmutableRadixInsert/Insert_(iradix)_within_100000_hosts-8     1000000000               0.584 ns/op
BenchmarkImmutableRadixInsert/Insert_(iradix)_within_1000000_hosts-8           1        6009170489 ns/op
BenchmarkImmutableRadixLookup/Lookup_(iradix)_within_1_hosts-8          1000000000               0.000000 ns/op
BenchmarkImmutableRadixLookup/Lookup_(iradix)_within_10_hosts-8         1000000000               0.000001 ns/op
BenchmarkImmutableRadixLookup/Lookup_(iradix)_within_100_hosts-8        1000000000               0.000001 ns/op
BenchmarkImmutableRadixLookup/Lookup_(iradix)_within_1000_hosts-8       1000000000               0.000001 ns/op
BenchmarkImmutableRadixLookup/Lookup_(iradix)_within_10000_hosts-8      1000000000               0.000001 ns/op
BenchmarkImmutableRadixLookup/Lookup_(iradix)_within_100000_hosts-8     1000000000               0.000004 ns/op
BenchmarkImmutableRadixLookup/Lookup_(iradix)_within_1000000_hosts-8    1000000000               0.000006 ns/op
BenchmarkLookupSlice/Lookup_(slice)_within_1_hosts-8                    1000000000               0.000000 ns/op
BenchmarkLookupSlice/Lookup_(slice)_within_10_hosts-8                   1000000000               0.000000 ns/op
BenchmarkLookupSlice/Lookup_(slice)_within_100_hosts-8                  1000000000               0.000001 ns/op
BenchmarkLookupSlice/Lookup_(slice)_within_1000_hosts-8                 1000000000               0.000005 ns/op
BenchmarkLookupSlice/Lookup_(slice)_within_10000_hosts-8                1000000000               0.000044 ns/op
BenchmarkLookupSlice/Lookup_(slice)_within_100000_hosts-8               1000000000               0.000660 ns/op
BenchmarkLookupSlice/Lookup_(slice)_within_1000000_hosts-8              1000000000               0.00565 ns/op
```
