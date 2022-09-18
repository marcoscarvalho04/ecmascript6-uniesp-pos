package list

import (
	"server/internal/app/db"
	"testing"
)

func TestMigrateSucessfully(t *testing.T) {
	listRepository := ListRepository{
		Model: TodoListModel{
			Name:   "Teste salvar",
			Status: "PENDING",
		},
		Connection: &db.ConfigurationDB{},
	}
	err := listRepository.Connection.InitializeDatabase()
	if err != nil {
		t.Errorf("error initializing database: %s", err.Error())
		return
	}
	err = listRepository.Migrate()
	if err != nil {
		t.Errorf("error migrating table: %s, erro: %s", listRepository.Model.TableName(), err.Error())
		return
	}
	t.Logf("TestMigrateSucessfully passed")
}
func TestSaveSucessfully(t *testing.T) {
	listRepository := ListRepository{
		Model: TodoListModel{
			Name:   "Teste salvar",
			Status: "PENDING",
		},
		Connection: &db.ConfigurationDB{},
	}
	listRepository.Connection.InitializeDatabase()
	err := listRepository.Save()
	if err != nil {
		t.Errorf("error saving model: %s", err.Error())
		return
	}
	t.Logf("saved sucessfully")
}
