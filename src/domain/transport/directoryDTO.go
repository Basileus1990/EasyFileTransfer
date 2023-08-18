package transport

type DirectoryDTO struct {
	Name        string         `json:"name"`
	Directories []DirectoryDTO `json:"directories"`
	Files       []FileDTO      `json:"files"`
}
