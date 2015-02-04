package main
import ("fmt"; "math")

type Circle struct{
    x, y, r float64
}

func (c *Circle) perimeter() float64{
    return math.Pi * c.r * 2
}

type Rectangle struct {
    l, w float64
}

func (r *Rectangle) perimeter() float64 {
  return (r.l * 2) + (r.w * 2)
}

type Shape interface {
    perimeter() float64
}

func main() {
    c := Circle{0, 0, 1}
    r := Rectangle{2, 4}
    fmt.Println(c.perimeter())
    fmt.Println(r.perimeter())
    fmt.Println(totalPerimeter(&c, &r))
}

func totalPerimeter(shapes ...Shape) float64 {
    var perimeter float64
    for _, s := range shapes {
        perimeter += s.perimeter()
    }
    return perimeter
}
