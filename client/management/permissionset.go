// This file was auto-generated by Fern from our API Definition.

package management

import (
	fmt "fmt"
)

type GetPermissionSetResponse struct {
	Result *PermissionSet `json:"result,omitempty"`
}

type ListPermissionSetsResponse struct {
	Result []*PermissionSet `json:"result,omitempty"`
}

type Permissions string

const (
	PermissionsAdministrator  Permissions = "administrator"
	PermissionsViewer         Permissions = "viewer"
	PermissionsMember         Permissions = "member"
	PermissionsAccountManager Permissions = "account-manager"
	PermissionsConnectUi      Permissions = "connect-ui"
)

func NewPermissionsFromString(s string) (Permissions, error) {
	switch s {
	case "administrator":
		return PermissionsAdministrator, nil
	case "viewer":
		return PermissionsViewer, nil
	case "member":
		return PermissionsMember, nil
	case "account-manager":
		return PermissionsAccountManager, nil
	case "connect-ui":
		return PermissionsConnectUi, nil
	}
	var t Permissions
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (p Permissions) Ptr() *Permissions {
	return &p
}