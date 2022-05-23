package api

import (
	"fmt"
	"github.com/furkansahinfs/AutoOrder-Backend/pkg/api/response"
	"net/http"
)

func (a *API) GetOrderHistory(w http.ResponseWriter, r *http.Request) {
	user, err := a.controlUser(w, r)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error getting getImage : %v", err), http.StatusBadRequest, err.Error())
		return
	}

	orders, err := a.service.GetOrdersService().GetOrders(user.ID)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error getting orders history : %v", err), http.StatusBadRequest, err.Error())
		return
	}

	response.Write(w, r, orders)
	return
}
