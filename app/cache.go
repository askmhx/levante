package app

type Cache struct {
	maxEntries int
	entriesMap map[string]interface{}
}

func NewCache(maxEntries int) *Cache {
	return &Cache{maxEntries: maxEntries}
}

func (this *Cache) Get(key string) (value interface{}) {
	return this.entriesMap[key]
}

func (this *Cache) Set(key string, value interface{}, expire int64) (err error) {

	curtSize := len(this.entriesMap)
	if (curtSize < this.maxEntries) {
		this.entriesMap[key] = value
	}else{

	}
	return nil
}

func (this *Cache) Del(key string) (affected bool) {
	if this.entriesMap[key] != nil {
		delete(this.entriesMap, key)
		return true
	}
	return false
}
