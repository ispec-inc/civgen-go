package value

import (
	"fmt"
	"strings"
)

const (
	PackageOS      Package = "os"
	PackageTesting Package = "testing"
	PackageFmt     Package = "fmt"

	PackageGorm    Package = "gorm.io/gorm"
	PackageAssert  Package = "github.com/stretchr/testify/assert"
	PackageSqlfile Package = "github.com/tanimutomo/sqlfile"
)

var (
	PackageEntity     Package = ""
	PackageModel      Package = ""
	PackageView       Package = ""
	PackageRepository Package = ""
	PackageDao        Package = ""
	PackageError      Package = ""
	PackageDatabase   Package = ""
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
