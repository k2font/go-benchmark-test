# go-benchmark-test
## Description
- Goの標準ベンチマークツールと`testing`パッケージを用いて、コードのパフォーマンスを計測するテスト

## Usage
### プロファイルの測定
```bash
$ go test -bench .

goos: darwin
goarch: arm64
pkg: example.com/go-benchmark-test
BenchmarkItoaByFmt-10           11068140                94.80 ns/op
BenchmarkItoaByStrconv1-10      27095732                48.18 ns/op
BenchmarkItoaByStrconv2-10      35214151                35.13 ns/op
PASS
ok      example.com/go-benchmark-test   4.129s
```

左から「プログラムを実行した回数」「1回あたりの実行にかかった時間」が表示されている

### メモリ割り当て量の出力
```bash
$ go test -bench . -benchmem

goos: darwin
goarch: arm64
pkg: example.com/go-benchmark-test
BenchmarkItoaByFmt-10           12576721                88.22 ns/op           98 B/op          1 allocs/op
BenchmarkItoaByStrconv1-10      33708693                46.22 ns/op          101 B/op          0 allocs/op
BenchmarkItoaByStrconv2-10      28513947                40.17 ns/op           96 B/op          0 allocs/op
PASS
ok      example.com/go-benchmark-test   4.132s
```

左から「プログラムを実行した回数」「1回あたりの実行にかかった時間」「1回あたりのallocateで確保したメモリの容量」「1回あたりのallocate回数」が表示されている

※[公式ドキュメント](https://pkg.go.dev/cmd/go#hdr-Description_of_testing_flags)のオプションの説明がある
```txt
-benchmem
    Print memory allocation statistics for benchmarks.
```

### ベンチマーク測定中に使用するCPU数を変更する
```bash
$ go test -bench . -cpu 2

goos: darwin
goarch: arm64
pkg: example.com/go-benchmark-test
BenchmarkItoaByFmt-2            10471420                97.08 ns/op
BenchmarkItoaByStrconv1-2       23852037                49.82 ns/op
BenchmarkItoaByStrconv2-2       23069085                43.41 ns/op
PASS
ok      example.com/go-benchmark-test   3.778s
```

※公式ドキュメントの説明
```txt
-cpu 1,2,4
    Specify a list of GOMAXPROCS values for which the tests, benchmarks or
    fuzz tests should be executed. The default is the current value
    of GOMAXPROCS. -cpu does not apply to fuzz tests matched by -fuzz.
```

### プロファイリングをWebで確認可能
```bash
$ go test -bench . -cpuprofile cpu.out

goos: darwin
goarch: arm64
pkg: example.com/go-benchmark-test
BenchmarkItoaByFmt-10           11547079                92.85 ns/op
BenchmarkItoaByStrconv1-10      31387045                43.58 ns/op
BenchmarkItoaByStrconv2-10      31025254                33.99 ns/op
PASS
ok      example.com/go-benchmark-test   4.070s

$ go tool pprof -http=":8989" cpu.out
Serving web UI on http://localhost:8888
```

localhostにアクセスすることで、プログラムの各ステップでどれくらい時間がかかっているかを確認できる。