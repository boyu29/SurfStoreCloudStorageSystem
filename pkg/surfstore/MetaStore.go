package surfstore

import (
	context "context"
	sync "sync"

	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

var locker sync.Mutex

type MetaStore struct {
	FileMetaMap    map[string]*FileMetaData
	BlockStoreAddr string
	UnimplementedMetaStoreServer
}

func (m *MetaStore) GetFileInfoMap(ctx context.Context, _ *emptypb.Empty) (*FileInfoMap, error) {
	// panic("todo")

	// lock := sync.Mutex{}
	locker.Lock()
	defer locker.Unlock()

	fileinfomap := &FileInfoMap{}
	fileinfomap.FileInfoMap = m.FileMetaMap
	return fileinfomap, nil
}

func (m *MetaStore) UpdateFile(ctx context.Context, fileMetaData *FileMetaData) (*Version, error) {
	// panic("todo")

	// lock := sync.Mutex{}
	locker.Lock()
	defer locker.Unlock()

	filename := fileMetaData.Filename
	newVersion := &Version{}
	if _, check := m.FileMetaMap[filename]; check {
		if fileMetaData.Version == m.FileMetaMap[filename].Version+1 {
			m.FileMetaMap[filename] = fileMetaData
			newVersion.Version = fileMetaData.Version
			return newVersion, nil
		} else {
			newVersion.Version = -1
			return newVersion, nil
		}
	} else {
		m.FileMetaMap[filename] = fileMetaData
		newVersion.Version = fileMetaData.Version
		return newVersion, nil
	}
}

func (m *MetaStore) GetBlockStoreAddr(ctx context.Context, _ *emptypb.Empty) (*BlockStoreAddr, error) {
	// panic("todo")

	// lock := sync.Mutex{}
	locker.Lock()
	defer locker.Unlock()

	return &BlockStoreAddr{Addr: m.BlockStoreAddr}, nil
}

// This line guarantees all method for MetaStore are implemented
var _ MetaStoreInterface = new(MetaStore)

func NewMetaStore(blockStoreAddr string) *MetaStore {
	return &MetaStore{
		FileMetaMap:    map[string]*FileMetaData{},
		BlockStoreAddr: blockStoreAddr,
	}
}
