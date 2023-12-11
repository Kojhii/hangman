package main

import "fmt"

// fonction qui va nous permettre de choisir le mode de jeux , ici on en a 3 : Français , anglais , et algerien (inspiré de l'algerie mais en français)
func ChoiceGame() {
	//on crée au passage un objet p de type hangman et lui attribue des valeurs 
	var p Hangman
	p.Init(8,0)

	//print du logo et affichage du tableau a choix
	logo()
	fmt.Println("\n\n                                                          _____________________________________________________\n                                                         |                  French  (press 1)                 |\n                                                         |                  English (press 2)                 |\n                                                         |____________________________________________________|\n\n\n   ")

	//scan de l'imput du joueur et redirection en fonction de son choix
	var imput string
	fmt.Scan(&imput)

	switch imput {
	case "1":
		fmt.Print("\033[H\033[2J")
		GameFrench(&p)

	case "2":
		fmt.Print("\033[H\033[2J")
		GameEnglish(&p)
	case "dzsupremacy":
		fmt.Print("\033[H\033[2J")
		GameAlgerian(&p)
	default:
		fmt.Print("\033[H\033[2J")
		ChoiceGame()
	}
}





























