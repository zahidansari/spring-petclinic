package main

import (
	"context"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"log"
	"time"
)

// UserManager provides user management functionality
// Migrated from Java LegacyUserManager.java to Go with modern security practices
type UserManager struct {
	// Add any necessary fields here
}

// NewUserManager creates a new instance of UserManager
func NewUserManager() *UserManager {
	return &UserManager{}
}

// HashPassword securely hashes a password using SHA-256
// Replaces the insecure MD5 hashing from the original Java implementation
func (um *UserManager) HashPassword(password string) (string, error) {
	if password == "" {
		return "", fmt.Errorf("password cannot be empty")
	}

	// Use SHA-256 instead of the deprecated MD5
	hasher := sha256.New()
	hasher.Write([]byte(password))
	digest := hasher.Sum(nil)

	// Use standard base64 encoding instead of sun.misc.BASE64Encoder
	encodedHash := base64.StdEncoding.EncodeToString(digest)
	return encodedHash, nil
}

// GetExpirationDate calculates expiration date (30 days from now)
// Replaces the legacy Calendar/Date API with Go's time package
func (um *UserManager) GetExpirationDate() string {
	// Use Go's time package instead of Calendar/Date
	expirationDate := time.Now().AddDate(0, 0, 30) // Add 30 days
	return expirationDate.Format("2006-01-02") // ISO 8601 format
}

// GetExpirationDateTime returns expiration date with time
func (um *UserManager) GetExpirationDateTime() time.Time {
	return time.Now().AddDate(0, 0, 30)
}

// StopBackgroundWorker gracefully stops a background worker using context cancellation
// Replaces the dangerous Thread.stop() method with proper Go concurrency patterns
func (um *UserManager) StopBackgroundWorker(ctx context.Context, cancel context.CancelFunc) {
	// Cancel the context to signal goroutines to stop
	cancel()
	
	// In Go, we use context cancellation instead of thread.stop()
	// Goroutines should check ctx.Done() and exit gracefully
	log.Println("Background worker stop signal sent")
}

// BackgroundWorkerExample demonstrates how to create a cancellable background worker
func (um *UserManager) BackgroundWorkerExample(ctx context.Context) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Println("Background worker stopped gracefully")
				return
			default:
				// Do background work here
				time.Sleep(1 * time.Second)
				log.Println("Background worker running...")
			}
		}
	}()
}

// PrintCertificateInfo prints X.509 certificate information
// Replaces the deprecated javax.security.cert.X509Certificate with Go's x509 package
func (um *UserManager) PrintCertificateInfo(cert *x509.Certificate) error {
	if cert == nil {
		return fmt.Errorf("certificate cannot be nil")
	}

	// Print certificate subject and issuer using Go's x509 package
	fmt.Printf("Certificate Subject: %s\n", cert.Subject.String())
	fmt.Printf("Certificate Issuer: %s\n", cert.Issuer.String())
	fmt.Printf("Certificate Serial Number: %s\n", cert.SerialNumber.String())
	fmt.Printf("Certificate Valid From: %s\n", cert.NotBefore.Format(time.RFC3339))
	fmt.Printf("Certificate Valid Until: %s\n", cert.NotAfter.Format(time.RFC3339))

	return nil
}

// ValidatePassword performs basic password validation
func (um *UserManager) ValidatePassword(password string) error {
	if len(password) < 8 {
		return fmt.Errorf("password must be at least 8 characters long")
	}
	return nil
}

// ComparePasswords securely compares a plain password with a hashed password
func (um *UserManager) ComparePasswords(plainPassword, hashedPassword string) (bool, error) {
	hashedInput, err := um.HashPassword(plainPassword)
	if err != nil {
		return false, err
	}
	return hashedInput == hashedPassword, nil
}

// Example usage and main function for demonstration
func main() {
	um := NewUserManager()

	// Example: Password hashing
	password := "mySecurePassword123"
	hashedPassword, err := um.HashPassword(password)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
	} else {
		fmt.Printf("Hashed Password: %s\n", hashedPassword)
	}

	// Example: Get expiration date
	expirationDate := um.GetExpirationDate()
	fmt.Printf("Expiration Date: %s\n", expirationDate)

	// Example: Background worker with context cancellation
	ctx, cancel := context.WithCancel(context.Background())
	um.BackgroundWorkerExample(ctx)

	// Stop the background worker after 5 seconds
	time.Sleep(5 * time.Second)
	um.StopBackgroundWorker(ctx, cancel)

	// Wait a bit to see the graceful shutdown
	time.Sleep(2 * time.Second)

	fmt.Println("UserManager Go migration completed successfully!")
}

/*
Migration Notes:
================

1. Security Improvements:
   - Replaced MD5 with SHA-256 for password hashing
   - Replaced sun.misc.BASE64Encoder with standard base64 encoding
   - Added password validation

2. Modern API Usage:
   - Replaced Calendar/Date with Go's time package
   - Replaced Thread.stop() with context cancellation
   - Replaced javax.security.cert.X509Certificate with Go's x509 package

3. Go Best Practices:
   - Used proper error handling with error returns
   - Implemented graceful goroutine shutdown using context
   - Added proper logging instead of printStackTrace
   - Used Go naming conventions (PascalCase for exported functions)

4. Additional Features:
   - Added password validation
   - Added password comparison functionality
   - Provided example usage in main function
   - Added comprehensive documentation

This Go implementation provides the same functionality as the original Java
LegacyUserManager.java but with modern, secure, and Go-idiomatic approaches.
*/