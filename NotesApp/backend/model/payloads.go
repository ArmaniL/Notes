package model

type UpdateNotePayload struct {
	NoteID string `json:"noteID"`
	// Either the header or content
	Header string `json:"header"`
	// Is the actual data for the parameter
	Content string `json:"content"`
}

type ShareNotePayload struct {
	NoteID string `json:"noteID"`
	Email  string `json:"email"`
}
