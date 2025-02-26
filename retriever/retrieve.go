package retriever

type Retriever interface {
	Retrieve(string) string
}
