# Agents guide for this repository

Overview
--------
This is a Go 1.26 tic-tac-toe application ("Jogo da Velha") built with Fyne GUI framework:
- **main.go** — Fyne app entrypoint
- **jogos/jogoDaVelha.go** — Game logic (board, moves, win detection)
- **ui/jogoDaVelhaBoard.go** — Fyne GUI widget for the game

**CRITICAL: Windows Build Error (Expected)**
On Windows, `go build ./...` fails with:
```
imports github.com/go-gl/gl/v2.1/gl: build constraints exclude all Go files...
```
This is **expected and fixable**—Fyne's OpenGL painter requires CGO (C bindings) and a C compiler (GCC/MinGW or MSVC). See "Windows Build Setup Instructions" below to resolve.

Quick checklist for an AI coding agent (what to do first)
- Read `main.go`, `jogos/jogoDaVelha.go`, and `ui/jogoDaVelhaBoard.go` to understand the architecture.
- **On Windows**: Install MinGW-w64 (GCC) or MSVC compiler before building (see Setup Instructions).
- **On macOS/Linux**: `go build ./...` should work if system OpenGL/GLFW libraries are installed.
- Run `go vet ./...` and `go test ./...` to check correctness.

Architecture & big-picture
--------------------------
- **Game Logic (Pure Go)**: `jogos/jogoDaVelha.go` models the board state and game rules, independent of UI.
- **GUI Layer**: `ui/jogoDaVelhaBoard.go` uses Fyne to render a 3×3 button grid, status display, and reset/resume buttons.
- **Entrypoint**: `main.go` launches the Fyne app and calls `ui.Build()` to construct the UI.
- **Platform-specific dependency**: Fyne's OpenGL painter transitively requires `github.com/go-gl/gl` (OpenGL bindings) and `github.com/go-gl/glfw/v3.3/glfw` (window/input). On Windows, these require a C compiler and CGO to build.

Key files to inspect
--------------------
- **main.go** (19 lines)
  - Calls `app.New()` to create Fyne app instance.
  - Calls `ui.Build(a)` to construct and show the UI.
  - Calls `a.Run()` to start the event loop.

- **jogos/jogoDaVelha.go** (56 lines)
  - `JogoDaVelha` struct: holds a `[3][3]string` board, current player ("X"/"O"), game-over flag, and winner.
  - `NewJDV()`: constructor initializing board with 3×3 empty strings and CurrentPlayer="X".
  - `PlayJDV(r, c)`: place mark at row r, col c; returns false if out-of-bounds/already occupied/game over.
  - `CheckWinner()`: (TODO: incomplete—currently just toggles CurrentPlayer without checking win conditions).
  - `ResetJDV()`: clears the board and resets game state.

- **ui/jogoDaVelhaBoard.go** (125 lines)
  - `Build(a fyne.App)`: main UI constructor.
  - Creates a 3×3 grid of buttons bound to `g.PlayJDV(r, c)`.
  - Status label shows current player or win message ("X venceu", "O venceu", or "Empate").
  - "Reiniciar" and "Nova Partida" buttons both call `g.ResetJDV()` and refresh display.
  - Local `atualizar()` closure re-renders board buttons and updates status on each move.

Windows Build Setup Instructions — FIX THE BUILD ERROR
-------------------------------------------------------
**Root cause**: Fyne's OpenGL rendering on Windows needs a C compiler. Without one, the go-gl packages fail to compile.

