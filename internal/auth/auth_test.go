package auth

import (
	"net/http"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestCheckpasswordHash(t *testing.T) {
	password1 := "correctPassword123!"
	password2 := "correctPassword452!"
	hash1, _ := HashPassword(password1)
	hash2, _ := HashPassword(password2)

	cases := []struct {
		name     string
		password string
		hash     string
		wantErr  bool
	}{
		{
			name:     "Correct password",
			password: password1,
			hash:     hash1,
			wantErr:  false,
		},
		{
			name:     "Incorrect password",
			password: password2,
			hash:     hash2,
			wantErr:  false,
		},
		{
			name:     "Password doesn't match different hash",
			password: password1,
			hash:     hash2,
			wantErr:  true,
		},
		{
			name:     "Empty password",
			password: "",
			hash:     hash1,
			wantErr:  true,
		},
		{
			name:     "Invalid hash",
			password: password1,
			hash:     "invalidhash",
			wantErr:  true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := CheckPasswordHash(c.password, c.hash)
			if (err != nil) != c.wantErr {
				t.Errorf("CheckPasswordHash() error = %v, wantErr %v", err, c.wantErr)
			}
		})
	}
}

func TestValidateJWT(t *testing.T) {
	userID := uuid.New()
	validToken, _ := MakeJWT(userID, "secret", time.Hour)

	tests := []struct {
		name        string
		tokenString string
		tokenSecret string
		wantUserID  uuid.UUID
		wantErr     bool
	}{
		{
			name:        "Valid token",
			tokenString: validToken,
			tokenSecret: "secret",
			wantUserID:  userID,
			wantErr:     false,
		},
		{
			name:        "Invalid token",
			tokenString: "invalid.token.string",
			tokenSecret: "secret",
			wantUserID:  uuid.Nil,
			wantErr:     true,
		},
		{
			name:        "Wrong secret",
			tokenString: validToken,
			tokenSecret: "wrong_secret",
			wantUserID:  uuid.Nil,
			wantErr:     true,
		},
	}

	for _, c := range tests {
		t.Run(c.name, func(t *testing.T) {
			gotUserId, err := ValidateJWT(c.tokenString, c.tokenSecret)
			if (err != nil) != c.wantErr {
				t.Errorf("ValidateJWT() error = %v, wantErr %v", err, c.wantErr)
				return
			}
			if gotUserId != c.wantUserID {
				t.Errorf("ValidateJWT() gotUserId = %v, want %v", gotUserId, c.wantUserID)
				return
			}
		})
	}
}

func TestGetBearerToken(t *testing.T) {
	cases := []struct {
		name      string
		headers   http.Header
		wantToken string
		wantErr   bool
	}{
		{
			name: "Valid bearer token",
			headers: http.Header{
				"Authorization": []string{"Bearer valid_token"},
			},
			wantToken: "valid_token",
			wantErr:   false,
		},
		{
			name:      "Missing Authorization header",
			headers:   http.Header{},
			wantToken: "",
			wantErr:   true,
		},
		{
			name: "Malformed Authorization header",
			headers: http.Header{
				"Authorization": []string{"InvalidBearer token"},
			},
			wantToken: "",
			wantErr:   true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gotToken, err := GetBearerToken(c.headers)

			if (err != nil) != c.wantErr {
				t.Errorf("GetBearerToken() error = %v, wantErr %v", err, c.wantErr)
				return
			}

			if gotToken != c.wantToken {
				t.Errorf("GetBearerToken() gotToken = %v, want %v", gotToken, c.wantToken)
				return
			}

		})
	}
}
