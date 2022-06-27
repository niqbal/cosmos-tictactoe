package types

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRules_ValidateWinner(t *testing.T) {
	tests := []struct {
		game   string
		winner Mark
	}{
		{
			game:   "XXX|***|***",
			winner: CROSS_MARK,
		}, {
			game:   "XXX|*O*|*O*",
			winner: CROSS_MARK,
		}, {
			game:   "XO*|*X*|*OX",
			winner: CROSS_MARK,
		}, {
			game:   "**X|*X*|XOO",
			winner: CROSS_MARK,
		}, {
			game:   "***|XXX|*OO",
			winner: CROSS_MARK,
		}, {
			game:   "O**|*O*|XXX",
			winner: CROSS_MARK,
		}, {
			game:   "OOO|*X*|*XX",
			winner: CIRCLE_MARK,
		}, {
			game:   "XXO|XOX|XOO",
			winner: CROSS_MARK,
		}, {
			game:   "XOX|XOX|OXO",
			winner: EMPTY_MARK,
		},
	}

	for _, tt := range tests {
		game, err := Parse(tt.game)
		if err != nil {
			panic("parse failure" + tt.game)
		}

		t.Run(fmt.Sprintf("%v Winner should be: %v", tt.game, tt.winner.mark), func(t *testing.T) {
			require.Equal(t, tt.winner, game.FindWinner())
		})
	}
}

func TestRules_ValidateMarkSpace(t *testing.T) {
	tests := []struct {
		game string
		mark Mark
		pos  Pos
		err  string
	}{
		{
			game: "XO*|***|***",
			mark: CROSS_MARK,
			pos:  Pos{X: 0, Y: 0},
			err:  "Position is not empty: {0 0}",
		}, {
			game: "XOX|***|***",
			mark: CROSS_MARK,
			pos:  Pos{X: 1, Y: 1},
			err:  "Not your turn: {X}",
		}, {
			game: "XOX|*OX|*O*",
			mark: CROSS_MARK,
			pos:  Pos{X: 2, Y: 2},
			err:  "game is closed: XOX|*OX|*O*",
		}, {
			game: "XOX|***|***",
			mark: CIRCLE_MARK,
			pos:  Pos{X: 2, Y: 2},
			err:  "",
		},
	}

	for _, tt := range tests {
		game, err := Parse(tt.game)
		if err != nil {
			panic("parse failure" + tt.game)
		}

		t.Run(fmt.Sprintf("Mark %v on board %v at position %v", tt.mark.mark, tt.game, tt.pos), func(t *testing.T) {
			if tt.err != "" {
				require.Equal(t, tt.err, game.MarkSpace(tt.pos, tt.mark).Error())
				return
			}
			require.NoError(t, game.MarkSpace(tt.pos, tt.mark))
		})
	}
}
