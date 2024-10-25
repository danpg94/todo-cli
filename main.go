package main

import "fmt"

func main() {
	fmt.Println("Inside main in todo app")

	todos := Todos{}
	storage := NewStorage[Todos]("todos.json")
	storage.Load(&todos)
	todos.print()
	storage.Save(todos)

}
