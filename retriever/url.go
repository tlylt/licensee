package retriever

import (
	"fmt"
	"io"
	"net/http"

	"github.com/tlylt/licensee/analyzer"
	license "github.com/tlylt/licensee/domain"
)

type URLRetriever struct{}

func NewURLRetriever() *URLRetriever {
	return &URLRetriever{}
}

func (u *URLRetriever) Retrieve(url string) (*license.License, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching the LICENSE file:", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Failed to fetch file, status code:", resp.StatusCode)
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, err
	}
	rawText := string(body)

	licenseType := analyzer.LicenseType(rawText)

	license := license.NewLicense(url, licenseType, rawText)

	return license, nil
}
