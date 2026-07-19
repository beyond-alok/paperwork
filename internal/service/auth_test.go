// package service

// import "testing"

// func TestRegister(t *testing.T) {
// 	t.R
// }

package service

import (
	"errors"
	"net/http"
	"testing"

	"github.com/beyond-alok/paperwork/internal/models"
	"github.com/beyond-alok/paperwork/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

// mockUserStore is a test double for repository.UserStore. Only GetByEmail
// and Create are exercised by Register; the rest are no-op stubs required
// to satisfy the interface.
type mockUserStore struct {
	getByEmailFunc func(email string) (*models.User, error)
	createFunc     func(user *models.User) error
	createdUser *models.User
}

var _ repository.UserStore = (*mockUserStore)(nil)

func (m *mockUserStore) GetByEmail(email string) (*models.User, error) {
	if m.getByEmailFunc != nil {
		return m.getByEmailFunc(email)
	}
	return nil, nil
}

func (m *mockUserStore) Create(user *models.User) error {
	if m.createFunc != nil {
		return m.createFunc(user)
	}
	m.createdUser = user
	return nil
}

// Unused by Register, stubbed only to satisfy repository.UserStore.
func (m *mockUserStore) GetAll()          {}
func (m *mockUserStore) GetById(id string) {}
func (m *mockUserStore) Update()          {}
func (m *mockUserStore) Delete()          {}

func TestAuthService_Register(t *testing.T) {
	t.Run("returns validation error for invalid input", func(t *testing.T) {
		store := &mockUserStore{}
		s := NewAuthService(store)

		req := &UserRegisterReq{
			Password: "short",
		}

		_, err := s.Register(req)
		if err == nil {
			t.Fatal("expected an error, got nil")
		}

		var svcErr Error
		if !errors.As(err, &svcErr) {
			t.Fatalf("expected error of type Error, got %T", err)
		}
		if svcErr.Code != http.StatusBadRequest {
			t.Errorf("expected code %d, got %d", http.StatusBadRequest, svcErr.Code)
		}
	})

	t.Run("returns error when user already exists", func(t *testing.T) {
		existing := &models.User{ID: "1", Email: "alok@example.com"}
		store := &mockUserStore{
			getByEmailFunc: func(email string) (*models.User, error) {
				return existing, nil
			},
		}
		s := NewAuthService(store)

		req := &UserRegisterReq{
			Name:     "Alok",
			Email:    "alok@example.com",
			Password: "Passw0rd!",
		}

		_, err := s.Register(req)
		if err == nil {
			t.Fatal("expected an error, got nil")
		}

		var svcErr Error
		if !errors.As(err, &svcErr) {
			t.Fatalf("expected error of type Error, got %T", err)
		}
		if svcErr.Code != http.StatusBadRequest {
			t.Errorf("expected code %d, got %d", http.StatusBadRequest, svcErr.Code)
		}
		if svcErr.Msg != "user already exist" {
			t.Errorf("expected msg %q, got %q", "user already exist", svcErr.Msg)
		}
	})

	t.Run("returns error when store.Create fails", func(t *testing.T) {
		store := &mockUserStore{
			getByEmailFunc: func(email string) (*models.User, error) {
				return nil, nil
			},
			createFunc: func(user *models.User) error {
				return errors.New("db down")
			},
		}
		s := NewAuthService(store)

		req := &UserRegisterReq{
			Name:     "Alok",
			Email:    "alok@example.com",
			Password: "Passw0rd!",
		}

		_, err := s.Register(req)
		if err == nil {
			t.Fatal("expected an error, got nil")
		}

		var svcErr Error
		if !errors.As(err, &svcErr) {
			t.Fatalf("expected error of type Error, got %T", err)
		}
		if svcErr.Code != http.StatusInternalServerError {
			t.Errorf("expected code %d, got %d", http.StatusInternalServerError, svcErr.Code)
		}
	})

	t.Run("registers a valid user successfully", func(t *testing.T) {
		store := &mockUserStore{
			getByEmailFunc: func(email string) (*models.User, error) {
				return nil, nil
			},
		}
		s := NewAuthService(store)

		req := &UserRegisterReq{
			Name:     "Alok",
			Email:    "alok@example.com",
			Password: "Passw0rd!",
		}

		msg, err := s.Register(req)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if msg != "user registered successfull" {
			t.Errorf("unexpected success message: %q", msg)
		}

		if store.createdUser == nil {
			t.Fatal("expected store.Create to be called")
		}
		if store.createdUser.Email != req.Email {
			t.Errorf("expected email %q, got %q", req.Email, store.createdUser.Email)
		}
		if store.createdUser.Name != req.Name {
			t.Errorf("expected name %q, got %q", req.Name, store.createdUser.Name)
		}
		if store.createdUser.PasswordHash == req.Password {
			t.Error("password was stored in plaintext, expected a bcrypt hash")
		}
		if err := bcrypt.CompareHashAndPassword(
			[]byte(store.createdUser.PasswordHash), []byte(req.Password),
		); err != nil {
			t.Errorf("stored hash does not match original password: %v", err)
		}
	})
}
