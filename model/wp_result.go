package model

type WpPost struct {
	Date string
	DateGmt string
	Guid string `json,format:'rendered:%s'`
	Id int
	Link string
	Modified string
	ModifiedGmt string
	Slug string
	Status string
	Type string `json,name:'type'`
	Password string
	Title string `json,format:'rendered:%s'`
	Content string `json,format:'rendered:%s;protected:false'`
	Author int
	Excerpt string `json,format:'rendered:%s;protected:false'`
	FeaturedMedia int
	CommentStatus string `json,default:'open'`
	PingStatus string `json,default:'open'`
	Format string
	Meta []string
	Sticky bool
	Template string
	Categories []int
	Tags []string
}

type WpPostRevision struct {
	Id int
	Date string
	DateGmt string
	Guid string `json,format:'rendered:%s'`
	Modified string
	ModifiedGmt string
	Slug string
	Parent int
	Title string `json,format:'rendered:%s'`
	Content string `json,format:'rendered:%s;protected:false'`
	Author int
	Excerpt string `json,format:'rendered:%s;protected:false'`
}

type WpCategory struct {
	Id int
	Count string
	Description string
	Link string `json,format:'rendered:%s'`
	Name string
	Taxonomy string
	Slug string
	Parent int
	Meta []string
}

type WpTag struct {
	Id int
	Count string
	Description string
	Link string `json,format:'rendered:%s'`
	Name string
	Taxonomy string
	Slug string
	Parent int
	Meta []string
}


type WpPage struct {
	Id string
	Date string
	DateGmt string
	Guid string `json,format:'rendered:%s'`
	Modified string
	ModifiedGmt string
	Slug string
	Status string
	PostType string `json,name:type`
	Link string
	Title string `json,format:'rendered:%s'`
	Content string `json,rendered:%s;protected:false`
	Excerpt string `json,rendered:%s;protected:false`
	Author int
	FeaturedMedia int
	Parent int
	MenuOrder string
	CommentStatus string `json,default:'closed'`
	PingStatus string `json,default:'open'`
	Template string `json,default:''`
	Meta []string
}

type WpComment struct {
	Id int
	Author int
	AuthorEmail string
	AuthorIp string
	AuthorName string
	AuthorUrl string
	AuthorUserAgent string
	Content string `json,format:'rendered:%s;protected:false'`
	Date string
	DateGmt string
	Link string
	Parent int
	Post int
	Status string
	Type string
	AuthorAvatarUrls []string
	Meta []string
}

type WpTaxonomy struct {
	Capabilities []string
	Description string
	Hierarchical bool
	Labels []string
	Name string
	Slug string
	ShowCloud bool
	Types []string
	RestBase string
}

type WpMedia struct {
	Date string
	Date_gmt string
	Guid string `json,format:'rendered:%s'`
	Id int
	Link string
	Modified string
	ModifiedGmt string
	Slug string
	Status string
	AltType string  `json,name:type`
	Title string `json,format:'rendered:%s'`
	Author int
	CommentStatus string
	PingStatus string
	Meta []string
	Template string
	AltText string
	Caption string `json,format:'rendered:%s'`
	Description string `json,format:'rendered:%s'`
	MediaType string
	MimeType string
	MediaDetails []string
	Post int
	SourceUrl string
}

type WpUser struct {
	Id int
	Username string
	Name string
	FirstName string
	LastName string
	Email string
	Url string
	Description string
	Link string
	Locale string
	Nickname string
	Slug string
	RegisteredDate string
	Roles []string
	Password string
	Capabilities []string
	ExtraCapabilities []string
	AvatarUrls []string
	Meta []string
}

type WpPostType struct {
	Capabilities []string
	Description string
	Hierarchical bool
	Labels []string
	Name string
	Slug string
	Supports []string
	Taxonomies []string
	RestBase string
}

type WpPostStatus struct {
	Name string
	Private bool
	Protected bool
	Public bool
	Queryable bool
	ShowInList bool
	Slug string
}

type WpSetting struct {
	Title string
	Description string
	Timezone string
	DateFormat string
	TimeFormat string
	StartOfWeek int
	Language string
	UseSmilies bool
	DefaultCategory int
	DefaultPostFormat string
	PostsPerPage int
	DefaultPingStatus string
	DefaultCommentStatus string
}

