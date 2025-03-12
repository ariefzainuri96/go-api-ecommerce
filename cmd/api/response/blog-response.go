package response

import "encoding/json"

type BlogsResponse struct {
	BaseResponse
	Blogs []Blog `json:"blogs"`
}

type BlogResponse struct {
	BaseResponse
	Blog Blog `json:"blog"`
}

func (r BlogsResponse) MarshalBlogsResponse() ([]byte, error) {
	marshal, err := json.Marshal(r)

	if err != nil {
		return nil, err
	}

	return marshal, nil
}

func (r *BlogsResponse) UnmarshalBlogsResponse(data []byte) error {
	return json.Unmarshal(data, &r)
}

func (r BlogResponse) MarshalBlogResponse() ([]byte, error) {
	marshal, err := json.Marshal(r)

	if err != nil {
		return nil, err
	}

	return marshal, nil
}

func (r *BlogResponse) UnmarshalBlogResponse(data []byte) error {
	return json.Unmarshal(data, &r)
}

type Blog struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
}
