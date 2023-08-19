package repository

import (
	"github.com/Basileus1990/EasyFileTransfer.git/src/domain/transport"
)

type DatabaseInterface interface {

	// Directory methods
	AddDirectory(dir transport.DirectoryDTO) error
	GetDirectory(dirID int) (transport.DirectoryDTO, error)
	GetDirectories() ([]transport.DirectoryDTO, error)

	// User methods
	AddUser(user transport.UserDTO) error
	GetUser(email string) (transport.UserDTO, error)
}
