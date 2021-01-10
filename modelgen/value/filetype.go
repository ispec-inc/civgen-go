package value

const (
	FiletypeEntity     = Filetype("entity")
	FiletypeModel      = Filetype("model")
	FiletypeView       = Filetype("view")
	FiletypeRepository = Filetype("repository")
	FiletypeDao        = Filetype("dao")
	FiletypeDaoTest    = Filetype("dao_test")
)

type Filetype string

func (t Filetype) String() string {
	return string(t)
}
