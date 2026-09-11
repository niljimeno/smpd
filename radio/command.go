package radio

import ()

func (r *Radio) Execute(command []string) string {
	args := command[1:]

	switch command[0] {
	case "pause":
		return simpleOp(r.stereo.Pause)
	case "resume":
		return simpleOp(r.stereo.Resume)
	case "stop":
		return simpleOp(r.stereo.Stop)
	case "play":
		return r.play(args)
	}

	return ""
}

func simpleOp(input func() bool) string {
	success := input()
	if success {
		return "ok"
	} else {
		return "err"
	}
}

func (r *Radio) play(args []string) string {
	if len(args) < 1 {
		return "err"
	}

	songName := args[0]
	if songName == "venture.mp3" || songName == "venture" {
		r.stereo.Play("./venture.mp3")
		return "ok"
	}

	return "not found"
}
