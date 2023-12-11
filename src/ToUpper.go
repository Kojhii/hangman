package main

//fonction permettant de transformer les imput du joueur en majuscule si il ne le sont pas deja
func ToUpper(s string) string {
	//on crée un tableau de rune de taille s nommé c
	c := []rune(s)
	//pour la longueur de notre tableau c
	for i := range c {
		//si entre a et z , sois une lettre minuscule 
		if c[i] >= 'a' && c[i] <= 'z' {
			// on lui retire -32 , cela permet de passer de miniscule a majuscule en se basant sur la table ascii
			c[i] = c[i] - 32
		}
	}
	//on retourne le tableau de rune casté en string
	return string(c)
}

