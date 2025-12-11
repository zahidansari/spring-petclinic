package main

import (
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"log"
	"sync"
	"time"
)

// LegacyUserManager provides user management utilities converted from Java to Go
// This is a Go conversion of the original LegacyUserManager.java class
type LegacyUserManager struct {
	mu sync.RWMutex