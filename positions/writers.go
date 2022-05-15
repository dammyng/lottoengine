package positions

import (
	"lottoengine/model"
)

type GameWriters interface {
	LogGame(model.GamePlayLog) (string, error)
	ArchiveGame(model.GamePlayLog) error
}
