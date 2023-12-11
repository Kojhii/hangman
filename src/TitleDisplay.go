package main

import "github.com/common-nighthawk/go-figure"

//ensemble de fonction permettant d'afficher des titres grace a une librairie preablement importée
func logo() {
	myFigure := figure.NewFigure("                            HANGMAN ", "", true)
	myFigure.Print()
}

func logolooser() {
	myFigure := figure.NewFigure("                        LOOSER ", "", true)
	myFigure.Print()
}

func logowin() {
	myFigure := figure.NewFigure("                                WINNER ", "", true)
	myFigure.Print()
}
