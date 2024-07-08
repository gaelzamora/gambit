package auth
     
import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)
 
type TokenJSON struct {
	Sub      string
	EventID  string `json:"event_id"`
	TokenUse string `json:"token_use"`
	Scope    string
	AuthTime int `json:"auth_time"`
	Iss      string
	Exp      int
	Iat      int
	ClientID string `json:"client_id"`
	Username string
}
 
func ValidoToken(token string) (bool, error, string) {
	parts := strings.Split(token, ".")
	if len(parts) < 2 {
		return false, fmt.Errorf("Invalid token format"), ""
	}
 
	userInfo, err := base64.RawStdEncoding.DecodeString(parts[1])
	if err != nil {
		return false, fmt.Errorf("Failed to decode token: %s", err.Error()), ""
	}
 
	var tkj TokenJSON
	err = json.Unmarshal(userInfo, &tkj)
	if err != nil {
		return false, fmt.Errorf("Failed to unmarshal JSON: %s", err.Error()), ""
	}
 
	ahora := time.Now()
	tm := time.Unix(int64(tkj.Exp), 0)
 
	if tm.Before(ahora) {
		return false, fmt.Errorf("Token has expired"), ""
	}
 
	return true, nil, tkj.Username
} 