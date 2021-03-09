package main

import "fmt"

// @desc go 接口类型判断

type Pet interface {
	Call() string
}

type Kitten struct {
}

func (k *Kitten) Taps() {
	fmt.Println("There can't be anything on the table")
}

type Puppy struct {
}

func (p *Puppy) WhoIsGoodBoy() {
	fmt.Println("Who is Good Boy? Is Kitten!!")
}

// Kitten implement Pet
func (k *Kitten) Call() string {
	return "Miao~"
}

// Puppy implement Pet
func (p *Puppy) Call() string {
	return "Wow!"
}

func main() {
	cat := new(Kitten)
	dog := new(Puppy)

	// 接口类型判断、类型转换：接口转化为实现类
	var p Pet = cat
	kitten, ok := p.(*Kitten)
	if ok {
		fmt.Println("is Kitten", kitten.Call())
		kitten.Taps()
	}

	// 接口类型判断
	whatFuckIsIt(cat)
	whatFuckIsIt2(dog)

}

// 接口类型判断，使用类型断言
func whatFuckIsIt(pet Pet) {
	if _, ok := pet.(*Kitten); ok {
		fmt.Println("is Kitten")
	} else if _, ok := pet.(*Puppy); ok {
		fmt.Println("is Puppy")
	} else {
		fmt.Println("who care?")
	}
}

// 接口类型判断，使用 type-switch
func whatFuckIsIt2(pet Pet) {
	switch pet.(type) {
	case *Kitten:
		fmt.Println("is cat")
		k, _ := pet.(*Kitten)
		k.Taps()
	case *Puppy:
		fmt.Println("is dog")
		d, _ := pet.(*Puppy)
		d.WhoIsGoodBoy()
	default:
		fmt.Println("who care")
	}
}
