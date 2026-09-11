package stereo

import (
	"os"
	"os/exec"
	"syscall"
)

const (
	Playing = 0
	Paused  = 1
	Stopped = 2
)

const (
	Play = 0
	End  = 1
	Stop = 2
)

type Current struct {
	Name string
	Cmd  *exec.Cmd
}

type Stereo struct {
	State   uint8
	Current *Current
	Events  *chan uint8
}

func NewStereo() Stereo {
	eventsChannel := make(chan uint8)
	return Stereo{
		State:   Stopped,
		Current: &Current{},
		Events:  &eventsChannel,
	}
}

func (s *Stereo) Play(path string) error {
	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	s.State = Playing

	cmd := exec.Command("ffplay", "-nodisp", "-autoexit", "-nostats", path)
	if err := cmd.Start(); err != nil {
		return err
	}

	s.Current.Cmd = cmd
	s.Current.Name = path

	go func() {
		cmd.Wait()
		*s.Events <- End
	}()

	return nil
}

func (s *Stereo) isProcessValid() bool {
	return (s.Current == nil || s.Current.Cmd == nil || s.Current.Cmd.Process == nil)
}

func (s *Stereo) Pause() bool {
	if s.isProcessValid() {
		return false
	}

	err := s.Current.Cmd.Process.Signal(syscall.SIGSTOP)
	if err != nil {
		return false
	}

	s.State = Paused
	return true
}

func (s *Stereo) Resume() bool {
	if s.isProcessValid() {
		return false
	}

	err := s.Current.Cmd.Process.Signal(syscall.SIGCONT)
	if err != nil {
		return false
	}

	s.State = Playing
	return true
}

func (s *Stereo) Stop() bool {
	if s.isProcessValid() {
		return false
	}

	s.Current.Cmd.Process.Kill()
	s.Current.Cmd.Process.Signal(syscall.SIGSTOP)
	s.State = Stopped
	return true
}
