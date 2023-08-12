package compression

import (
	"github.com/google/uuid"
	"go.uber.org/fx"
	"path"
	"time"
)

type Compressor interface {
	Mimetype() string
	Extension() string
	Compress(targetPath string) (*State, error)
}

type State struct {
	Id           uuid.UUID
	Path         string
	Progress     float32
	CreatedTime  time.Time
	FinishedTime time.Time
}

func (s State) HasSucceeded() bool {
	return s.Progress >= 1
}
func (s State) HasFailed() bool {
	return s.Progress < 0
}
func (s State) IsDone() bool {
	return s.HasSucceeded() || s.HasFailed()
}
func newState(c Compressor) *State {
	id := uuid.New()
	return &State{
		Id:          id,
		Path:        path.Join("archives", id.String()+"."+c.Extension()),
		Progress:    0,
		CreatedTime: time.Now(),
	}
}

func AsCompressor(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(Compressor)),
		fx.ResultTags(`group:"compressors"`),
	)
}
