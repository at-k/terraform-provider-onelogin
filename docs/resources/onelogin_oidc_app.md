---
layout: "onelogin"
page_title: "OneLogin: onelogin_oidc_apps"
sidebar_current: "docs-onelogin-resource-oidc-apps"
description: |-
  Creates a OIDC Application.
---

# onelogin_oidc_apps

Creates an OIDC Application.

This resource allows you to create and configure an OIDC Application.

## Example Usage

```hcl
resource onelogin_oidc_apps my_oidc_app {
  name = "my OIDC APP"
	notes = "example"
  visible = true
	allow_assumed_signin = false
	connector_id = 123456
	description = "example OIDC app"

	configuration = {
		login_url = "https://www.example.com"
		oidc_application_type = 0
		redirect_uri = "https://example.com/example"
		refresh_token_expiration_minutes = 1
		token_endpoint_auth_method = 1
		access_token_expiration_minutes = 1
	}

	provisioning = {
		enabled = false
	}

	parameters = {
		provisioned_entitlements = false
		user_attribute_macros = ""
		user_attribute_mappings = ""
		values = ""
		attributes_transformations = ""
		default_values = ""
		include_in_saml_assertion = false
		label = "example app parameter "
		param_key_name = "example"
		safe_entitlements_enabled = false
		skip_if_blank = false
	}
}
```

## Argument Reference

The following arguments are supported:
The following arguments are supported:
* `name` - (Required) The app's name.

* `connector_id` - (Required) The ID for the app connector, dictates the type of app (e.g. AWS Multi-Role App).

* `description` - (Optional) App description.

* `notes` - (Optional) Notes about the app.

* `visible` - (Optional) Determine if app should be visible in OneLogin portal. Defaults to `true`.

* `allow_assumed_signin` - (Optional) Enable sign in when user has been assumed by the account owner. Defaults to `false`.

* `provisioning` - (Optional) Settings regarding the app's provisioning ability.
  * `enabled` - (Required) Indicates if provisioning is enabled for this app.


* `parameters` - (Optional) a list of custom parameters for this app.
  * `param_key_name` - (Required) Name to represent the parameter in OneLogin.

  * `safe_entitlements_enabled` - (Optional) Indicates that the parameter is used to support creating entitlements using OneLogin Mappings. Defaults to `false`.

  * `user_attribute_mappings` - (Optional) A user attribute to map values from. For custom attributes prefix the name of the attribute with `custom_attribute_`.

  * `provisioned_entitlements` -  (Optional) Provisioned access entitlements for the app. Defaults to `false`.

  * `skip_if_blank` - (Optional)  Flag to let the SCIM provisioner know not include this value if it's blank. Defaults to `false`.

  * `user_attribute_macros` - (Optional) When `user_attribute_mappings` is set to `_macro_` this macro will be used to assign the parameter value.

  * `attributes_transformations` - (Optional) Describes how the app's attributes should be transformed.

  * `default_values` - (Optional) Default parameter values.

  * `include_in_saml_assertion` - (Optional) When true, this parameter will be included in a SAML assertion payload.

  * `label` - (Optional) The can only be set when creating a new parameter. It can not be updated.

  * `values` - (Optional) Parameter values.


* `configuration` - OIDC settings that control the authentication flow e.g. redirect urls and token settings.
  * `redirect_uri` - (Optional) The redirect_uri for the OIDC flow. Will be computed by API if not given.

  * `refresh_token_expiration_minutes` - (Optional) Number of minutes for the refresh token to be valid. Defaults to 1 minute.

  * `login_url` - (Optional) The login_url for the OIDC flow. Will be computed by API if not given.

  * `oidc_application_type` - (Optional) Must be one of one of `0` (Web) or `1` (Native/Mobile). Defaults to 0.

  * `token_endpoint_auth_method` - (Optional) Must be one of one of `0` (Basic) `1` (POST) `2` (Nonce/PKCE).

## Attributes Reference
* `id` - App's unique ID in OneLogin.

* `allow_assumed_signin` - App sign in allowed when user assumed by account administrator.

* `auth_method` - The apps auth method. Refer to the [OneLogin Apps Documentation](https://developers.onelogin.com/api-docs/2/apps/app-resource) for a comprehensive list of available auth methods.

* `connector_id` - ID of the apps underlying connector. Dictates the type of app (e.g. AWS Multi-Role App).

* `description` - App description.

* `icon_url` - The url for the app's icon.

* `name` - The app's name.

* `notes` - Notes about the app.

* `tab_id` - The tab in which to display in OneLogin portal.

* `updated_at` - Timestamp for app's last update.

* `created_at` - Timestamp for app's creation.

* `policy_id` - The security policy assigned to the app.

* `visible` - Indicates if the app is visible in the OneLogin portal.

* `parameters` - The parameters section contains parameterized attributes that have defined at the connector level as well as custom attributes that have been defined specifically for this app. Regardless of how they are defined, all parameters have the following attributes.
    * `attributes_transformations` - Describes how the app's attributes should be transformed.

    * `default_values` - Default parameter values.

    * `include_in_saml_assertion` - Dictates if the parameter needs to be included in a SAML assertion

    * `label` - The attribute label

    * `param_id` - The parameter ID.

    * `param_key_name` - The name of the parameter stored in OneLogin.

    * `provisioned_entitlements` -  Provisioned access entitlements for the app.

    * `safe_entitlements_enabled` -  Indicates whether entitlements can be created.

    * `skip_if_blank` - Flag to let the SCIM provisioner know not include this value if it's blank.

    * `user_attribute_macros` - When `user_attribute_mappings` is set to `_macro_` this macro will be used to assign the parameter value.

    * `user_attribute_mappings` - A user attribute to map values from. For custom attributes the name of the attribute is prefixed with `custom_attribute_`.

    * `values` - Parameter values.

* `provisioning` -  Settings regarding the app's provisioning ability.
    * `enabled` - Indicates if provisioning is enabled for this app.


* `configuration`
  * `redirect_uri` - The redirect_uri for the OIDC flow.

  * `refresh_token_expiration_minutes` - Number of minutes for the refresh token to be valid.

  * `login_url` - The login_url for the OIDC flow.

  * `oidc_application_type` - Indicates OIDC app type.

  * `token_endpoint_auth_method` - Indicates the token endpoint authentication method.

## Import

A OIDC App can be imported via the OneLogin App ID.

```
$ terraform import onelogin_oidc_apps.my_oidc_app <app id>
```
