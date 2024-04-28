package signals

// golang implementation of javascript signals (like in solidjs) implemented with generics
func createSignal[T any](value T) (get func() T, set func(T), subscribe func(func(T))) {
	var value_ T
	var subscribers []func(T)

	get = func() T {
		return value_
	}
	set = func(value T) {
		value_ = value

		// notify subscribers
		for _, subscriber := range subscribers {
			subscriber(value)
		}
	}
	set(value)

	subscribe = func(subscriber func(T)) {
		subscribers = append(subscribers, subscriber)
	}
	return
}

// create a function that receives a signal and a function that returns a new value based on the signal - including a subscription to the new signal
func createDerivedSignal[T any, U any](signalSubscribe func(func(T)), derive func() U) (get func() U, subscribe func(func(U))) {
	var value U
	var subscribers []func(U)

	get = func() U {
		return value
	}
	set := func() U {
		value = derive()

		// notify subscribers
		for _, subscriber := range subscribers {
			subscriber(value)
		}

		return value
	}
	set()

	subscribe = func(subscriber func(U)) {
		subscribers = append(subscribers, subscriber)
	}

	signalSubscribe(func(T) {
		set()
	})

	return
}
