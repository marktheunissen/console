// This file is part of MinIO Console Server
// Copyright (c) 2021 MinIO, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package oauth2

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/minio/console/pkg/auth/token"
	"github.com/minio/console/pkg/auth/utils"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/minio/minio-go/v7/pkg/set"
	"github.com/minio/pkg/v3/env"
	"golang.org/x/oauth2"
	xoauth2 "golang.org/x/oauth2"
)

type Configuration interface {
	Exchange(ctx context.Context, code string, opts ...xoauth2.AuthCodeOption) (*xoauth2.Token, error)
	AuthCodeURL(state string, opts ...xoauth2.AuthCodeOption) string
	PasswordCredentialsToken(ctx context.Context, username, password string) (*xoauth2.Token, error)
	Client(ctx context.Context, t *xoauth2.Token) *http.Client
	TokenSource(ctx context.Context, t *xoauth2.Token) xoauth2.TokenSource
}

type Config struct {
	xoauth2.Config
}

// DiscoveryDoc - parses the output from openid-configuration
// for example https://accounts.google.com/.well-known/openid-configuration
type DiscoveryDoc struct {
	Issuer                           string   `json:"issuer,omitempty"`
	AuthEndpoint                     string   `json:"authorization_endpoint,omitempty"`
	TokenEndpoint                    string   `json:"token_endpoint,omitempty"`
	UserInfoEndpoint                 string   `json:"userinfo_endpoint,omitempty"`
	RevocationEndpoint               string   `json:"revocation_endpoint,omitempty"`
	JwksURI                          string   `json:"jwks_uri,omitempty"`
	ResponseTypesSupported           []string `json:"response_types_supported,omitempty"`
	SubjectTypesSupported            []string `json:"subject_types_supported,omitempty"`
	IDTokenSigningAlgValuesSupported []string `json:"id_token_signing_alg_values_supported,omitempty"`
	ScopesSupported                  []string `json:"scopes_supported,omitempty"`
	TokenEndpointAuthMethods         []string `json:"token_endpoint_auth_methods_supported,omitempty"`
	ClaimsSupported                  []string `json:"claims_supported,omitempty"`
	CodeChallengeMethodsSupported    []string `json:"code_challenge_methods_supported,omitempty"`
}

func (ac Config) Exchange(ctx context.Context, code string, opts ...xoauth2.AuthCodeOption) (*xoauth2.Token, error) {
	return ac.Config.Exchange(ctx, code, opts...)
}

func (ac Config) AuthCodeURL(state string, opts ...xoauth2.AuthCodeOption) string {
	return ac.Config.AuthCodeURL(state, opts...)
}

func (ac Config) PasswordCredentialsToken(ctx context.Context, username, password string) (*xoauth2.Token, error) {
	return ac.Config.PasswordCredentialsToken(ctx, username, password)
}

func (ac Config) Client(ctx context.Context, t *xoauth2.Token) *http.Client {
	return ac.Config.Client(ctx, t)
}

func (ac Config) TokenSource(ctx context.Context, t *xoauth2.Token) xoauth2.TokenSource {
	return ac.Config.TokenSource(ctx, t)
}

// Provider is a wrapper of the oauth2 configuration and the oidc provider
type Provider struct {
	// oauth2Config is an interface configuration that contains the following fields
	// Config{
	// 	 IDPName string
	//	 ClientSecret string
	//	 RedirectURL string
	//	 Endpoint oauth2.Endpoint
	//	 Scopes []string
	// }
	// - IDPName is the public identifier for this application
	// - ClientSecret is a shared secret between this application and the authorization server
	// - RedirectURL is the URL to redirect users going through
	//   the OAuth flow, after the resource owner's URLs.
	// - Endpoint contains the resource server's token endpoint
	//   URLs. These are constants specific to each server and are
	//   often available via site-specific packages, such as
	//   google.Endpoint or github.Endpoint.
	// - Scopes specifies optional requested permissions.
	IDPName string
	// if enabled means that we need extrace access_token as well
	UserInfo     bool
	RefreshToken string
	oauth2Config Configuration
	client       *http.Client
}

const (
	schemeHTTP  = "http"
	schemeHTTPS = "https"
)

func getLoginCallbackURL(r *http.Request) string {
	scheme := getSourceScheme(r)
	if scheme == "" {
		if r.TLS != nil {
			scheme = schemeHTTPS
		} else {
			scheme = schemeHTTP
		}
	}

	redirectURL := scheme + "://" + r.Host + "/oauth_callback"
	_, err := url.Parse(redirectURL)
	if err != nil {
		panic(err)
	}
	return redirectURL
}

var requiredResponseTypes = set.CreateStringSet("code")

var defaultScopes = []string{"openid", "profile", "email"}

