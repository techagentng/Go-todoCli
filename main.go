package main

import "fmt"

type todoApp struct {
	todoList map[string]bool
	unDoneTask  []string
	doneTask []string
}

//An instance of a new todoApp struct
var newTodo1 = todoApp{
	todoList: map[string]bool{},
	unDoneTask: nil,
	doneTask:   nil,
}

//addTask receiver method, setting up the structure of map with valid params
func (todo *todoApp) addTodo(todoTask string, todoValue bool) *todoApp{
	todo.todoList[todoTask] = todoValue
	//I am returning a todo struct here
	return todo
}

//readTask method:: Check the string bool value and mark task as done
func (todo *todoApp) readTask(singleTodoVal string) *todoApp{
	if todo.todoList[singleTodoVal] == false {
		todo.todoList[singleTodoVal] = true
	}
	return todo
}

//
//deleteTask method
func (todo todoApp) deleteTask(s string){
	//loop through todoList
	//check if the value of todoList key is truthy
	//copy that key into  todo.donTask
	// delete that key by pointing to its address in the map

}

//markAsDone receiver function implementation, mark task as done
func (todo *todoApp) markAsDone(todoKey string, todoValue bool) *todoApp{
	//changing values and pointing to the original struct with (*)
	newTodo1.todoList[todoKey] = true
	//fmt.Println(newTodo1.todoList[todoKey])
	return todo
}

func main(){
	newTodo1.addTodo("Read about maps in golang", false)
	newTodo1.addTodo("Read about struct in golang", false)
	newTodo1.addTodo("Read algorithm", false)
	newTodo1.addTodo("xxxxxxxxxx", false)
	fmt.Println(newTodo1.todoList)
	//newTodo1.readTask("")
	//fmt.Println(newTodo1.readTask("Read about struct in golang"))
}