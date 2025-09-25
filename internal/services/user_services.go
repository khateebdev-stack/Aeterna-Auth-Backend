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

	query := `INSERT INTO users(email, hashed_password, first_name, last_name, profile_picture) VALUES ($1, $2, $3, $4, $5)`
	_, err = s.DB.Exec(query, user.Email, hashedPassword, user.FirstName, user.LastName, user.ProfilePicture)
	if err != nil {
		return fmt.Errorf("Failed to insert user into database: %w", err)
	}

	return nil
}

// this function authenticates a user and returns their data
func (s *UserService) LoginUser(email, password string) (*models.User, error) {
	// 1. Fetch the user from the database by email
	var user models.User
	query := `SELECT id, email, hashed_password, first_name, last_name, profile_picture,  created_at FROM users WHERE email = $1` // Corrected 'if' to 'id'
	row := s.DB.QueryRow(query, email)

	err := row.Scan(&user.ID, &user.Email, &user.HashedPassword, &user.FirstName, &user.LastName, &user.ProfilePicture, &user.CreatedAt)
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

func (s *UserService) DeleteUser(email string) error {
	query := `DELETE FROM users WHERE email = $1`
	result, err := s.DB.Exec(query, email)

	if err != nil {
		return fmt.Errorf("Failed to delete user: %w", err)

	}

	rowAffected, _ := result.RowsAffected()
	if rowAffected == 0 {
		return fmt.Errorf("User not Found")

	}
	return nil
}

func (s *UserService) UpdateUserEmail(currentEmail, newEmail string) error{

	query:=`UPDATE users SET email = $1 WHere email = $2`
	_, err:= s.DB.Exec(query, newEmail, currentEmail)
	if err !=nil{
		return fmt.Errorf("Failed to update email; %W", err)
	}

	return nil
}



func (s* UserService) UpdateUserPassword(email, newPassword string) error{

	hashedPassword, err:= utils.HashPassword(newPassword)
	if err != nil{
		return fmt.Errorf("Failed to hash new Password: %w", err)
	}

	query:=`UPDATE users SET hashed_password = $1 WHERE email = $2`

	result, err := s.DB.Exec(query, hashedPassword, email)
	if err != nil {
		return fmt.Errorf("Failed to update password: %w", err)

	}

	rowsAffected, _ := result.RowsAffected()
if rowsAffected == 0 {
	return fmt.Errorf("user not found")
}
	return nil
}



func (s *UserService) CheckEmail( email string) (bool, error){
	var exists bool
	query:= `SELECT EXIST(SELECT 1 FROM users WHERE email = $1)`
	err:=s.DB.QueryRow(query, email).Scan(&exists)
	if err != nil && err!=sql.ErrNoRows{
		return false, fmt.Errorf("Failed to check existing email: %w" , err)
	}

	return exists, nil
}