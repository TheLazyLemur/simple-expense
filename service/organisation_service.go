package service

import (
	db "TheLazyLemur/simple-expense/db/sqlc"
	"context"
	"errors"
)

var UserExistsInOrganisationError = errors.New("User already exists in organisation")

func AddUserToOrganisation(userId int64, organisationId int64, store *db.Store) (*db.UserOrganisationsAccess, error) {
	arg := db.CreateUserOrganisationAccessParams{
		UserID:         userId,
		OrganisationID: organisationId,
	}

	userExistsInOrg, _ := CheckIfUserExistsInOrganisation(userId, organisationId, store)
	if userExistsInOrg {
		return nil, UserExistsInOrganisationError
	}

	response, err := store.CreateUserOrganisationAccess(context.Background(), arg)

	return &response, err
}

func GetOrganisation(organisationId int64, store *db.Store) (db.Organisation, error) {
	return store.GetOrganisation(context.Background(), organisationId)
}

func CheckIfUserExistsInOrganisation(userId int64, organisationId int64, store *db.Store) (bool, error) {
	arg := db.GetUserOrganisationAccessParams{
		UserID:         userId,
		OrganisationID: organisationId,
	}

	_, err := store.GetUserOrganisationAccess(context.Background(), arg)
	return err == nil, err
}

func CreateOrganisation(ownerId int64, organisationName string, store *db.Store) (db.Organisation, error) {
	arg := db.CreateOrganisationParams{
		Name:  organisationName,
		Owner: ownerId,
	}

	return store.CreateOrganisation(context.Background(), arg)
}
