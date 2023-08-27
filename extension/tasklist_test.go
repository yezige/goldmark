package extension

import (
	"testing"

	"github.com/yezige/goldmark"
	"github.com/yezige/goldmark/renderer/html"
	"github.com/yezige/goldmark/testutil"
)

func TestTaskList(t *testing.T) {
	markdown := goldmark.New(
		goldmark.WithRendererOptions(
			html.WithUnsafe(),
		),
		goldmark.WithExtensions(
			TaskList,
		),
	)
	testutil.DoTestCaseFile(markdown, "_test/tasklist.txt", t, testutil.ParseCliCaseArg()...)
}
