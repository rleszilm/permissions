package permissions

import "context"

// Allower determines whether the given user has the necessary permissions to
// act on a resource.
type Allower struct {
	roles RoleProvider
}

// Allow returns true if the requested permissions are owned by the user and
// are allowed on the given resource.
func (a *Allower) Allow(ctx context.Context, u User, r Resource, req Permission) (bool, error) {
	if a == nil {
		return false, ErrNoAllower
	}

	select {
	case <-ctx.Done():
		return false, ctx.Err()
	default:
		// If the user owns this resource they have all permissions to act on it.
		if u.PermissionID() == r.OwningUserID() {
			return allow(Permission_All, r.Permissions().GetUser(), req), nil
		}

		// If the user does not own the resource they either have a role within its
		// group or they are an other. Either way we need to get the roles to determine
		// which they are.
		roles, err := a.roles.Roles(ctx, u.PermissionID(), r.OwningGroupID())
		if err != nil {
			return false, err
		}

		// If the user has no roles they are an other and have all other permissions to act.
		if len(roles) == 0 {
			return allow(Permission_All, r.Permissions().GetOther(), req), nil
		}

		// The user has at least 1 role so get hte union of all and check for permissions.
		has := Permission_None
		for _, role := range roles {
			has = has | role.Permissions()
		}

		return allow(has, r.Permissions().GetGroup(), req), nil
	}
}

func allow(u Permission, r Permission, req Permission) bool {
	return u&r&req == req
}

// NewAllower returns a new Allower.
func NewAllower(rp RoleProvider) *Allower {
	return &Allower{
		roles: rp,
	}
}
