package auth

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"fmt"
	"github.com/vynious/go-travel/pkg"
	"google.golang.org/api/option"
	"log"
)

type Client struct {
	fbClient *auth.Client
}

func NewFirebaseClient() (*Client, error) {
	opt := option.WithCredentialsFile("../firebase.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %w", err)
	}

	fireAuth, err := app.Auth(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error connecting to firebase: %w", err)
	}
	return &Client{
		fbClient: fireAuth,
	}, nil
}

func (client *Client) CreateNewUser(ctx context.Context, name string, email string, password string) (*auth.UserRecord, error) {
	params := (&auth.UserToCreate{}).
		Email(email).
		EmailVerified(false).
		Password(password).
		DisplayName(name).
		Disabled(false)
	u, err := client.fbClient.CreateUser(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("error creating user: %w", err)
	}
	pkg.Log.Info("Successfully created user: %+v\n", *u)
	return u, nil
}

func (client *Client) GetUser(ctx context.Context, email string, password string) (*auth.UserRecord, error) {
	return nil, nil
}

func (client *Client) UpdateUserEmail(ctx context.Context, uid string, email string) (*auth.UserRecord, error) {
	params := (&auth.UserToUpdate{}).
		Email(email)
	u, err := client.fbClient.UpdateUser(ctx, uid, params)
	if err != nil {
		return nil, fmt.Errorf("error updating user: %v\n", err)
	}
	log.Printf("Successfully updated user: %v\n", u)
	return u, nil
}

func (client *Client) UpdateUserPassword(ctx context.Context, uid string, password string) (*auth.UserRecord, error) {
	params := (&auth.UserToUpdate{}).
		Password(password)
	u, err := client.fbClient.UpdateUser(ctx, uid, params)
	if err != nil {
		return nil, fmt.Errorf("error updating user: %v\n", err)
	}
	log.Printf("Successfully updated user: %v\n", u)
	return u, nil
}

func (client *Client) UpdateUserName(ctx context.Context, uid string, name string) (*auth.UserRecord, error) {
	params := (&auth.UserToUpdate{}).
		DisplayName(name)
	u, err := client.fbClient.UpdateUser(ctx, uid, params)
	if err != nil {
		return nil, fmt.Errorf("error updating user: %v\n", err)
	}
	log.Printf("Successfully updated user: %v\n", u)
	return u, nil
}

func (client *Client) DeleteUser(ctx context.Context, uid string) error {
	if err := client.fbClient.DeleteUser(ctx, uid); err != nil {
		return fmt.Errorf("error deleting user: %w", err)
	}
	pkg.Log.Info("Successfully deleted user")
	return nil
}

func (client *Client) CreateCustomToken(ctx context.Context, uid string) (string, error) {
	token, err := client.fbClient.CustomToken(ctx, uid)
	if err != nil {
		return "", fmt.Errorf("error creating custom token: %w", err)
	}
	return token, nil
}
