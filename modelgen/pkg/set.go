package pkg

import "fmt"

type SetPkgsInput struct {
	ProjectRoot string
	EntityPath  string
	ModelPath   string
	ViewPath    string
}

func SetPkgs(ipt SetPkgsInput) error {
	ep, err := newpkg(ipt.ProjectRoot, ipt.EntityPath)
	if err != nil {
		return fmt.Errorf("Invalid entity_path: %v\n", err)
	}
	mp, err := newpkg(ipt.ProjectRoot, ipt.ModelPath)
	if err != nil {
		return fmt.Errorf("Invalid model_path: %v\n", err)
	}
	vp, err := newpkg(ipt.ProjectRoot, ipt.ViewPath)
	if err != nil {
		return fmt.Errorf("Invalid view_path: %v\n", err)
	}

	Pkgs = pkgs{
		Entity: ep,
		Model:  mp,
		View:   vp,
	}
	return nil
}
