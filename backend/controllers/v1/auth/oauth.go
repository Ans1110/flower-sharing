package auth_controller

import (
	"context"
	"encoding/json"
	"flower-backend/libs"
	"flower-backend/models"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/google"
	"gorm.io/gorm"
)

// GoogleLogin initiates Google OAuth flow
// @Summary Google OAuth Login
// @Description Redirects to Google OAuth consent screen
// @Tags auth
// @Success 302 {string} string "Redirect to Google"
// @Router /auth/google [get]
func (ctrl *authController) GoogleLogin(c *gin.Context) {
	googleOauthConfig := &oauth2.Config{
		ClientID:     ctrl.cfg.GoogleClientID,
		ClientSecret: ctrl.cfg.GoogleClientSecret,
		RedirectURL:  ctrl.cfg.GoogleRedirectURL,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}

	// Generate state token for CSRF protection
	state := libs.GenerateRandomString(32)

	// Store state in session or cache (simplified here)
	c.SetCookie("oauth_state", state, 600, "/", "", false, true)

	url := googleOauthConfig.AuthCodeURL(state, oauth2.AccessTypeOffline)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

// GoogleCallback handles Google OAuth callback
// @Summary Google OAuth Callback
// @Description Handles callback from Google OAuth
// @Tags auth
// @Param code query string true "Authorization code"
// @Param state query string true "State token"
// @Success 302 {string} string "Redirect to frontend"
// @Router /auth/google/callback [get]
func (ctrl *authController) GoogleCallback(c *gin.Context) {
	// Verify state token
	state := c.Query("state")
	savedState, err := c.Cookie("oauth_state")
	if err != nil || state != savedState {
		ctrl.logger.Error("Invalid state token")
		c.Redirect(http.StatusTemporaryRedirect, ctrl.cfg.FrontendURL+"/login?error=invalid_state")
		return
	}

	// Clear state cookie
	c.SetCookie("oauth_state", "", -1, "/", "", false, true)

	code := c.Query("code")
	if code == "" {
		ctrl.logger.Error("No authorization code")
		c.Redirect(http.StatusTemporaryRedirect, ctrl.cfg.FrontendURL+"/login?error=no_code")
		return
	}

	googleOauthConfig := &oauth2.Config{
		ClientID:     ctrl.cfg.GoogleClientID,
		ClientSecret: ctrl.cfg.GoogleClientSecret,
		RedirectURL:  ctrl.cfg.GoogleRedirectURL,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}

	// Exchange code for token
	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		ctrl.logger.Errorf("Failed to exchange token: %v", err)
		c.Redirect(http.StatusTemporaryRedirect, ctrl.cfg.FrontendURL+"/login?error=token_exchange_failed")
		return
	}

	// Get user info from Google
	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		ctrl.logger.Errorf("Failed to get user info: %v", err)
		c.Redirect(http.StatusTemporaryRedirect, ctrl.cfg.FrontendURL+"/login?error=user_info_failed")
		return
	}
	defer resp.Body.Close()

	data, _ := io.ReadAll(resp.Body)
	var googleUser struct {
		ID            string `json:"id"`
		Email         string `json:"email"`
		VerifiedEmail bool   `json:"verified_email"`
		Name          string `json:"name"`
		Picture       string `json:"picture"`
	}

	if err := json.Unmarshal(data, &googleUser); err != nil {
		ctrl.logger.Errorf("Failed to parse user info: %v", err)
		c.Redirect(http.StatusTemporaryRedirect, ctrl.cfg.FrontendURL+"/login?error=parse_failed")
		return
	}

	// Find or create user
	user, err := ctrl.handleOAuthUser(googleUser.Email, googleUser.ID, "google", googleUser.Name, googleUser.Picture, string(data))
	if err != nil {
		ctrl.logger.Errorf("Failed to handle OAuth user: %v", err)
		c.Redirect(http.StatusTemporaryRedirect, ctrl.cfg.FrontendURL+"/login?error=user_creation_failed")
		return
	}

	// Generate JWT tokens
	accessToken := libs.GenerateAccessToken(user.ID)
	if accessToken == "" {
		ctrl.logger.Error("Failed to generate access token")
		c.Redirect(http.StatusTemporaryRedirect, ctrl.cfg.FrontendURL+"/login?error=token_generation_failed")
		return
	}

	refreshToken := libs.GenerateRefreshToken(user.ID)
	if refreshToken == "" {
		ctrl.logger.Error("Failed to generate refresh token")
		c.Redirect(http.StatusTemporaryRedirect, ctrl.cfg.FrontendURL+"/login?error=token_generation_failed")
		return
	}

	// Save refresh token
	tokenModel := &models.Token{
		UserID: user.ID,
		Token:  refreshToken,
	}
	if err := ctrl.svc.CreateToken(tokenModel); err != nil {
		ctrl.logger.Errorf("Failed to save refresh token: %v", err)
		c.Redirect(http.StatusTemporaryRedirect, ctrl.cfg.FrontendURL+"/login?error=token_save_failed")
		return
	}

	// Set cookies
	c.SetSameSite(http.SameSiteStrictMode)
	c.SetCookie("refreshToken", refreshToken, 7*24*60*60, "/", "", ctrl.cfg.GO_ENV == "production", true)
	c.SetCookie("role", user.Role, 7*24*60*60, "/", "", ctrl.cfg.GO_ENV == "production", true)

	// Redirect to frontend with tokens
	redirectURL := fmt.Sprintf("%s/auth/callback?access_token=%s&refresh_token=%s", ctrl.cfg.FrontendURL, accessToken, refreshToken)
	c.Redirect(http.StatusTemporaryRedirect, redirectURL)
}

