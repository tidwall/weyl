Middle Square Weyl Sequence RNG for Go
======================================
<a href="https://godoc.org/github.com/tidwall/weyl"><img src="https://img.shields.io/badge/api-reference-blue.svg?style=flat-square" alt="GoDoc"></a>

A new implementation of John von Neumann's middle square random number generator (RNG).
A Weyl sequence is utilized to keep the generator running through a long period.

[Paper](https://arxiv.org/pdf/1704.00358.pdf)

Using
-----

To start using Weyl, install Go and run `go get`:

```
$ go get -u github.com/tidwall/weyl
```

Create a new RNG

```go
rng := rand.New(weyl.NewSource(time.Now().UnixNano()))
```

Use the RNG just as you would use the standard [math/rand](https://golang.org/pkg/math/rand/).

```go
rng.Int63()  // random int64
rng.Float64() // random float64
```


Performance
-----------

Weyl RNG:

```
BenchmarkInt63Threadsafe-8     	100000000	        16.3 ns/op
BenchmarkInt63Unthreadsafe-8   	300000000	         5.32 ns/op
BenchmarkIntn1000-8            	100000000	        16.0 ns/op
BenchmarkInt63n1000-8          	50000000	        25.7 ns/op
BenchmarkInt31n1000-8          	100000000	        14.6 ns/op
BenchmarkFloat32-8             	100000000	        11.9 ns/op
BenchmarkFloat64-8             	200000000	         8.44 ns/op
BenchmarkPerm3-8               	20000000	        78.0 ns/op
BenchmarkPerm30-8              	 2000000	       600 ns/op
BenchmarkRead3-8               	100000000	        14.8 ns/op
BenchmarkRead64-8              	10000000	       135 ns/op
BenchmarkRead1000-8            	 1000000	      1801 ns/op
```

Go standard [math/rand](https://golang.org/pkg/math/rand/):

```
BenchmarkInt63Threadsafe-8     	100000000	        21.4 ns/op
BenchmarkInt63Unthreadsafe-8   	200000000	         6.74 ns/op
BenchmarkIntn1000-8            	100000000	        17.0 ns/op
BenchmarkInt63n1000-8          	50000000	        26.4 ns/op
BenchmarkInt31n1000-8          	100000000	        14.4 ns/op
BenchmarkFloat32-8             	100000000	        12.9 ns/op
BenchmarkFloat64-8             	200000000	         9.77 ns/op
BenchmarkPerm3-8               	20000000	        80.6 ns/op
BenchmarkPerm30-8              	 2000000	       658 ns/op
BenchmarkRead3-8               	100000000	        14.6 ns/op
BenchmarkRead64-8              	10000000	       147 ns/op
BenchmarkRead1000-8            	 1000000	      2112 ns/op
```

Contact
-------
Josh Baker [@tidwall](http://twitter.com/tidwall)

License
-------
Weyl source code is available under the ISC [License](/LICENSE).

