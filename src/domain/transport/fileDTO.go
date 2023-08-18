package transport

type FileDTO struct {
	Name         string `json:"name"`
	Extension    string `json:"extension"`
	Size         int    `json:"size"`
	CreationDate string `json:"creation_date"` // for now
	EditionDate  string `json:"edition_date"`
}
