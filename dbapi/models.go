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
)
