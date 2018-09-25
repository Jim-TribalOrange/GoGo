package main

import (
	"io/ioutil"
	"log"
	"os"
)

type sgfFile struct {
	Ginfo      sgfGameInfo
	Moves      []sgfMove
	filePath   string
	fileString string //the content of the sgf file
}

type sgfMove struct {
	Comment string
	colour  rune
	move    []rune //x then y location of the move in a..boardsize formate
}

type sgfGameInfo struct {
	AB string //Add Black: locations of Black stones to be placed on the board prior to the first move
	AW string //Add White: locations of White stones to be placed on the board prior to the first move.
	AN string //Annotations: name of the person commenting the game.
	AP string //Application: application that was used to create the SGF file (e.g. CGOban2,...).
	BR string //Black Rank: rank of the Black player.
	BT string //Black Team: name of the Black team.
	CP string //Copyright: copyright information.
	DT string //Date: date of the game.
	EV string //Event: name of the event (e.g. 58th Honinbō Title Match).
	FF string //File format: version of SGF specification governing this SGF file.
	GM string //Game: type of game represented by this SGF file. A property value of 1 refers to Go.
	GN string //Game Name: name of the game record.
	HA string //Handicap: the number of handicap stones given to Black. Placement of the handicap stones are set using the AB property.
	KM string //Komi: komi.
	ON string //Opening: information about the opening (Fuseki), rarely used in any file.
	OT string //Overtime: overtime system.
	PB string //Black Name: name of the black player.
	PC string //Place: place where the game was played (e.g.: Tokyo).
	PL string //Player: color of player to start.
	PW string //White Name: name of the white player.
	RE string //Result: result, usually in the format "B+R" (Black wins by resign) or "B+3.5" (black wins by 3.5).
	RO string //Round: round (e.g.: 5th game).
	RU string //Rules: ruleset (e.g.: Japanese).
	SO string //Source: source of the SGF file.
	SZ string //Size: size of the board, non-square boards are supported.
	TM string //Time limit: time limit in seconds.
	US string //User: name of the person who created the SGF file.
	WR string //White Rank: rank of the White player.
	WT string //White Team: name of the White team.

}

func (s *sgfFile) loadFile() {

	file, err := os.Open(s.filePath)
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	s.fileString = string(data[:])

}
