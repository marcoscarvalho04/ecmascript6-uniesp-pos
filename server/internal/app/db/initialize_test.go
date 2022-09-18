package db

import "testing"

func TestInitializeDatabaseSucessfully(t *testing.T) {
	database := ConfigurationDB{}
	database.InitializeDatabase()
	if len(database.ConnectionString) == 0 {
		t.Errorf("unexpected empty database connection")
		return
	}
	t.Logf("TestInitializeDatabaseSucessfully passed")

}
