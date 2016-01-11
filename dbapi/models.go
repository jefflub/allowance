package dbapi

type (
	// Family represents the structure of our resource
	Family struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	// Parent is the parent of some kids
	Parent struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	// Kid is a kid with an allowance
	Kid struct {
		ID      int      `json:"id"`
		Name    string   `json:"name"`
		Buckets []Bucket `json:"buckets"`
	}

	// Bucket is a category of allowance savings
	Bucket struct {
		ID                int     `json:"id"`
		Name              string  `json:"name"`
		DefaultAllocation int     `json:"defaultAllocation"`
		Total             float64 `json:"total"`
	}
)
