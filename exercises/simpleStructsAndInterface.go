package main

import (
	"fmt"
	"math"
)

// Shape method sets to be implemented by different shapes
type Shape interface {
	area() float64
	perimeter() float64
}

// Circle field set need for the circle struct
type Circle struct {
	xDim, yDim, rRadius float64
}

// Rectangle field set need for the rectangle struct
type Rectangle struct {
	xDimOne, yDimOne, xDimTwo, yDimTwo float64
}

// Distance field set need for the Distance struct
type Distance struct {
	xDimOne, yDimOne, xDimTwo, yDimTwo float64
}

// coordinate distance calculation
func (d *Distance) distance() float64 {
	xDimDiff := d.xDimTwo - d.xDimOne                       // calculate the side base
	yDimDiff := d.yDimTwo - d.yDimOne                       // calulate the side
	return math.Sqrt(xDimDiff*xDimDiff + yDimDiff*yDimDiff) // calculate the hypothenus i.s. the distance the two points
}

// area method sets
func (r *Rectangle) area() float64 {
	length := Distance{r.xDimOne, r.yDimOne, r.xDimTwo, r.yDimOne} // lenght is calculated by making y constant but x variable
	width := Distance{r.xDimOne, r.yDimOne, r.xDimOne, r.yDimTwo}  // width is calculated by making x contant but y variable
	return length.distance() * width.distance()
}

func (c *Circle) area() float64 {
	return math.Pi * c.rRadius * c.rRadius
}

// perimeter method sets
func (r *Rectangle) perimeter() float64 {
	length := Distance{r.xDimOne, r.yDimOne, r.xDimTwo, r.yDimOne} // lenght is calculated by making y constant but x variable
	width := Distance{r.xDimOne, r.yDimOne, r.xDimOne, r.yDimTwo}  // width is calculated by making x contant but y variable
	return 2 * (length.distance() + width.distance())
}

func (c *Circle) perimeter() float64 {
	return 2 * math.Pi * c.rRadius
}

// Shape interface being used as a argument [allows us to extract aggregated method values]
func aggregrateMethodValues(shapes ...Shape) (shapesTotalArea float64, shapesTotalPerimeter float64) {
	for _, s := range shapes {
		shapesTotalArea += s.area()
		shapesTotalPerimeter += s.perimeter()
	}
	return shapesTotalArea, shapesTotalPerimeter
}

func main() {

	var shapesTotalAreaMain float64
	var shapesTotalPerimeterMain float64

	r := Rectangle{0, 0, 10, 10}
	c := Circle{0, 0, 2}

	fmt.Println("Rectangle Area:", r.area())
	fmt.Println("Rectangle Perimeter:", r.perimeter())
	fmt.Println(" ")
	fmt.Println("==== ==== ====")
	fmt.Println(" ")
	fmt.Println("Circle Radius:", c.rRadius)
	fmt.Println("Circle Area:", c.area())
	fmt.Println("Circle Perimeter:", c.perimeter())
	fmt.Println(" ")
	fmt.Println("==== ==== ====")
	fmt.Println(" ")

	// call
	shapesTotalAreaMain, shapesTotalPerimeterMain = aggregrateMethodValues(&r, &c)
	fmt.Println("Aggregated Area across all shapes is:", shapesTotalAreaMain)
	fmt.Println("Aggregated Perimeter across all shapes is:", shapesTotalPerimeterMain)
	fmt.Println(" ")
	fmt.Println("==== ==== ====")
	fmt.Println(" ")
}
