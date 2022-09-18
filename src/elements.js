export function getTodoItem({id, name, status, onDelete, onUpdate}) { 
  const element = document.createElement('div');
  element.append(getCheckbox({status,onUpdate}))
  element.innerHTML = name; 
  element.append(getDeleteButton(onDelete))
  element.append(getCheckbox({status, onClick: onUpdate}))

  return element
}

export function getDeleteButton(onClick) { 
  const buttonDelete = document.createElement('button'); 
  buttonDelete.innerHTML = 'Remover'
  buttonDelete.onclick = () => {
    buttonDelete.disabled = 'disabled';
    onClick();
  }
  return buttonDelete
}

export function getCheckbox({status, onClick}){ 

  const button = document.createElement('button')
  button.innerHTML = status === 'pending' ? 'complete': 'undo'
  button.onclick = () => { 
    onClick()
  }

  return button 
}

