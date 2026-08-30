package cli

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	ModeAI  bool
	Color   bool
	Big     bool
	Verbose bool
	First   string
	NameX   string
	NameO   string
	Size    int
}

func PrintUsage() {
	fmt.Println("Usage: go run main.go (--players | --ai) [options]")
	fmt.Println()
	fmt.Println("Modes (exactly one required):")
	fmt.Println("  --players        two human players take turns")
	fmt.Println("  --ai             play against the computer (you are X)")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  --color          enable colored output (default: plain)")
	fmt.Println("  --big            render the board with large glyphs")
	fmt.Println("  --verbose        show extended statistics")
	fmt.Println("  --first X|O      who moves first (default: X)")
	fmt.Println("  --name A,B       custom names: X=A, O=B (e.g. --name Alice,Bob)")
	fmt.Println("  --size N         board is N×N, win = N in a row (default: 3)")
	fmt.Println("  --help, -h       print this help and exit 0")
}

func Parse(args []string) (Config, error) {
	cfg := Config{
		First: "X",
		NameX: "X",
		NameO: "O",
	}

	mode := ""

	for i := 0; i < len(args); i++ {
		arg := args[i]

		switch arg {
		case "--help", "-h":
			PrintUsage()
			os.Exit(0)

		case "--players":
			if mode != "" {
				return cfg, errors.New("choose exactly one of --players or --ai")
			}
			mode = "players"

		case "--ai":
			if mode != "" {
				return cfg, errors.New("choose exactly one of --players or --ai")
			}
			mode = "ai"
			cfg.ModeAI = true

		case "--color":
			cfg.Color = true

		case "--big":
			cfg.Big = true

		case "--verbose":
			cfg.Verbose = true

		case "--first":
			if i+1 >= len(args) {
				return cfg, errors.New("missing value for --first")
			}
			i++
			cfg.First = strings.ToUpper(args[i])

		case "--name":
			if i+1 >= len(args) {
				return cfg, errors.New("missing value for --name")
			}
			i++
			names := strings.Split(args[i], ",")
			if len(names) != 2 || names[0] == "" || names[1] == "" {
				return cfg, errors.New("invalid --name format, expected Name1,Name2")
			}
			cfg.NameX = names[0]
			cfg.NameO = names[1]

		case "--size":
			if i+1 >= len(args) {
				return cfg, errors.New("missing value for --size")
			}
			i++
			val, err := strconv.Atoi(args[i])
			if err != nil || val < 3 {
				return cfg, errors.New("--size must be an integer >= 3")
			}
			cfg.Size = val

		default:
			return cfg, fmt.Errorf("unknown option %s", arg)
		}
	}

	if mode == "" {
		return cfg, errors.New("choose exactly one of --players or --ai")
	}

	if mode == "ai" && cfg.Size > 0 {
		return cfg, errors.New("--ai and --size cannot be combined (AI is 3×3 only)")
	}

	if cfg.Size == 0 {
		cfg.Size = 3
	}

	if cfg.First != "X" && cfg.First != "O" {
		return cfg, errors.New("--first accepts only X or O")
	}

	return cfg, nil
}