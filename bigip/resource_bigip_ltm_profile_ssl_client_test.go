package bigip

import (
	"fmt"
	"testing"

	"github.com/f5devcentral/go-bigip"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

var TEST_CLIENTSSL_NAME = fmt.Sprintf("/%s/test-ClientSsl", TEST_PARTITION)

var TEST_CLIENTSSL_RESOURCE = `
resource "bigip_ltm_profile_client_ssl" "test-ClientSsl" {
	name                                = "/Common/test-ClientSsl"
	partition                           = "Common"
	defaults_from                       = "/Common/clientssl"
	alert_timeout                       = "indefinite"
	allow_non_ssl                       = "disabled"
	authenticate                        = "once"
	authenticate_depth                  = 9
	ca_file                             = "none"
	cache_size                          = 262144
	cache_timeout                       = 3600
	cert                                = "/Common/default.crt"
	cert_key_chain {
		cert = "/Common/default.crt"
		key  = "/Common/default.key"
		name = "default"
	}
	cert_extension_includes             = [
		"basic-constraints",
		"subject-alternative-name",
	  ]
	cert_life_span                      = 30
	cert_lookup_by_ipaddr_port          = "disabled"
	chain                               = "none"
	ciphers                             = "ALL"
	client_cert_ca                      = "none"
	crl_file                            = "none"
	forward_proxy_bypass_default_action = "intercept"
	generic_alert                       = "enabled"
	handshake_timeout                   = "10"
	inherit_cert_keychain               = "false"
	key                                 = "/Common/default.key"
	mod_ssl_methods                     = "disabled"
	mode                                = "enabled"
	peer_cert_mode                      = "ignore"
	proxy_ca_cert                       = "none"
	proxy_ca_key                        = "none"
	proxy_ca_passphrase                 = ""
	proxy_ssl                           = "disabled"
	proxy_ssl_passthrough               = "disabled"
	renegotiate_period                  = "indefinite"
	renegotiate_size                    = "indefinite"
	renegotiation                       = "enabled"
	retain_certificate                  = "true"
	secure_renegotiation                = "require"
	server_name                         = "none"
	session_mirroring                   = "disabled"
	session_ticket                      = "disabled"
	sni_default                         = "false"
	sni_require                         = "false"
	ssl_forward_proxy                   = "disabled"
	ssl_forward_proxy_bypass            = "disabled"
	ssl_sign_hash                       = "any"
	strict_resume                       = "disabled"
	tm_options                          = [
		"dont-insert-empty-fragments",
	  ]
	unclean_shutdown                    = "enabled"
  }
`

func TestAccBigipLtmProfileClientSsl_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAcctPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testCheckClientSslDestroyed,
		Steps: []resource.TestStep{
			{
				Config: TEST_CLIENTSSL_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					testCheckClientSslExists(TEST_CLIENTSSL_NAME, true),
					resource.TestCheckResourceAttr("bigip_ltm_profile_client_ssl.test-ClientSsl", "name", "/Common/test-ClientSsl"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_client_ssl.test-ClientSsl", "partition", "Common"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_client_ssl.test-ClientSsl", "defaults_from", "/Common/clientssl"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_client_ssl.test-ClientSsl", "alert_timeout", "indefinite"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_client_ssl.test-ClientSsl", "allow_non_ssl", "disabled"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_client_ssl.test-ClientSsl", "authenticate", "once"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_client_ssl.test-ClientSsl", "authenticate_depth", "9"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_client_ssl.test-ClientSsl", "ca_file", "none"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_client_ssl.test-ClientSsl", "cache_size", "262144"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_client_ssl.test-ClientSsl", "cache_timeout", "3600"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_client_ssl.test-ClientSsl", "cert", "/Common/default.crt"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_client_ssl.test-ClientSsl",
						fmt.Sprintf("cert_extension_includes.%d", schema.HashString("basic-constraints")),
						"basic-constraints"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_client_ssl.test-ClientSsl",
						fmt.Sprintf("cert_extension_includes.%d", schema.HashString("subject-alternative-name")),
						"subject-alternative-name"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_client_ssl.test-ClientSsl", "cert_life_span", "30"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_client_ssl.test-ClientSsl", "cert_lookup_by_ipaddr_port", "disabled"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_client_ssl.test-ClientSsl", "chain", "none"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_client_ssl.test-ClientSsl", "ciphers", "ALL"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_client_ssl.test-ClientSsl", "client_cert_ca", "none"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_client_ssl.test-ClientSsl", "crl_file", "none"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_client_ssl.test-ClientSsl", "defaults_from", "/Common/clientssl"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_client_ssl.test-ClientSsl", "forward_proxy_bypass_default_action", "intercept"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_client_ssl.test-ClientSsl", "generic_alert", "enabled"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_client_ssl.test-ClientSsl", "handshake_timeout", "10"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_client_ssl.test-ClientSsl", "inherit_cert_keychain", "false"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_client_ssl.test-ClientSsl", "key", "/Common/default.key"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_client_ssl.test-ClientSsl", "mod_ssl_methods", "disabled"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_client_ssl.test-ClientSsl", "mode", "enabled"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_client_ssl.test-ClientSsl", "peer_cert_mode", "ignore"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_client_ssl.test-ClientSsl", "proxy_ca_cert", "none"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_client_ssl.test-ClientSsl", "proxy_ca_key", "none"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_client_ssl.test-ClientSsl", "proxy_ca_passphrase", ""),
					resource.TestCheckResourceAttr("bigip_ltm_profile_client_ssl.test-ClientSsl", "proxy_ssl", "disabled"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_client_ssl.test-ClientSsl", "proxy_ssl_passthrough", "disabled"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_client_ssl.test-ClientSsl", "renegotiate_period", "indefinite"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_client_ssl.test-ClientSsl", "renegotiate_size", "indefinite"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_client_ssl.test-ClientSsl", "renegotiation", "enabled"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_client_ssl.test-ClientSsl", "retain_certificate", "true"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_client_ssl.test-ClientSsl", "secure_renegotiation", "require"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_client_ssl.test-ClientSsl", "server_name", "none"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_client_ssl.test-ClientSsl", "session_mirroring", "disabled"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_client_ssl.test-ClientSsl", "session_ticket", "disabled"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_client_ssl.test-ClientSsl", "sni_default", "false"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_client_ssl.test-ClientSsl", "sni_require", "false"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_client_ssl.test-ClientSsl", "ssl_forward_proxy", "disabled"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_client_ssl.test-ClientSsl", "ssl_forward_proxy_bypass", "disabled"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_client_ssl.test-ClientSsl", "ssl_sign_hash", "any"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_client_ssl.test-ClientSsl", "strict_resume", "disabled"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_client_ssl.test-ClientSsl",
						fmt.Sprintf("tm_options.%d", schema.HashString("dont-insert-empty-fragments")),
						"dont-insert-empty-fragments"),

					resource.TestCheckResourceAttr("bigip_ltm_profile_client_ssl.test-ClientSsl", "unclean_shutdown", "enabled"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_client_ssl.test-ClientSsl", "cert_key_chain.0.name", "default"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_client_ssl.test-ClientSsl", "cert_key_chain.0.cert", "/Common/default.crt"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_client_ssl.test-ClientSsl", "cert_key_chain.0.key", "/Common/default.key"),
				),
			},
		},
	})
}

