export const URL_BASE = 'http://localhost:8080'

export async function listTodos() { 
  const response = await fetch(`${URL_BASE}/todo`, {
    method: "GET"
  });
  return  await response.json()
}

export async function saveToDo(todo) { 
  const response = await fetch(`${URL_BASE}/todo`, {
    method: "POST",
    body: JSON.stringify(todo),
    headers: {
      'Content-Type': 'application/json'
    },
  });
  return await response.json()
  
}

export async function deleteTodo(id) { 
  const response = await fetch(`${URL_BASE}/todo/${id}`, { 
    method: "DELETE", 
  })
  return await  response
}


export async function updateTodo(todo) { 
  const response = await fetch(`${URL_BASE}/todo/${todo.id}`, {
    method: "PUT",
    body: JSON.stringify(todo),
    headers: {
      'Content-Type': 'application/json'
    },
  });
  return await response.json()
  
}