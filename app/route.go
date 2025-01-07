package app

import (
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"levante/controllers"
	"levante/repositories"
	"levante/services"
)

const url_index = "/"
const url_favicon = "/favicon.ico"
const url_sitemap = "/sitemap.xml"
const url_start = "/start"
const url_about = "/about"
const url_photo = "/photo"
const url_post = "/post"
const url_comment = "/comment"

func InitRoute(application *iris.Application, config *AppConfig, database *gorm.DB) {
	application.Favicon(config.Home+"assets/statics/img/favicon.ico", url_favicon)

	indexCtrl := new(controllers.IndexController)
	postCtrl := new(controllers.PostController)
	commentCtrl := new(controllers.CommentController)
	photoCtrl := new(controllers.PhotoController)
	aboutCtrl := new(controllers.AboutController)
	startCtrl := new(controllers.StartController)
	sitemapCtrl := new(controllers.SiteMapController)

	cacheRepository := repositories.NewMemCacheRepository(2048)

	postRepository := repositories.NewPostRepository(database)
	linkRepository := repositories.NewLinkRepository(database)

	postService := services.NewPostService(postRepository, cacheRepository)
	linkService := services.NewLinkService(linkRepository)

	indexMvc := mvc.New(application.Party(url_index))
	indexMvc.Register(postService)
	indexMvc.Handle(indexCtrl)

	postMvc := mvc.New(application.Party(url_post))
	postMvc.Register(postService)
	postMvc.Handle(postCtrl)

	startMvc := mvc.New(application.Party(url_start))
	startMvc.Register(linkService)
	startMvc.Handle(startCtrl)

	commentMvc := mvc.New(application.Party(url_comment))
	commentMvc.Handle(commentCtrl)

	photoMvc := mvc.New(application.Party(url_photo))
	photoMvc.Handle(photoCtrl)

	aboutMvc := mvc.New(application.Party(url_about))
	aboutMvc.Handle(aboutCtrl)

	sitemapMvc := mvc.New(application.Party(url_sitemap))
	sitemapMvc.Register(postService)
	sitemapMvc.Handle(sitemapCtrl)
}
