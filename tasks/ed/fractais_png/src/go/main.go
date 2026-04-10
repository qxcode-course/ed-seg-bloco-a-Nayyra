package main

import (
	"fmt"
	"math/rand/v2"
)

// func randInt(min, max int) int {
// 	return min + rand.Intn(max-min+1)
// // }

func estrela(pen *Pen, dis float64){

	r := rand.Float64() * 250
	g := rand.Float64() * 250
	b := rand.Float64() * 250

	if dis < 5{
		return
	}
	pen.Left(90)
	pen.SetRGB(r, g, b)
	pen.SetLineWidth(4)
	pen.Walk(dis)
	estrela(pen, dis - 10)

}

func main() {
	pen := NewPen(800, 800)
	pen.SetPosition(700, 700)
	pen.SetHeading(0)
	estrela(pen, 600)
	pen.SavePNG("tree.png")
	fmt.Println("PNG file created successfully.")
}
