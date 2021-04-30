package pkg

// FanOut reads all values from a given input channel and sends each value to a resulting channel
// The FAN-OUT pattern states that multiple invocation of this function
// will generate multiple go routines to read from the same channel till the input channel is closed
// This allows for better work distribution
func FanOut(input <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for v := range input {
			out <- v
		}
		close(out)
	}()
	return out
}
