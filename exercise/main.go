package exercise

import (
	"fmt"
)

type Item struct {
	Name string
	Type string
}

type Player struct {
	Name      string
	Inventory []Item
}

func (p *Player) PickUpItem(item Item) {
	p.Inventory = append(p.Inventory, item)
}

func (p *Player) DropItem(itemName string) {

	// index := -1
	// for i, item := range p.Inventory {
	// 	if item.Name == itemName {
	// 		index = i
	// 		break
	// 	}
	// }
	// if index != -1 {
	// 	p.Inventory = slices.Delete(p.Inventory, index, index+1)
	// }
	fmt.Println("Inventory before:", p.Inventory)

	// newItemArray := p.Inventory
	for index, item := range p.Inventory {
		if item.Name != itemName {
			p.Inventory = append(p.Inventory[:index], p.Inventory[index+1:]...)
			fmt.Printf("%s dropped %s\n", p.Name, itemName)
			return
		}

	}
	// p.Inventory = newItemArray
	fmt.Println("new Inventory:", p.Inventory)
}

// Use an item (if potion, remove it after use)
func (p *Player) UseItem(itemName string) {
	switch itemName {
	case "potion":
		fmt.Println("Used potion!")
		p.DropItem("potion")
		fmt.Println(p)
	}
}
