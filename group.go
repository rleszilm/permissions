package permissions

// GroupID is an integer value used to reference a group within the permissions system.
type GroupID int64

// Group defines a grouping of users. A group defines roles that it assigns its members.
// A Users group permissions is the union of all of its rolws within a group.
type Group interface {
	PermissionID() GroupID
}
