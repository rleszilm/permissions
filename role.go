package permissions

import "context"

// RoleID is an integer value used to reference a role within the permissions system.
type RoleID int64

// Role defines the interface for an entity that can request
// permissions on a resource.
type Role interface {
	Permissions() Permission
}

// RoleProvider is an interface that returns the roles for a user within a group.
type RoleProvider interface {
	Roles(context.Context, UserID, GroupID) ([]Role, error)
}
