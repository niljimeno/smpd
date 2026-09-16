package radio

import (
	"fmt"
	"smpd/radio/stereo"
)

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

func trackSongEnd(ev *chan uint8) {
	for {
		<-*ev
		fmt.Println("Song ended")
	}
}

func NewRadio() Radio {
	s := stereo.NewStereo()
	go trackSongEnd(s.Events)
	return Radio{stereo: s}
}
