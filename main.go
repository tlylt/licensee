package main

import (
	"context"
	"fmt"
	"runtime"

	"github.com/tlylt/licensee/retriever"
	"golang.org/x/sync/semaphore"
)

var (
	maxWorkders = runtime.GOMAXPROCS(0)
	sem         = semaphore.NewWeighted(int64(maxWorkders))
	urls        = []string{
		"https://raw.githubusercontent.com/gohugoio/hugo/master/LICENSE",
		"https://raw.githubusercontent.com/moby/moby/master/LICENSE",
		"https://raw.githubusercontent.com/junegunn/fzf/master/LICENSE",
		"https://raw.githubusercontent.com/google/go-licenses/master/LICENSE",
	}
	ctx = context.TODO()
)

func main() {
	urlRetriever := retriever.NewURLRetriever()

	for idx, url := range urls {
		if err := sem.Acquire(ctx, 1); err != nil {
			fmt.Println("Failed to acquire semaphore:", err)
			break
		}
		go func(idx int, url string) {
			defer sem.Release(1)
			content, err := urlRetriever.Retrieve(url)
			if err != nil {
				fmt.Println("Error retrieving the LICENSE file:", err)
				return
			}
			// print the first line of the content
			fmt.Println(idx)
			content = string(content)[:100]
			fmt.Println(content)
		}(idx, url)
	}
	if err := sem.Acquire(ctx, int64(maxWorkders)); err != nil {
		fmt.Println("Failed to acquire semaphore:", err)
	}
	fmt.Println("Done")
}
