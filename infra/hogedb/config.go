package hogedb

// Env Env
type Env map[string]string

// AConfig AConfig
type AConfig struct {
	ID  string
	Key string
}

// BConfig BConfig
type BConfig struct {
	ID  string
	Key string
}

// NewConfigA NewConfigA
func NewConfigA(e Env) *AConfig {
	id, _ := e["A_ID"]
	key, _ := e["A_KEY"]
	return &AConfig{
		ID:  id,
		Key: key,
	}
}

// NewConfigB NewConfigB
func NewConfigB(e Env) *BConfig {
	id, _ := e["B_ID"]
	key, _ := e["B_KEY"]
	return &BConfig{
		ID:  id,
		Key: key,
	}
}
