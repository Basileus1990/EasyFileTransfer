package transport

type FileDTO struct {
	Name      string `json:"name"`
	Extension string `json:"extension"`
	Size      int    `json:"size"`
}
