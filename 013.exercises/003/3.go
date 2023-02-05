package main

import "fmt"

type vehicle struct {
	// doors string
	doors int
	color string
}

type truck struct {
	vehicle
	fourWhell bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	t1 := truck{
		vehicle: vehicle{
			// doors: "four",
			doors: 4,
			color: "brown",
		},
		fourWhell: true,
	}
	s1 := sedan{
		vehicle: vehicle{
			// doors: "four",
			doors: 4,
			color: "gray",
		},
		luxury: false,
	}

	fmt.Println("truck1")
	fmt.Println(t1.doors, t1.color, t1.fourWhell)

	fmt.Println("sedan1")
	fmt.Println(s1.doors, s1.color, s1.luxury)

}
