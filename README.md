## What is this?
This is a cli application I built to ask my girlfriend to be my valentine. 
It shows a note with a yes or no prompt beneath. When you click on a yes, it plays the [Cupcake song](https://youtu.be/wAgZVLk6J4M?si=L_ey_0tOIImyKVoS) while showing the lyrics of the song on the screen.

## How can I run this?
1. Clone this repo by running `git clone https://github.com/quamejnr/Val.git` on your terminal
2. Navigate to the repo by running `cd Val` then run `go run .`. This should run the code. 
Ensure you have Golang already installed 

## How can I use it? 
### Change the notes shown
1. You can change the note that's shown with the prompt by changing the `note` variable from line 19 to 26.

### Change the song and lyrics
1. Delete the cupcake song in the `assets/` folder and copy your song to the `assets` folder.
2. Change `go playSong("assets/cupcake.mp3")` on line 56 to the name of your new file `go playSong("assets/<file_name>")`
3. Change the lyrics assigned to the `lyrics` variable from line 30 to 40 to your new lyrics.
4. Change the title that appears when your song is being played from "The Cuppycake Song" to the name of your song on line 57. `huh.NewNote().Title("Name of your new song").Description(lyrics).Run()`
5. Save your file and [run](#how-can-i-run-this) the code.


