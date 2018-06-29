package main

import ( "golang.org/x/tour/pic"
	"os"
	//	"strings"
	"fmt"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Pic(dx, dy int) [][]uint8 {
	dy =256
	dx =256
	image := make([][]uint8,dy)
	sdy :=make([]uint8,dy)
	//s1 := rand.NewSource(time.Now().UnixNano())
	for x :=0; x < dx; x++{
		for y :=0 ; y < dy; y++{
			//r1 := rand.New(s1)
			//sdy[y] = uint8(r1.Intn(255))
			sdy[y] = uint8((x+y)/2)
		}
		image[x] = sdy
	}
	//fmt.Printf("%v",image)
	return image
}



func main() {
	mypic := pic.Show(Pic)
	f, err := os.Create("image.png")
	check(err)
	defer f.Close()
	//d2 := strings.Replace(mypic,"IMAGE:","",1)
	n2, err := f.Write(mypic)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)
}

