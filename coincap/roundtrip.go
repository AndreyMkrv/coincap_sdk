package coincap

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

const URL string = "https://api.coincap.io/v2/assets/"

type LoggingRoundTripper struct {
	logger io.Writer
	next   http.RoundTripper
}

func (l LoggingRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	fmt.Fprintf(l.logger, "[%s] %s %s\n", time.Now().Format(time.ANSIC), r.Method, r.URL)
	time.Sleep(time.Second * 2)
	return l.next.RoundTrip(r)
}
