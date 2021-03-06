package fixtures

import (
	"path/filepath"
	"testing"

	"github.com/bblfsh/go-driver/driver/golang"
	"github.com/bblfsh/go-driver/driver/normalizer"
	"github.com/bblfsh/sdk/v3/driver"
	"github.com/bblfsh/sdk/v3/driver/fixtures"
)

const projectRoot = "../../"

var Suite = &fixtures.Suite{
	Lang: "go",
	Ext:  ".go",
	Path: filepath.Join(projectRoot, fixtures.Dir),
	NewDriver: func() driver.Native {
		return golang.NewDriver()
	},
	Transforms: normalizer.Transforms,
	BenchName:  "json",
	Semantic: fixtures.SemanticConfig{
		BlacklistTypes: []string{
			"Ident",
			"Comment",
			"BlockStmt",
			"ImportSpec",
			"FuncDecl",
			"FuncType",
		},
	},
}

func TestGoDriver(t *testing.T) {
	Suite.RunTests(t)
}

func BenchmarkGoDriver(b *testing.B) {
	Suite.RunBenchmarks(b)
}
