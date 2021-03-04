package main

import "fmt"

var yfcount int

type yfdatasql struct {
	Id       int
	Name     string
	Count    int
	Otherinf string
}

var yfdatamap []yfdatasql

func refreshyfdata() {
	userdb.Table("yfdatasql")
	userdb.Find(&yfdatamap)
	fmt.Println(yfdatamap)
	userdb.Model(yfdatasql{}).Count(&yfcount)
	userdb.Table("users")

}
