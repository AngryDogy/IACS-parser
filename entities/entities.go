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

type ChangedFile struct {
	Name              string `json:"name"`
	Description       string `json:"description"`
	Notes             string `json:"notes"`
	Link              string `json:"link"`
	FutureName        string `json:"futureName"`
	FutureDescription string `json:"futureDescription"`
	FutureNotes       string `json:"futureNotes"`
	FutureLink        string `json:"futureLink"`
	Changes           string `json:"changes"`
}

type Email struct {
	Email string `gorm:"primaryKey"`
}

func NewChangedFile(fileDB *FileDB, changes string) *ChangedFile {
	changedFile := new(ChangedFile)
	changedFile.Name = fileDB.Name
	changedFile.Description = fileDB.Description
	changedFile.Notes = fileDB.Notes
	changedFile.Link = fileDB.Link
	changedFile.FutureName = fileDB.FutureName
	changedFile.FutureDescription = fileDB.FutureDescription
	changedFile.FutureNotes = fileDB.FutureNotes
	changedFile.FutureLink = fileDB.FutureLink
	changedFile.Changes = changes
	return changedFile

}

func NewFileJSON(name, link string) *FileJSON {
	file := new(FileJSON)
	file.RenderedName = new(Rendered)
	file.RenderedName.Name = name
	file.ACF = new(ACF)
	file.ACF.Link = link
	return file
}

func NewFileDB(fileJSON *FileJSON) *FileDB {
	fileDB := new(FileDB)
	fileDB.Name = fileJSON.RenderedName.Name
	fileDB.Description = fileJSON.ACF.Description
	fileDB.Notes = fileJSON.ACF.Notes
	fileDB.Link = fileJSON.ACF.Link
	fileDB.FutureName = fileJSON.ACF.FutureName
	fileDB.FutureDescription = fileJSON.ACF.FutureDescription
	fileDB.FutureNotes = fileJSON.ACF.FutureNotes
	fileDB.FutureLink = fileJSON.ACF.FutureLink
	return fileDB
}
