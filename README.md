# ollamaembed

[![Go Reference](https://pkg.go.dev/badge/github.com/richardwooding/ollamaembed.svg)](https://pkg.go.dev/github.com/richardwooding/ollamaembed)
[![CI](https://github.com/richardwooding/ollamaembed/actions/workflows/ci.yml/badge.svg)](https://github.com/richardwooding/ollamaembed/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/richardwooding/ollamaembed)](https://goreportcard.com/report/github.com/richardwooding/ollamaembed)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

A small, **zero-dependency** Go client for text→vector embeddings via [Ollama](https://ollama.com)'s
local HTTP API — plus the vector helpers and curated model catalog that go with it. The
building blocks for semantic / RAG search in Go, with **no SDK, no CGo, no bundled model**.

- **`Embedder`** interface + **`OllamaEmbedder`** — `Embed(ctx, text) ([]float32, error)`,
  context-aware, with model **list / pull** management built in.
- **Vector math** — `Cosine`, `Dot`, `Normalize` (operate on `[]float32`).
- **Curated `Catalog`** — recommended embedding models with dimensions/size, plus
  `CatalogLookup` / `BareName` helpers.
- **stdlib only** (`net/http`, `encoding/json`, `math`), **go 1.21+**.

```sh
go get github.com/richardwooding/ollamaembed
```

> Named `ollamaembed` (not `embed`) so it never collides with the standard library's
> `embed` package.

## Embed text

```go
emb := ollamaembed.NewOllama("http://localhost:11434", "nomic-embed-text")

vec, err := emb.Embed(ctx, "the quick brown fox")
if err != nil {
    log.Fatal(err) // clear errors when Ollama is down or the model isn't pulled
}
ollamaembed.Normalize(vec) // unit vector → cosine == dot
```

Set up Ollama once: install from [ollama.com](https://ollama.com), then
`ollama pull nomic-embed-text` (or `mxbai-embed-large`, `all-minilm`, …).

## Rank by similarity

```go
query := mustEmbed(emb, ctx, "seafaring adventure")
for _, doc := range docs {
    v := mustEmbed(emb, ctx, doc.Text)
    ollamaembed.Normalize(v)
    doc.Score = ollamaembed.Cosine(query, v) // higher = more similar
}
```

(For sub-linear top-K at scale, feed the vectors into your ANN index of choice; this package
is the embedding + vector-math layer.)

## Discover & manage models

```go
// What's installed locally?
models, _ := emb.ListLocal(ctx)

// Curated recommendations (no network):
for _, m := range ollamaembed.Catalog {
    fmt.Printf("%s — %dd, %s\n", m.Name, m.Dimensions, m.Description)
}

// Pull one, streaming progress:
err := emb.Pull(ctx, "nomic-embed-text", func(p ollamaembed.PullProgress) {
    fmt.Printf("\r%s %d/%d", p.Status, p.Completed, p.Total)
})
```

## License

MIT — see [LICENSE](LICENSE).

---

Extracted from [file-search-on](https://github.com/richardwooding/file-search-on), where it
powers the `similarity` ranking and the `search_semantic` tool.
