package entities

type File struct {
	RenderedName *Rendered `json:"title"`
	ReleaseDate  string    `json:"date"`
	ModifiedDate string    `json:"modified"`
	ACF          *ACF      `json:"acf"`
}
type Rendered struct {
	Name string `json:"rendered"`
}
type ACF struct {
	Link        string `json:"file"`
	Description string `json:"publication_description"`
	Notes       string `json:"publication_notes"`
}

type FileDB struct {
	Name         string `gorm:"primaryKey"`
	ReleaseDate  string
	ModifiedDate string
	Description  string
	Notes        string
	Link         string
}

type MessageChange struct {
	Name        string
	Description string
	Notes       string
	Link        string
	Changes     string
}

func NewMessageChange(name, description, notes, link, changes string) *MessageChange {
	message := new(MessageChange)
	message.Name = name
	message.Description = description
	message.Notes = notes
	message.Link = link
	message.Changes = changes
	return message
}

func NewFileDB(name, releaseDate, modifiedDate, description, notes, link string) *FileDB {
	fileDB := new(FileDB)
	fileDB.Name = name
	fileDB.ReleaseDate = releaseDate
	fileDB.ModifiedDate = modifiedDate
	fileDB.Description = description
	fileDB.Notes = notes
	fileDB.Link = link
	return fileDB
}

func NewFile(name, link string) *File {
	file := new(File)
	file.RenderedName = new(Rendered)
	file.RenderedName.Name = name
	file.ACF = new(ACF)
	file.ACF.Link = link
	return file

}
