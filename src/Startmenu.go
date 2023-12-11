package main 

import "fmt"
//premier menu qui permet de laisser le choix au joueur de jouer ou non , a pour but un coté humoristique dans le cas ou le joeuur choisi non 
func startmenu() {
	//affichage logo et tableau
	logo()
	fmt.Println("\n\n                                                          _____________________________________________________\n                                                         |               Play hangman ? (press yes)           |\n                                                         |____________________________________________________|\n\n\n   ")

	//Scan de l'imput du joueur
	var imput string
	fmt.Scan(&imput)

	//si imput = no , on lance l"affichage du logo avec écrit "looser", si il est egal a "yes" on lance le choix du mode de jeux
	switch imput {
	case "no":
		fmt.Print("\033[H\033[2J")
		logolooser()
	case "yes":
		fmt.Print("\033[H\033[2J")
		ChoiceGame()
	//si le joeur rentre un choix autre que yes/no , il est renvoyé a cette fonction permettant de refaire un choix
	default:
		fmt.Print("\033[H\033[2J")
		main()
	}
}

