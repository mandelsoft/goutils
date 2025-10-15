package pool

import "reflect"

// DoneFor provides a done channel wgich is closed
// when one of any number of other done channels is closed.
func DoneFor(channels ...<-chan struct{}) <-chan struct{} {
	out := make(chan struct{})

	go func() {
		defer close(out) // close output once any channel is done
		cases := make([]reflect.SelectCase, len(channels))
		for i, ch := range channels {
			cases[i] = reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(ch)}
		}

		reflect.Select(cases) // blocks until one channel is closed
		// once any channel is closed, we reach here and close(out)
	}()

	return out
}
