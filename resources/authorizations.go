package types

import (
  "context"
  "github.com/carlmjohnson/requests"
  "log"
)


type AuthorizationsResources struct {}



// **Deprecation Notice:** GitHub Enterprise Server will discontinue the [OAuth Authorizations API](https://docs.github.com/enterprise-server@3.7/rest/reference/oauth-authorizations), which is used by integrations to create personal access tokens and OAuth tokens, and you must now create these tokens using our [web application flow](https://docs.github.com/enterprise-server@3.7/developers/apps/authorizing-oauth-apps#web-application-flow). The [OAuth Authorizations API](https://docs.github.com/enterprise-server@3.7/rest/reference/oauth-authorizations) will be removed on November, 13, 2020. For more information, including scheduled brownouts, see the [blog post](https://developer.github.com/changes/2020-02-14-deprecating-oauth-auth-endpoint/).

**Warning:** Apps must use the [web application flow](https://docs.github.com/enterprise-server@3.7/apps/building-oauth-apps/authorizing-oauth-apps/#web-application-flow) to obtain OAuth tokens that work with GitHub Enterprise Server SAML organizations. OAuth tokens created using the Authorizations API will be unable to access GitHub Enterprise Server SAML organizations. For more information, see the [blog post](https://developer.github.com/changes/2019-11-05-deprecated-passwords-and-authorizations-api).

Creates OAuth tokens using [Basic Authentication](https://docs.github.com/enterprise-server@3.7/rest/overview/other-authentication-methods#basic-authentication). If you have two-factor authentication setup, Basic Authentication for this endpoint requires that you use a one-time password (OTP) and your username and password instead of tokens. For more information, see "[Working with two-factor authentication](https://docs.github.com/enterprise-server@3.7/rest/overview/other-authentication-methods#working-with-two-factor-authentication)."

To create tokens for a particular OAuth application using this endpoint, you must authenticate as the user you want to create an authorization for and provide the app's client ID and secret, found on your OAuth application's settings page. If your OAuth application intends to create multiple tokens for one user, use `fingerprint` to differentiate between them.

You can also create tokens on GitHub Enterprise Server from the [personal access tokens settings](https://github.com/settings/tokens) page. Read more about these tokens in [the GitHub Help documentation](https://docs.github.com/enterprise-server@3.7/articles/creating-an-access-token-for-command-line-use).

Organizations that enforce SAML SSO require personal access tokens to be allowed. Read more about allowing tokens in [the GitHub Help documentation](https://docs.github.com/enterprise-server@3.7/articles/about-identity-and-access-management-with-saml-single-sign-on). 
func (oauthauthorizationscreateauthorization *AuthorizationsResources) OauthAuthorizationsCreateAuthorization()  {

    err := requests.
        URL("/authorizations").Fetch(context.Background())

    if err != nil {
         log.Println("error sending request", err)
    }}


type OauthAuthorizationsGetOrCreateAuthorizationForAppAndFingerprintParameters struct {
}



// **Deprecation Notice:** GitHub Enterprise Server will discontinue the [OAuth Authorizations API](https://docs.github.com/enterprise-server@3.7/rest/reference/oauth-authorizations/), which is used by integrations to create personal access tokens and OAuth tokens, and you must now create these tokens using our [web application flow](https://docs.github.com/enterprise-server@3.7/developers/apps/authorizing-oauth-apps#web-application-flow). The [OAuth Authorizations API](https://docs.github.com/enterprise-server@3.7/rest/reference/oauth-authorizations) will be removed on November, 13, 2020. For more information, including scheduled brownouts, see the [blog post](https://developer.github.com/changes/2020-02-14-deprecating-oauth-auth-endpoint/).

**Warning:** Apps must use the [web application flow](https://docs.github.com/enterprise-server@3.7/apps/building-oauth-apps/authorizing-oauth-apps/#web-application-flow) to obtain OAuth tokens that work with GitHub Enterprise Server SAML organizations. OAuth tokens created using the Authorizations API will be unable to access GitHub Enterprise Server SAML organizations. For more information, see the [blog post](https://developer.github.com/changes/2019-11-05-deprecated-passwords-and-authorizations-api).

This method will create a new authorization for the specified OAuth application, only if an authorization for that application and fingerprint do not already exist for the user. The URL includes the 20 character client ID for the OAuth app that is requesting the token. `fingerprint` is a unique string to distinguish an authorization from others created for the same client ID and user. It returns the user's existing authorization for the application if one is present. Otherwise, it creates and returns a new one.

If you have two-factor authentication setup, Basic Authentication for this endpoint requires that you use a one-time password (OTP) and your username and password instead of tokens. For more information, see "[Working with two-factor authentication](https://docs.github.com/enterprise-server@3.7/rest/overview/other-authentication-methods#working-with-two-factor-authentication)." 
func (oauthauthorizationsgetorcreateauthorizationforappandfingerprint *AuthorizationsResources) OauthAuthorizationsGetOrCreateAuthorizationForAppAndFingerprint(params OauthAuthorizationsGetOrCreateAuthorizationForAppAndFingerprintParameters)  {

    err := requests.
        URL("/authorizations/clients/{client_id}/{fingerprint}").Fetch(context.Background())

    if err != nil {
         log.Println("error sending request", err)
    }}


type OauthAuthorizationsGetAuthorizationParameters struct {
}



// **Deprecation Notice:** GitHub Enterprise Server will discontinue the [OAuth Authorizations API](https://docs.github.com/enterprise-server@3.7/rest/reference/oauth-authorizations), which is used by integrations to create personal access tokens and OAuth tokens, and you must now create these tokens using our [web application flow](https://docs.github.com/enterprise-server@3.7/apps/building-oauth-apps/authorizing-oauth-apps/#web-application-flow). The [OAuth Authorizations API](https://docs.github.com/enterprise-server@3.7/rest/reference/oauth-authorizations) will be removed on November, 13, 2020. For more information, including scheduled brownouts, see the [blog post](https://developer.github.com/changes/2020-02-14-deprecating-oauth-auth-endpoint/). 
func (oauthauthorizationsgetauthorization *AuthorizationsResources) OauthAuthorizationsGetAuthorization(params OauthAuthorizationsGetAuthorizationParameters)  {

    err := requests.
        URL("/authorizations/{authorization_id}").Fetch(context.Background())

    if err != nil {
         log.Println("error sending request", err)
    }}


type OauthAuthorizationsDeleteAuthorizationParameters struct {
}



// **Deprecation Notice:** GitHub Enterprise Server will discontinue the [OAuth Authorizations API](https://docs.github.com/enterprise-server@3.7/rest/reference/oauth-authorizations), which is used by integrations to create personal access tokens and OAuth tokens, and you must now create these tokens using our [web application flow](https://docs.github.com/enterprise-server@3.7/apps/building-oauth-apps/authorizing-oauth-apps/#web-application-flow). The [OAuth Authorizations API](https://docs.github.com/enterprise-server@3.7/rest/reference/oauth-authorizations) will be removed on November, 13, 2020. For more information, including scheduled brownouts, see the [blog post](https://developer.github.com/changes/2020-02-14-deprecating-oauth-auth-endpoint/). 
func (oauthauthorizationsdeleteauthorization *AuthorizationsResources) OauthAuthorizationsDeleteAuthorization(params OauthAuthorizationsDeleteAuthorizationParameters)  {

    err := requests.
        URL("/authorizations/{authorization_id}").Fetch(context.Background())

    if err != nil {
         log.Println("error sending request", err)
    }}


type OauthAuthorizationsUpdateAuthorizationParameters struct {
}



// **Deprecation Notice:** GitHub Enterprise Server will discontinue the [OAuth Authorizations API](https://docs.github.com/enterprise-server@3.7/rest/reference/oauth-authorizations/), which is used by integrations to create personal access tokens and OAuth tokens, and you must now create these tokens using our [web application flow](https://docs.github.com/enterprise-server@3.7/developers/apps/authorizing-oauth-apps#web-application-flow). The [OAuth Authorizations API](https://docs.github.com/enterprise-server@3.7/rest/reference/oauth-authorizations) will be removed on November, 13, 2020. For more information, including scheduled brownouts, see the [blog post](https://developer.github.com/changes/2020-02-14-deprecating-oauth-auth-endpoint/).

If you have two-factor authentication setup, Basic Authentication for this endpoint requires that you use a one-time password (OTP) and your username and password instead of tokens. For more information, see "[Working with two-factor authentication](https://docs.github.com/enterprise-server@3.7/rest/overview/other-authentication-methods#working-with-two-factor-authentication)."

You can only send one of these scope keys at a time. 
func (oauthauthorizationsupdateauthorization *AuthorizationsResources) OauthAuthorizationsUpdateAuthorization(params OauthAuthorizationsUpdateAuthorizationParameters)  {

    err := requests.
        URL("/authorizations/{authorization_id}").Fetch(context.Background())

    if err != nil {
         log.Println("error sending request", err)
    }}


type OauthAuthorizationsGetOrCreateAuthorizationForAppParameters struct {
}



// **Deprecation Notice:** GitHub Enterprise Server will discontinue the [OAuth Authorizations API](https://docs.github.com/enterprise-server@3.7/rest/reference/oauth-authorizations/), which is used by integrations to create personal access tokens and OAuth tokens, and you must now create these tokens using our [web application flow](https://docs.github.com/enterprise-server@3.7/developers/apps/authorizing-oauth-apps#web-application-flow). The [OAuth Authorizations API](https://docs.github.com/enterprise-server@3.7/rest/reference/oauth-authorizations) will be removed on November, 13, 2020. For more information, including scheduled brownouts, see the [blog post](https://developer.github.com/changes/2020-02-14-deprecating-oauth-auth-endpoint/).

**Warning:** Apps must use the [web application flow](https://docs.github.com/enterprise-server@3.7/apps/building-oauth-apps/authorizing-oauth-apps/#web-application-flow) to obtain OAuth tokens that work with GitHub Enterprise Server SAML organizations. OAuth tokens created using the Authorizations API will be unable to access GitHub Enterprise Server SAML organizations. For more information, see the [blog post](https://developer.github.com/changes/2019-11-05-deprecated-passwords-and-authorizations-api).

Creates a new authorization for the specified OAuth application, only if an authorization for that application doesn't already exist for the user. The URL includes the 20 character client ID for the OAuth app that is requesting the token. It returns the user's existing authorization for the application if one is present. Otherwise, it creates and returns a new one.

If you have two-factor authentication setup, Basic Authentication for this endpoint requires that you use a one-time password (OTP) and your username and password instead of tokens. For more information, see "[Working with two-factor authentication](https://docs.github.com/enterprise-server@3.7/rest/overview/other-authentication-methods#working-with-two-factor-authentication)."

**Deprecation Notice:** GitHub Enterprise Server will discontinue the [OAuth Authorizations API](https://docs.github.com/enterprise-server@3.7/rest/reference/oauth-authorizations/), which is used by integrations to create personal access tokens and OAuth tokens, and you must now create these tokens using our [web application flow](https://docs.github.com/enterprise-server@3.7/developers/apps/authorizing-oauth-apps#web-application-flow). The [OAuth Authorizations API](https://docs.github.com/enterprise-server@3.7/rest/reference/oauth-authorizations) will be removed on November, 13, 2020. For more information, including scheduled brownouts, see the [blog post](https://developer.github.com/changes/2020-02-14-deprecating-oauth-auth-endpoint/). 
func (oauthauthorizationsgetorcreateauthorizationforapp *AuthorizationsResources) OauthAuthorizationsGetOrCreateAuthorizationForApp(params OauthAuthorizationsGetOrCreateAuthorizationForAppParameters)  {

    err := requests.
        URL("/authorizations/clients/{client_id}").Fetch(context.Background())

    if err != nil {
         log.Println("error sending request", err)
    }}

