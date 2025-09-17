package utils

import (
    "crypto/rand"
    "encoding/base64"
    "fmt"
    "strings"

    "golang.org/x/crypto/argon2"
)

// HashPassword generates a secure password hash.
func HashPassword(password string) (string, error) { // Corrected the function name
    salt := make([]byte, 16)
    if _, err := rand.Read(salt); err != nil {
        return "", err
    }

    time := uint32(1)
    memory := uint32(64 * 1024)
    threads := uint8(4)
    keyLen := uint32(32)

    hash := argon2.IDKey([]byte(password), salt, time, memory, threads, keyLen)

    encodedHash := base64.StdEncoding.EncodeToString(hash)
    encodedSalt := base64.StdEncoding.EncodeToString(salt)

    return fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, memory, time, threads, encodedSalt, encodedHash), nil
}

// HashPasswordWithSalt is a helper function to hash with a specific salt.
func HashPasswordWithSalt(password string, salt []byte) (string, error) {
    time := uint32(1)
    memory := uint32(64 * 1024)
    threads := uint8(4)
    keyLen := uint32(32)

    hash := argon2.IDKey([]byte(password), salt, time, memory, threads, keyLen)

    encodedHash := base64.StdEncoding.EncodeToString(hash)
    encodedSalt := base64.StdEncoding.EncodeToString(salt)

    return fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
        argon2.Version, memory, time, threads, encodedSalt, encodedHash), nil
}

// CheckPasswordHash compares a raw password with its Argon2 hash.
func CheckPasswordHash(password, hash string) bool {
    parts := strings.Split(hash, "$")
    if len(parts) != 6 {
        return false
    }

    salt, err := base64.StdEncoding.DecodeString(parts[4])
    if err != nil {
        return false
    }

    hashedPassword, err := HashPasswordWithSalt(password, salt)
    if err != nil {
        return false
    }

    return hashedPassword == hash
}