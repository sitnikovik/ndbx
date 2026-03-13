package doc

// Document represents a MongoDB document.
type Document struct {
	// kvs is a slice of key-value pairs representing the fields of the document.
	kvs []KV
	// id is the string representation of the document's ObjectID.
	id string
}

// NewDocument creates a new Document with the provided ID and key-value pairs.
func NewDocument(id string, kvs ...KV) Document {
	return Document{
		id:  id,
		kvs: kvs,
	}
}

// ID returns the string representation of the document's ObjectID.
func (d Document) ID() string {
	return d.id
}

// KVs returns the slice of key-value pairs representing the fields of the document.
func (d Document) KVs() KVs {
	return d.kvs
}
