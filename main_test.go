package main

import (
	"fmt"
	"testing"
)

func TestQuickFind(t *testing.T) {
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
}
func TestQuickUnion(t *testing.T) {
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
