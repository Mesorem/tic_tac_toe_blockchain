package rules

import (
	"bytes"
	"container/list"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	CROSS       = "cross"
	CIRCLE      = "circle"
	CROSS_MARK  = "x"
	CIRCLE_MARK = "o"
	VOID_MARK   = "-"
	SEP         = ";"
)

// PLAYERS
type Mark = struct {
	Sign string
}

var CircleMark = Mark{CIRCLE_MARK}
var CrossMark = Mark{CROSS_MARK}
var VoidMark = Mark{VOID_MARK}

type Player struct {
	Kind string
	Mark Mark
}

var CirclePlayer = Player{CIRCLE, CircleMark}
var CrossPlayer = Player{CROSS, CrossMark}

var marksPlayerCorrespondence = map[Mark]Player{
	CrossMark:  CrossPlayer,
	CircleMark: CirclePlayer,
}

// POSITION
type Pos struct {
	X int
	Y int
}

var NO_POS = Pos{-1, -1}

// GAME
type Game struct {
	Representation string
	Boxes          map[int]map[int]Mark
	MovesHist      list.List
	TurnNumber     int
	CurrentPlayer  Player
	EndedGame      bool
}

type MoveHist struct {
	Position Pos
	Player   Player
}

func (game *Game) Represent() {
	var buf bytes.Buffer
	for y := 0; y < 3; y++ {
		if y > 0 {
			buf.WriteString("|")
		}
		for x := 0; x < 3; x++ {
			buf.WriteString(game.Boxes[y][x].Sign)
			if x < 2 {
				buf.WriteString(",")
			}
		}
	}
	buf.WriteString(SEP)
	buf.WriteString(strconv.Itoa(game.TurnNumber))
	buf.WriteString(SEP)
	buf.WriteString(game.CurrentPlayer.Mark.Sign)
	buf.WriteString(SEP)
	buf.WriteString(strconv.FormatBool(game.EndedGame))
}

func (game *Game) Move(destPos Pos, playerMoving Player) (Pos, error) {
	if playerMoving != game.CurrentPlayer {
		return NO_POS, errors.New(fmt.Sprintf("Invalid movement, turn is for player %v, moved player %v", game.CurrentPlayer, playerMoving))
	}

	if (destPos.X > 2 || destPos.X < 0) || (destPos.Y > 2 || destPos.Y < 0) {
		return NO_POS, errors.New(fmt.Sprintf("Invalid movement, index of movement out of bounds %v", destPos))
	}

	if game.EndedGame {
		return NO_POS, errors.New(fmt.Sprintf("Invalid movement, game has been already ended"))
	}

	endPos, err := game.move(destPos)
	if err != nil {
		return endPos, err
	}

	if game.CurrentPlayer == CirclePlayer {
		game.CurrentPlayer = CrossPlayer
	} else {
		game.CurrentPlayer = CirclePlayer
	}

	game.TurnNumber = game.TurnNumber + 1

	return endPos, err
}

func (game *Game) move(destPos Pos) (confirmedPos Pos, err error) {
	confirmedPos = NO_POS
	err = nil

	stateDestPos := game.Boxes[destPos.X][destPos.Y]
	if stateDestPos != VoidMark {
		return NO_POS, errors.New(fmt.Sprintf("Invaild movement, position %v has already been marked by player %v", destPos, stateDestPos))
	}

	confirmedPos = destPos
	game.Boxes[destPos.X][destPos.Y] = game.CurrentPlayer.Mark

	return
}

func (game *Game) Winner() (isWinnerSomePlayer bool, winner Player) {
	isWinnerSomePlayer = false
	winner = CirclePlayer

	for y := 0; y < 3; y++ {
		for x := 0; x < 3; x++ {
			currValue := game.Boxes[y][x]
			if currValue != VoidMark {

				// validating in x axis
				foundEquals := true
				equalsCount := 0
				for n := x; n < 3 && foundEquals; n++ {
					if game.Boxes[n][y] != currValue {
						foundEquals = false
					} else {
						equalsCount++
						if equalsCount == 3 {
							return true, marksPlayerCorrespondence[currValue]
						}
					}
				}

				// validating in y axis, but only from the first line
				if y == 0 {
					equalsCount = 0
					foundEquals = true
					for n := y; n < 3 && foundEquals; n++ {
						if game.Boxes[x][n] != currValue {
							foundEquals = false
						} else {
							equalsCount++
							if equalsCount == 3 {
								return true, marksPlayerCorrespondence[currValue]
							}
						}
					}
				}

				// validating xy axis, but only from the first line
				if x == 0 {
					equalsCount = 0
					foundEquals = true
					for ny := y; ny < 3 && foundEquals; {
						for nx := x; nx < 3 && foundEquals; nx++ {
							if game.Boxes[nx][ny] != currValue {
								foundEquals = false
							} else {
								equalsCount++
								if equalsCount == 3 {
									return true, marksPlayerCorrespondence[currValue]
								}
							}
							ny++
						}
					}

					foundEquals = true
					equalsCount = 0
					for ny := y; ny >= 0 && foundEquals; {
						for nx := x; nx >= 0 && foundEquals; nx++ {
							//fmt.Printf("Comparing (%v, %v)=%v to (%v, %v)=%v\n", x, y, currValue, nx, ny, game.Boxes[nx][ny])
							if game.Boxes[nx][ny] != currValue {
								foundEquals = false
							} else {
								equalsCount++
								if equalsCount == 3 {
									return true, marksPlayerCorrespondence[currValue]
								}
							}
							ny--
						}
					}
				}
			}
		}
	}

	return
}

func New() *Game {

	var newGame = &Game{
		Representation: "",
		Boxes:          *initBoxes(),
		MovesHist:      *list.New(),
		TurnNumber:     0,
		CurrentPlayer:  CrossPlayer,
	}

	return newGame
}

func initBoxes() *map[int]map[int]Mark {
	initBoxes := make(map[int]map[int]Mark)
	for y := 0; y < 3; y++ {
		var xMap = make(map[int]Mark)
		for x := 0; x < 3; x++ {
			xMap[x] = VoidMark
		}
		initBoxes[y] = xMap
	}
	return &initBoxes
}

func Parse(gameText string) (*Game, error) {
	fields := strings.Split(gameText, SEP)
	boardLines := strings.Split(fields[0], "|")
	game := New()
	for indexX, valueX := range boardLines {
		boxes := strings.Split(valueX, ",")
		for indexY, valueY := range boxes {
			if valueY != VOID_MARK {
				game.Boxes[indexX][indexY] = Mark{valueY}
			}
		}
	}
	conv, err := strconv.Atoi(fields[1])
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Invaild turn %v, it has to be a valid uint", fields[1]))
	}
	game.TurnNumber = conv

	currentPlayerSign := fields[2]
	if currentPlayerSign == CROSS_MARK {
		game.CurrentPlayer = CrossPlayer
	} else {
		game.CurrentPlayer = CirclePlayer
	}

	convBool, err := strconv.ParseBool(fields[3])
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Invaild ended game value %v, it has to be a valid boolean", fields[3]))
	}
	game.EndedGame = convBool
	game.Representation = gameText

	return game, nil
}
