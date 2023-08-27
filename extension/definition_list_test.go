package extension

import (
	"testing"

	"github.com/yezige/goldmark"
	"github.com/yezige/goldmark/renderer/html"
	"github.com/yezige/goldmark/testutil"
)

func TestDefinitionList(t *testing.T) {
	markdown := goldmark.New(
		goldmark.WithRendererOptions(
			html.WithUnsafe(),
		),
		goldmark.WithExtensions(
			DefinitionList,
		),
	)
	testutil.DoTestCaseFile(markdown, "_test/definition_list.txt", t, testutil.ParseCliCaseArg()...)
}
