package utils

import (
	"fmt"
	"io"
	"net/http"
)

// ดึงข้อมูลเนื้อผ่าน api
func RequestExternalApi(url string) (io.ReadCloser, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return resp.Body, nil
}
