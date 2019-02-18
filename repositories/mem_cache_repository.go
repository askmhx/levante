package repositories

type MemCacheRepository struct {
	maxEntries int
	entriesMap map[string]interface{}
}

func NewMemCacheRepository(maxEntries int) MemCacheRepository {
	var mem = MemCacheRepository{maxEntries: maxEntries}
	mem.entriesMap = make(map[string]interface{},maxEntries)
	return mem
}

func (this *MemCacheRepository) Get(key string) (value interface{}) {
	return this.entriesMap[key]
}

func (this *MemCacheRepository) Set(key string, value interface{}, expire int64) (err error) {
	curtSize := len(this.entriesMap)
	if (curtSize < this.maxEntries) {
		this.entriesMap[key] = value
	}else{

	}
	return nil
}

func (this *MemCacheRepository) Del(key string) (affected bool) {
	if this.entriesMap[key] != nil {
		delete(this.entriesMap, key)
		return true
	}
	return false
}
