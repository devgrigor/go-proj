package main

import (
	"fmt"
	"math"
	"time"
	"math/rand"
)

type Circle struct {
	x float64
	y float64
	radius float64
}

type Rect struct {
	leftSide, rightSide float64
}

func changeInt(a *int) {
	*a = 8;
}

type Shape interface {
	area() float64
}

func circleArea(c *Circle) float64 {
	return math.Pi * c.radius*c.radius
}

func (c *Circle) area() float64 {
	return math.Pi * c.radius*c.radius
}

func (r *Rect) area() float64 {
	return r.leftSide * r.rightSide
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s:= range shapes {
		area += s.area()
	}

	return area
}

type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hi ma name is ", p.Name)
}

type Android struct {
	Person
	Model string	
}

func asynch() {
	for i := 0; i <= 15; i++ {
		fmt.Println("Some ind", i)
		seconds := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * seconds)
	}
}

func main() {
	// Type is not defined but is comming from right side statement
	a := 1
	changeInt(&a)
	fmt.Println(a)

	// You can't write Android{Name: 'something'} instead of new(Android)
	andr := new(Android)	
	andr.Name = "asdsadadd"
	andr.Talk()

	var b float64
	go asynch();
	fmt.Println("this supposed to be after asynch")
	fmt.Scanf("%f", &b)

	
	// emptyC := new(Circle)
	c := Circle{x: 0, y: 1, radius: 15}
	fmt.Println(c.x, c.y, c.radius, circleArea(&c), c.area())
	r := Rect{leftSide: 40, rightSide: 40}

	fmt.Println(totalArea(&c, &r))
	// Associated arrays - maps
	x := make(map[string]int)
	x["key"] = 10
	fmt.Println(x)

	arr := []float64{1,2,3,4,5,6}
	
	// array slice
	slice := arr[2:5]
	fmt.Println(slice)

	arr[4] = 84
	fmt.Println(len(arr))

	for _, value := range arr {
		fmt.Println(value)
	}

	var (
		a1 = 8
		b1 = 8
		c1 = 9
	)

	i := 0

	for i <=18 {
		fmt.Println(i)
		switch i {
			case 0: {
				fmt.Println("zero")
			}
			default: fmt.Println("Fuck you")
		}

		i++
	}

	for i1 := 1; i1 <= 10; i1++ {
		if i1 % 2 == 0 {
			fmt.Println("odd", i1)	
		} else if i1 % 3 == 0 {
			fmt.Println("diivisible by 3", i1)
		} else {
			fmt.Println("even", i1)
		}
		
	}

    fmt.Println("hello, world", b*8)
    fmt.Println("hello, world :", a1+b1+c1+a)
    fmt.Scanf("%f", &b)
}