package server

type Page struct {
	Title string
	Path  string
	Data  interface{}
}

func (m *Page) New(title string, path string, data interface{}) Page {
	if data == nil {
		data = func() {}
	}
	return Page{
		Title: title,
		Path:  path,
		Data:  data,
	}
}
