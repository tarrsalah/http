package mm

import (
	"net/http"
	"testing"
)

func dummyMiddleware() Middleware {
	return func(h http.Handler) http.Handler {
		return h
	}
}

var tracking []string

func dummyTrackedHttpHandler() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		tracking = append(tracking, "handler")
	})
}

func dummyTrackedMiddleware(identifier string) Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			tracking = append(tracking, "before-"+identifier)
			h.ServeHTTP(writer, request)
			tracking = append(tracking, "after-"+identifier)
		})
	}
}

// No test really needed, I run this code instead of writiing main
func TestNew(t *testing.T) {
	var c Chain

	c = New()
	AssertLog(t, len(c) == 0, "Empty chain is not if length zero")

	c = New(dummyMiddleware())
	AssertLogf(t, len(c) == 1, "Chain with 1 mware was of length %d not 1", len(c))
}

func TestChainAppend(t *testing.T) {
	var c Chain

	c = New().Append(dummyMiddleware())
	AssertLogf(t, len(c) == 1, "Append to empty Chain gave len of %d not 1", len(c))
	AssertLog(t, c[0] != nil, "First middleware in chain exists")

	c = New().Append(dummyMiddleware(), dummyMiddleware()).Append(dummyMiddleware())
	AssertLogf(t, len(c) == 3, "Append two times to empty Chain gave len of %d not 3", len(c))
}

func TestChainThen(t *testing.T) {
	tracking = []string{}

	// Create a new middleware chain with 3 tracked middlewares that will log before
	// and after they execute, having in the middle a http handle
	c := New(
		dummyTrackedMiddleware("1"),
		dummyTrackedMiddleware("2"),
		dummyTrackedMiddleware("3"),
	)
	h := c.Then(dummyTrackedHttpHandler())

	h.ServeHTTP(nil, nil)

	AssertLogf(t, tracking[0] == "before-1", "Expected 1st trace to be 'before-1' got '%s'", tracking[0])
	AssertLogf(t, tracking[1] == "before-2", "Expected 2nd trace to be 'before-2' got '%s'", tracking[1])
	AssertLogf(t, tracking[2] == "before-3", "Expected 3rd trace to be 'before-3' got '%s'", tracking[2])
	AssertLogf(t, tracking[3] == "handler", "Expected 4th trace to be 'handler' got '%s'", tracking[3])
	AssertLogf(t, tracking[4] == "after-3", "Expected 5th trace to be 'after-3' got '%s'", tracking[4])
	AssertLogf(t, tracking[5] == "after-2", "Expected 6th trace to be 'after-2' got '%s'", tracking[5])
	AssertLogf(t, tracking[6] == "after-1", "Expected 7th trace to be 'after-1' got '%s'", tracking[6])
}

func AssertLog(t *testing.T, worked bool, mess string) {
	if !worked {
		t.Error(mess)
	}
}

func AssertLogf(t *testing.T, worked bool, mess string, args ...interface{}) {
	if !worked {
		t.Errorf(mess, args...)
	}
}
