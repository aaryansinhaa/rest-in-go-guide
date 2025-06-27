package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/aaryansinhaa/miyuki/pkg/storage"
	"github.com/aaryansinhaa/miyuki/pkg/types"
	"github.com/aaryansinhaa/miyuki/pkg/utils/response"
	"github.com/go-playground/validator/v10"
)

func CreateBlock(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Decode the JSON body into a Miyuki struct
		var miyuki types.Miyuki
		err := json.NewDecoder(r.Body).Decode(&miyuki)
		if errors.Is(err, io.EOF) {
			slog.Error("Empy Json Body", "error", err)
			response.ErrorResponseWriter(w, http.StatusBadRequest, "Empty JSON body")
			return
		}

		if err != nil {
			response.ErrorResponseWriter(w, http.StatusBadRequest, "Unexpected error while decoding JSON body")
			return
		}

		//request validation
		if err := validator.New().Struct(miyuki); err != nil {
			response.ErrorResponseWriter(w, http.StatusBadRequest, "Validation error: "+err.Error())
			return
		}
		// Save the Miyuki struct to the storage
		id, err := storage.CreateBlock(miyuki.Name, miyuki.Age, miyuki.Address, miyuki.Email)
		if err != nil {
			slog.Error("Failed to create block", "error", err)
			response.ErrorResponseWriter(w, http.StatusInternalServerError, "Failed to create block: "+err.Error())
			return
		}

		slog.Info("Block created successfully at id", slog.Int64("id", id))
		message := "Block created successfully at id: " + strconv.Itoa(int(id))
		response.ResponseWriter(w, http.StatusCreated, message)
	}
}

// Handlers for block operations
func GetBlock()    {}
func UpdateBlock() {}
func DeleteBlock() {}
func ListBlocks()  {}
