package security

import (
	"aahframework.org/config.v0"
	"aahframework.org/security.v0/authc"
	"aahframework.org/security.v0/authz"
)

var _ authz.Authorizer = (*AuthorizationProvider)(nil)

// AuthorizationProvider struct implements `authz.Authorizer` interface.
type AuthorizationProvider struct {}

// Init method initializes the AuthorizationProvider, this method gets called
// during server start up.
func (a *AuthorizationProvider) Init(cfg *config.Config) error {
	return nil
}

// GetAuthorizationInfo method is `authz.Authorizer` interface.
//
// GetAuthorizationInfo method gets called after authentication is successful
// to get Subject's (aka User) access control information such as roles and permissions.
// It is called by Security Manager.
func (a *AuthorizationProvider) GetAuthorizationInfo(authcInfo *authc.AuthenticationInfo) *authz.AuthorizationInfo {
	authzInfo := authz.NewAuthorizationInfo()
	authzInfo.AddRole("user")
	// authzInfo.AddPermissionString(authorities.Permissions...)

	return authzInfo
}
