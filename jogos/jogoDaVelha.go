package jogos

type JogoDaVelha struct {
	Board [3][3]string

	CurrentPlayer string

	GameOver bool

	Winner string
}

func NewJDV() *JogoDaVelha {

	return &JogoDaVelha{
		CurrentPlayer: "X",
	}
}

func (v *JogoDaVelha) PlayJDV(r int, c int) bool {
	if v.GameOver {
		return false
	}
	if v.Board[r][c] != "" {
		return false
	}

	v.Board[r][c] = v.CurrentPlayer

	v.CheckWinner()

	if !v.GameOver {

		if v.CurrentPlayer == "X" {
			v.CurrentPlayer = "O"
		} else {
			v.CurrentPlayer = "X"
		}
	}
	return true
}

func (v *JogoDaVelha) ResetJDV() {
	*v = JogoDaVelha{
		CurrentPlayer: "X",
	}
}

func (v *JogoDaVelha) CheckWinner() {
	if v.CurrentPlayer == "X" {
		v.CurrentPlayer = "O"
	} else {
		v.CurrentPlayer = "X"
	}
}
