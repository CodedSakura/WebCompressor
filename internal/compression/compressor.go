package compression

import (
	"WebCompressor/internal/utils"
	"github.com/google/uuid"
	"time"
)

type Compressor interface {
	Mimetype() string
	Extension() string
	Compress(targetPath string) State
}
type compressorBase struct {
	utils *utils.Utils
}

type State struct {
	Id           uuid.UUID
	Path         string
	Progress     float32
	CreatedTime  time.Time
	FinishedTime time.Time
}

func (s State) IsDone() bool {
	return s.Progress >= 1
}
func newState(c Compressor) State {
	id := uuid.New()
	return State{
		Id:          id,
		Path:        id.String() + "." + c.Extension(),
		Progress:    0,
		CreatedTime: time.Now(),
	}
}
