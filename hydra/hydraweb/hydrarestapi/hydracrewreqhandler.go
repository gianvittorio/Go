package hydrarestapi

import (
	"encoding/json"
	"fmt"
	"hydra/hydra/hydrablayer"
	"net/http"
	"strconv"
)

type HydraCrewReqHandler struct {
	dbConn hydrablayer.DBLayer
}

func NewHydraCrewReqHandler() *HydraCrewReqHandler {
	return new(HydraCrewReqHandler)
}

func (hcwreq *HydraCrewReqHandler) connect(o, conn string) error {
	dblayer, err := hydrablayer.ConnectToDatabase(o, conn)
	if err != nil {
		return err
	}
	hcwreq.dbConn = dblayer
	return nil
}

func (hcwreq *HydraCrewReqHandler) handleHydraCrewRequests(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		ids := r.RequestURI[len("/hydracrew/"):]
		id, err := strconv.Atoi(ids)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "id %s provided is not valid. \n", ids)
			return
		}
		cm, err := hcwreq.dbConn.FindMember(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "Error %s occurred when searching for id %d \n", err.Error(), id)
			return
		}
		json.NewEncoder(w).Encode(&cm)
	case "POST":
		cm := new(hydrablayer.CrewMember)
		err := json.NewDecoder(r.Body).Decode(cm)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "Error %s occurred", err)
			return
		}
		err = hcwreq.dbConn.AddMember(cm)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "Error %s occurred while adding a crew member to the Hydra database", err)
			return
		}
		fmt.Fprint(w, "Successfully inserted id %d \n", cm.ID)
	}
}
