package entities

type File struct {
	RenderedName *Rendered `json:"title"`
	ReleaseDate  string    `json:"date"`
	ModifiedDate string    `json:"modified"`
	ACFLink      *ACF      `json:"acf"`
}
type Rendered struct {
	Name string `json:"rendered"`
}
type ACF struct {
	Link string `json:"file"`
}

func NewFile(name, link string) *File {
	file := new(File)
	file.RenderedName = new(Rendered)
	file.RenderedName.Name = name
	file.ACFLink = new(ACF)
	file.ACFLink.Link = link
	return file

}
