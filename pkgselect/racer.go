package pkgselect

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func Racer(a, b string) (winner string, err error) {
	ctxWithTimeout, ctxCancelFunc := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer ctxCancelFunc()
	select {
	case <-ping(ctxWithTimeout, a):
		return a, nil
	case <-ping(ctxWithTimeout, b):
		return b, nil
	case <-ctxWithTimeout.Done():
		return "", fmt.Errorf("timed out waiting for %s and %s ", a, b)
	}
}

func ping(ctx context.Context, url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		var body io.Reader
		req, _ := http.NewRequestWithContext(ctx, "GET", url, body)
		res, err := http.DefaultClient.Do(req)
		if err == nil {
			res.Body.Close()
		}
		close(ch)
	}()
	return ch
}
