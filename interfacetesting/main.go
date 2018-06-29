package main

import "fmt"

type Record struct {
	x int
	y int
}

type  Records interface {
	Operate(r Record)int
}

type Ops []Records

func (o Ops) Operate(r Record)int{
	for _, rec := range o {
		val := rec.Operate(r)
		fmt.Println(val)
	}
	return 0
}

type Sumer struct {}

var varSumer Sumer

func (Sumer) Operate(r Record) int {
	return r.x + r.y
}


type Rester struct {}

var varRester Rester

func (Rester) Operate(r Record) int {
	return r.x - r.y
}

func main ()  {
	r := Record{4,5}
	//r.x =2
	//r.y = 4
	var opts  Records = Ops{&Sumer{},varRester}
	opts.Operate(r)
}


