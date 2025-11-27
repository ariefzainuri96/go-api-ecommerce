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

func (r *UpdateCartRequest) Unmarshal(data []byte) error {
	return json.Unmarshal(data, &r)
}

func (r UpdateCartRequest) ToMap() (map[string]any, error) {
	data, err := r.Marshal()

	if err != nil {
		return nil, err
	}

	var dataMap map[string]any
	
	if err = json.Unmarshal(data, &dataMap); err != nil {
		return nil, err
	}

	return dataMap, nil	
}