-- name: CreateUserOrganisationAccess :one
INSERT INTO user_organisations_access (user_id, organisation_id)
VALUES ($1, $2) 
RETURNING *;

-- name: GetUserOrganisationAccess :one
SELECT * FROM user_organisations_access
WHERE user_id = $1 
AND organisation_id = $2;

-- name: DeleteUserOrganisationAccess :exec
DELETE FROM user_organisations_access
WHERE user_id = $1
AND organisation_id = $2;
