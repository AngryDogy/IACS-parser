package entities

type FileJSON struct {
	RenderedName *Rendered `json:"title"`
	ACF          *ACF      `json:"acf"`
}
type Rendered struct {
	Name string `json:"rendered"`
}
type ACF struct {
	Description       string `json:"publication_description"`
	Notes             string `json:"publication_notes"`
	Link              string `json:"file"`
	FutureName        string `json:"cl_future_publication_name"`
	FutureDescription string `json:"cl_future_publication_description"`
	FutureNotes       string `json:"cl_future_publication_notes"`
	FutureLink        string `json:"cl_future_publication_file"`
}

type FileDB struct {
	Name              string `gorm:"primaryKey"`
	Description       string
	Notes             string
	Link              string
	FutureName        string
	FutureDescription string
	FutureNotes       string
	FutureLink        string
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

func NewFileDB(fileJSON *FileJSON) *FileDB {
	fileDB := new(FileDB)
	fileDB.Name = fileJSON.RenderedName.Name
	fileDB.Description = fileJSON.ACF.Description
	fileDB.Notes = fileJSON.ACF.Notes
	fileDB.Link = fileJSON.ACF.Link
	fileDB.FutureDescription = fileJSON.ACF.FutureDescription
	fileDB.FutureNotes = fileJSON.ACF.FutureNotes
	fileDB.FutureLink = fileJSON.ACF.FutureLink
	return fileDB
}

func NewFileJSON(name, link string) *FileJSON {
	file := new(FileJSON)
	file.RenderedName = new(Rendered)
	file.RenderedName.Name = name
	file.ACF = new(ACF)
	file.ACF.Link = link
	return file

}
