package permissions

import "context"

var (
	// DefaultAllower is the Allower that calls to Allow uses.
	defaultAllower *Allower
)

// UseAllower updates the Allower used by Allow.
func UseAllower(a *Allower) {
	defaultAllower = a
}

// Allow is a convenience method that calls Allow of the DefaultAllower
func Allow(ctx context.Context, u User, r Resource, req Permission) (bool, error) {
	return defaultAllower.Allow(ctx, u, r, req)
}
