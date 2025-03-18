package handler

import (
	"encoding/json"
	"intro-to-go/internal/model"
	"net/http"
)

type Handler struct {
	todos map[int]string
}

func (h *Handler) GetTodos(writer http.ResponseWriter, request *http.Request) {
	h.todos = map[int]string{
		1: "zbud se",
		2: "pejt srat",
	}
	response := &model.TodosResponse{}
	//naloga 1: if stavek, ki returna status code 404 Not Found,
	// in message "no todos in memory", če v h.todos ni ničesar
	if len(h.todos) == 0 {
		response.Message = "no todos in memory"

		writer.WriteHeader(http.StatusNotFound)

	} else {
		response.Todos = h.todos
		b, err := json.MarshalIndent(response, "", "  ")
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte(err.Error()))
			return
		}
		writer.WriteHeader(http.StatusOK)
		writer.Write(b)

	}

}
