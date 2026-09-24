package main

import "fmt"

type Clothing struct{
	name string
	color string 
	state bool
}

type Basket struct{
	name string
	color string
	content []Clothing
}

func displayBasket(d *Basket,c *Clothing){
	//afficher le contenant du panier, le vetement,"color",etat, afficher msg si panier vide 
	if d.content == [0]Clothing{
		fmt.Println("Panier vide")
	}
	fmt.Println("===="+ d.name + d.color + "====")
	fmt.Println(c.name+c.color)
	fmt.Println(d.color+c.color)
	fmt.Print(c.state)
}
func  addToBasket(c *Clothing,d *Basket)  {
	//add un vetement sale et du meme type que la corbeille, afficher message dans chaque cas de figure
	if c.state == false && c.color == d.color{
		append()
		fmt.Println("Un vêtement a été ajouté à" d.name + ":" + c.name)
	}else {
		fmt.Println("Oupss ! Ajout impossible")
	}
}
func cleanBasket(){
	//changement etat vetement sale to propre, si panier vide afficher un msg, else lavage finis
}
func emptyCleanLaundry(){
	//remove les vetement, affiche les vetements retiré, si pas propre ou panier vider afficher message erreur
}

func main(){
	basketColor := Basket{"Panier couleur","color",[]Clothing{}}
	basketWhite := Basket{"Panier blanch","white",[]Clothing{}}
	basketBlack := Basket{"Panier noir","black",[]Clothing{}}

	clot01 := Clothing{"t-shirt de couleur", "color", false}
	clot02 := Clothing{"short de couleur", "color", false}
}