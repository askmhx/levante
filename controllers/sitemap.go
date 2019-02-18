package controllers

import (
	"bytes"
	"github.com/kataras/iris"
	"iosxc.com/levante/services"
	"text/template"
)

const (
	siteMapXmlTpl = `<?xml version="1.0" encoding="UTF-8"?>
<urlset>
    {{ range $index , $link := .SiteUrls }}
        <url>
            <loc>{{ $link.Loc }}</loc>
            <lastmod>{{ $link.LastMod }}</lastmod>
            <priority>{{ $link.Priority }}</priority>
        </url>
    {{ end }}
</urlset>`
)

type SiteMapController struct {
	BaseController
	PostService services.PostService
}

func (this *SiteMapController) Get() string {
	var scheme = "https"
	if (this.Ctx.Request().TLS==nil) {
		scheme = "http"
	}
	var host  = scheme+"://"+this.Ctx.Host()
	links := this.PostService.GetLinks(host)
	var t = template.Must(template.New("sitemap").Parse(siteMapXmlTpl))
	var doc bytes.Buffer
	var _ = t.Execute(&doc, iris.Map{
		"SiteUrls": links,
	})
	this.Ctx.ContentType("application/xml")
	return doc.String()
}