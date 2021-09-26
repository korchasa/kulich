package posix

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

func Download(uri *url.URL, out *os.File) (int64, string, error) {
	resp, err := http.Get(uri.String())
	if err != nil {
		return 0, "", fmt.Errorf("can't get uri `%s`: %v", uri.String(), err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, "", fmt.Errorf("bad status `%s` for `%s`", resp.Status, uri.String())
	}

	wb, err := io.Copy(out, resp.Body)

	return wb, out.Name(), err
}
