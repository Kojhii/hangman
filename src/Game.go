package main

import (
	"fmt"
	"time"
	"math/rand"
)

func GameFrench(p *Hangman) {
	fmt.Print("\033[H\033[2J")

	
	// on va attribuer la valeur false a la variable "good", puis on crée la liste de mot
	good := false
	worldliste := []string{"TEST","CHOCOLAT","ABRICOT","ABDELRENDMOIMONARGENT","OUI","NON","POMME","MUR","GUIMAUVE","CHAT","MAISON","SOLEIL","ÉTOILE","POMME","JARDIN","RIRE","LIVRE","CAFÉ","MUSIQUE","FLEUR","VÉLO","MONTAGNE","RIVIÈRE","PLUIE","NUAGE","FORÊT","ROUTE","OISEAU","MER","PLAGE","SOURIRE","CLÉ","FENÊTRE","ÉCOLE","VÉRITÉ","SILENCE","CADEAU","COULEUR","NUIT","JOUR","HEURE","ÉNERGIE","ESPOIR","LIBERTÉ","RÊVE","VOYAGE","AVENTURE","MYSTÈRE","COURAGE","PASSION","ÉQUILIBRE","HARMONIE","BEAUTÉ","MAGIE","SAGESSE","ÉMOTION","POÉSIE","TENDRESSE","BONHEUR","TRÉSOR","ÉTOILE FILANTE","PARFUM","TENDANCE","FANTAISIE","SAVANE","CASCADE","PARAPLUIE","CÂLIN","BALADE","GÂTEAU","MIROIR","RÊVERIE","COQUELICOT","ÉCLIPSE","PÉTALE","PLUME","MONTGOLFIÈRE","MÉDITATION","ÉCLAIR","GLACIER","BIBLIOTHÈQUE","BISOU","ÉCLAT","DOUCEUR","ÉTINCELLE","CLAIR DE LUNE","GAZOUILLIS","FESTIN","ÉVASION","SUCCULENT","SÉRÉNITÉ","COMPLICITÉ","MÉLODIE","BULLE","DANSE","RAFALE","ÉCORCE","TOURNESOL","ARC-EN-CIEL","AQUARELLE","HARMONICA","AILE","SYMPHONIE","ZÉNITUDE"}
	
	//on va crée un nombre aleatoire grace a la fonction rand (de la librarie math), puis on choisis un nombre aleatoire de 0 a 109 (la taille de la liste de mot)
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	randomIndex := r1.Intn(109)
	
	//notre mot est alors egal a l'element aleatoire de la liste de mot , imaginons que le nombre est 15 , et bien notre mot sera le 15 mot de la liste,wordliste[15]
	world := worldliste[randomIndex]

	//on crée la variable lenworld egal a la longueur du mot 
	lenworld := len(string(world))
	//on utilise la fonction hide pour crée un tableau de tiret egal a la longueur du mot
	tiret := Hide(lenworld)
	//boucle tant que le joueur a des essaie
	for i := 2; i <= p.Try; {
		fmt.Print("\033[H\033[2J")
		//print du logo et affichage du texte "rentrer une lettre"
		logo()
		fmt.Println("                                                                               [Rentre une lettre]")

		good = false
		//on affiche le mot en tirét
		fmt.Println("\n\n\n                                                                            ",tiret)
		//on affichage le hangman
		displayhangman(*p)
		//on scan l'imput du joueur
		var imput string
		fmt.Scan(&imput)

		//on transforme l'imput du joueur en majuscule obligatoirement
		imput = ToUpper(imput)
		
		//nouvelle boucle de la longueur du mot
		for j := 0; j < len(world); j++ {
			//si la lettre rentré par le joueur est egal a la lettre que nous somme entrain de parcourir
			if string(imput) == string(world[j]) {
				//le mot en tiret est egal a la lettre rentré par le joueur
				tiret[j] = imput
				//il gagne une bonne lettre en plus , on place dans la valeur true a good
				p.GoodLetter +=1
				good = true
				//si le nombre de bonne lettre du joueur est egal au nombre de lettre du mot , il gagne !
				if p.GoodLetter == len(world) {
					WINfr()
					
					
				}
			} 
			
		}
		//si good n'est pas true , donc le joueur n'a pas trouvé de bonne lettre , alors il perd un essaie
		if !good {
			p.Try -=1
		}
	}
	//si on sort de la boucle car le joueur n'a plus d'essaie il perd
	LOOSEFR(*p)
}


//aucune difference avec le mode français , juste la liste de mot est anglaise et le texte aussi
func GameEnglish(p *Hangman) {
	fmt.Print("\033[H\033[2J")

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	worldliste := []string{"SWORD", "APPLE", "ROCK","CAT","HOUSE","SUN","STAR","APPLE","GARDEN","LAUGH","BOOK","COFFEE","MUSIC","FLOWER","BIKE","MOUNTAIN","RIVER","RAIN","CLOUD","FOREST","ROAD","BIRD","SEA","BEACH","SMILE","KEY","WINDOW","SCHOOL","TRUTH","SILENCE","GIFT","COLOR","NIGHT","DAY","HOUR","ENERGY","HOPE","FREEDOM","DREAM","TRIP","ADVENTURE","MYSTERY","COURAGE","PASSION","BALANCE","HARMONY","BEAUTY","MAGIC","WISDOM","EMOTION","POETRY","TENDERNESS","HAPPINESS","TREASURE","SHOOTING STAR","PERFUME","TREND","FANTASY","SAVANNAH","WATERFALL","UMBRELLA","HUG","WALK","CAKE","MIRROR","REVERIE","POPPY","ECLIPSE","PETAL","FEATHER","HOT AIR BALLOON","MEDITATION","FLASH","GLACIER","LIBRARY","KISS","GLARE","GENTLENESS","SPARKLE","MOONLIGHT","TWITTER","FEAST","ESCAPE","DELICIOUS","SERENITY","COMPLICITY","MELODY","BUBBLE","DANCE","GUST","BARK","SUNFLOWER","RAINBOW","WATERCOLOR","HARMONICA","WING","SYMPHONY","ZENITUDE"}
	
	good := false
	
	randomIndex := r1.Intn(103)
	
	Wordtofind := []string{}
	Wordtofind = append(Wordtofind, worldliste[randomIndex])
	
	lenworld := len(string(Wordtofind[0]))
	tiret := Hide(lenworld)
	
	test := []rune(Wordtofind[0])
	
	for i := 2; i <= p.Try; {
	
		fmt.Print("\033[H\033[2J")
		logo()
	
		fmt.Println("                                                                               [ENTER A LETTER]")
		good = false

		fmt.Println("\n\n\n                                                                            ",tiret)

		displayhangman(*p)

		var imput string
		fmt.Scan(&imput)

		imput = ToUpper(imput)

		for j := 0; j < len(test); j++ {
			if string(imput) == string(test[j]) {
				tiret[j] = imput
				p.GoodLetter +=1
				good = true
				if p.GoodLetter == len(test) {
					WINEN()
					
					
				}
			} 
			
		}
		if !good {
			p.Try -=1
		}
	}
	LOOSEEN(*p)
}