package permissions_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/rleszilm/permissions"
	"github.com/rleszilm/permissions/permissionsfakes"
)

func TestAllow(t *testing.T) {
	user100 := &permissionsfakes.FakeUser{}
	user100.PermissionIDReturns(100)

	user101 := &permissionsfakes.FakeUser{}
	user101.PermissionIDReturns(101)

	user102 := &permissionsfakes.FakeUser{}
	user102.PermissionIDReturns(102)

	user103 := &permissionsfakes.FakeUser{}
	user103.PermissionIDReturns(103)

	user104 := &permissionsfakes.FakeUser{}
	user104.PermissionIDReturns(104)

	role1 := &permissionsfakes.FakeRole{}
	role1.PermissionsReturns(permissions.Permission_Read)

	role2 := &permissionsfakes.FakeRole{}
	role2.PermissionsReturns(permissions.Permission_Write)

	role3 := &permissionsfakes.FakeRole{}
	role3.PermissionsReturns(permissions.Permission_Execute)

	errBadUserID := errors.New("bad user id")

	roleprovider := &permissionsfakes.FakeRoleProvider{}
	roleprovider.RolesStub = func(_ context.Context, u permissions.UserID, g permissions.GroupID) ([]permissions.Role, error) {
		switch u {
		case 100:
			return nil, nil
		case 101:
			return []permissions.Role{role1, role2, role3}, nil
		case 102:
			return []permissions.Role{role1, role2}, nil
		case 103:
			return nil, nil
		default:
			return nil, errBadUserID
		}
	}

	group200 := &permissionsfakes.FakeGroup{}
	group200.PermissionIDReturns(200)

	resource100200700 := &permissionsfakes.FakeResource{}
	resource100200700.OwningUserIDReturns(100)
	resource100200700.OwningGroupIDReturns(200)
	resource100200700.PermissionsReturns(&permissions.Permissions{
		User:  permissions.Permission_None,
		Group: permissions.Permission_All,
		Other: permissions.Permission_All,
	})

	resource100200070 := &permissionsfakes.FakeResource{}
	resource100200070.OwningUserIDReturns(100)
	resource100200070.OwningGroupIDReturns(200)
	resource100200070.PermissionsReturns(&permissions.Permissions{
		User:  permissions.Permission_None,
		Group: permissions.Permission_All,
		Other: permissions.Permission_All,
	})

	resource100200007 := &permissionsfakes.FakeResource{}
	resource100200007.OwningUserIDReturns(100)
	resource100200007.OwningGroupIDReturns(200)
	resource100200007.PermissionsReturns(&permissions.Permissions{
		User:  permissions.Permission_None,
		Group: permissions.Permission_None,
		Other: permissions.Permission_All,
	})

	resource100200077 := &permissionsfakes.FakeResource{}
	resource100200077.OwningUserIDReturns(100)
	resource100200077.OwningGroupIDReturns(200)
	resource100200077.PermissionsReturns(&permissions.Permissions{
		User:  permissions.Permission_None,
		Group: permissions.Permission_All,
		Other: permissions.Permission_All,
	})

	resource100200777 := &permissionsfakes.FakeResource{}
	resource100200777.OwningUserIDReturns(100)
	resource100200777.OwningGroupIDReturns(200)
	resource100200777.PermissionsReturns(&permissions.Permissions{
		User:  permissions.Permission_All,
		Group: permissions.Permission_All,
		Other: permissions.Permission_All,
	})

	allower := permissions.NewAllower(roleprovider)

	testcases := []struct {
		desc     string
		timeout  time.Duration
		allower  *permissions.Allower
		user     *permissionsfakes.FakeUser
		resource *permissionsfakes.FakeResource
		request  permissions.Permission
		expect   bool
		err      error
	}{
		{
			desc:    "expired context",
			timeout: -1 * time.Second,
			allower: allower,
			expect:  false,
			err:     context.DeadlineExceeded,
		},
		{
			desc:     "has user level read permissions",
			timeout:  1 * time.Second,
			allower:  allower,
			user:     user100,
			resource: resource100200777,
			request:  permissions.Permission_Read,
			expect:   true,
		},
		{
			desc:     "has user level write permissions",
			timeout:  1 * time.Second,
			allower:  allower,
			user:     user100,
			resource: resource100200777,
			request:  permissions.Permission_Write,
			expect:   true,
		},
		{
			desc:     "has user level execute permissions",
			timeout:  1 * time.Second,
			allower:  allower,
			user:     user100,
			resource: resource100200777,
			request:  permissions.Permission_Execute,
			expect:   true,
		},
		{
			desc:     "has user level all permissions",
			timeout:  1 * time.Second,
			allower:  allower,
			user:     user100,
			resource: resource100200777,
			request:  permissions.Permission_All,
			expect:   true,
		},
		{
			desc:     "has no user level read permissions",
			timeout:  1 * time.Second,
			allower:  allower,
			user:     user100,
			resource: resource100200077,
			request:  permissions.Permission_Read,
			expect:   false,
		},
		{
			desc:     "has no user level write permissions",
			timeout:  1 * time.Second,
			allower:  allower,
			user:     user100,
			resource: resource100200077,
			request:  permissions.Permission_Write,
			expect:   false,
		},
		{
			desc:     "has no user level execute permissions",
			timeout:  1 * time.Second,
			allower:  allower,
			user:     user100,
			resource: resource100200077,
			request:  permissions.Permission_Execute,
			expect:   false,
		},
		{
			desc:     "has no user level permissions",
			timeout:  1 * time.Second,
			allower:  allower,
			user:     user100,
			resource: resource100200077,
			request:  permissions.Permission_All,
			expect:   false,
		},
		{
			desc:     "has group level permissions",
			timeout:  1 * time.Second,
			allower:  allower,
			user:     user101,
			resource: resource100200070,
			request:  permissions.Permission_All,
			expect:   true,
		},
		{
			desc:     "has some group level permissions",
			timeout:  1 * time.Second,
			allower:  allower,
			user:     user102,
			resource: resource100200070,
			request:  permissions.Permission_ReadWrite,
			expect:   true,
		},
		{
			desc:     "has no group level execute permissions",
			timeout:  1 * time.Second,
			allower:  allower,
			user:     user102,
			resource: resource100200070,
			request:  permissions.Permission_All,
			expect:   false,
		},
		{
			desc:     "has other level permissions",
			timeout:  1 * time.Second,
			allower:  allower,
			user:     user103,
			resource: resource100200007,
			request:  permissions.Permission_All,
			expect:   true,
		},
		{
			desc:     "no allower",
			timeout:  1 * time.Second,
			user:     user100,
			resource: resource100200777,
			request:  permissions.Permission_All,
			expect:   false,
			err:      permissions.ErrNoAllower,
		},
		{
			desc:     "bad user id when getting roles",
			timeout:  1 * time.Second,
			allower:  allower,
			user:     user104,
			resource: resource100200777,
			request:  permissions.Permission_All,
			expect:   false,
			err:      errBadUserID,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.desc, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), tc.timeout)
			defer cancel()

			permissions.UseAllower(tc.allower)

			actual, err := permissions.Allow(ctx, tc.user, tc.resource, tc.request)
			if err != tc.err {
				t.Errorf("Incorrect error received from Allow. Expected: %v Actual: %v", tc.err, err)
			}

			if actual != tc.expect {
				t.Errorf("Incorrect permission result received from Allow. Expected: %v Actual: %v", tc.expect, actual)
			}
		})
	}
}