func TestAccBigipLtmProfileClientSsl_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAcctPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testCheckClientSslDestroyed,
		Steps: []resource.TestStep{
			{
				Config: TEST_CLIENTSSL_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					testCheckClientSslExists(TEST_CLIENTSSL_NAME, true),
				),
				ResourceName:      TEST_CLIENTSSL_NAME,
				ImportState:       false,
				ImportStateVerify: true,
			},
		},
	})
}

func testCheckClientSslExists(name string, exists bool) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*bigip.BigIP)
		p, err := client.GetClientSSLProfile(name)
		if err != nil {
			return err
		}
		if exists && p == nil {
			return fmt.Errorf("ClientSsl Profile %s was not created.", name)
		}
		if !exists && p == nil {
			return fmt.Errorf("ClientSsl Profile %s still exists.", name)
		}
		return nil
	}
}

func testCheckClientSslDestroyed(s *terraform.State) error {
	client := testAccProvider.Meta().(*bigip.BigIP)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "bigip_ltm_profile_clientssl" {
			continue
		}

		name := rs.Primary.ID
		ClientSsl, err := client.GetClientSSLProfile(name)
		if err != nil {
			return err
		}
		if ClientSsl != nil {
			return fmt.Errorf("ClientSsl Profile %s not destroyed.", name)
		}
	}
	return nil
}