// GithubLogin initiates GitHub OAuth flow
// @Summary GitHub OAuth Login
// @Description Redirects to GitHub OAuth consent screen
// @Tags auth
// @Success 302 {string} string "Redirect to GitHub"
// @Router /auth/github [get]
func (ctrl *authController) GithubLogin(c *gin.Context) {
	githubOauthConfig := &oauth2.Config{
		ClientID:     ctrl.cfg.GithubClientID,
		ClientSecret: ctrl.cfg.GithubClientSecret,
		RedirectURL:  ctrl.cfg.GithubRedirectURL,
		Scopes:       []string{"user:email"},
		Endpoint:     github.Endpoint,
	}

	// Generate state token for CSRF protection
	state := libs.GenerateRandomString(32)

	// Store state in session or cache (simplified here)
	c.SetCookie("oauth_state", state, 600, "/", "", false, true)

	url := githubOauthConfig.AuthCodeURL(state)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

// GithubCallback handles GitHub OAuth callback
// @Summary GitHub OAuth Callback
// @Description Handles callback from GitHub OAuth
// @Tags auth
// @Param code query string true "Authorization code"
// @Param state query string true "State token"
// @Success 302 {string} string "Redirect to frontend"
// @Router /auth/github/callback [get]
func (ctrl *authController) GithubCallback(c *gin.Context) {
	// Verify state token
	state := c.Query("state")
	savedState, err := c.Cookie("oauth_state")
	if err != nil || state != savedState {
		ctrl.logger.Error("Invalid state token")
		c.Redirect(http.StatusTemporaryRedirect, ctrl.cfg.FrontendURL+"/login?error=invalid_state")
		return
	}

	// Clear state cookie
	c.SetCookie("oauth_state", "", -1, "/", "", false, true)

	code := c.Query("code")
	if code == "" {
		ctrl.logger.Error("No authorization code")
		c.Redirect(http.StatusTemporaryRedirect, ctrl.cfg.FrontendURL+"/login?error=no_code")
		return
	}

	githubOauthConfig := &oauth2.Config{
		ClientID:     ctrl.cfg.GithubClientID,
		ClientSecret: ctrl.cfg.GithubClientSecret,
		RedirectURL:  ctrl.cfg.GithubRedirectURL,
		Scopes:       []string{"user:email"},
		Endpoint:     github.Endpoint,
	}

	// Exchange code for token
	token, err := githubOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		ctrl.logger.Errorf("Failed to exchange token: %v", err)
		c.Redirect(http.StatusTemporaryRedirect, ctrl.cfg.FrontendURL+"/login?error=token_exchange_failed")
		return
	}

	// Get user info from GitHub
	req, _ := http.NewRequest("GET", "https://api.github.com/user", nil)
	req.Header.Set("Authorization", "Bearer "+token.AccessToken)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		ctrl.logger.Errorf("Failed to get user info: %v", err)
		c.Redirect(http.StatusTemporaryRedirect, ctrl.cfg.FrontendURL+"/login?error=user_info_failed")
		return
	}
	defer resp.Body.Close()

	data, _ := io.ReadAll(resp.Body)
	var githubUser struct {
		ID        int64  `json:"id"`
		Login     string `json:"login"`
		Email     string `json:"email"`
		Name      string `json:"name"`
		AvatarURL string `json:"avatar_url"`
	}

	if err := json.Unmarshal(data, &githubUser); err != nil {
		ctrl.logger.Errorf("Failed to parse user info: %v", err)
		c.Redirect(http.StatusTemporaryRedirect, ctrl.cfg.FrontendURL+"/login?error=parse_failed")
		return
	}

	// If email is not public, fetch it separately
	if githubUser.Email == "" {
		emailReq, _ := http.NewRequest("GET", "https://api.github.com/user/emails", nil)
		emailReq.Header.Set("Authorization", "Bearer "+token.AccessToken)

		emailResp, err := client.Do(emailReq)
		if err == nil {
			defer emailResp.Body.Close()
			emailData, _ := io.ReadAll(emailResp.Body)

			var emails []struct {
				Email    string `json:"email"`
				Primary  bool   `json:"primary"`
				Verified bool   `json:"verified"`
			}

			if json.Unmarshal(emailData, &emails) == nil {
				for _, email := range emails {
					if email.Primary && email.Verified {
						githubUser.Email = email.Email
						break
					}
				}
			}
		}
	}

	if githubUser.Email == "" {
		ctrl.logger.Error("No email found for GitHub user")
		c.Redirect(http.StatusTemporaryRedirect, ctrl.cfg.FrontendURL+"/login?error=no_email")
		return
	}

	// Use login as name if name is empty
	if githubUser.Name == "" {
		githubUser.Name = githubUser.Login
	}

	// Find or create user
	user, err := ctrl.handleOAuthUser(githubUser.Email, fmt.Sprintf("%d", githubUser.ID), "github", githubUser.Name, githubUser.AvatarURL, string(data))
	if err != nil {
		ctrl.logger.Errorf("Failed to handle OAuth user: %v", err)
		c.Redirect(http.StatusTemporaryRedirect, ctrl.cfg.FrontendURL+"/login?error=user_creation_failed")
		return
	}

	// Generate JWT tokens
	accessToken := libs.GenerateAccessToken(user.ID)
	if accessToken == "" {
		ctrl.logger.Error("Failed to generate access token")
		c.Redirect(http.StatusTemporaryRedirect, ctrl.cfg.FrontendURL+"/login?error=token_generation_failed")
		return
	}

	refreshToken := libs.GenerateRefreshToken(user.ID)
	if refreshToken == "" {
		ctrl.logger.Error("Failed to generate refresh token")
		c.Redirect(http.StatusTemporaryRedirect, ctrl.cfg.FrontendURL+"/login?error=token_generation_failed")
		return
	}

	// Save refresh token
	tokenModel := &models.Token{
		UserID: user.ID,
		Token:  refreshToken,
	}
	if err := ctrl.svc.CreateToken(tokenModel); err != nil {
		ctrl.logger.Errorf("Failed to save refresh token: %v", err)
		c.Redirect(http.StatusTemporaryRedirect, ctrl.cfg.FrontendURL+"/login?error=token_save_failed")
		return
	}

	// Set cookies
	c.SetSameSite(http.SameSiteStrictMode)
	c.SetCookie("refreshToken", refreshToken, 7*24*60*60, "/", "", ctrl.cfg.GO_ENV == "production", true)
	c.SetCookie("role", user.Role, 7*24*60*60, "/", "", ctrl.cfg.GO_ENV == "production", true)

	// Redirect to frontend with tokens
	redirectURL := fmt.Sprintf("%s/auth/callback?access_token=%s&refresh_token=%s", ctrl.cfg.FrontendURL, accessToken, refreshToken)
	c.Redirect(http.StatusTemporaryRedirect, redirectURL)
}

