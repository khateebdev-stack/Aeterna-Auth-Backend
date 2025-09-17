package services

import (
    "aeterna-auth/internal/models"
    "aeterna-auth/pkg/utils"
    "database/sql"
    "fmt"
)

type UserService struct {
    DB *sql.DB
}

// this function creates and returns a new UserService
func NewUserService(db *sql.DB) *UserService {
    return &UserService{DB: db}
}

// this function is used to register the user by hashing the password and inserting into the database
func (s *UserService) RegisterUser(user *models.RegisterRequest) error {
    hashedPassword, err := utils.HashPassword(user.Password) // Call the corrected function name
    if err != nil {
        return fmt.Errorf("Failed to hash password: %w", err)
    }

    query := `INSERT INTO users(email, hashed_password) VALUES ($1, $2)`
    _, err = s.DB.Exec(query, user.Email, hashedPassword)
    if err != nil {
        return fmt.Errorf("Failed to insert user into database: %w", err)
    }

    return nil
}

// this function authenticates a user and returns their data
func (s *UserService) LoginUser(email, password string) (*models.User, error) {
    // 1. Fetch the user from the database by email
    var user models.User
    query := `SELECT id, email, hashed_password, created_at FROM users WHERE email = $1` // Corrected 'if' to 'id'
    row := s.DB.QueryRow(query, email)

    err := row.Scan(&user.ID, &user.Email, &user.HashedPassword, &user.CreatedAt)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, fmt.Errorf("user not found")
        }
        return nil, fmt.Errorf("failed to fetch user: %w", err) // Added a return statement here
    }

    // 2. Compare the provided password with the stored hash
    if !utils.CheckPasswordHash(password, user.HashedPassword) {
        return nil, fmt.Errorf("invalid password")
    }

    // 3. Return the authenticated user
    return &user, nil
}