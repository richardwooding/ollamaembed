package ollamaembed_test

import (
	"fmt"

	"github.com/richardwooding/ollamaembed"
)

// Example_cosine shows the vector helpers: normalise two vectors, then
// measure their cosine similarity (== dot product on unit vectors).
func Example_cosine() {
	a := []float32{1, 0, 0}
	b := []float32{1, 1, 0}
	ollamaembed.Normalize(a)
	ollamaembed.Normalize(b)
	fmt.Printf("%.3f\n", ollamaembed.Cosine(a, b))
	// Output: 0.707
}

// ExampleBareName strips an Ollama model tag down to its bare name.
func ExampleBareName() {
	fmt.Println(ollamaembed.BareName("nomic-embed-text:latest"))
	// Output: nomic-embed-text
}

// ExampleCatalogLookup looks up a curated model's metadata (here, the
// vector dimensions) without contacting Ollama.
func ExampleCatalogLookup() {
	if e := ollamaembed.CatalogLookup("nomic-embed-text"); e != nil {
		fmt.Println(e.Dimensions)
	}
	// Output: 768
}
