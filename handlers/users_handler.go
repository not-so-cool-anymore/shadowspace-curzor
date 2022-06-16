package handlers

import (
	"encoding/base64"
	"net/http"
	"encoding/json"
	"fmt"
)

func ListUsers(w http.ResponseWriter, r *http.Request) {
	auth := r.Header.Get("Authorization")
    decodedAuthToken, _ := base64.StdEncoding.DecodeString(auth)
	
	var authToken map[string]interface{}
	json.Unmarshal([]byte(decodedAuthToken), &authToken)

	var userID int = int(authToken["user_id"].(float64))

	if userID == 0 && authToken["user_type"] == "admin"{
		fmt.Fprint(w, "\"users\": [\"admin\", \"ivan\", \"alexander\", \"master\", \"DavidLightman\", \"MajorMalfunction\"]")
	}else{

		fmt.Fprint(w, "Unauthorized User Type and ID.")
	}
	return

}