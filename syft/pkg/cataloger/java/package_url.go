package java

import (
	"github.com/anchore/packageurl-go"
	"github.com/anchore/syft/syft/pkg"
	"github.com/anchore/syft/syft/pkg/cataloger/common/cpe"
)

// PackageURL returns the PURL for the specific java package (see https://github.com/package-url/purl-spec)
func packageURL(name, version string, metadata pkg.JavaMetadata) string {
	var groupID = name
	if gID := cpe.GroupIDFromJavaMetadata(name, metadata); gID != "" {
		groupID = gID
	}

	pURL := packageurl.NewPackageURL(
		packageurl.TypeMaven, // TODO: should we filter down by package types here?
		groupID,
		name,
		version,
		nil, // TODO: there are probably several qualifiers that can be specified here
		"")
	return pURL.ToString()
}
