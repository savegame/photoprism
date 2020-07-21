package entity

import (
	"fmt"

	"github.com/photoprism/photoprism/pkg/txt"
)

type Duplicates []Duplicate
type DuplicatesMap map[string]Duplicate

// Duplicate represents an exact file duplicate.
type Duplicate struct {
	FileName string `gorm:"type:varbinary(768);primary_key;" json:"Name" yaml:"Name"`
	FileRoot string `gorm:"type:varbinary(16);primary_key;default:'/';" json:"Root" yaml:"Root,omitempty"`
	FileHash string `gorm:"type:varbinary(128);default:'';index" json:"Hash" yaml:"Hash,omitempty"`
	FileSize int64  `json:"Size" yaml:"Size,omitempty"`
	ModTime  int64  `json:"ModTime" yaml:"-"`
}

func AddDuplicate(fileName, fileRoot, fileHash string, fileSize, modTime int64) error {
	if fileName == "" {
		return fmt.Errorf("duplicate: file name must not be empty (add)")
	} else if fileHash == "" {
		return fmt.Errorf("duplicate: file hash must not be empty (add)")
	} else if modTime == 0 {
		return fmt.Errorf("duplicate: mod time must not be empty (add)")
	} else if fileRoot == "" {
		return fmt.Errorf("duplicate: file root must not be empty (add)")
	}

	duplicate := &Duplicate{
		FileName: fileName,
		FileRoot: fileRoot,
		FileHash: fileHash,
		FileSize: fileSize,
		ModTime:  modTime,
	}

	if err := duplicate.Create(); err == nil {
		return nil
	} else if err := duplicate.Save(); err != nil {
		return err
	}

	return nil
}

// Find returns a photo from the database.
func (m *Duplicate) Find() error {
	return UnscopedDb().First(m, "file_name = ?", m.FileName).Error
}

// Create inserts a new row to the database.
func (m *Duplicate) Create() error {
	if m.FileName == "" {
		return fmt.Errorf("duplicate: file name must not be empty (create)")
	} else if m.FileHash == "" {
		return fmt.Errorf("duplicate: file hash must not be empty (create)")
	} else if m.ModTime == 0 {
		return fmt.Errorf("duplicate: mod time must not be empty (create)")
	} else if m.FileRoot == "" {
		return fmt.Errorf("duplicate: file root must not be empty (create)")
	}

	return UnscopedDb().Create(m).Error
}

// Saves the duplicates in the database.
func (m *Duplicate) Save() error {
	if m.FileName == "" {
		return fmt.Errorf("duplicate: file name must not be empty (save)")
	} else if m.FileHash == "" {
		return fmt.Errorf("duplicate: file hash must not be empty (save)")
	} else if m.ModTime == 0 {
		return fmt.Errorf("duplicate: mod time must not be empty (save)")
	} else if m.FileRoot == "" {
		return fmt.Errorf("duplicate: file root must not be empty (save)")
	}

	if err := UnscopedDb().Save(m).Error; err != nil {
		log.Errorf("duplicate: %s (save %s)", err, txt.Quote(m.FileName))
		return err
	}

	return nil
}