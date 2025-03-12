package request

import (
	"encoding/json"
)

type LoginRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r LoginRequest) Marshal() ([]byte, error) {
	marshal, err := json.Marshal(r)

	if err != nil {
		return nil, err
	}

	return marshal, nil
}

func (r *LoginRequest) Unmarshal(data []byte) error {
	return json.Unmarshal(data, &r)
}
