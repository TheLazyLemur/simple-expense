package api

import (
	"TheLazyLemur/simple-expense/auth"
	"TheLazyLemur/simple-expense/service"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (s *Server) newOrganisation(w http.ResponseWriter, r *http.Request) {
	reqBody, readErr := io.ReadAll(r.Body)
	if readErr != nil {
		log.Fatal(readErr)
		return
	}

	var orgReq createOrganisationRequest
	jsonErr := json.Unmarshal(reqBody, &orgReq)
	if jsonErr != nil {
		log.Fatal(jsonErr)
		return
	}

	org, err := service.CreateOrganisation(orgReq.OwnerID, orgReq.Name, s.store)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	pl, err := json.Marshal(org)
	if err != nil {
		log.Fatal(err)
		return
	}

	_, err = w.Write(pl)
	if err != nil {
		log.Fatal(err)
		return
	}

}

func (s *Server) getOrganisation(w http.ResponseWriter, r *http.Request) {
	muxVars := mux.Vars(r)
	id := muxVars["id"]
	i, _ := strconv.Atoi(id)

	ownerId := auth.GetClaimsProperty(r, "id").(float64)

	org, getOrgErr := service.GetOrganisation(int64(i), s.store)
	if getOrgErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(getOrgErr.Error()))
		return
	}

	if org.Owner != int64(ownerId) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	pl, jsonErr := json.Marshal(org)
	if jsonErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(jsonErr.Error()))
		return
	}

	_, _ = w.Write(pl)
}

func (s *Server) addUserToOrganisation(w http.ResponseWriter, r *http.Request) {
	ownerId := auth.GetClaimsProperty(r, "id").(float64)

	var orgAccessReq giveOrganisationAccessRequest

	reqBody, readErr := io.ReadAll(r.Body)
	if readErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonErr := json.Unmarshal(reqBody, &orgAccessReq)
	if jsonErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	existingOrg, getOrgErr := service.GetOrganisation(orgAccessReq.OrganisationID, s.store)
	if getOrgErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(getOrgErr.Error()))
		return
	}

	if existingOrg.Owner != int64(ownerId) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	accessRequestResponse, addToOrgErr := service.AddUserToOrganisation(orgAccessReq.UserID, orgAccessReq.OrganisationID, s.store)
	if addToOrgErr != nil && addToOrgErr == service.UserExistsInOrganisationError {
		w.WriteHeader(http.StatusConflict)
		_, _ = w.Write([]byte(addToOrgErr.Error()))
		return
	}

	pl, jsonErr := json.Marshal(accessRequestResponse)
	if jsonErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(jsonErr.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	_, wErr := w.Write(pl)
	if wErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(wErr.Error()))
		return
	}
}
