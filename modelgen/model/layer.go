package model

const (
	LayerEntity Layer = "entity"
	LayerModel  Layer = "model"
	LayerView   Layer = "view"
)

type Layer string

func (t Layer) String() string {
	return string(t)
}
