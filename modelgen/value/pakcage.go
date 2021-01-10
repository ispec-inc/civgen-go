package value

import (
	"fmt"
	"strings"
)

const (
	PackageGorm    Package = "gorm.io/gorm"
	PackageTesting Package = "testing"
	PackageAssert  Package = "github.com/stretchr/testify/assert"
)

var (
	PackageEntity     Package = ""
	PackageModel      Package = ""
	PackageView       Package = ""
	PackageRepository Package = ""
	PackageDao        Package = ""
	PackageError      Package = ""
)

const (
	pkgSep = "/"
)

type Package string

func (p Package) Path() string {
	return string(p)
}

func (p Package) Name() string {
	if string(p) == "" {
		return ""
	}
	ps := strings.Split(string(p), pkgSep)
	return ps[len(ps)-1]
}

func NewLocalPackage(root, path string) Package {
	return Package(fmt.Sprintf("%s/%s", root, path))
}
