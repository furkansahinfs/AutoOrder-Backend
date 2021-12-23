package api

import (
	"fmt"
	"net/http"

	"github.com/furkansahinfs/AutoOrder-Backend/pkg/api/response"
	"github.com/furkansahinfs/AutoOrder-Backend/pkg/model"
	"github.com/mitchellh/mapstructure"
)

//Endpoint for Front Item enums
func (a *API) GetItemsFront(w http.ResponseWriter, r *http.Request) {
	var frontItems []model.Item
	for _, y := range a.config.Enums.Front {
		var item model.Item
		err := mapstructure.Decode(y, &item)
		if err != nil {
			response.Errorf(w, r, fmt.Errorf("error getting front Item enum info: %v", err), http.StatusBadRequest, err.Error())
			return
		}
		frontItems = append(frontItems, item)

	}
	response.Write(w, r, frontItems)
	return

}

//Endpoint for back Item enums
func (a *API) GetItemsBack(w http.ResponseWriter, r *http.Request) {
	var backItems []model.Item
	for _, y := range a.config.Enums.Back {
		var item model.Item
		err := mapstructure.Decode(y, &item)
		if err != nil {
			response.Errorf(w, r, fmt.Errorf("error getting back Item enum info: %v", err), http.StatusBadRequest, err.Error())
			return
		}
		backItems = append(backItems, item)

	}
	response.Write(w, r, backItems)
	return
}

//Endpoint for all Item enums
func (a *API) GetItemsAll(w http.ResponseWriter, r *http.Request) {
	var allItems []model.Item
	for _, y := range a.config.Enums.Front {
		var item model.Item
		err := mapstructure.Decode(y, &item)
		if err != nil {
			response.Errorf(w, r, fmt.Errorf("error getting front-all Item enum info: %v", err), http.StatusBadRequest, err.Error())
			return
		}
		allItems = append(allItems, item)

	}
	for _, y := range a.config.Enums.Back {
		var item model.Item
		err := mapstructure.Decode(y, &item)
		if err != nil {
			response.Errorf(w, r, fmt.Errorf("error getting back-all Item enum info: %v", err), http.StatusBadRequest, err.Error())
			return
		}
		allItems = append(allItems, item)

	}

	response.Write(w, r, allItems)
	return

}
