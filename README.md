你好！
很冒昧用这样的方式来和你沟通，如有打扰请忽略我的提交哈。我是光年实验室（gnlab.com）的HR，在招Golang开发工程师，我们是一个技术型团队，技术氛围非常好。全职和兼职都可以，不过最好是全职，工作地点杭州。
我们公司是做流量增长的，Golang负责开发SAAS平台的应用，我们做的很多应用是全新的，工作非常有挑战也很有意思，是国内很多大厂的顾问。
如果有兴趣的话加我微信：13515810775  ，也可以访问 https://gnlab.com/，联系客服转发给HR。
# benchtrie

Benchmarking name lookups, e.g. given a number of hostnames like

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
*slices* are ok for up to 100 items, maybe.

## Details

```shell
$ go test -bench .
goos: linux
goarch: amd64
pkg: github.com/miku/benchtrie
BenchmarkTreeInsert/Insert-tree-1-8             1000000000               0.000007 ns/op
BenchmarkTreeInsert/Insert-tree-10-8            1000000000               0.000015 ns/op
BenchmarkTreeInsert/Insert-tree-100-8           1000000000               0.000140 ns/op
BenchmarkTreeInsert/Insert-tree-1000-8          1000000000               0.000830 ns/op
BenchmarkTreeInsert/Insert-tree-10000-8         1000000000               0.00807 ns/op
BenchmarkTreeInsert/Insert-tree-100000-8        1000000000               0.0476 ns/op
BenchmarkTreeInsert/Insert-tree-1000000-8              1        2474794588 ns/op
BenchmarkTreeLookup/Lookup-tree-1-8             1000000000               0.000003 ns/op
BenchmarkTreeLookup/Lookup-tree-10-8            1000000000               0.000003 ns/op
BenchmarkTreeLookup/Lookup-tree-100-8           1000000000               0.000002 ns/op
BenchmarkTreeLookup/Lookup-tree-1000-8          1000000000               0.000003 ns/op
BenchmarkTreeLookup/Lookup-tree-10000-8         1000000000               0.000006 ns/op
BenchmarkTreeLookup/Lookup-tree-100000-8        1000000000               0.000007 ns/op
BenchmarkTreeLookup/Lookup-tree-1000000-8       1000000000               0.000007 ns/op
BenchmarkImmutableRadixInsert/Insert-iradix-1-8                 1000000000               0.000008 ns/op
BenchmarkImmutableRadixInsert/Insert-iradix-10-8                1000000000               0.000027 ns/op
BenchmarkImmutableRadixInsert/Insert-iradix-100-8               1000000000               0.000255 ns/op
BenchmarkImmutableRadixInsert/Insert-iradix-1000-8              1000000000               0.00371 ns/op
BenchmarkImmutableRadixInsert/Insert-iradix-10000-8             1000000000               0.127 ns/op
BenchmarkImmutableRadixInsert/Insert-iradix-100000-8            1000000000               0.580 ns/op
BenchmarkImmutableRadixInsert/Insert-iradix-1000000-8                  1        5890982774 ns/op
BenchmarkImmutableRadixLookup/Lookup-iradix-1-8                 1000000000               0.000001 ns/op
BenchmarkImmutableRadixLookup/Lookup-iradix-10-8                1000000000               0.000001 ns/op
BenchmarkImmutableRadixLookup/Lookup-iradix-100-8               1000000000               0.000001 ns/op
BenchmarkImmutableRadixLookup/Lookup-iradix-1000-8              1000000000               0.000001 ns/op
BenchmarkImmutableRadixLookup/Lookup-iradix-10000-8             1000000000               0.000001 ns/op
BenchmarkImmutableRadixLookup/Lookup-iradix-100000-8            1000000000               0.000004 ns/op
BenchmarkImmutableRadixLookup/Lookup-iradix-1000000-8           1000000000               0.000006 ns/op
BenchmarkLookupSlice/Lookup-slice-1-8                           1000000000               0.000000 ns/op
BenchmarkLookupSlice/Lookup-slice-10-8                          1000000000               0.000000 ns/op
BenchmarkLookupSlice/Lookup-slice-100-8                         1000000000               0.000001 ns/op
BenchmarkLookupSlice/Lookup-slice-1000-8                        1000000000               0.000005 ns/op
BenchmarkLookupSlice/Lookup-slice-10000-8                       1000000000               0.000043 ns/op
BenchmarkLookupSlice/Lookup-slice-100000-8                      1000000000               0.00136 ns/op
BenchmarkLookupSlice/Lookup-slice-1000000-8                     1000000000               0.00583 ns/op
PASS
ok      github.com/miku/benchtrie       53.498s
```
