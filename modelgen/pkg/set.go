package pkg

import "fmt"

type SetPkgsInput struct {
	ProjectRoot    string
	EntityPath     string
	ModelPath      string
	ViewPath       string
	RepositoryPath string
	DaoPath        string
	ErrorPath      string
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
	rp, err := newpkg(ipt.ProjectRoot, ipt.RepositoryPath)
	if err != nil {
		return fmt.Errorf("Invalid repository_path: %v\n", err)
	}
	dp, err := newpkg(ipt.ProjectRoot, ipt.DaoPath)
	if err != nil {
		return fmt.Errorf("Invalid dao_path: %v\n", err)
	}
	erp, err := newpkg(ipt.ProjectRoot, ipt.ErrorPath)
	if err != nil {
		return fmt.Errorf("Invalid error_path: %v\n", err)
	}

	Pkgs = pkgs{
		Entity:     ep,
		Model:      mp,
		View:       vp,
		Repository: rp,
		Dao:        dp,
		Error:      erp,
	}
	return nil
}
