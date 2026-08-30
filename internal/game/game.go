package game

import (
	"fmt"
	"strings"

	"github.com/dynamicmines/tic-tac-toe-arena/internal/ai"
	"github.com/dynamicmines/tic-tac-toe-arena/internal/board"
	"github.com/dynamicmines/tic-tac-toe-arena/internal/cli"
)

func ReadMove(name string, maxCell int) int {
	for {
		fmt.Printf("%s move: ", name)
		var cell int
		if _, err := fmt.Scan(&cell); err == nil && cell >= 1 && cell <= maxCell {
			return cell
		}
		var discard string
		fmt.Scan(&discard)
		fmt.Printf("Error: enter a number 1-%d\n", maxCell)
	}
}

func AskPlayAgain() bool {
	for {
		fmt.Print("Play again? (y/n): ")
		var resp string
		fmt.Scan(&resp)
		resp = strings.ToLower(resp)
		if resp == "y" {
			return true
		}
		if resp == "n" {
			return false
		}
	}
}

func PrintStats(total, winsX, winsO, draws, moves int, verbose bool) {
	fmt.Println("=== Stats ===")
	fmt.Printf("Games: %d   X: %d   O: %d   Draws: %d\n", total, winsX, winsO, draws)
	if verbose && moves > 0 && total > 0 {
		fmt.Printf("Moves this game: %d   Win rate — X: %d%%  O: %d%%\n", moves, (winsX*100)/total, (winsO*100)/total)
	}
}

func Run(cfg cli.Config) {
	totalGames, winsX, winsO, draws := 0, 0, 0, 0

	for {
		totalGames++
		moves := 0
		b := board.NewBoard(cfg.Size)
		turn := cfg.First

		for {
			b.Render(cfg.Color, cfg.Big)

			name := cfg.NameX
			if turn == "O" {
				name = cfg.NameO
			}

			var move int
			if cfg.ModeAI && turn == "O" {
				var grid [9]string
				copy(grid[:], b.Cells)
				move = ai.FindMove(grid, "O", "X")
				if cfg.Verbose {
					fmt.Printf("AI: move at %d\n", move)
				} else {
					fmt.Printf("%s plays %d\n", name, move)
				}
			} else {
				for {
					move = ReadMove(name, cfg.Size*cfg.Size)
					if b.Cells[move-1] == "" {
						break
					}
					fmt.Printf("Error: cell %d is taken\n", move)
				}
			}

			err := b.PlaceMark(move, turn)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			moves++

			if won, winner := b.CheckWin(); won {
				b.Render(cfg.Color, cfg.Big)
				if winner == "X" {
					winsX++
					fmt.Printf("%s wins!\n", cfg.NameX)
				} else {
					winsO++
					fmt.Printf("%s wins!\n", cfg.NameO)
				}
				break
			}

			if b.IsDraw() {
				b.Render(cfg.Color, cfg.Big)
				draws++
				fmt.Println("Draw!")
				break
			}

			if turn == "X" {
				turn = "O"
			} else {
				turn = "X"
			}
		}

		PrintStats(totalGames, winsX, winsO, draws, moves, cfg.Verbose)

		if !AskPlayAgain() {
			break
		}
	}
}