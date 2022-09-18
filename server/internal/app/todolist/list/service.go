package list

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"server/internal/app/todolist/request"
	"strconv"
)

type ListDTO struct {
	Id     int32  `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

type ListService struct {
	ListRepository *ListRepository
	ListDTO        ListDTO
}

func New(respository *ListRepository) *ListService {
	return &ListService{
		ListRepository: respository,
		ListDTO:        ListDTO{},
	}
}

func (s ListService) Save(w http.ResponseWriter, r *http.Request) {
	requestContent := &ListDTO{}
	bodyRawContent, err := io.ReadAll(r.Body)
	if err != nil {
		msg := fmt.Sprintf("error reading body request: %s", err.Error())
		request.WriteStatusCodeAndMessage(w, request.BAD_REQUEST, msg)
		return
	}
	err = json.Unmarshal(bodyRawContent, &requestContent)
	if err != nil {
		msg := fmt.Sprintf("error unmarshall body content: %s", err.Error())
		request.WriteStatusCodeAndMessage(w, request.BAD_REQUEST, msg)
		return
	}

	s.ListRepository.Model = TodoListModel{
		Status: requestContent.Status,
		Name:   requestContent.Name,
	}
	err = s.ListRepository.Save()
	if err != nil {
		msg := fmt.Sprintf("error saving list: %s", err.Error())
		request.WriteStatusCodeAndMessage(w, request.INTERNAL_ERROR_SERVER, msg)
		return
	}

	result, err := json.Marshal(s.ListRepository.Model)
	if err != nil {
		msg := fmt.Sprintf("error marshelling result: %s", err.Error())
		request.WriteStatusCodeAndMessage(w, request.INTERNAL_ERROR_SERVER, msg)
		return
	}

	w.WriteHeader(request.CREATED)
	w.Write(result)

}

func (s ListService) ListAll(w http.ResponseWriter, r *http.Request) {

	allListDTO := make([]ListDTO, 0)
	allList, err := s.ListRepository.FindAll()
	if err != nil {
		msg := fmt.Sprintf("erro getting all list: %s", err.Error())
		request.WriteStatusCodeAndMessage(w, request.INTERNAL_ERROR_SERVER, msg)
		return
	}
	for _, value := range allList {
		allListDTOElement := &ListDTO{
			Id:     value.Id,
			Status: value.Status,
			Name:   value.Name,
		}
		allListDTO = append(allListDTO, *allListDTOElement)
	}

	contentResult, err := json.Marshal(&allListDTO)
	if err != nil {
		msg := fmt.Sprintf("erro getting all list: %s", err.Error())
		request.WriteStatusCodeAndMessage(w, request.INTERNAL_ERROR_SERVER, msg)
		return
	}
	w.WriteHeader(request.OK)
	w.Write(contentResult)
}

func (s ListService) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	if len(id) == 0 {
		request.WriteStatusCodeAndMessage(w, request.BAD_REQUEST, "missing id ")
		return
	}
	convertedId, err := strconv.Atoi(id)
	if err != nil {
		request.WriteStatusCodeAndMessage(w, request.BAD_REQUEST, "expected integer id")
		return
	}

	s.ListRepository.Model = TodoListModel{
		Id: int32(convertedId),
	}
	err = s.ListRepository.Delete()
	if err != nil {
		msg := fmt.Sprintf("erro deleting id: %s", err.Error())
		request.WriteStatusCodeAndMessage(w, request.INTERNAL_ERROR_SERVER, msg)
		return
	}
	w.WriteHeader(request.DELETED)
}

func (s ListService) Update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	if len(id) == 0 {
		request.WriteStatusCodeAndMessage(w, request.BAD_REQUEST, "missing id ")
		return
	}
	convertedId, err := strconv.Atoi(id)
	if err != nil {
		request.WriteStatusCodeAndMessage(w, request.BAD_REQUEST, "expected integer id")
		return
	}
	requestContent := &ListDTO{}
	bodyRawContent, err := io.ReadAll(r.Body)
	if err != nil {
		msg := fmt.Sprintf("error reading body content: %s", err.Error())
		request.WriteStatusCodeAndMessage(w, request.BAD_REQUEST, msg)
		return
	}
	err = json.Unmarshal(bodyRawContent, &requestContent)
	if err != nil {
		msg := fmt.Sprintf("error unmarshall body content: %s", err.Error())
		request.WriteStatusCodeAndMessage(w, request.BAD_REQUEST, msg)
		return
	}

	s.ListRepository.Model = TodoListModel{
		Id:     int32(convertedId),
		Status: requestContent.Status,
		Name:   requestContent.Name,
	}
	err = s.ListRepository.Update()
	if err != nil {
		msg := fmt.Sprintf("error saving list: %s", err.Error())
		request.WriteStatusCodeAndMessage(w, request.INTERNAL_ERROR_SERVER, msg)
		return
	}

	result, err := json.Marshal(s.ListRepository.Model)
	if err != nil {
		msg := fmt.Sprintf("error marshelling result: %s", err.Error())
		request.WriteStatusCodeAndMessage(w, request.INTERNAL_ERROR_SERVER, msg)
		return
	}

	w.WriteHeader(request.CREATED)
	w.Write(result)

}
