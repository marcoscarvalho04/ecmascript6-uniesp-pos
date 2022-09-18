package list

import "server/internal/app/db"

type ListRepository struct {
	Connection *db.ConfigurationDB
	Model      TodoListModel
}

func (s ListRepository) Migrate() error {
	err := s.Connection.DatabaseConnection.AutoMigrate(s.Model)
	if err != nil {
		return err
	}
	return nil
}
func (s ListRepository) Save() error {

	result := s.Connection.DatabaseConnection.Save(&s.Model)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s ListRepository) Update() error {
	result := s.Connection.DatabaseConnection.Updates(&s.Model)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s ListRepository) Delete() error {
	result := s.Connection.DatabaseConnection.Delete(s.Model)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s ListRepository) FindAll() ([]TodoListModel, error) {
	allList := make([]TodoListModel, 0)
	result := s.Connection.DatabaseConnection.Find(&allList)
	if result != nil {
		return allList, result.Error
	}
	return allList, nil
}

func (s ListRepository) GetById(id int32) (TodoListModel, error) {
	listById := TodoListModel{Id: id}
	result := s.Connection.DatabaseConnection.Model(listById).Where("id = ? ", id).Find(&listById)
	if result != nil {
		return listById, result.Error
	}
	return listById, nil

}
