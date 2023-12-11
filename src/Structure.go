package main

//structure du hangman avec nos deux valeurs primordiale , le nombre d'essaie et le nombre de bonne lettre
type Hangman struct {
	//creation des deux variable de type int
	Try int
	GoodLetter int
}

//fonction pour permettre d'iniatiliser les valeurs
func (p *Hangman) Init(Try int, GoodLetter int) {
	p.Try = Try
	p.GoodLetter = GoodLetter
}
