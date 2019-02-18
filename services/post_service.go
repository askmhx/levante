package services

import (
	"fmt"
	"iosxc.com/levante/model"
	"iosxc.com/levante/orm"
	"iosxc.com/levante/repositories"
	"reflect"
)

const (
	site_map_cache_key = "post_service_get_links_site_map_key"
	site_map_cache_time = 60*60*24
)

type PostService interface {
	GetByID(id uint64) (orm.Post, bool)
	GetByCatalog(catalog string) ([]orm.Post , bool)
	GetMonthList() []model.ColCount
	GetTagList() []model.ColCount
	GetList(start,limit uint64) []orm.Post
	GetLinks(host string) []model.SiteUrl

}

func NewPostService(repo repositories.PostRepository,cache repositories.MemCacheRepository) PostService {
	return &postService{
		repo: repo,
		cache: cache,
	}
}

type postService struct {
	repo repositories.PostRepository
	cache repositories.MemCacheRepository
}

func (s *postService) GetLinks(host string) []model.SiteUrl {
	var cachedSiteURLs = s.cache.Get(site_map_cache_key)
	if cachedSiteURLs != nil {
		return reflect.ValueOf(cachedSiteURLs).Interface().([]model.SiteUrl)
	}
	var sql = "select id,created_at,updated_at from levante.posts order by updated_at desc"
	var posts = s.repo.QueryList(sql)
	var links = []model.SiteUrl{}
	timeFormat := "2006-01-02T15:04:05+07:00"
	for _,post := range posts  {
		links = append(links, model.SiteUrl{LastMod:post.UpdatedAt.Format(timeFormat),Loc:fmt.Sprintf("%s/post/%d",host,post.ID),Priority:8})
	}
	timeMax := posts[0].UpdatedAt.Format(timeFormat)
	var homeURL = model.SiteUrl{Loc:host,LastMod:timeMax,Priority:10}
	var aboutURL = model.SiteUrl{Loc:host+"/about",LastMod:timeMax,Priority:9}
	var startURL = model.SiteUrl{Loc:host+"/start",LastMod:timeMax,Priority:9}
	links = append(links,homeURL)
	links = append(links,aboutURL)
	links = append(links,startURL)
	_ = s.cache.Set(site_map_cache_key,links,site_map_cache_time)
	return links
}

func (s *postService) GetList(start,limit uint64) []orm.Post {
	return s.repo.SelectMany(start,limit)
}

func (s *postService) GetMonthList() []model.ColCount {
	return s.repo.QueryColCount("select date_format(updated_at, '%Y-%m') as col_name,count(1) as count from levante.posts group by col_name order by col_name desc")
}

func (s *postService) GetTagList() []model.ColCount {
	return s.repo.QueryColCount("select catalog as col_name,count(1) as count from levante.posts group by col_name order by col_name desc")
}

func (s *postService) GetByID(id uint64) (orm.Post, bool) {
	return s.repo.SelectOne(id)
}

// GetByUsernameAndPassword returns a user based on its username and passowrd,
// used for authentication.
func (s *postService) GetByCatalog(catalog string) ([]orm.Post, bool) {
	if catalog == "" {
		return nil, false
	}

	return nil, false

	//return s.repo.Select(func(m orm.Post) bool {
	//	if m.Username == username {
	//		hashed := m.HashedPassword
	//		if ok, _ := datamodels.ValidatePassword(userPassword, hashed); ok {
	//			return true
	//		}
	//	}
	//	return false
	//})
}