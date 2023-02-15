package main

import "fmt"

func main() {
	fmt.Println("Starting method receivers topic....")
	d := Dimension{5, 4, 3}
	fmt.Println("############################")
	fmt.Println("Area: ", d.AreaPointer())
	fmt.Printf("AreaPointer: %v \n", d)
	fmt.Println("Volume: ", d.VolumePointer())
	fmt.Printf("VolumePointer: %v \n", d)
	fmt.Println("############################")
	fmt.Println("Area: ", d.Area())
	fmt.Printf("Area: %v \n", d)
	fmt.Println("Volume: ", d.Volume())
	fmt.Printf("Volume: %v \n", d)
}

type Dimension struct {
	length  int
	breadth int
	height  int
}

func (d Dimension) Area() int {
	d.length = 10
	return d.length * d.breadth
}

func (d Dimension) Volume() int {
	d.length = 20
	return d.length * d.breadth * d.height
}

func (d *Dimension) AreaPointer() int {
	d.length = 6
	return d.length * d.breadth
}

func (d *Dimension) VolumePointer() int {
	d.length = 7
	return d.length * d.breadth * d.height
}
