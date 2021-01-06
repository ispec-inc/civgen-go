package pkg

import (
	"errors"
	"fmt"
	"strings"
)

const (
	pkgSep = "/"
)

type pkg string

func (p pkg) Path() string {
	return string(p)
}

func (p pkg) Name() string {
	if string(p) == "" {
		return ""
	}
	ps := strings.Split(string(p), pkgSep)
	return ps[len(ps)-1]
}

func newpkg(root, path string) (pkg, error) {
	switch {
	case path == "":
		return "", nil
	case path[:2] == "./":
		return "", errors.New("path should start from directory name not '.'")
	default:
		return pkg(fmt.Sprintf("%s/%s", root, path)), nil
	}
}
