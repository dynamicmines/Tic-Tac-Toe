package board

import (
	"errors"
	"fmt"
	"strconv" 
	"strings"
)

const (
	ColorReset = "\033[0m"
	ColorRed   = "\033[91m"
	ColorBlue  = "\033[94m"
	ColorGreen = "\033[1;32m"
	ColorDim   = "\033[2m"
)


type Board struct {
	Size     int
	Cells    []string 
	WinLine  []int
	IsWinner bool
}

func NewBoard(size int) *Board {
	return &Board{
		Size:  size,
		
		Cells: make([]string, size*size),
	}
}

func (b *Board) Copy() *Board {
	newCells := make([]string, len(b.Cells))
	for i, v := range b.Cells {
		newCells[i] = v
	}

	newWinLine := make([]int, len(b.WinLine))
	for i, v := range b.WinLine {
		newWinLine[i] = v
	}

	return &Board{
		Size:     b.Size,
		Cells:    newCells,
		WinLine:  newWinLine,
		IsWinner: b.IsWinner,
	}
}

func (b *Board) PlaceMark(cellNum int, symbol string) error {
	if cellNum < 1 || cellNum > b.Size*b.Size {
		return fmt.Errorf("выход за пределы поля: введите от 1 до %d", b.Size*b.Size)
	}

	idx := cellNum - 1
	if b.Cells[idx] != "" {
		return errors.New("эта клетка уже занята")
	}

	b.Cells[idx] = symbol
	return nil 
}

func (b *Board) CheckWin() (bool, string) {
	n := b.Size

	for i := 0; i < n; i++ {
		line := make([]int, 0)
		for j := 0; j < n; j++ {
			line = append(line, i*n+j) 
		}
		if won, sym := b.checkLine(line); won {
			return won, sym
		}
	}

	for j := 0; j < n; j++ {
		line := make([]int, 0)
		for i := 0; i < n; i++ {
			line = append(line, i*n+j)
		}
		if won, sym := b.checkLine(line); won {
			return won, sym
		}
	}

	diag1 := make([]int, 0)
	for i := 0; i < n; i++ {
		diag1 = append(diag1, i*n+i)
	}
	if won, sym := b.checkLine(diag1); won {
		return won, sym
	}

	diag2 := make([]int, 0)
	for i := 0; i < n; i++ {
		diag2 = append(diag2, i*n+(n-1-i))
	}
	if won, sym := b.checkLine(diag2); won {
		return won, sym
	}

	return false, ""
}

func (b *Board) checkLine(indices []int) (bool, string) {
	first := b.Cells[indices[0]]
	if first == "" {
		return false, ""
	}
	
	for _, idx := range indices {
		if b.Cells[idx] != first {
			return false, ""
		}
	}
	
	b.WinLine = indices
	b.IsWinner = true
	return true, first
}

func (b *Board) IsDraw() bool {
	for _, cell := range b.Cells {
		if cell == "" {
			return false
		}
	}
	won, _ := b.CheckWin()
	return !won 
}

func (b *Board) Render(useColor, useBig bool) {
	if useBig {
		b.renderBig(useColor)
	} else {
		b.renderStandard(useColor)
	}
}

func (b *Board) renderStandard(useColor bool) {
	for i := 0; i < b.Size; i++ {
		row := make([]string, 0)
		
		for j := 0; j < b.Size; j++ {
			idx := i*b.Size + j
			sym := b.Cells[idx]

			if sym == "" {
				sym = strconv.Itoa(idx + 1)
				if useColor {
					sym = ColorDim + sym + ColorReset
				}
			} else {
				if useColor {
					sym = b.colorize(sym, idx)
				}
			}
			formatted := fmt.Sprintf(" %2s ", sym)
			row = append(row, formatted)
		}

		fmt.Println(strings.Join(row, "|"))
		
		if i < b.Size-1 {
			separator := make([]string, b.Size)
			for s := range separator {
				separator[s] = "----"
			}
			fmt.Println(strings.Join(separator, "+"))
		}
	}
}

func (b *Board) renderBig(useColor bool) {
	for i := 0; i < b.Size; i++ {
		line1 := make([]string, b.Size)
		line2 := make([]string, b.Size)
		line3 := make([]string, b.Size)

		for j := 0; j < b.Size; j++ {
			idx := i*b.Size + j
			sym := b.Cells[idx]
			
			var glyph [3]string 
			
			switch sym {
			case "X":
				glyph = [3]string{"X   X", "  X  ", "X   X"}
			case "O":
				glyph = [3]string{" OOO ", "O   O", " OOO "}
			default:
				numStr := strconv.Itoa(idx + 1)
				if len(numStr) == 1 {
					glyph = [3]string{"     ", fmt.Sprintf("  %s  ", numStr), "     "}
				} else {
					glyph = [3]string{"     ", fmt.Sprintf(" %s  ", numStr), "     "}
				}
			}

			if useColor {
				for l := 0; l < 3; l++ {
					if sym == "" {
						glyph[l] = ColorDim + glyph[l] + ColorReset
					} else {
						glyph[l] = b.colorize(glyph[l], idx)
					}
				}
			}

			line1[j] = glyph[0]
			line2[j] = glyph[1]
			line3[j] = glyph[2]
		}

		fmt.Println(strings.Join(line1, "|"))
		fmt.Println(strings.Join(line2, "|"))
		fmt.Println(strings.Join(line3, "|"))

		if i < b.Size-1 {
			separator := make([]string, b.Size)
			for s := range separator {
				separator[s] = "-----"
			}
			fmt.Println(strings.Join(separator, "+"))
		}
	}
}

func (b *Board) colorize(sym string, idx int) string {
	isWin := false
	if b.IsWinner {
		for _, wIdx := range b.WinLine {
			if idx == wIdx {
				isWin = true
				break
			}
		}
	}

	if isWin {
		return ColorGreen + sym + ColorReset
	}
	if strings.Contains(sym, "X") {
		return ColorRed + sym + ColorReset
	}
	if strings.Contains(sym, "O") {
		return ColorBlue + sym + ColorReset
	}
	return sym
}
