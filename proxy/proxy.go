package main

type server interface {
	handleRequest(string, string) (int, string)
}

type app struct{}

func (a *app) handleRequest(url, method string) (int, string) {
	if url == stateURL && method == "GET" {
		return 200, "Done"
	}

	if url == createURL && method == "POST" {
		return 201, "Created"
	}
	return 404, "Not Ok"
}

type proxy struct {
	maxRequest int
	limiter    map[string]int
	app        *app
}

func newProxy() *proxy {
	return &proxy{
		app:        &app{},
		maxRequest: 2,
		limiter:    make(map[string]int),
	}
}

func (n *proxy) handleRequest(url, method string) (int, string) {
	if !n.checkRateLimiting(url) {
		return 403, "Not Allowed"
	}
	return n.app.handleRequest(url, method)
}

func (n *proxy) checkRateLimiting(url string) bool {
	if n.limiter[url] == 0 {
		n.limiter[url] = 1
	}
	if n.limiter[url] > n.maxRequest {
		return false
	}
	n.limiter[url]++
	return true
}
