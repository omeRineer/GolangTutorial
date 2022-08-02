package chapter3

import "fmt"

type DataBase interface {
	Add()
}

type SqlManager struct{}
type OracleManager struct{}

func (s SqlManager) Add() {
	fmt.Println("Product added with SqlManager")
}

func (o OracleManager) Add() {
	fmt.Println("Product added with OracleManager")
}

func InterfaceDemo() {
	var dataBase DataBase

	var sqlManager = SqlManager{}
	var oracleManager = OracleManager{}

	dataBase = &sqlManager
	dataBase = &oracleManager
	dataBase.Add()

}
