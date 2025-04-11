package models

import (
	"context"

	"cloud.google.com/go/firestore"
)

// FirebaseCompany represents a company in Firebase
type FirebaseCompany struct {
	ID          string `json:"id" firestore:"id"`
	Name        string `json:"name" firestore:"name"`
	Address     string `json:"address" firestore:"address"`
	Phone       string `json:"phone" firestore:"phone"`
	Email       string `json:"email" firestore:"email"`
	Description string `json:"description" firestore:"description"`
}

// CompanyFirebase represents the Firebase operations for companies
type CompanyFirebase struct {
	Client *firestore.Client
}

// List returns all companies
func (u *CompanyFirebase) List(ctx context.Context) ([]FirebaseCompany, error) {
	var companies []FirebaseCompany
	iter := u.Client.Collection("companies").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err != nil {
			break
		}
		var company FirebaseCompany
		if err := doc.DataTo(&company); err != nil {
			return nil, err
		}
		company.ID = doc.Ref.ID
		companies = append(companies, company)
	}
	return companies, nil
}

// Get returns a company by ID
func (u *CompanyFirebase) Get(ctx context.Context, id string) (*FirebaseCompany, error) {
	doc, err := u.Client.Collection("companies").Doc(id).Get(ctx)
	if err != nil {
		return nil, err
	}
	var company FirebaseCompany
	if err := doc.DataTo(&company); err != nil {
		return nil, err
	}
	company.ID = doc.Ref.ID
	return &company, nil
}

// Create creates a new company
func (u *CompanyFirebase) Create(ctx context.Context, company *FirebaseCompany) (string, error) {
	doc, _, err := u.Client.Collection("companies").Add(ctx, company)
	if err != nil {
		return "", err
	}
	return doc.ID, nil
}

// Update updates a company
func (u *CompanyFirebase) Update(ctx context.Context, id string, company *FirebaseCompany) error {
	_, err := u.Client.Collection("companies").Doc(id).Set(ctx, company)
	return err
}

// Delete deletes a company
func (u *CompanyFirebase) Delete(ctx context.Context, id string) error {
	_, err := u.Client.Collection("companies").Doc(id).Delete(ctx)
	return err
}

// NewCompanyFirebase creates a new CompanyFirebase instance
func NewCompanyFirebase(client *firestore.Client) *CompanyFirebase {
	return &CompanyFirebase{
		Client: client,
	}
}

// Join joins a company
func (u *CompanyFirebase) Join(ctx context.Context, id string, userID string) error {
	_, err := u.Client.Collection("companies").Doc(id).Update(ctx, []firestore.Update{
		{
			Path:  "users",
			Value: firestore.ArrayUnion(userID),
		},
	})
	return err
}
