package mem

import (
	"github.com/viant/afs/base"
	"github.com/viant/afs/file"
	"github.com/viant/afs/storage"
	"github.com/viant/afs/url"
)

type storager struct {
	base.Storager
	scheme string
	Root   *Folder
}

func (s *storager) Close() error {
	return nil
}

//NewStorager create a new in memeory storage service
func NewStorager(baseURL string) storage.Storager {
	result := &storager{
		Root:   NewFolder(baseURL, file.DefaultDirOsMode),
		scheme: url.Scheme(baseURL, Scheme),
	}
	result.Storager.List = result.List
	return result
}
