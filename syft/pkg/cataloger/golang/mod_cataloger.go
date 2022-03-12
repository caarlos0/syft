/*
Package golang provides a concrete Cataloger implementation for go.mod files.
*/
package golang

import (
	"github.com/anchore/syft/syft/pkg/cataloger/generic"
)

// NewGoModFileCataloger returns a new Go module cataloger object.
func NewGoModFileCataloger() *generic.Cataloger {
	globParsers := map[string]generic.Parser{
		"**/go.mod": parseGoMod,
	}

	return generic.NewCataloger(nil, globParsers, "go-mod-file-cataloger")
}
