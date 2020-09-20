package modlib

import (
	"path/filepath"
	"strings"
	"os"
	"time"

	. "github.com/NjinN/RML/go/core"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/wav"
	"github.com/faiface/beep/speaker"
)


func AudioPlay(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos():es.LastEndPos()]

	if args[1].Tp == FILE {

		filePath, err := filepath.Abs(strings.ReplaceAll(args[1].Str(), `"`, ``))
		if err != nil {
			return &Token{ERR, err.Error()}, nil
		}

		return PlayAudioFromFile(filePath, false)
		
	}

	return &Token{ERR, "Type Mismatch"}, nil
}

func AudioPlayLoop(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos():es.LastEndPos()]

	if args[1].Tp == FILE {

		filePath, err := filepath.Abs(strings.ReplaceAll(args[1].Str(), `"`, ``))
		if err != nil {
			return &Token{ERR, err.Error()}, nil
		}

		return PlayAudioFromFile(filePath, true)
		
	}

	return &Token{ERR, "Type Mismatch"}, nil
}


func PlayAudioFromFile(filePath string, loop bool) (*Token, error){
	if len(filePath) <= 4 || (filePath[len(filePath)-4:] != ".mp3" && filePath[len(filePath)-4:] != ".wav") {
		return &Token{ERR, "Error file path of " + filePath}, nil
	}

	f, err := os.Open(filePath)
	if err != nil {
		return &Token{ERR, err.Error()}, nil
	}

	var streamer beep.StreamSeekCloser
	var format beep.Format
	if filePath[len(filePath)-4:] == ".mp3" {
		streamer, format, err = mp3.Decode(f)
	}else{
		streamer, format, err = wav.Decode(f)
	}

	if err != nil {
		return &Token{ERR, err.Error()}, nil
	}
	
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	go goPlayAudio(streamer, loop)

	return &Token{LOGIC, true}, nil
}


func goPlayAudio(streamer beep.StreamSeekCloser, loop bool){
	defer streamer.Close()

	for {
		streamer.Seek(0)
		done := make(chan bool)
		speaker.Play(beep.Seq(streamer, beep.Callback(func() {
			done <- true
		})))

		select {
		case <- done:
			if loop {
				break
			}else{
				return
			}
		}
	}
	
}


func AudioStop(es *EvalStack, ctx *BindMap) (*Token, error) {
	
	speaker.Close()
	return &Token{LOGIC, true}, nil
}