# About

Package [math/rand][mathrand] provides support for generating pseudo-random numbers.

Here is how to generate a random integer between `0` and `99`:

```go
import "math/rand"

n := rand.Intn(100) // n is a random int, 0 <= n < 100
```

Function `rand.Float64` returns a random floating point number between `0.0` and `1.0`:

```go
f := rand.Float64() // f is a random float64, 0.0 <= f < 1.0
```

There is also support for shuffling a slice (or other data structures):

```go
x := []string{"a", "b", "c", "d", "e"}
// shuffling the slice put its elements into a random order
rand.Shuffle(len(x), func(i, j int) {
	x[i], x[j] = x[j], x[i]
})
```

## Seeds

The number sequences generated by package `math/rand` are not truly random.
Given a specific "seed" value, the results are entirely deterministic.

In Go 1.20+ the seed is automatically picked at random so you will see a different sequences of random numbers each time you run your program. 

In prior versions of Go, the seed was `1` by default.
So to get different sequences for various runs of the program, you had to manually seed the random number generator, e.g. with the current time, before retrieving any random numbers.

```go
rand.Seed(time.Now().UnixNano())
```

[mathrand]: https://pkg.go.dev/math/rand