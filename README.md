# Golang Signal Function Implementation

Golang implementation of JavaScript signals, similar to those in SolidJs, using generics.

The repository includes two core functions createSignal and createDerievedSignal.

## createSignal function:
The createSignal function initializes a signal and offers methods to get the current value, set a new value, and subscribe to the signal.

```go
func createSignal[T any](value T) (get func() T, set func(T), subscribe func(func(T)))
```

## createDerivedSignal function:
The createDerivedSignal function creates a new signal derived from an existing signal and a derive function, it also includes a function to subscribe to the new signal.

```go
func createDerivedSignal[T any, U any](signalSubscribe func(func(T)), derive func() U) (get func() U, subscribe func(func(U)))
```

## Usage

Here are some example use cases:

```go
func main() {
    cats, setCats, _ := createSignal([]Cat{
        {
            ID:   "1",
            Name: "Kitty",
        },
        {
            ID:   "2",
            Name: "Whiskers",
        },
    });

    for _, cat := range cats() {
        fmt.Println(cat);
    }

    setCats(append(cats(), Cat{
        ID:   "3",
        Name: "Fluffy",
    }));

    for _, cat := range cats() {
        fmt.Println(cat);
    }

    count, setCount, subscribeCount := createSignal(0);
    doubleCount, subscribeDoubleCount := createDerivedSignal(subscribeCount, func() int {
        return count() * 2;
    });

    subscribeCount(func(value int) {
        fmt.Println("count changed to", value);
    });

    subscribeDoubleCount(func(value int) {
        fmt.Println("double count changed to", value);
    });

    fmt.Println(doubleCount());
    setCount(1);
    fmt.Println(doubleCount());
    setCount(2);
    fmt.Println(doubleCount());

}
```

I don't know why would you need to use this in real golang scenarios, but now you can.

### Does not support
- Unsubscribing from signals
- Lazy evaluation of derived signals