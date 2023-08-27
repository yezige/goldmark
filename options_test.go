package goldmark_test

import (
	"testing"

	. "github.com/yezige/goldmark"
	"github.com/yezige/goldmark/parser"
	"github.com/yezige/goldmark/testutil"
)

func TestAttributeAndAutoHeadingID(t *testing.T) {
	markdown := New(
		WithParserOptions(
			parser.WithAttribute(),
			parser.WithAutoHeadingID(),
		),
	)
	testutil.DoTestCaseFile(markdown, "_test/options.txt", t, testutil.ParseCliCaseArg()...)
}
