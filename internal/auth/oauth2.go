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
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		RedirectURL:  "http://localhost:3000/auth/google/callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}

	OAuthStateString = "random_string"
)

func GoogleLogin(c *fiber.Ctx) error {
	c.Redirect(GoogleOAuthConfig.AuthCodeURL(OAuthStateString))
	return nil
}

func GoogleCallback(c *fiber.Ctx) error {
	if c.Query("state") != OAuthStateString {
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