// NewOauth2ProviderClientByName returns a provider if present specified by the input name of the provider.
func (ois OpenIDPCfg) NewOauth2ProviderClientByName(name string, scopes []string, r *http.Request, clnt *http.Client) (provider *Provider, err error) {
	oi, ok := ois[name]
	if !ok {
		return nil, fmt.Errorf("%s IDP provider does not exist", name)
	}
	return oi.GetOauth2Provider(name, scopes, r, clnt)
}

// NewOauth2ProviderClient instantiates a new oauth2 client using the
// `OpenIDPCfg` configuration struct. It returns a *Provider object that
// contains the necessary configuration to initiate an oauth2 authentication
// flow.
//
// We only support Authentication with the Authorization Code Flow - spec:
// https://openid.net/specs/openid-connect-core-1_0.html#CodeFlowAuth
func (ois OpenIDPCfg) NewOauth2ProviderClient(scopes []string, r *http.Request, clnt *http.Client) (provider *Provider, providerCfg ProviderConfig, err error) {
	for name, oi := range ois {
		provider, err = oi.GetOauth2Provider(name, scopes, r, clnt)
		if err != nil {
			// Upon error look for the next IDP.
			continue
		}
		// Upon success return right away.
		providerCfg = oi
		break
	}
	return provider, providerCfg, err
}

type User struct {
	AppMetadata       map[string]interface{} `json:"app_metadata"`
	Blocked           bool                   `json:"blocked"`
	CreatedAt         string                 `json:"created_at"`
	Email             string                 `json:"email"`
	EmailVerified     bool                   `json:"email_verified"`
	FamilyName        string                 `json:"family_name"`
	GivenName         string                 `json:"given_name"`
	Identities        []interface{}          `json:"identities"`
	LastIP            string                 `json:"last_ip"`
	LastLogin         string                 `json:"last_login"`
	LastPasswordReset string                 `json:"last_password_reset"`
	LoginsCount       int                    `json:"logins_count"`
	MultiFactor       string                 `json:"multifactor"`
	Name              string                 `json:"name"`
	Nickname          string                 `json:"nickname"`
	PhoneNumber       string                 `json:"phone_number"`
	PhoneVerified     bool                   `json:"phone_verified"`
	Picture           string                 `json:"picture"`
	UpdatedAt         string                 `json:"updated_at"`
	UserID            string                 `json:"user_id"`
	UserMetadata      map[string]interface{} `json:"user_metadata"`
	Username          string                 `json:"username"`
}

// StateKeyFunc - is a function that returns a key used in OAuth Authorization
// flow state generation and verification.
type StateKeyFunc func() []byte

// VerifyIdentity will contact the configured IDP to the user identity based on the authorization code and state
// if the user is valid, then it will contact MinIO to get valid sts credentials based on the identity provided by the IDP
func (client *Provider) VerifyIdentity(ctx context.Context, code, state, roleARN string, keyFunc StateKeyFunc) (*credentials.Credentials, error) {
	// verify the provided state is valid (prevents CSRF attacks)
	if err := validateOauth2State(state, keyFunc); err != nil {
		return nil, err
	}
	getWebTokenExpiry := func() (*credentials.WebIdentityToken, error) {
		customCtx := context.WithValue(ctx, oauth2.HTTPClient, client.client)
		oauth2Token, err := client.oauth2Config.Exchange(customCtx, code)
		if err != nil {
			return nil, err
		}
		if !oauth2Token.Valid() {
			return nil, errors.New("invalid token")
		}
		client.RefreshToken = oauth2Token.RefreshToken

		envStsDuration := env.Get(token.ConsoleSTSDuration, "")
		stsDuration, err := time.ParseDuration(envStsDuration)

		expiration := 12 * time.Hour

		if err == nil && stsDuration > 0 {
			expiration = stsDuration
		} else {
			// Use the expiration configured in the token itself if it is closer than the configured value
			if exp := oauth2Token.Expiry.Sub(time.Now().UTC()); exp < expiration {
				expiration = exp
			}
		}

		// Minimum duration in S3 spec is 15 minutes, do not bother returning
		// an error to the user and force the minimum duration instead
		if expiration < 900*time.Second {
			expiration = 900 * time.Second
		}

		idToken := oauth2Token.Extra("id_token")
		if idToken == nil {
			return nil, errors.New("missing id_token")
		}
		token := &credentials.WebIdentityToken{
			Token:  idToken.(string),
			Expiry: int(expiration.Seconds()),
		}
		if client.UserInfo { // look for access_token only if userinfo is requested.
			accessToken := oauth2Token.Extra("access_token")
			if accessToken == nil {
				return nil, errors.New("missing access_token")
			}
			token.AccessToken = accessToken.(string)
			refreshToken := oauth2Token.Extra("refresh_token")
			if refreshToken != nil {
				token.RefreshToken = refreshToken.(string)
			} else { //nolint:revive,staticcheck
				// TODO in Nov 2026 : add an error when the refresh token is not found.
				// This is not done yet because users may not have access_offline scope
				// and this may break their deployments
			}

		}
		return token, nil
	}
	stsEndpoint := GetSTSEndpoint()

	sts := credentials.New(&credentials.STSWebIdentity{
		Client:              client.client,
		STSEndpoint:         stsEndpoint,
		GetWebIDTokenExpiry: getWebTokenExpiry,
		RoleARN:             roleARN,
	})
	return sts, nil
}

