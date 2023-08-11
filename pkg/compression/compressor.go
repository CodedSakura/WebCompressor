package compression

import "time"

type Compressor interface {
	Mimetype() string
	Extension() string
	Compress(targetPath string) State
}

type State struct {
	Path         string
	Progress     float32
	CreatedTime  time.Time
	FinishedTime time.Time
}

func (s State) IsDone() bool {
	return s.Progress >= 1
}
func newState(c Compressor, path string) State {
	return State{
		Path:        path + "." + c.Extension(),
		Progress:    0,
		CreatedTime: time.Now(),
	}
}
