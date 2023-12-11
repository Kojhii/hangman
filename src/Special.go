package main

import (
	"fmt"
	"time"
	"math/rand"
	"github.com/TheZoraiz/ascii-image-converter/aic_package"

)

//fonction identique a la fonction GameFrench() et GameEnglish() , se referer a elle pour comprendre le code
//seule la liste de mot et l'image en cas de defaite change
func GameAlgerian(p *Hangman) {
	fmt.Print("\033[H\033[2J")

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	
	worldliste := []string{"ALGÉRIE","CASBAH","SAHARA","CITRON","CAMELIA","HARIRA","BLED","MÉDITERRANÉE","KABYLIE","HARKI","DÉSERT","ORAN","CONSTANTINE","ALGER","TIMGAD","COUSCOUS","MAGHREB","BAGUETTE","MAQUIS","ZIRCON","TASSILI","SID-IBOU-SAÏD","CARTHAGE","CÈDRE","TAGINE","JASMINE","MOSQUÉE","CHAMEAU","FÊTE","CINÉMA ALGÉRIEN","RAÏ","AMAZIGH","TLEMCEN","ANNABA","DJELFA","TIARET","GUELMA","TAMAZIGHT","COLOMBE","TRAVAILLEUR","DIASPORA","DRAPEAU","FLN","FRATERNITÉ","GÉRANIUM","HORIZON","IMZAD","JARDIN","KALAA-DES-BÉNI-HAMDANE","LION","MAGRÉBINE","NUIT","OLIVIER","PAVILLON-NOIR","QUARTZ","RÉSISTANCE","SABLE","TRADITION","USM-ALGER","VOYOU","WILAYA","XÉNON","YOUYOU","ZAOUÏA","ASSIA-DJEBAR","BELLE-VIE","CHAOUIA","DALILA","EL-DJAZAIR","FANTAISIE","GOURAYA","HAMMAM","ISLAM","JUGURTHA","KHMISSA","LOTFI","MAURITANIE","NADOR","ORACLE","PORTE","QUARTZITE","RÉVOLUTION","SOUK-AHRAS","TÉMOIGNAGE","UNIVERSITÉ","VALLÉE","WADI","XÉNOPHOBE","YAKOUT","ZÉNITH","AGADIR","BOUCHEROUITE","CHOTT","DROMADAIRE","EL-BIAR","FONTAINE","GRANADA","HAUTEUR","ISTIKLAL","JOIE","KELT","LAKHDAR","MÉDAILLE","NACRE","OASIS","PÉTROLE","QUINCAILLERIE","ROSEAU","SIDI-BEL-ABBÈS","TAFSUT","URANUS","VAGUE","WARGLA","XÉROPHYTE","YENNAYER","ZÉPHYR"}
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
					WinAlgerian(*p)
				}
			} 
			
		}
		if !good {
			p.Try -=1
		}
	}
	LOOSEEN(*p)
}

//fonction permettant d'afficher une image de riyad mahrez en cas de victoire du joueur , ainsi qu'un texte le rendant fier de lui
func WinAlgerian(p Hangman){
	fmt.Print("\033[H\033[2J")

	logowin()
	fmt.Println("                                                                    [Tu es la fierté de l'algerie , bravo ]")

	filePath := "../image/riyad.jpg"
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