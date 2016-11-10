package main

import (
    "fmt"
    app "./app"
)

func main() {
    fmt.Println(app.DetermineTriangleType(-1, 2, 3))
    fmt.Println(app.DetermineTriangleType(1, 2, 3))
    fmt.Println(app.DetermineTriangleType(1, 2, 4))
    fmt.Println(app.DetermineTriangleType(3, 4, 5))
    fmt.Println(app.DetermineTriangleType(4, 4, 5))
    fmt.Println(app.DetermineTriangleType(5, 5, 5))
}