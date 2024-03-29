package chan_unique

import (
	array_unique "github.com/FosenTS/unique-types/array-unique"
)

func ConvertToUniqueClosedChan[T comparable](in chan T) chan T {
	var a array_unique.UniqueArray[T]
	var el T
	var ok bool
	for {
		select {
		case el, ok = <-in:
		}
		if !ok {
			break
		}
		a.Add(el)
	}
	out := make(chan T)
	for _, el := range a.Elements() {
		out <- el
	}
	close(out)

	return out
}
