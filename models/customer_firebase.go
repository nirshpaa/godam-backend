package models

import (
	"context"

	"cloud.google.com/go/firestore"
)

// CustomerFirebase represents a customer in Firebase
type CustomerFirebase struct {
	*FirebaseModel
	client *firestore.Client
}

// NewCustomerFirebase creates a new Firebase customer model
func NewCustomerFirebase(client *firestore.Client) *CustomerFirebase {
	return &CustomerFirebase{
		FirebaseModel: NewFirebaseModel("customers", client),
		client:        client,
	}
}

// FirebaseCustomer represents a customer in Firebase
type FirebaseCustomer struct {
	ID        string `json:"id"`
	CompanyID string `json:"company_id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Address   string `json:"address"`
	Phone     string `json:"phone"`
}

// List retrieves all customers
func (c *CustomerFirebase) List(ctx context.Context) ([]FirebaseCustomer, error) {
	var customers []FirebaseCustomer
	err := c.FirebaseModel.List(ctx, &customers)
	return customers, err
}

// Get retrieves a customer by ID
func (c *CustomerFirebase) Get(ctx context.Context, id string) (*FirebaseCustomer, error) {
	var customer FirebaseCustomer
	err := c.FirebaseModel.Get(ctx, id, &customer)
	if err != nil {
		return nil, err
	}
	customer.ID = id
	return &customer, nil
}

// Create creates a new customer
func (c *CustomerFirebase) Create(ctx context.Context, customer *FirebaseCustomer) (string, error) {
	return c.FirebaseModel.Create(ctx, customer)
}

// Update updates an existing customer
func (c *CustomerFirebase) Update(ctx context.Context, id string, customer *FirebaseCustomer) error {
	return c.FirebaseModel.Update(ctx, id, customer)
}

// Delete removes a customer
func (c *CustomerFirebase) Delete(ctx context.Context, id string) error {
	return c.FirebaseModel.Delete(ctx, id)
}

// FindByCompany retrieves all customers for a specific company
func (c *CustomerFirebase) FindByCompany(ctx context.Context, companyID string) ([]FirebaseCustomer, error) {
	query := c.ref.Where("company_id", "==", companyID)

	var customers []FirebaseCustomer
	err := c.FirebaseModel.Query(ctx, &query, &customers)
	return customers, err
}
