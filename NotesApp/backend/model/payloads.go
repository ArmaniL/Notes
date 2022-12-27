package model

type UpdateNotePayload struct {
	NoteID string `json:"noteID"`
	// Either the header or content
	Parameter string `json:"parameter"`
	// Is the actual data for the parameter
	Data string `json:"data"`
}

type ShareNotePayload struct {
	NoteID string `json:"noteID"`
	Email  string `json:"email"`
}
