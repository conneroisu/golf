package domain

// WorkspaceFolder is a workspace folder.
type WorkspaceFolder struct {
	// The associated URI for this workspace folder.
	URI TextDocumentURI `json:"uri"`

	// The name of the workspace folder. Used to refer to this
	// workspace folder in the user interface.
	Name string `json:"name"`
}
