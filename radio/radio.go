package radio

import "smpd/radio/stereo"

type Track struct {
	name string
}

type Playlist struct {
	items   []*Track
	current *Track
}

type State struct {
}

type Radio struct {
	playlist Playlist
	state    State
	stereo   stereo.Stereo
}

func NewRadio() Radio {
	s := stereo.NewStereo()
	return Radio{stereo: s}
}