// handleOAuthUser finds or creates a user from OAuth provider
func (ctrl *authController) handleOAuthUser(email, providerID, provider, name, avatar, providerData string) (*models.User, error) {
	// Try to find existing user by email
	user, err := ctrl.svc.GetUserByEmail(email)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// User doesn't exist, create new one

			// Determine user role based on whitelist
			role := "user"
			for _, adminEmail := range ctrl.cfg.WhiteListAdminEmails {
				if email == adminEmail {
					role = "admin"
					break
				}
			}

			newUser := models.User{
				Email:        email,
				Username:     name,
				Avatar:       avatar,
				Provider:     provider,
				ProviderID:   providerID,
				ProviderData: providerData,
				Role:         role,
				CreatedAt:    time.Now(),
			}

			createdUser, err := ctrl.svc.CreateUser(newUser)
			if err != nil {
				return nil, err
			}

			return createdUser, nil
		}
		return nil, err
	}

	// User exists, update OAuth info if needed
	if user.Provider == "local" || user.Provider == "" {
		// Link OAuth account to existing local account
		user.Provider = provider
		user.ProviderID = providerID
		user.ProviderData = providerData
		if user.Avatar == "" {
			user.Avatar = avatar
		}

		// Update user using the existing update method
		updates := map[string]any{
			"provider":      provider,
			"provider_id":   providerID,
			"provider_data": providerData,
		}
		if user.Avatar == "" {
			updates["avatar"] = avatar
		}

		updatedUser, err := ctrl.svc.UpdateUserByIDWithSelect(user.ID, updates, nil, []string{"id", "email", "username", "avatar", "provider", "provider_id", "provider_data", "role"})
		if err != nil {
			return nil, err
		}
		return updatedUser, nil
	}

	return user, nil
}
