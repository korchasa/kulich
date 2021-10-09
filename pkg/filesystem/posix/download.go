package posix

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

func (fs *Posix) download(out *os.File, uri *url.URL) (int64, error) {
	httpClient := &http.Client{
		Timeout: time.Minute,
	}

	req, err := http.NewRequestWithContext(context.TODO(), "GET", uri.String(), nil)
	if err != nil {
		return 0, fmt.Errorf("can't create request: %w", err)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return 0, fmt.Errorf("can't get uri `%s`: %w", uri.String(), err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("bad status `%s` for `%s`", resp.Status, uri.String())
	}

	wb, err := io.Copy(out, resp.Body)

	return wb, err
}
