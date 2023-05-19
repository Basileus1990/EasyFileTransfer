// model for creating a structure for sent directiories and files
// the files and directories have to have names
// files have to have non empty type and size

package dataStructs

import (
	"errors"
)

type File struct {
	Name string `json:"name"`
	Size int    `json:"size"`
}

func (f File) IsValidFile() error {
	if f.Name == "" {
		return errors.New("an invalid file was given")
	}
	return nil
}

type Directory struct {
	Name    string      `json:"name"`
	SubDirs []Directory `json:"subdirs"`
	Files   []File      `json:"files"`
}

// loops all over the directory and its sub directories and its files and checks if they are valid
func (d Directory) IsValidDirectory() error {
	if d.Name == "" {
		return errors.New("an empty directory was given")
	}
	for _, dir := range d.SubDirs {
		err := dir.IsValidDirectory()
		if err != nil {
			return err
		}
	}
	for _, f := range d.Files {
		err := f.IsValidFile()
		if err != nil {
			return err
		}
	}
	return nil
}
