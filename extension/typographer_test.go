package extension

import (
	"testing"

	"github.com/yezige/goldmark"
	"github.com/yezige/goldmark/renderer/html"
	"github.com/yezige/goldmark/testutil"
)

func TestTypographer(t *testing.T) {
	markdown := goldmark.New(
		goldmark.WithRendererOptions(
			html.WithUnsafe(),
		),
		goldmark.WithExtensions(
			Typographer,
		),
	)
	testutil.DoTestCaseFile(markdown, "_test/typographer.txt", t, testutil.ParseCliCaseArg()...)
}
