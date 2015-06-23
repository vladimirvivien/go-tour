package main

import "fmt"


func main() {
    type Point struct {
    	X, Y int
    }

    // creates map pointers using 
    points := make(map[string]Point)

    points["top"] = Point{1,2}

    fmt.Printf ("Point top %v of type %T\n", points["top"], points)

    // create map literals
    // wath out for (annoying) trailing commas, ????
    points2 := map[string]Point  {
    	"Point1" : Point{
    		12, 14,
    	},
    	// Note that type name optional
    	"Point2" : {X:12, Y:29},
    }

    fmt.Printf ("My points %v\n", points2)

    v, ok := points2["Point2"]

    if ok {
    	fmt.Println ("Maps contains Point2 with value ", v)
    }else{
    	fmt.Println ("No key found.") 
    }
}