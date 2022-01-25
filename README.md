# go-crypto

根据基础加解密算法包做出的一定性能优化

## 优化方向

- 减少了字符串与字节数组转化的内存损耗
- 对加解密过程中申请的字节空间使用`sync.Pool`进行管理

## Benchmark 基准测试

基准测试命令：`go test -bench=. -benchtime=10s -benchmem -cpu=4 -timeout=100m`

```text
--- dataSize: 10
goos: darwin
goarch: amd64
pkg: github.com/frank-yf/go-crypto
cpu: Intel(R) Core(TM) i5-1038NG7 CPU @ 2.00GHz
BenchmarkAesECBCrypto-4                          7462260              1608 ns/op            1112 B/op         40 allocs/op
BenchmarkAesECBCrypto_WithPool-4                 4858204              2447 ns/op            1496 B/op         53 allocs/op
BenchmarkAesECBCrypto_WithBase64-4               3404462              3520 ns/op            1912 B/op         60 allocs/op
BenchmarkAesECBCrypto_WithBase64_WithPool-4      2129854              5618 ns/op            2856 B/op         99 allocs/op
BenchmarkAesECBCrypto_WithHex-4                  3260694              3668 ns/op            2264 B/op         60 allocs/op
BenchmarkAesECBCrypto_WithHex_WithPool-4         1979689              6013 ns/op            3208 B/op         99 allocs/op
BenchmarkNone-4                                  1690008              7094 ns/op           10072 B/op        120 allocs/op
PASS
ok      github.com/frank-yf/go-crypto   114.702s

--- dataSize: 10
goos: darwin
goarch: amd64
pkg: github.com/frank-yf/go-crypto
cpu: Intel(R) Core(TM) i5-1038NG7 CPU @ 2.00GHz
BenchmarkAesECBCrypto-4                           211598             58158 ns/op           84792 B/op        400 allocs/op
BenchmarkAesECBCrypto_WithPool-4                  172875             69764 ns/op           88818 B/op        525 allocs/op
BenchmarkAesECBCrypto_WithBase64-4                 91771            130877 ns/op          135544 B/op        600 allocs/op
BenchmarkAesECBCrypto_WithBase64_WithPool-4        77180            156315 ns/op          127207 B/op        999 allocs/op
BenchmarkAesECBCrypto_WithHex-4                    73759            165117 ns/op          169784 B/op        600 allocs/op
BenchmarkAesECBCrypto_WithHex_WithPool-4           62107            192252 ns/op          161457 B/op        999 allocs/op
BenchmarkNone-4                                   111423            110126 ns/op          174392 B/op       1200 allocs/op
PASS
ok      github.com/frank-yf/go-crypto   94.274s

--- dataSize: 10
goos: darwin
goarch: amd64
pkg: github.com/frank-yf/go-crypto
cpu: Intel(R) Core(TM) i5-1038NG7 CPU @ 2.00GHz
BenchmarkAesECBCrypto-4                             2990           4037089 ns/op         7070559 B/op       4000 allocs/op
BenchmarkAesECBCrypto_WithPool-4                    2905           4166075 ns/op         7113762 B/op       5260 allocs/op
BenchmarkAesECBCrypto_WithBase64-4                  1269           9184116 ns/op        12067350 B/op       6001 allocs/op
BenchmarkAesECBCrypto_WithBase64_WithPool-4         1285           9496406 ns/op        10070904 B/op      10016 allocs/op
BenchmarkAesECBCrypto_WithHex-4                      969          12509609 ns/op        15620470 B/op       6002 allocs/op
BenchmarkAesECBCrypto_WithHex_WithPool-4             952          12481948 ns/op        13632622 B/op      10024 allocs/op
BenchmarkNone-4                                     2650           4529414 ns/op         7966664 B/op      12000 allocs/op
PASS
ok      github.com/frank-yf/go-crypto   90.392s

```

- 线性增加数据长度以及元素数量后，基准结果也呈线性增长；
- 当数据长度超过一定范围后，`hex/base64`等编码函数更耗资源；

---

> 欢迎提出意见或提交 PR。
