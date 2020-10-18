package permissions

// Resource defines the interface for an entity that requires permissions
// to act upon it.
type Resource interface {
	OwningUserID() UserID
	OwningGroupID() GroupID
	Permissions() *Permissions
}
