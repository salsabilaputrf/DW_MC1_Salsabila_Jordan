package main

import "fmt"

type Profile struct {
	Name string
	Health int
	Power int
	Exp int
} 

func main(){

	profile := MakeProfile("Goku")
	fmt.Println("Name : ", profile.Name)
	fmt.Println("Health : ", profile.Health)
	fmt.Println("Power : ", profile.Power)
	fmt.Println("Exp : ", profile.Exp)
	fmt.Println("===HEROES POWER UP===")
	powerUp := PowerUp(profile, 2)
	fmt.Println("Name : ", powerUp.Name)
	fmt.Println("Health : ", powerUp.Health)
	fmt.Println("Power : ", powerUp.Power)
	fmt.Println("Exp : ", powerUp.Exp)

}

func MakeProfile(name string) Profile{
	data := Profile{}
	if name == "Sasuke" {
		
		data = Profile{
			Name: name,
			Health: 200,
			Power: 100,
			Exp: 0,
		}
		
	}else if name == "Goku"{
		data = Profile{
			Name: "Goku",
			Health: 400,
			Power: 300,
			Exp: 100,
		}
	}else if name == "Naruto"{
		data = Profile{
			Name: "Naruto",
			Health: 150,
			Power: 200,
			Exp: 50,
		}
	}


	return data
}

func PowerUp(data Profile, addPower int) Profile{
	totalHealth := data.Health + (data.Health * addPower)
	totalPower := data.Power + (data.Power * addPower)
	totalExp := data.Exp + (data.Exp * addPower)

	newData := Profile{
		Name: data.Name,
		Health: totalHealth,
		Power: totalPower,
		Exp: totalExp,
	}

	return newData
}