// VerifyIdentityForOperator will contact the configured IDP and validate the user identity based on the authorization code and state
func (client *Provider) VerifyIdentityForOperator(ctx context.Context, code, state string, keyFunc StateKeyFunc) (*xoauth2.Token, error) {
	// verify the provided state is valid (prevents CSRF attacks)
	if err := validateOauth2State(state, keyFunc); err != nil {
		return nil, err
	}
	customCtx := context.WithValue(ctx, oauth2.HTTPClient, client.client)
	oauth2Token, err := client.oauth2Config.Exchange(customCtx, code)
	if err != nil {
		return nil, err
	}
	if !oauth2Token.Valid() {
		return nil, errors.New("invalid token")
	}
	return oauth2Token, nil
}

// validateOauth2State validates the provided state was originated using the same
// instance (or one configured using the same secrets) of Console, this is basically used to prevent CSRF attacks
// https://security.stackexchange.com/questions/20187/oauth2-cross-site-request-forgery-and-state-parameter
func validateOauth2State(state string, keyFunc StateKeyFunc) error {
	// state contains a base64 encoded string that may ends with "==", the browser encodes that to "%3D%3D"
	// query unescape is need it before trying to decode the base64 string
	encodedMessage, err := url.QueryUnescape(state)
	if err != nil {
		return err
	}
	// decode the state parameter value
	message, err := base64.StdEncoding.DecodeString(encodedMessage)
	if err != nil {
		return err
	}
	s := strings.Split(string(message), ":")
	// Validate that the decoded message has the right format "message:hmac"
	if len(s) != 2 {
		return fmt.Errorf("invalid number of tokens, expected only 2, got %d instead", len(s))
	}
	// extract the state and hmac
	incomingState, incomingHmac := s[0], s[1]
	// validate that hmac(incomingState + pbkdf2(secret, salt)) == incomingHmac
	if calculatedHmac := utils.ComputeHmac256(incomingState, keyFunc()); calculatedHmac != incomingHmac {
		return fmt.Errorf("oauth2 state is invalid, expected %s, got %s", calculatedHmac, incomingHmac)
	}
	return nil
}

// parseDiscoveryDoc parses a discovery doc from an OAuth provider
// into a DiscoveryDoc struct that have the correct endpoints
func parseDiscoveryDoc(ctx context.Context, ustr string, httpClient *http.Client) (DiscoveryDoc, error) {
	d := DiscoveryDoc{}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ustr, nil)
	if err != nil {
		return d, err
	}
	clnt := http.Client{
		Transport: httpClient.Transport,
	}
	resp, err := clnt.Do(req)
	if err != nil {
		return d, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return d, err
	}
	dec := json.NewDecoder(resp.Body)
	if err = dec.Decode(&d); err != nil {
		return d, err
	}
	return d, nil
}

// GetRandomStateWithHMAC computes message + hmac(message, pbkdf2(key, salt)) to be used as state during the oauth authorization
func GetRandomStateWithHMAC(length int, keyFunc StateKeyFunc) string {
	state := utils.RandomCharString(length)
	hmac := utils.ComputeHmac256(state, keyFunc())
	return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", state, hmac)))
}

type LoginURLParams struct {
	State   string `json:"state"`
	IDPName string `json:"idp_name"`
}

// GenerateLoginURL returns a new login URL based on the configured IDP
func (client *Provider) GenerateLoginURL(keyFunc StateKeyFunc, iDPName string) string {
	// generates random state and sign it using HMAC256
	state := GetRandomStateWithHMAC(25, keyFunc)

	configureID := "_"

	if iDPName != "" {
		configureID = iDPName
	}

	lgParams := LoginURLParams{
		State:   state,
		IDPName: configureID,
	}

	jsonEnc, err := json.Marshal(lgParams)
	if err != nil {
		return ""
	}

	stEncode := base64.StdEncoding.EncodeToString(jsonEnc)
	loginURL := client.oauth2Config.AuthCodeURL(stEncode)

	return strings.TrimSpace(loginURL)
}
