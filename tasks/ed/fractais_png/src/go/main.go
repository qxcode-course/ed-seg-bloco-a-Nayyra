package main

import (
	"fmt"
	"math/rand"
)

func randInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func estrela(pen *Pen, dis float64){

	if dis < 2{
		return
	}

	for range 5{
		pen.Walk(dis)
		estrela(pen, dis * 0.4)
		pen.Walk(-dis)
		pen.Left(72)
	}
}

func main() {
	pen := NewPen(800, 800)
	pen.SetPosition(400, 400)
	pen.SetHeading(50)
	estrela(pen, 150)
	pen.SavePNG("tree.png")
	fmt.Println("PNG file created successfully.")
}
