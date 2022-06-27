package Auth

import (
	"context"
	"fmt"
	"os"
	"time"
	"cloud.google.com/go/pubsub"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
)

// oauthClient shows how to use an OAuth client ID to authenticate as an end-user.
func OauthClient(c*fiber.Ctx) error {
	ctx := context.Background()

	// Please make sure the redirect URL is the same as the one you specified when you
	// created the client ID.
	redirectURL := os.Getenv("OAUTH2_CALLBACK")
	if redirectURL == "" {
		redirectURL = "http://localhost:1234/google/callback"
	}

	config := &oauth2.Config{
		ClientID:     "238983540674-0i6lp85nrk7mvdai5ojpdhbg6updm00k.apps.googleusercontent.com",
		ClientSecret: "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",
		RedirectURL:  redirectURL,
		Scopes:       []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint:     google.Endpoint,
	}


	// Dummy authorization flow to read auth code from stdin.
	authURL := config.AuthCodeURL("randomstate")
	fmt.Println(authURL)

	// Read the authentication code from the command line
	var code string
	fmt.Scanln(&code)

	// Exchange auth code for OAuth token.
	token, err := config.Exchange(ctx, code)
	if err != nil {
		return fmt.Errorf("config.Exchange: %v", err)
	}
	client, err := pubsub.NewClient(ctx, "My First Project", option.WithTokenSource(config.TokenSource(ctx, token)))
	if err != nil {
		return fmt.Errorf("pubsub.NewClient: %v", err)
	}
	defer client.Close()

	// Use the authenticated client.
	_ = client

	return c.JSON(authURL)
}
func Genetare(c*fiber.Ctx) error{

	state:=Randstring(10)
	cookie := fiber.Cookie{

		Name:     "outstate",
		Value:    state,
		Expires:  time.Now().Add(time.Minute * 2),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	return nil

}
