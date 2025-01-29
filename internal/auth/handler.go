package auth

import (
	"encoding/json"
	"io"
	"os"

	"github.com/Mariano-JR/auth/internal/user"
	"github.com/gofiber/fiber/v2"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"gorm.io/gorm"
)

var (
	GoogleOAuthConfig = &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  "http://localhost:3000/auth/google/callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
	GoogleOAuthState = "random_string"

	GitHubOAuthConfig = &oauth2.Config{
		ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
		ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
		RedirectURL:  "http://localhost:3000/auth/github/callback",
		Scopes:       []string{"user:email"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://github.com/login/oauth/authorize",
			TokenURL: "https://github.com/login/oauth/access_token",
		},
	}
	GitHubOAuthState = "random_github_state"
)

func GoogleLogin(c *fiber.Ctx) error {
	c.Redirect(GoogleOAuthConfig.AuthCodeURL(GoogleOAuthState))
	return nil
}

func GoogleCallback(c *fiber.Ctx) error {
	if c.Query("state") != GoogleOAuthState {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid OAuth state")
	}

	code := c.Query("code")
	token, err := GoogleOAuthConfig.Exchange(c.Context(), code)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to exchange token: " + err.Error())
	}

	client := GoogleOAuthConfig.Client(c.Context(), token)
	userInfo, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to get user info: " + err.Error())
	}

	body, err := io.ReadAll(userInfo.Body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to read user info: " + err.Error())
	}

	var googleUser GoogleUser
	if err := json.Unmarshal(body, &googleUser); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to unmarshal user info: " + err.Error())
	}

	_, err = user.GetUser(googleUser.Email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			newUser := user.User{
				Email: googleUser.Email,
				Name:  googleUser.Name,
			}

			if _, err := user.Save(newUser.Email, newUser.Name, ""); err != nil {
				return c.Status(fiber.StatusInternalServerError).SendString("Failed to save user: " + err.Error())
			}
		} else {
			return c.Status(fiber.StatusInternalServerError).SendString("Database error: " + err.Error())
		}
	}

	u, _ := user.GetUser(googleUser.Email)

	c.Cookie(&fiber.Cookie{
		Name:  "user_id",
		Value: u.ID,
		Path:  "/",
	})
	c.Cookie(&fiber.Cookie{
		Name:  "user_email",
		Value: u.Email,
		Path:  "/",
	})
	c.Cookie(&fiber.Cookie{
		Name:  "user_name",
		Value: u.Name,
		Path:  "/",
	})

	c.Redirect("/home.html")
	return nil
}

func GithubLogin(c *fiber.Ctx) error {
	c.Redirect(GitHubOAuthConfig.AuthCodeURL(GitHubOAuthState))
	return nil
}

func GithubCallback(c *fiber.Ctx) error {
	if c.Query("state") != GitHubOAuthState {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid OAuth state")
	}

	code := c.Query("code")
	token, err := GitHubOAuthConfig.Exchange(c.Context(), code)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to exchange token: " + err.Error())
	}

	client := GitHubOAuthConfig.Client(c.Context(), token)
	userInfo, err := client.Get("https://api.github.com/user")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to get user info: " + err.Error())
	}

	body, err := io.ReadAll(userInfo.Body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to read user info: " + err.Error())
	}

	var githubUser GithubUser
	if err := json.Unmarshal(body, &githubUser); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to unmarshal user info: " + err.Error())
	}

	u, err := user.GetUser(githubUser.Email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			newUser := user.User{
				Email: githubUser.Email,
				Name:  githubUser.Name,
			}

			if _, err := user.Save(newUser.Email, newUser.Name, ""); err != nil {
				return c.Status(fiber.StatusInternalServerError).SendString("Failed to save user: " + err.Error())
			}
		} else {
			return c.Status(fiber.StatusInternalServerError).SendString("Database error: " + err.Error())
		}
	}

	c.Cookie(&fiber.Cookie{
		Name:  "user_id",
		Value: u.ID,
		Path:  "/",
	})
	c.Cookie(&fiber.Cookie{
		Name:  "user_email",
		Value: u.Email,
		Path:  "/",
	})
	c.Cookie(&fiber.Cookie{
		Name:  "user_name",
		Value: u.Name,
		Path:  "/",
	})

	c.Redirect("/home.html")
	return nil
}
