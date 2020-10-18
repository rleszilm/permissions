package permissions

// UserID is an integer value used to reference a user within the permissions system.
type UserID int64

// User defines the interface for an entity that can request
// permissions on a resource.
type User interface {
	PermissionID() UserID
}
