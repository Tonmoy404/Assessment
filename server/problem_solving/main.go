package main

import (
	"container/list"
	"errors"
	"fmt"
	"sync"
)

type Song struct {
	ID   string
	Name string
}

type MusicPlayer struct {
	songs       map[string]*Song
	recentSongs *list.List
	cacheSize   int
	mu          sync.Mutex
}

func NewMusicPlayer(cacheSize int) *MusicPlayer {
	return &MusicPlayer{
		songs:       make(map[string]*Song),
		recentSongs: list.New(),
		cacheSize:   cacheSize,
	}
}

func (p *MusicPlayer) AddNewSong(song *Song) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if _, ok := p.songs[song.ID]; ok {
		return errors.New("song already exists")
	}

	p.songs[song.ID] = song

	p.recentSongs.PushFront(song.ID)

	if p.recentSongs.Len() > p.cacheSize {
		last := p.recentSongs.Back()
		if last != nil {
			p.recentSongs.Remove(last)
		}
	}

	return nil
}

func (p *MusicPlayer) GetSong(id string) (*Song, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	song, ok := p.songs[id]
	if !ok {
		return nil, fmt.Errorf("song not found")
	}

	for e := p.recentSongs.Front(); e != nil; e = e.Next() {
		if e.Value == id {
			p.recentSongs.MoveToFront(e)
			break
		}
	}

	return song, nil
}

func (p *MusicPlayer) GetRecentSongs() []*Song {
	p.mu.Lock()
	defer p.mu.Unlock()

	var recent []*Song
	for e := p.recentSongs.Front(); e != nil; e = e.Next() {
		id := e.Value.(string)
		if song, ok := p.songs[id]; ok {
			recent = append(recent, song)
		}
	}

	return recent
}

func main() {
	player := NewMusicPlayer(5)

	song1 := &Song{ID: "1", Name: "Song 1"}
	song2 := &Song{ID: "2", Name: "Song 2"}
	song3 := &Song{ID: "3", Name: "Song 3"}
	song4 := &Song{ID: "4", Name: "Song 4"}
	song5 := &Song{ID: "5", Name: "Song 5"}
	song6 := &Song{ID: "6", Name: "Song 6"}
	player.AddNewSong(song1)
	player.AddNewSong(song2)
	player.AddNewSong(song3)
	player.AddNewSong(song4)
	player.AddNewSong(song5)
	player.AddNewSong(song6)

	song, err := player.GetSong("3")
	if err == nil {
		fmt.Println("Found song:", song.Name)
	} else {
		fmt.Println("Error:", err)
	}

	recentSongs := player.GetRecentSongs()
	fmt.Println("Recent songs:")
	for _, song := range recentSongs {
		fmt.Println(song.Name)
	}
}
