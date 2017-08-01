package simra

import (
	"io"

	mp3 "github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto"
	"golang.org/x/mobile/asset"
)

type audio struct {
	isClosed chan bool
}

// Audioer is an interface for treating audio
type Audioer interface {
	Play(resource asset.File, loop bool, doneCallback func(err error)) error
	Stop() error
}

// NewAudio returns new audio instance that implements Audioer interface
func NewAudio() Audioer {
	return &audio{
		isClosed: make(chan bool, 1),
	}
}

func (a *audio) Play(resource asset.File, loop bool, doneCallback func(err error)) error {
	dec, err := mp3.NewDecoder(resource)
	if err != nil {
		return err
	}

	player, err := oto.NewPlayer(dec.SampleRate(), 2, 2, 8192)
	if err != nil {
		return err
	}

	doneChan := make(chan error)
	go func() {
		doneChan <- a.doPlay(player, dec, loop)
	}()

	go func() {
		defer func() {
			err := dec.Close()
			if err != nil {
				LogError(err.Error())
			}
			err = player.Close()
			if err != nil {
				LogError(err.Error())
			}
		}()

		<-doneChan
		doneCallback(err)
	}()

	return nil
}

func (a *audio) doPlay(player *oto.Player, r io.ReadSeeker, loop bool) error {
	var offset, written int64
	var err error
	readByte := (int64)(4096 * 30)
loop:
	for {
	playback:
		for {
			r.Seek(offset, io.SeekStart)
			select {
			case <-a.isClosed:
				readByte = 0
				r.Seek(0, io.SeekEnd)
			default:
				// for non-blocking
			}
			written, err = io.CopyN(player, r, readByte)
			if err != nil || written == 0 {
				// error or EOF
				break playback
			}
			offset += written
		}
		if !loop || err != io.EOF {
			break loop
		}
		offset = 0
	}
	return err
}

func (a *audio) Stop() error {
	a.isClosed <- true
	return nil
}
