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
		id, err := a.service.GetImageService().SaveImagePath(path, user.ID)
		if err != nil {
			response.Errorf(w, r, fmt.Errorf("error getting GetImage info: %v", err), http.StatusBadRequest, err.Error())
			return
		}

		//TODO
		//Python backende image gönder

		response.Write(w, r, id)
	} else {
		response.Errorf(w, r, fmt.Errorf("error getting GetImage info: %v", err), http.StatusBadRequest, errors.New("File extension error").Error())
		return
	}

}

/*
func (a *API) send(filePath string, config []model.Item) (*http.Response, error, int) {
	client := &http.Client{
		Timeout: time.Second * 60,
	}
	// New multipart writer.
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	fw, err := writer.CreateFormField("data")
	if err != nil {
		return nil, err, 0
	}
	_, err = io.Copy(fw, strings.NewReader(fmt.Sprintf("%v", config)))
	if err != nil {
		return nil, err, 0
	}
	fw, err = writer.CreateFormFile("image", filePath)
	if err != nil {
		return nil, err, 0
	}
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err, 0
	}
	_, err = io.Copy(fw, file)
	if err != nil {
		return nil, err, 0
	}
	// Close multipart writer.
	writer.Close()
	req, err := http.NewRequest("POST", a.config.PythonBackendAddress, bytes.NewReader(body.Bytes()))
	if err != nil {
		return nil, err, 0
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	rsp, _ := client.Do(req)
	if rsp.StatusCode != http.StatusOK {
		return nil, err, rsp.StatusCode
	}
	return rsp, nil, rsp.StatusCode
}
*/
