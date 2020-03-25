package main

import "github.com/headfirstgo/gadget"

type Player interface {
        Play(string)
        Stop()
}

func playList(device Player, songs []string) {
        for _, song := range songs {
                device.Play(song)
        }
        device.Stop()
}

func TryOut(device Player) {
        device.Play("Test Track")
        device.Stop()
        recorder, ok := device.(gadget.TapeRecorder)
        if ok {
                recorder.Record()
        }
}

func main() {
        player := gadget.TapePlayer{}
        mixtape := []string{"Jessies Girl", "Whip It", "Can I Pet that Dog"}
        playList(player, mixtape)
        recorder := gadget.TapeRecorder{}
        recordingTape := []string{"A little bit closer", "dead of night"}
        playList(recorder, recordingTape)

        TryOut(gadget.TapeRecorder{})
        TryOut(gadget.TapePlayer{})
}
