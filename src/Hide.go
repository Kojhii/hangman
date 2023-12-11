package main

//fonction qui va nous permettre de "cacher le mot" , en créant un tableau de tiré equivalant a la longueur du mot rentré en parametre

func Hide(len int) []string {

	//creation du tableau
	dash := []string{}
	
	//boucle de taille du mot , pour chaque passage le tableau prend un tiret en plus
	for i := 0; i < len; i++ {
		dash = append(dash, "-")
	}
	return dash //et pas daesh haha
}