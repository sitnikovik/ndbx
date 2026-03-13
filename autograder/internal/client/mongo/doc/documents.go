package doc

// Documents represents a list of Document,
// providing convenience methods for working with multiple documents.
type Documents []Document

// NewDocuments creates a new list of Documents from the provided ones.
func NewDocuments(docs ...Document) Documents {
	return docs
}

// IDs returns a list of string representations of the ObjectIDs of the documents.
func (dd Documents) IDs() []string {
	if dd.Empty() {
		return nil
	}
	ids := make([]string, len(dd))
	for i, doc := range dd {
		ids[i] = doc.ID()
	}
	return ids
}

// First returns the first document in the list of documents.
//
// If the list is empty, it returns an empty Document.
func (dd Documents) First() Document {
	if dd.Empty() {
		return Document{}
	}
	return dd[0]
}

// Last returns the last document in the list of documents.
//
// If the list is empty, it returns an empty Document.
func (dd Documents) Last() Document {
	if dd.Empty() {
		return Document{}
	}
	return dd[len(dd)-1]
}

// Empty checks if the list of documents is empty.
func (dd Documents) Empty() bool {
	return dd.Len() == 0
}

// Len returns the number of documents in the list.
func (dd Documents) Len() int {
	return len(dd)
}
