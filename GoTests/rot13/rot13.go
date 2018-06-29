package main

import (
	"io"
	"os"
	"strings"
//	"fmt"
	"fmt"
)

type rot13Reader struct {
	r io.Reader
}

func (ro rot13Reader) Read([]byte) (int, error) {
	a_list := []string{"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z"}
	b_list := []string{"N","O","P","Q","R","S","T","U","V","W","X","Y","Z","A","B","C","D","E","F","G","H","I,","J","K","L","M"}

	p := make([]byte,24)
	n, err := ro.r.Read(p)

	new_s := make([]string, n)

	///fmt.Printf("Type: %T %T %T ",a_list, b_list, s)
	for i:=0; i<len(p);i++{
		for h:=0; h<len(a_list);h++{
			if string(p[i]) == a_list[h]{
				//println(b_list[h])
				//p[i] = byte(b_list[h])
				new_s[i] = b_list[h]
			}
			//println(new_s[i] ,b_list[i])
		}
	}
	//fmt.Println("###",new_s,n, err)
	//for i := 0; i<n;i++{
	//	fmt.Printf("%s \n",p[i])
	//}
	return fmt.Printf("%s %v\n",new_s, err)
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}

	io.Copy(os.Stdout, &r)
}