**Fix Option A: Install MinGW-w64 (GCC for Windows)** ★ Recommended
1. Install MinGW-w64:
   - **MSYS2** (easiest): Download from [msys2.org](https://www.msys2.org)
     - Run installer, then in MSYS2 MinGW64 terminal: `pacman -S mingw-w64-x86_64-gcc`
     - Add to PATH: `C:\msys64\mingw64\bin`
   - **Chocolatey** (if admin): `choco install mingw -y`
   - **Manual**: Download [MinGW-w64 standalone](https://www.mingw-w64.org/downloads) and add to PATH.
2. Verify installation:
   ```powershell
   gcc --version
   ```
3. Build the project with CGO enabled:
   ```powershell
   $env:CGO_ENABLED=1; go build ./...
   ```

**Fix Option B: Install Visual Studio Community (MSVC compiler)**
1. Download and install [Visual Studio Community](https://visualstudio.microsoft.com/downloads).
2. Select workload "Desktop development with C++".
3. Restart your machine.
4. Build: `$env:CGO_ENABLED=1; go build ./...`

**Fix Option C: Cross-compile from Windows to Linux (if Windows compilation blocked)**
1. `$env:GOOS=linux; $env:GOARCH=amd64; $env:CGO_ENABLED=0; go build -o awesomeProject-linux ./...`
2. Deploy and run on Linux where C libraries are available.

**Verify build success**:
```powershell
$env:CGO_ENABLED=1; go build ./...
if ($?) { Write-Host "Build successful!" } else { Write-Host "Build failed" }
```

Developer workflows (commands)
------------------------------
- **Run**: `$env:CGO_ENABLED=1; go run main.go` (on Windows; macOS/Linux: `CGO_ENABLED=1 go run main.go`)
- **Build**: `$env:CGO_ENABLED=1; go build ./...`
- **Test**: `go test ./...` (no tests currently; add `main_test.go`, `jogos/*_test.go`)
- **Vet**: `go vet ./...` (checks for common errors)
- **Tidy**: `go mod tidy` (removes unused dependencies)
- **Debug**: Use GoLand IDE breakpoints, or CLI `dlv debug ./main.go`

Project-specific patterns and pitfalls
--------------------------------------
- **Incomplete win detection**: `CheckWinner()` in `jogos/jogoDaVelha.go` (lines 49–55) only toggles CurrentPlayer. Must check for three in a row (horizontal, vertical, diagonal) and set `GameOver=true`, `Winner="X"/"O"`. Also detect draw when board full.
- **UI refresh on every move**: `atualizar()` in `ui/jogoDaVelhaBoard.go` redraws all 9 buttons every move; acceptable for 3×3 but scale carefully for larger grids.
- **Buttons remain active after game over**: Buttons don't disable after win/draw (though `PlayJDV()` returns false). Consider calling `button.Disable()` in the UI.
- **"Reiniciar" vs "Nova Partida"**: Both buttons call `g.ResetJDV()` identically; they could have different behaviors (e.g., restart vs. go to menu).

Integration points and external dependencies
---------------------------------------------
- **fyne.io/fyne/v2**: GUI framework
  - Transitive dependencies: `github.com/go-gl/gl` (OpenGL), `github.com/go-gl/glfw/v3.3/glfw` (windowing)
  - **Windows**: Requires CGO + C compiler (MinGW or MSVC)
  - **macOS**: Typically works with system frameworks
  - **Linux**: Requires system OpenGL and GLFW dev packages: `apt install libgl1-mesa-dev libxcursor-dev libxrandr-dev libxinerama-dev libxi-dev libxext-dev`

Examples (concrete edits an agent might perform)
-----------------------------------------------
1. **Fix win detection** (highest priority):
   ```go
   // In jogos/jogoDaVelha.go, rewrite CheckWinner():
   func (v *JogoDaVelha) CheckWinner() {
       // Check all rows, cols, diagonals for v.CurrentPlayer
       // If found, set v.GameOver=true and v.Winner=v.CurrentPlayer
       // After toggling player, if board is full, set v.GameOver=true and v.Winner="Empate"
   }
   ```

2. **Disable buttons after game over**:
   ```go
   // In ui/jogoDaVelhaBoard.go, atualizar() function:
   if g.GameOver {
       for _, btn := range buttons {
           btn.Disable()
       }
   }
   ```

3. **Add unit tests**:
   - Create `jogos/jogoDaVelha_test.go`
   - Test win patterns, draw, edge cases

4. **Add a CLI mode** (optional, for users without C compiler):
   - Create `main-cli.go` that reads moves from stdin, prints ASCII board

What not to change without human confirmation
----------------------------------------------
- Module name (`awesomeProject` in `go.mod`) — affects import paths.
- Fyne version (may require C compiler version sync on Windows).
- Core game rules in `jogos/jogoDaVelha.go` — confirm intended behavior first.

Where to look next
-------------------
- If adding more games: new files in `jogos/` (e.g., `jogoDaMemoria.go`)
- If adding UI views: `ui/` package or create `ui/menu.go`
- If deploying: consider `fyne` command-line tool for packaging

Contact points for humans
--------------------------
- **Cannot install C compiler (admin blocked)**: Escalate for alternate build strategy (cross-compile, pre-built binaries, etc.)
- **Fyne version conflicts**: Changes to `fyne.io/fyne/v2` may require C compiler or system library updates.
- **Game logic changes**: Confirm desired win/draw rules before implementing.

