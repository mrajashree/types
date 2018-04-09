package client

const (
	OpenLDAPTestAndApplyInputType                = "openLDAPTestAndApplyInput"
	OpenLDAPTestAndApplyInputFieldEnabled        = "enabled"
	OpenLDAPTestAndApplyInputFieldOpenLDAPConfig = "openLdapConfig"
	OpenLDAPTestAndApplyInputFieldPassword       = "password"
	OpenLDAPTestAndApplyInputFieldUsername       = "username"
)

type OpenLDAPTestAndApplyInput struct {
	Enabled        bool            `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	OpenLDAPConfig *OpenLDAPConfig `json:"openLdapConfig,omitempty" yaml:"openLdapConfig,omitempty"`
	Password       string          `json:"password,omitempty" yaml:"password,omitempty"`
	Username       string          `json:"username,omitempty" yaml:"username,omitempty"`
}
