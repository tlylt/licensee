package retriever

import license "github.com/tlylt/licensee/domain"

type Retriever interface {
	Retrieve(url string) *license.License
}
