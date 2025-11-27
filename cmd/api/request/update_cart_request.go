package request

import "encoding/json"

type UpdateCartRequest struct {
	Quantity int `json:"quantity" validate:"required"`
}

func (r UpdateCartRequest) Marshal() ([]byte, error) {
	marshal, err := json.Marshal(r)

	if err != nil {
		return nil, err
	}

	return marshal, nil
}

// func (r *UpdateCartRequest) Unmarshal(data []byte, dataMap *map[string]any) error {
// 	return json.Unmarshal(data, &dataMap)
// }

func (r UpdateCartRequest) Unmarshal(dataMap *map[string]any) error {
	data, err := r.Marshal()

	if err != nil {
		return err
	}

	return json.Unmarshal(data, &dataMap)
}