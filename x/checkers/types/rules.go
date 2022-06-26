package types

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
)

const (
	BOARD_DIM = 3
	CROSS     = "X"
	CIRCLE    = "O"
	EMPTY     = "*"
	ROW_SEP   = "|"
)

type Mark struct {
	mark string
}

var CROSS_MARK = Mark{mark: CROSS}
var CIRCLE_MARK = Mark{mark: CIRCLE}
var EMPTY_MARK = Mark{mark: EMPTY}

type Pos struct {
	X int
	Y int
}

type Game struct {
	Pieces map[Pos]string
}

func New() *Game {
	pieces := make(map[Pos]string)
	game := &Game{pieces}
	game.initializeSpace()
	return game
}

func (game *Game) initializeSpace() {
	for i := 0; i < BOARD_DIM; i++ {
		for j := 0; j < BOARD_DIM; j++ {
			game.Pieces[Pos{X: i, Y: j}] = EMPTY
		}
	}
}

func (game *Game) GetTurn() Mark {
	emptyCount := game.getEmptyCount()

	if emptyCount%2 == 1 {
		return CROSS_MARK
	} else {
		return CIRCLE_MARK
	}
}

func (game *Game) getEmptyCount() int {
	emptyCount := 0
	for i := 0; i < BOARD_DIM; i++ {
		for j := 0; j < BOARD_DIM; j++ {

			if game.Pieces[Pos{X: i, Y: j}] == EMPTY {
				emptyCount++
			}

		}
	}
	return emptyCount
}

func (game *Game) FindWinner() Mark {
	for i := 0; i < BOARD_DIM; i++ {
		winnerMark := Mark{game.Pieces[Pos{X: i, Y: 0}]}
		for j := 0; j < BOARD_DIM; j++ {
			if game.Pieces[Pos{X: i, Y: j}] != winnerMark.mark {
				winnerMark = EMPTY_MARK
				break
			}
		}

		if winnerMark != EMPTY_MARK {
			return winnerMark
		}
	}

	for j := 0; j < BOARD_DIM; j++ {
		winnerMark := Mark{game.Pieces[Pos{X: 0, Y: j}]}
		for i := 0; i < BOARD_DIM; i++ {
			// fmt.Println(fmt.Sprintln("x:%v y:%v value:%v", i, j, game.Pieces[Pos{X: i, Y: j}]))
			if game.Pieces[Pos{X: i, Y: j}] != winnerMark.mark {
				winnerMark = EMPTY_MARK
				break
			}
		}

		if winnerMark != EMPTY_MARK {
			return winnerMark
		}
	}

	winnerMark := Mark{game.Pieces[Pos{X: 0, Y: 0}]}
	for i := 0; i < BOARD_DIM; i++ {
		if game.Pieces[Pos{X: i, Y: i}] != game.Pieces[Pos{X: 0, Y: 0}] {
			winnerMark = EMPTY_MARK
			break
		}
	}

	if winnerMark != EMPTY_MARK {
		return winnerMark
	}

	winnerMark = Mark{game.Pieces[Pos{X: BOARD_DIM - 1, Y: 0}]}

	for i := 0; i < BOARD_DIM; i++ {
		if game.Pieces[Pos{X: BOARD_DIM - 1 - i, Y: i}] != game.Pieces[Pos{X: BOARD_DIM - 1, Y: 0}] {
			winnerMark = EMPTY_MARK
			break
		}
	}

	return winnerMark
}

func (game *Game) GameClosed() bool {
	if game.getEmptyCount() == 0 {
		fmt.Sprintf("No empty space")
		return true
	}

	return game.FindWinner() != EMPTY_MARK
}

func (game *Game) MarkSpace(pos Pos, mark Mark) error {
	if game.GameClosed() {
		return errors.New(fmt.Sprintf("game is closed: %v", game))
	}

	if game.Pieces[pos] == EMPTY {
		if mark == game.GetTurn() {
			game.Pieces[pos] = mark.mark
			return nil
		} else {
			return errors.New(fmt.Sprintf("invalid move: %v", mark))
		}
	}

	return errors.New(fmt.Sprintf("Position is not empty: %v", pos))
}

func (game *Game) IsCrossTurn() bool {
	return game.GetTurn() == CROSS_MARK
}

func (game *Game) IsCircleTurn() bool {
	return game.GetTurn() == CIRCLE_MARK
}

func (game *Game) String() string {
	var buf bytes.Buffer
	for i := 0; i < BOARD_DIM; i++ {
		for j := 0; j < BOARD_DIM; j++ {
			buf.WriteString(game.Pieces[Pos{X: i, Y: j}])
		}
		if i < (BOARD_DIM - 1) {
			buf.WriteString(ROW_SEP)
		}
	}
	return buf.String()
}

func Parse(s string) (*Game, error) {
	if len(s) != BOARD_DIM*BOARD_DIM+(BOARD_DIM-1) {
		return nil, errors.New(fmt.Sprintf("invalid board string: %v", s))
	}
	pieces := make(map[Pos]string)
	result := &Game{pieces}
	for y, row := range strings.Split(s, ROW_SEP) {
		for x, c := range strings.Split(row, "") {
			if x >= BOARD_DIM || y >= BOARD_DIM {
				return nil, errors.New(fmt.Sprintf("invalid board, mark out of bounds: %v, %v", x, y))
			}
			if isValidMark(c) {
				result.Pieces[Pos{y, x}] = c
			} else {
				return nil, errors.New(fmt.Sprintf("invalid board, invalid mark at %v, %v", x, y))
			}
		}
	}

	return result, nil
}

func isValidMark(c string) bool {
	return c == CROSS || c == CIRCLE || c == EMPTY
}

// func main() {
// 	g := New()
// 	fmt.Println(g.MarkSpace(Pos{0, 2}, CROSS_MARK))
// 	fmt.Println(g.FindWinner())

// 	fmt.Println(g.MarkSpace(Pos{2, 1}, CIRCLE_MARK))
// 	fmt.Println(g.FindWinner())
// 	fmt.Println(g.MarkSpace(Pos{0, 0}, CROSS_MARK))
// 	fmt.Println(g.FindWinner())
// 	fmt.Println(g.MarkSpace(Pos{2, 2}, CIRCLE_MARK))
// 	fmt.Println(g.FindWinner())
// 	fmt.Println(g.MarkSpace(Pos{0, 1}, CROSS_MARK))
// 	fmt.Println(g.FindWinner())
// 	fmt.Println(g.MarkSpace(Pos{2, 0}, CIRCLE_MARK))

// 	fmt.Println(Parse(g.String()))

// }
