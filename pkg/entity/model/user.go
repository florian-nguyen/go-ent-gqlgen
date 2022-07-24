package model

import "go-ent-gqlgen/ent"

// User is the model entity for the User schema.
type User = ent.User

// CreateUserInput represents a mutation input for creating users.
type CreateUserInput = ent.CreateUserInput

// UpdateUserInput represents a mutation input for updating users.
type UpdateUserInput = ent.UpdateUserInput

// UserConnection is the connection containing edges to User.
type UserConnection = ent.UserConnection

type UserWhereInput = ent.UserWhereInput
