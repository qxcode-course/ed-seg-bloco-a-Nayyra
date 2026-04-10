package main

import (
	"fmt"
	"math/rand"
)

func randInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func arvore(pen *Pen, dis float64){
	angle := 30.0

	if dis < 10{
		pen.SetRGB(0, 0, 0)
		return
	}

	pen.Walk(dis)
	pen.Left(angle)
	arvore(pen, dis - 5)
	pen.Right(2*angle)
	arvore(pen, dis - 5)
	pen.Left(angle)
	pen.Walk(-dis)

	
}

func main() {
	pen := NewPen(800, 600)
	pen.SetPosition(540, 300)
	pen.SetHeading(9)
	arvore(pen, 45.0)
	pen.Left(450)
	arvore(pen, 45.0)
	pen.Left(490)
	pen.Walk(170)
	pen.Right(-630)
	pen.Walk(170)
	pen.SetHeading(90)
	arvore(pen, 45.0)
	pen.Left(450)
	arvore(pen, 45.0)
	pen.Left(630)
	
	
	pen.SavePNG("tree.png")
	fmt.Println("PNG file created successfully.")
}
