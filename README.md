# Jogo da Velha (Tic-Tac-Toe) — Go + Fyne GUI

A cross-platform tic-tac-toe game built in Go 1.26 with the Fyne GUI framework. Play tic-tac-toe with a graphical interface.

## Quick Start

### Windows

1. **Install C Compiler** (required for Fyne's OpenGL support)

   ** Visual Studio Community**
   - Download from https://visualstudio.microsoft.com
   - Install "Desktop development with C++" workload
   - Restart your machine

2. **Verify GCC Installation**
   ```powershell
   gcc --version
   ```

3. **Build and Run**
   ```powershell
   .\build-setup.ps1
   ```
   
   Or build manually:
   ```powershell
   $env:CGO_ENABLED=1; go build ./...
   .\awesomeProject.exe
   ```

### macOS

```bash
go build ./...
./awesomeProject
```

### Linux

```bash
# Install dependencies (Ubuntu/Debian)
sudo apt-get install -y build-essential pkg-config libgl1-mesa-dev libxcursor-dev libxrandr-dev libxinerama-dev libxi-dev libxext-dev

# Build and run
go build ./...
./awesomeProject
```

## How to Play

1. Launch the application
2. Players alternate: **X** goes first, then **O**
3. Click any empty square on the 3×3 board to place your mark
4. Win by getting three marks in a row (horizontal, vertical, or diagonal)
5. Game ends in a draw if the board fills with no winner
6. Click **Reiniciar** or **Nova Partida** to reset and play again

## Build Error: `github.com/go-gl/gl: build constraints exclude all Go files`

This error occurs on Windows when building without a C compiler. **This is expected and fixable:**

**Root Cause**: Fyne's GUI framework uses OpenGL for rendering on Windows, which requires a C compiler (GCC or MSVC) and CGO to build C bindings.

**Solution**: Install MinGW-w64, Visual Studio, or Scoop (see "Quick Start" above).

After installing a C compiler:
```powershell
$env:CGO_ENABLED=1; go build ./...
```

## Project Structure

```
awesomeProject/
├── main.go                    # Fyne app entry point
├── go.mod                     # Go module definition
├── go.sum                     # Dependency checksums
├── README.md                  # This file
├── AGENTS.md                  # AI agent guide (architecture, patterns, workflows)
├── build-setup.ps1            # Windows build helper script
├── jogos/
│   └── jogoDaVelha.go        # Game logic (board, moves, win detection)
└── ui/
    └── jogoDaVelhaBoard.go    # Fyne GUI (buttons, status, layout)
```

## Key Components

- **`main.go`** (19 lines)  
  Fyne application entry point. Creates an app and builds the UI.

- **`jogos/jogoDaVelha.go`** (56 lines)  
  Game logic: board state, move validation, turn management. Note: `CheckWinner()` is incomplete and needs win/draw detection logic.

- **`ui/jogoDaVelhaBoard.go`** (125 lines)  
  Fyne GUI: 3×3 button grid, status label, reset buttons. Calls game logic on each button click.

## Troubleshooting

| Error | Cause | Fix |
|-------|-------|-----|
| `build constraints exclude all Go files in github.com/go-gl/gl` | No C compiler on Windows | Install MinGW-w64 or Visual Studio (see Quick Start) |
| `cgo: C compiler "gcc" not found: exec: "gcc": executable file not found in %PATH%` | GCC installed but not in PATH | Add MinGW `bin` directory to Windows PATH (e.g., `C:\msys64\mingw64\bin`) |
| Application crashes on startup | Graphics driver issue | Update GPU drivers; try on a different machine if persists |

## Development

### Run Tests
```bash
go test ./...
```

### Lint
```bash
go vet ./...
```

### Clean Rebuild
```bash
go clean -cache
$env:CGO_ENABLED=1; go build ./...
```

## Known Limitations

1. **Incomplete Win Detection**  
   `CheckWinner()` in `jogos/jogoDaVelha.go` only toggles the current player. It needs to:
   - Check for three in a row (horizontal, vertical, diagonal)
   - Detect a draw when the board is full
   - Set `GameOver=true` and `Winner` appropriately

2. **Buttons Remain Active After Game Over**  
   The game prevents moves but doesn't visually disable buttons. Consider adding `button.Disable()` in the UI.

3. **No AI Opponent**  
   Only supports human vs. human play. An AI using minimax would be a good addition.

## Dependencies

- **Go 1.26+** — Programming language
- **fyne.io/fyne/v2** — GUI framework
  - **On Windows**: Requires C compiler (MinGW-w64 GCC, MSVC, or Visual Studio)
  - **On macOS/Linux**: System OpenGL/GLFW libraries

## Next Steps for Developers

1. **Fix `CheckWinner()`** — Highest priority; game unplayable without it
2. **Add unit tests** — Create `jogos/jogoDaVelha_test.go`
3. **Add AI opponent** — Implement minimax algorithm
4. **Enhance UI** — Add score tracking, game menu, themes
5. **CLI mode** — Support non-GUI play for environments without C compiler

## For AI Coding Agents

See `AGENTS.md` for detailed documentation including:
- Big-picture architecture and design decisions
- Critical developer workflows and build commands
- Project-specific patterns and pitfalls (incomplete win detection, etc.)
- Integration points and platform-specific dependencies
- Concrete examples of improvements to make

