package main

import (
)

func main() {

}

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}

	// NOTE: It's tempting to use a second deferred call, to f.Close, to close
	// the local file, but this would be subtly wrong because os.Create opens a
	// file for writing, creating it as needed. On many file systems, notably NFS,
	// write errors are not reported immediately but may be postpond until the file
	// is closed. Failure to check the result of the close operation could cause
	// serious data loss to go unnoticed. However, if both io.Copy and f.Close
	// fail, we should prefer to report the error from io.Copy since it occurred
	// first and is more likely to tell use the root cause.
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	n, err = io.Copy(f, resp.Body)
	// Close file, but prefer error from Copy, if any.
	if closeErr := f.Close(); err == nil {
		err = closeErr
	}
	return local, n, err
}