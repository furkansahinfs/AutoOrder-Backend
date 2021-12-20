package api

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/furkansahinfs/AutoOrder-Backend/pkg/api/response"
)

//endpoint for store userInformation
func (a *API) GetImage(w http.ResponseWriter, r *http.Request) {
	user, err := a.controlUser(w, r)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error getting getImage : %v", err), http.StatusBadRequest, err.Error())
		return
	}
	err = r.ParseMultipartForm(32 << 20) // maxMemory 32MB
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error getting GetImage info: %v", err), http.StatusBadRequest, err.Error())
		return
	}
	//Access the photo key - First Approach
	file, h, err := r.FormFile("image")
	fileExtentions := strings.Split(h.Filename, ".")
	if len(fileExtentions) != 2 {
		response.Errorf(w, r, fmt.Errorf("error getting GetImage info: %v", err), http.StatusBadRequest, err.Error())
		return
	}
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error getting GetImage info: %v", err), http.StatusBadRequest, err.Error())
		return
	}
	if fileExtentions[1] == "png" || fileExtentions[1] == "jpg" || fileExtentions[1] == "jpeg" {
		timeNow := time.Now().String()
		path := a.config.ImagePath + user.Email + "_" + timeNow + "." + fileExtentions[1]
		tmpfile, err := os.Create(path)
		defer tmpfile.Close()
		if err != nil {
			response.Errorf(w, r, fmt.Errorf("error getting GetImage info: %v", err), http.StatusInternalServerError, err.Error())
			return
		}
		_, err = io.Copy(tmpfile, file)
		if err != nil {
			response.Errorf(w, r, fmt.Errorf("error getting GetImage info: %v", err), http.StatusInternalServerError, err.Error())
			return
		}
		fmt.Println(user.ID)
		id, err := a.service.GetImageService().SaveImagePath(path, user.ID)
		if err != nil {
			response.Errorf(w, r, fmt.Errorf("error getting GetImage info: %v", err), http.StatusBadRequest, err.Error())
			return
		}
		response.Write(w, r, id)
	} else {
		response.Errorf(w, r, fmt.Errorf("error getting GetImage info: %v", err), http.StatusBadRequest, errors.New("File extension error").Error())
		return
	}

}
