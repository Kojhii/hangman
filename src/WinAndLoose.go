package main

import (
	"fmt"
	"time"
	"github.com/TheZoraiz/ascii-image-converter/aic_package"
)

//fonction quand on gagné avec le mode fr
func WINfr(){
	fmt.Print("\033[H\033[2J")

	//affichage du titre avec écrit "win"
	logowin()
	
	fmt.Println("                                                      		              [bravo champion]")
	time.Sleep(1 * time.Second)
	fmt.Println("\n\n ")
	
	
	
}

//fonction quand on perd avec le mode fr
func LOOSEFR(p Hangman){
	fmt.Print("\033[H\033[2J")

	//affichage du titre avec écrit "looser"
	logolooser()
	fmt.Println("                                                          [Petit Kabyle va]")
	
	
	//affichage de la photo en cas de defaite grace a une librairie preablement importé , on lui attribue des headers de taille /couleur etc...
	fmt.Println("\n\n ")
	
	filePath := "../image/looser.webp"
	flags := aic_package.DefaultFlags()
	flags.Dimensions = []int{160, 40}
	flags.Colored = true
	flags.Braille = true
	flags.CustomMap = " .-=+#@"
	flags.SaveBackgroundColor = [4]int{50, 50, 50, 100}
	asciiArt, err := aic_package.Convert(filePath, flags)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%v\n", asciiArt)
	
}

//Fonction si on gagne avec le mode anglais
func WINEN(){
	fmt.Print("\033[H\033[2J")
	
	//affichage du logo avec écrit "win"
	logowin()
	
	fmt.Println("                                                  		                     [Good job bro]")
	time.Sleep(4 * time.Second)
	fmt.Print("\033[H\033[2J")

}
//Fonction si on perd avec le mode anglais
func LOOSEEN(p Hangman){
	fmt.Print("\033[H\033[2J")

	//affichage du titre avec écrit "looser"

	logolooser()
	fmt.Println("                                                                    [Hahaha fucking noob , you piece of shit]")
	
	//affichage d'une image en cas de defaite avec des parametre tel que la taille , la couleur ou bien le mode braille.
	
	filePath := "../image/looser.webp"
	flags := aic_package.DefaultFlags()
	flags.Dimensions = []int{160, 37}
	flags.Colored = true
	flags.Braille = true
	flags.CustomMap = " .-=+#@"
	flags.SaveBackgroundColor = [4]int{50, 50, 50, 100}
	asciiArt, err := aic_package.Convert(filePath, flags)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%v\n", asciiArt)
	
}