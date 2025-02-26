package retriever

import (
	"fmt"
	"io"
	"net/http"
)

type URLRetriever struct{}

func NewURLRetriever() *URLRetriever {
	return &URLRetriever{}
}

func (u *URLRetriever) Retrieve(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching the LICENSE file:", err)
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Failed to fetch file, status code:", resp.StatusCode)
		return "", err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return "", err
	}

	return string(body), nil
}
