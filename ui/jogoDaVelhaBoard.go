package ui

import (
	"awesomeProject/jogos"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Build(a fyne.App) {

	g := jogos.NewJDV()

	w := a.NewWindow("Jogo da Velha")

	w.Resize(fyne.NewSize(420, 500))

	status := widget.NewLabel(
		"Jogador X",
	)

	grid := container.NewGridWithColumns(3)

	buttons := make([]*widget.Button, 9)

	var atualizar func()

	atualizar = func() {

		for i := 0; i < 9; i++ {

			r := i / 3
			c := i % 3

			buttons[i].SetText(
				g.Board[r][c],
			)
		}

		if g.GameOver {

			if g.Winner == "Empate" {

				status.SetText("Empate")

			} else {

				status.SetText(
					fmt.Sprintf(
						"%s venceu",
						g.Winner,
					),
				)
			}

			return
		}

		status.SetText(
			fmt.Sprintf(
				"Jogador %s",
				g.CurrentPlayer,
			),
		)
	}

	for i := 0; i < 9; i++ {

		index := i

		button := widget.NewButton(
			"",
			func() {

				r := index / 3

				c := index % 3

				g.PlayJDV(r, c)

				atualizar()
			},
		)

		buttons[i] = button

		grid.Add(button)
	}

	reiniciar := widget.NewButton(
		"Reiniciar",
		func() {

			g.ResetJDV()

			atualizar()
		},
	)

	novaPartida := widget.NewButton(
		"Nova Partida",
		func() {

			g.ResetJDV()

			atualizar()
		},
	)

	content := container.NewVBox(
		status,
		grid,
		reiniciar,
		novaPartida,
	)

	w.SetContent(content)

	atualizar()

	w.Show()
}
