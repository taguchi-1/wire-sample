package hogedb

// HogeDB HogeDB
type HogeDB struct {
	AConfig *AConfig
	BConfig *BConfig
}

// Get Get
func (db *HogeDB) Get() (map[string]interface{}, error) {
	return map[string]interface{}{
		"a": 1,
		"b": 2,
		"c": 3,
	}, nil
}
