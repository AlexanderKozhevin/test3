package mediaserver

import (
	"errors"
	"fmt"

	native "github.com/notedit/media-server-go/wrapper"
)

type Player struct {
	player      native.PlayerFacade
	tracks      map[string]*IncomingStreamTrack
	endCallback playerEndCallback
}

type playerEndCallback interface {
	native.PlayerEndListener
	deletePlayerListener()
}

type goplayerEndCallback struct {
	native.PlayerEndListener
}

func (r *goplayerEndCallback) deletePlayerListener() {
	native.DeleteDirectorPlayerEndListener(r.PlayerEndListener)
}

type overwrittenEndCallback struct {
	p native.PlayerEndListener
}

func (p *overwrittenEndCallback) OnEnd() {
	fmt.Println("OnEnd ====================")
}

func NewPlayer(filename string) (*Player, error) {
	player := &Player{}
	player.player = native.NewPlayerFacade()
	player.tracks = make(map[string]*IncomingStreamTrack)

	if player.player.Open(filename) == 0 {
		native.DeletePlayerFacade(player.player)
		return nil, errors.New("player can not open filanme:" + filename)
	}

	if player.player.HasVideoTrack() {

		trackID := "video"
		source := player.player.GetVideoSource()

		incoming := newIncomingStreamTrack("video", trackID, nil, map[string]native.RTPIncomingSourceGroup{"": source})

		// todo event
		player.tracks[trackID] = incoming
	}

	if player.player.HasAudioTrack() {

		trackID := "audio"
		source := player.player.GetAudioSource()

		incoming := newIncomingStreamTrack("audio", trackID, nil, map[string]native.RTPIncomingSourceGroup{"": source})

		// todo
		player.tracks[trackID] = incoming

	}

	callback := &overwrittenEndCallback{}
	p := native.NewDirectorPlayerEndListener(callback)
	callback.p = p

	player.endCallback = &goplayerEndCallback{PlayerEndListener: p}

	player.player.SetPlayEndListener(player.endCallback)

	return player, nil
}

func (p *Player) GetTracks() []*IncomingStreamTrack {
	tracks := []*IncomingStreamTrack{}
	for _, track := range p.tracks {
		tracks = append(tracks, track)
	}
	return tracks
}

func (p *Player) GetAudioTracks() []*IncomingStreamTrack {
	tracks := []*IncomingStreamTrack{}
	for _, track := range p.tracks {
		if track.GetMedia() == "audio" {
			tracks = append(tracks, track)
		}
	}
	return tracks
}

func (p *Player) GetVideoTracks() []*IncomingStreamTrack {
	tracks := []*IncomingStreamTrack{}
	for _, track := range p.tracks {
		if track.GetMedia() == "video" {
			tracks = append(tracks, track)
		}
	}
	return tracks
}

func (p *Player) Play() {

	if p.player != nil {
		p.player.Play()
	}
}

func (p *Player) Resume() {

	if p.player != nil {
		p.player.Play()
	}
}

func (p *Player) Pause() {

	if p.player != nil {
		p.player.Stop()
	}
}

func (p *Player) Seek(time uint64) {

	if p.player != nil {
		p.player.Seek(time)
	}
}

func (p *Player) Stop() {

	if p.player == nil {
		return
	}

	if p.endCallback != nil {
		p.endCallback.deletePlayerListener()
	}

	for _, track := range p.tracks {
		track.Stop()
	}

	p.tracks = nil

	p.player.Close()

	p.player = nil
}
