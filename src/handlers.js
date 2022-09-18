import { deleteTodo, listTodos, saveToDo, updateTodo } from "./api";
import {element, getTodoItem} from "./elements"

function getNewTodoName(){ 
  const input = document.querySelector('[name=new-name]').value
  return input
}


export function listenSubmit(){
  refreshTodoList()
  const button = document.getElementById('new-button'); 
  button.onclick = async () => {
    const inputValue = getNewTodoName()
    console.log('button clicked', inputValue)
    await saveToDo({
      name: inputValue,
      status: "pending"
    })
    refreshTodoList()
  }
}

async function refreshTodoList() { 
  const todoList = await listTodos()
  const todoElements = document.getElementById('todo-list');
  todoElements.innerHTML = ""
  todoList.forEach(todo => {
    const onDelete = async () => { 
      await deleteTodo(todo.id)
      refreshTodoList()
    }

    const onUpdate = async () => {
      await updateTodo({
        ...todo, 
        status: todo.status === 'completed' ? 'pending': 'completed'
      })
      refreshTodoList()
    }
    todoElements.append(getTodoItem({
      ...todo,
      onDelete,
      onUpdate,
    }))
  })
    
  };


export function loadTodos() { 
  refreshTodoList()
}