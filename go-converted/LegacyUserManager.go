package main

import (
    "crypto/md5"
    "encoding/base64"
    "fmt"
    "time"
    "crypto/x509"
    "encoding/pem"
)

type LegacyUserManager struct{}

func (l *LegacyUserManager) HashPassword(password string) string {
    // INSECURE: MD5 should not be used in production!
    hash := md5.Sum([]byte(password))
    return base64.StdEncoding.EncodeToString(hash[:])
}

func (l *LegacyUserManager) GetExpirationDate() string {
    expiration := time.Now().AddDate(0, 0, 30)
    return expiration.String()
}

// WARNING: Go does not support forcibly stopping goroutines. This is unsafe and discouraged.
func (l *LegacyUserManager) StopBackgroundThread() {
    fmt.Println("Not supported: forcibly stopping goroutines is unsafe and not directly possible in Go.")
}

// For demo: print subject and issuer from a PEM encoded certificate
func (l *LegacyUserManager) PrintCertificateInfo(certPEM string) {
    block, _ := pem.Decode([]byte(certPEM))
    if block == nil {
        fmt.Println("Failed to parse certificate PEM.")
        return
    }
    cert, err := x509.ParseCertificate(block.Bytes)
    if err != nil {
        fmt.Println("Failed to parse certificate:", err)
        return
    }
    fmt.Println("Certificate Subject:", cert.Subject)
    fmt.Println("Certificate Issuer:", cert.Issuer)
}

func main() {
    manager := LegacyUserManager{}
    fmt.Println(manager.HashPassword("password123"))
    fmt.Println(manager.GetExpirationDate())
    manager.StopBackgroundThread()
    // manager.PrintCertificateInfo(certPEM) // Provide PEM string as needed
}