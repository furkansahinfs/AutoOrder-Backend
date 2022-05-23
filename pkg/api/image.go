package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/furkansahinfs/AutoOrder-Backend/pkg/model"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
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
	if fileExtentions[1] == "jpg" || fileExtentions[1] == "jpeg" {
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
		_, err = a.service.GetImageService().SaveImagePath(path, user.ID)
		if err != nil {
			response.Errorf(w, r, fmt.Errorf("error getting GetImage info: %v", err), http.StatusBadRequest, err.Error())
			return
		}
		var uConfigItems []string
		frontItems, err := a.service.GetUserConfigurationService().GetUserConfiguration(user.ID, "front")
		if err != nil {
			if err.Error() == "User dont have a configuration" {

			} else {
				response.Errorf(w, r, fmt.Errorf("error getting GetImage info: %v", err), http.StatusBadRequest, err.Error())
				return
			}

		}
		if len(frontItems) > 0 {
			for _, item := range frontItems {
				uConfigItems = append(uConfigItems, item.Name)
			}
		}

		backItems, err := a.service.GetUserConfigurationService().GetUserConfiguration(user.ID, "back")
		if err != nil {
			if err.Error() == "User dont have a configuration" {

			} else {
				response.Errorf(w, r, fmt.Errorf("error getting GetImage info: %v", err), http.StatusBadRequest, err.Error())
				return
			}

		}
		if len(backItems) > 0 {
			for _, item := range backItems {
				uConfigItems = append(uConfigItems, item.Name)
			}
		}

		if len(uConfigItems) > 0 {
			responseBodyString, err := a.sendToImageAnalyse(path, uConfigItems)
			if err != nil {
				response.Errorf(w, r, fmt.Errorf("error getting GetImage info: %v", err), http.StatusBadRequest, err.Error())
				return
			}

			results, err := a.sendToFakeApi(responseBodyString)
			if err != nil {
				response.Errorf(w, r, fmt.Errorf("error getting making order from fake api : %v", err), http.StatusBadRequest, err.Error())
				return
			}

			err = a.service.GetOrdersService().SaveOrders(results, user.ID, path)
			if err != nil {
				response.Errorf(w, r, fmt.Errorf("error getting saving orders to db info: %v", err), http.StatusBadRequest, err.Error())
				return
			}

			response.Write(w, r, "")
			return
		} else {
			response.Errorf(w, r, fmt.Errorf("error getting GetImage info: %v", err), http.StatusBadRequest, errors.New("User Dont have a configuration").Error())
			return
		}
	} else {
		response.Errorf(w, r, fmt.Errorf("error getting GetImage info: %v", err), http.StatusBadRequest, errors.New("File extension error").Error())
		return
	}

}

func (a *API) sendToImageAnalyse(filePath string, config []string) (string, error) {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	file, errFile1 := os.Open(filePath)
	defer file.Close()
	part1, errFile1 := writer.CreateFormFile("image", filepath.Base(filePath))
	_, errFile1 = io.Copy(part1, file)
	if errFile1 != nil {
		return "", errFile1
	}
	json, err := json.Marshal(config)
	if err != nil {
		return "", err
	}
	_ = writer.WriteField("config", string(json))
	err = writer.Close()
	if err != nil {
		return "", err
	}

	client := &http.Client{
		Timeout: time.Minute * 5,
	}
	req, err := http.NewRequest("POST", a.config.PythonBackendAddress, payload)

	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func (a *API) sendToFakeApi(items string) ([]model.OrderResponse, error) {
	fmt.Println(items)
	// make http post
	itemArray := strings.Split(items, ",")
	req := model.OrderRequestJava{
		OrderRequest: itemArray,
	}
	jsonData, err := json.Marshal(req)

	if err != nil {
		return nil, err
	}

	resp, err := http.Post(a.config.FakeApiAddress, "application/json",
		bytes.NewBuffer(jsonData))

	if err != nil {
		return nil, err
	}

	var res model.OrderResponseJava

	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return nil, err
	}
	fmt.Println(res.OrderResponse)

	return res.OrderResponse, nil

}
