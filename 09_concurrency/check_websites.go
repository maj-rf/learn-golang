package concurrency

/*
- Concurrency: "having more than one thing in progress"
- If we wait for a function to return something, it's a "blocking" operation
- goroutine: an operation that doesn't block in go and will run in a
  separate process
- go test -race : spot race conditions with built-in race detector
- channels, to help organize and control the communication between the different
  processes, allowing us to avoid a race condition bug.
- resultChannel <- result{u, wc(u)} // send statement
- r := <-resultChannel // receive expression
*/

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)
	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.string] = r.bool
	}
	return results
}
