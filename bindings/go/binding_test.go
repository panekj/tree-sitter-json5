package tree_sitter_json5_test

import (
	"testing"

	tree_sitter "github.com/smacker/go-tree-sitter"
	"github.com/tree-sitter/tree-sitter-json5"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_json5.Language())
	if language == nil {
		t.Errorf("Error loading Json5 grammar")
	}
}
