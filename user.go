// Package mytest Title
//
// This is the description.
//     Version: 1.0.0
// swagger:meta
package mytest

// swagger:route POST / createUser
// responses:
//   201: User

// swagger:parameters createUser
type createUserRequest struct {
	// Description of request body
	// in:body
	Body User
}

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
