package api

import (
	"TheLazyLemur/simple-expense/auth"
	db "TheLazyLemur/simple-expense/db/sqlc"
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

	arg := db.CreateOrganisationParams{
		Name:  orgReq.Name,
		Owner: orgReq.OwnerID,
	}

	org, err := s.store.CreateOrganisation(r.Context(), arg)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
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

	org, err := s.store.GetOrganisation(r.Context(), int64(i))
	if err != nil {
		log.Fatal(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
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

func (s *Server) addUserToOrganisation(w http.ResponseWriter, r *http.Request) {

	token := r.Header.Get("Token")
	claims, err := auth.DecodeJwt(token)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	ownerId := claims["id"].(float64)

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

	existingOrg, err := s.store.GetOrganisation(r.Context(), orgAccessReq.OrganisationID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	if existingOrg.Owner != int64(ownerId) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	arg := db.CreateUserOrganisationAccessParams{
		UserID:         orgAccessReq.UserID,
		OrganisationID: orgAccessReq.OrganisationID,
	}

	accessResponse, err := s.store.CreateUserOrganisationAccess(r.Context(), arg)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	pl, err := json.Marshal(accessResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	_, err = w.Write(pl)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
}
