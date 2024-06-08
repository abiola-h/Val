package main

import (
	"embed"
	"io"

	"github.com/charmbracelet/huh"
	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto"
)

var (
	//go:embed assets
	fs embed.FS

	ans bool

	note string = `
  You've shouldered so much responsibility at a young age, 
  And you have gracefully held your head high through it all.
  Your dreams are valid and I'm so proud of you and all you do.
  I love that you exist.

  Now I would like to ask the most beautiful woman in the universe, 
  The one whom Lilies blush at her sight,
  The one whom Roses are too embarrassed to bloom in her presence.
  `

	lyrics string = `
  You're my hunny bunch Sugar plum,
  Pumpy-umpy-umpkin
  You're my sweetie pie
  You're my cuppy cake
  Gum drops
  Snoogums boogums, you are...
  The apple of my eye
  And I love you so and I want you to know
  That I'll always be right here
  And I love to sing, sweet songs to you
  Because you are so dear
	`
)

func main() {

	huh.NewForm(
		huh.NewGroup(
			huh.NewNote().
				Description(note),
			huh.NewConfirm().
				Title("Will you be my valentine?").Value(&ans),
		),
	).Run()

	if ans == true {
		go playSong("assets/cupcake.mp3")
		huh.NewNote().Title("The Cuppycake Song").Description(lyrics).Run()
	}
}

func playSong(file string) error {
	f, err := fs.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	d, err := mp3.NewDecoder(f)
	if err != nil {
		return err
	}

	c, err := oto.NewContext(d.SampleRate(), 2, 2, 8192)
	if err != nil {
		return err
	}
	defer c.Close()

	p := c.NewPlayer()
	defer p.Close()

	if _, err := io.Copy(p, d); err != nil {
		return err
	}
	return nil
}
