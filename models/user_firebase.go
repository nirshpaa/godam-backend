package models

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"firebase.google.com/go/v4/auth"
)

// FirebaseUser represents a user in Firebase
type FirebaseUser struct {
	ID        string `firestore:"id" json:"id"`
	Email     string `firestore:"email" json:"email"`
	Name      string `firestore:"name" json:"name"`
	Role      string `firestore:"role" json:"role"`
	CompanyID string `firestore:"company_id" json:"company_id"`
}

// UserFirebase represents the Firebase user operations
type UserFirebase struct {
	Client *firestore.Client
	Auth   *auth.Client
}

// NewUserFirebase creates a new UserFirebase instance
func NewUserFirebase(client *firestore.Client, authClient *auth.Client) *UserFirebase {
	if client == nil {
		log.Fatal("Firestore client is nil")
	}
	if authClient == nil {
		log.Fatal("Auth client is nil")
	}

	return &UserFirebase{
		Client: client,
		Auth:   authClient,
	}
}

// List returns all users
func (u *UserFirebase) List(ctx context.Context) ([]*FirebaseUser, error) {
	docs, err := u.Client.Collection("users").Documents(ctx).GetAll()
	if err != nil {
		return nil, err
	}

	users := make([]*FirebaseUser, len(docs))
	for i, doc := range docs {
		var user FirebaseUser
		if err := doc.DataTo(&user); err != nil {
			return nil, err
		}
		user.ID = doc.Ref.ID
		users[i] = &user
	}

	return users, nil
}

// Get returns a user by ID
func (u *UserFirebase) Get(ctx context.Context, id string) (*FirebaseUser, error) {
	doc, err := u.Client.Collection("users").Doc(id).Get(ctx)
	if err != nil {
		return nil, err
	}

	var user FirebaseUser
	if err := doc.DataTo(&user); err != nil {
		return nil, err
	}
	user.ID = doc.Ref.ID

	return &user, nil
}

// Create creates a new user in both Firestore and Firebase Auth
func (u *UserFirebase) Create(ctx context.Context, user *FirebaseUser) (string, error) {
	// Create user in Firebase Auth
	params := (&auth.UserToCreate{}).
		Email(user.Email).
		Password("defaultPassword") // You should generate a secure password

	authUser, err := u.Auth.CreateUser(ctx, params)
	if err != nil {
		return "", err
	}

	// Create user in Firestore
	user.ID = authUser.UID
	_, err = u.Client.Collection("users").Doc(authUser.UID).Set(ctx, user)
	if err != nil {
		// If Firestore creation fails, delete the Auth user
		u.Auth.DeleteUser(ctx, authUser.UID)
		return "", err
	}

	return authUser.UID, nil
}

// Update updates a user in both Firestore and Firebase Auth
func (u *UserFirebase) Update(ctx context.Context, id string, user *FirebaseUser) error {
	// Update user in Firebase Auth
	params := (&auth.UserToUpdate{}).
		Email(user.Email)

	_, err := u.Auth.UpdateUser(ctx, id, params)
	if err != nil {
		return err
	}

	// Update user in Firestore
	_, err = u.Client.Collection("users").Doc(id).Set(ctx, user)
	return err
}

// Delete deletes a user from both Firestore and Firebase Auth
func (u *UserFirebase) Delete(ctx context.Context, id string) error {
	// First delete from Firestore
	_, err := u.Client.Collection("users").Doc(id).Delete(ctx)
	if err != nil {
		return err
	}

	// Then delete from Firebase Auth
	if err := u.Auth.DeleteUser(ctx, id); err != nil {
		// If Auth deletion fails, log the error but don't return it
		// since the Firestore deletion was successful
		log.Printf("Warning: Failed to delete user from Firebase Auth: %v", err)
	}

	return nil
}
