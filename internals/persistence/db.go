package persistence

import (
	"lottoengine/model"
)

type Handler interface {
	PersistGameLog(model.GamePlayLog) (string, error)
	GetGameLog(model.GamePlayLog) ([]model.GamePlayLog, error)
}
