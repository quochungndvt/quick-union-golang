package main

import (
	"fmt"
)

func main() {
	fmt.Println("[TEST QuickFind]")
	var qf QuickFind
	qf.QuickFind(10)
	qf.Unite(3, 4)
	qf.Unite(4, 9)
	qf.Unite(8, 0)
	qf.Unite(2, 3)
	qf.Unite(5, 6)
	qf.Unite(5, 9)
	qf.Unite(7, 3)
	qf.Unite(4, 8)
	qf.Unite(6, 1)
	qf.Show()

	fmt.Println("qf.Find(3, 8)", qf.Find(3, 8), "_____________________________")

	fmt.Println("[TEST QuickUnion]")
	var qu QuickUnion
	qu.QuickUnion(10)
	qu.Unite(3, 4)
	qu.Unite(4, 9)
	qu.Unite(8, 0)
	qu.Unite(2, 3)
	qu.Unite(5, 6)
	qu.Unite(5, 9)
	qu.Unite(7, 3)
	qu.Unite(4, 8)
	qu.Unite(6, 1)
	qu.Show()

	fmt.Println("qu.Find(3, 8)", qu.Find(3, 8), "_____________________________")

}

type QuickFind struct {
	id []int
}

func (this *QuickFind) QuickFind(N int) {
	this.id = make([]int, N)
	for i := 0; i < N; i++ {
		this.id[i] = i
	}
}
func (this *QuickFind) Find(p, q int) bool {
	return this.id[p] == this.id[q]
}
func (this *QuickFind) Unite(p, q int) {
	pid := this.id[p]
	for i := 0; i < len(this.id); i++ {
		if this.id[i] == pid {
			this.id[i] = this.id[q]
		}
	}
	fmt.Println("[QuickFind.Unite]", p, q, this.id)
}
func (this *QuickFind) Show() {
	fmt.Println("[QuickFind]", this.id)
}

//
type QuickUnion struct {
	id []int
	sz []int
}

func (this *QuickUnion) QuickUnion(N int) {
	this.id = make([]int, N)
	this.sz = make([]int, N)
	for i := 0; i < N; i++ {
		this.id[i] = i
		//init weight for each item
		this.sz[i] = 1
	}
}
func (this *QuickUnion) root(i int) int {
	for i != this.id[i] {
		//implement
		this.id[i] = this.id[this.id[i]]
		i = this.id[i]
	}
	return i
}
func (this *QuickUnion) Find(p, q int) bool {
	return this.root(p) == this.root(q)
}
func (this *QuickUnion) Unite(p, q int) {
	i := this.root(p)
	j := this.root(q)

	if this.sz[i] < this.sz[j] {
		this.id[i] = j
		this.sz[j] += this.sz[i]
	} else {
		this.id[j] = i
		this.sz[i] += this.sz[j]
	}

	//this.id[i] = j
	fmt.Println("[QuickUnion.Unite]", p, q, this.id)
}
func (this *QuickUnion) Show() {
	fmt.Println("[QuickUnion]", this.id)
}
