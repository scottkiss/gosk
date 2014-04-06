package gosk

import (
	"bufio"
	"fmt"
	"github.com/scottkiss/blackfriday"
	"github.com/scottkiss/go-gypsy/yaml"
	"html"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
    "regexp"
)

type RenderFactory struct{}

const (
	INDEX_TPL    = "index"
	TAG_TPL      = "tag"
	POSTS_TPL    = "posts"
	PAGES_TPL    = "pages"
	RSS_TPL      = "rss"
	CATEGORY_TPL = "category"
	ARCHIVE_TPL  = "archive"
)

const (
	POST_DIR     = "posts"
	PUBLICSH_DIR = "publish"
)

const (
	COMMON_HEADER_FILE = "header.tpl"
	COMMON_FOOTER_FILE = "footer.tpl"
)

var (
	articles        Artic
	articleListSize int = 5000
	rss             []RssConfig
	rssListSize     int = 10
	navBarList      []NavConfig
	allTags         map[string]Tag
	categories      map[string]Category
	pages           []*CustomPage
	archives        map[string]*YearArchive
	allArchive      YearArchives
)

func parseTemplate(root, tpl string, cfg *yaml.File) *template.Template {
	//get theme template
	themeFold, errt := cfg.Get("theme")
	if errt != nil {
		log.Println("get theme error!check config.yml")
		os.Exit(1)
	}

	file := root + "templates/" + themeFold + "/" + tpl + ".tpl"
	if !isExists(file) {
		log.Println(file + " can not be found!")
		os.Exit(1)
	}
	t := template.New(tpl + ".tpl")
	t.Funcs(template.FuncMap{"get": cfg.Get})
	t.Funcs(template.FuncMap{"unescaped": unescaped})

	headerTpl := root + "templates/" + themeFold + "/common/" + COMMON_HEADER_FILE
	footerTpl := root + "templates/" + themeFold + "/common/" + COMMON_FOOTER_FILE

	if !isExists(headerTpl) {
		log.Println(headerTpl + " can not be found!")
		os.Exit(1)
	}

	if !isExists(footerTpl) {
		log.Println(footerTpl + " can not be found!")
		os.Exit(1)
	}

	t, err := t.ParseFiles(file, headerTpl, footerTpl)
	if err != nil {
		log.Println("parse " + tpl + " Template error!" + err.Error())
		os.Exit(1)
	}

	log.Println("parse " + tpl + " Template complete!")
	return t
}

func parseXMLTemplate(root, tpl string, cfg *yaml.File) *template.Template {
	//get theme template
	themeFold, errt := cfg.Get("theme")
	if errt != nil {
		log.Println("get theme error!check config.yml")
		os.Exit(1)
	}
	file := root + "/templates/" + themeFold + "/" + tpl + ".tpl"
	if !isExists(file) {
		log.Println(file + " can not be found!")
		os.Exit(1)
	}
	t := template.New(tpl + ".tpl")
	t.Funcs(template.FuncMap{"get": cfg.Get})
	t.Funcs(template.FuncMap{"unescaped": unescaped})
	t.Funcs(template.FuncMap{"xmlheader": xmlHeader})

	t, err := t.ParseFiles(file)
	if err != nil {
		log.Println("parse " + tpl + " Template error!" + err.Error())
		os.Exit(1)
	}

	log.Println("parse " + tpl + " Template complete!")
	return t
}

func isExists(file string) bool {
	_, err := os.Stat(file)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}

func (self *RenderFactory) RenderIndex(root string, yamls map[string]interface{}) error {
	if !strings.HasSuffix(root, "/") {
		root += "/"
	}

	yCfg := yamls["config.yml"]
	var cfg = yCfg.(*yaml.File)
	t := parseTemplate(root, INDEX_TPL, cfg)

	targetFile := PUBLICSH_DIR + "/index.html"
	fout, err := os.Create(targetFile)
	if err != nil {
		log.Println("create file " + targetFile + " error!")
		os.Exit(1)
	}
	defer fout.Close()

	m := map[string]interface{}{"ar": articles, "nav": navBarList}
	exErr := t.Execute(fout, m)
	return exErr
}

//render tags
func (self *RenderFactory) RenderTag(root string, yamls map[string]interface{}) error {
	if !strings.HasSuffix(root, "/") {
		root += "/"
	}
	yCfg := yamls["config.yml"]
	var cfg = yCfg.(*yaml.File)

	t := parseTemplate(root, TAG_TPL, cfg)
	targetFile := PUBLICSH_DIR + "/tag.html"
	fout, err := os.Create(targetFile)
	if err != nil {
		log.Println("create file " + targetFile + " error!")
		os.Exit(1)
	}
	defer fout.Close()

	generateTags()
	//log.Println(allTags)
	m := map[string]interface{}{"tag": allTags, "nav": navBarList}
	exErr := t.Execute(fout, m)
	return exErr
}

//render categories
func (self *RenderFactory) RenderCategories(root string, yamls map[string]interface{}) error {
	if !strings.HasSuffix(root, "/") {
		root += "/"
	}

	yCfg := yamls["config.yml"]
	var cfg = yCfg.(*yaml.File)

	t := parseTemplate(root, CATEGORY_TPL, cfg)
	targetFile := PUBLICSH_DIR + "/category.html"
	fout, err := os.Create(targetFile)
	if err != nil {
		log.Println("create file " + targetFile + " error!")
		os.Exit(1)
	}
	defer fout.Close()

	generateCategories()
	//log.Println(categories)
	m := map[string]interface{}{"cats": categories, "nav": navBarList}
	exErr := t.Execute(fout, m)
	return exErr
}

//render rss page
func (self *RenderFactory) RenderRss(root string, yamls map[string]interface{}) error {
	if !strings.HasSuffix(root, "/") {
		root += "/"
	}
	yCfg := yamls["config.yml"]
	var cfg = yCfg.(*yaml.File)
	t := parseXMLTemplate(root, RSS_TPL, cfg)
	targetFile := PUBLICSH_DIR + "/rss.xml"
	fout, err := os.Create(targetFile)
	if err != nil {
		log.Println("create file " + targetFile + " error!")
		os.Exit(1)
	}
	defer fout.Close()
	rssCount, err := cfg.Get("rss.max")
	if err != nil {
		log.Println(err)
	}
	rssListSize, errconv := strconv.Atoi(rssCount)
	if errconv != nil {
		log.Println("rss max in config.yml is not a number!" + errconv.Error())
	}
	if len(articles) < rssListSize {
		rssListSize = len(articles)
	}
	ars := articles[:rssListSize]

	for _, ar := range ars {
		rss = append(rss, RssConfig{ar.Title, ar.Link, ar.Author,
			ar.Date, ar.Content})
	}

	r := Rss{time.Now().Format(time.RFC1123), rss}

	exErr := t.Execute(fout, r)
	return exErr
}

//render posts pages
func (self *RenderFactory) RenderPosts(root string, yamls map[string]interface{}) error {
	if !strings.HasSuffix(root, "/") {
		root += "/"
	}
	generateNavBar(yamls)
	articles = make([]*ArticleConfig, 0, articleListSize)
	fileInfos, err := ioutil.ReadDir(root + POST_DIR)
	if err != nil {
		log.Println("read posts dir error!")
	}

	yCfg := yamls["config.yml"]
	var cfg = yCfg.(*yaml.File)
	//log.Println(cfg.Get("title"))
	t := parseTemplate(root, POSTS_TPL, cfg)

	for _, fileInfo := range fileInfos {
		if !fileInfo.IsDir() {
			log.Println("begin process article -- " + fileInfo.Name())
			fileName := fileInfo.Name()
			mardownStr, fi, err := processArticleFile(root+POST_DIR+"/"+fileName, fileName)
			//create post html file

			trName := strings.TrimSuffix(fileName, ".md")
            fmt.Println(trName)

			//process url path => /articles/yyyy/MM/dd/{filename}.html
			p := processArticleUrl(fi)
			//create dir /yyyy/MM/dd
			if !isExists(PUBLICSH_DIR + "/articles/" + p) {
				os.MkdirAll(PUBLICSH_DIR+"/articles/"+p, 0777)
			}
			targetFile := PUBLICSH_DIR + "/articles/" + p + "/" + trName + ".html"
			fout, err := os.Create(targetFile)
			if err != nil {
				log.Println("create file " + targetFile + " error!")
				os.Exit(1)
			}
			defer fout.Close()

			//deal markdown
			htmlByte := blackfriday.MarkdownCommon([]byte(mardownStr))
			//init other article infos
			htmlStr := html.UnescapeString(string(htmlByte))
            re := regexp.MustCompile(`<pre><code>([\s\S]*?)</code></pre>`)
            htmlStr = re.ReplaceAllString(htmlStr, `<pre class="prettyprint linenums">${1}</pre>`)
			fi.Content = htmlStr
			fi.Link = p + trName + ".html"
			//if abstract is empty,auto gen it
			if fi.Abstract == "" {
				var limit int = 500
				rs := []rune(htmlStr)
				if len(rs) < 500 {
					limit = len(rs)
				}

				abstract := subStr(htmlStr, 0, limit)
				fi.Abstract = trimHTML(abstract)
			}
			if fi.Author == "" {
				author, cerr := cfg.Get("meta.author")
				if cerr != nil {
					log.Println(cerr)
				}
				fi.Author = author
			}
			//sort by date
			addAndSortArticles(fi)

			t.Execute(fout, fi)

		}
	}
	return nil
}

func processArticleUrl(ar ArticleConfig) string {
	y := strconv.Itoa(ar.Time.Year())
	m := strconv.Itoa(int(ar.Time.Month()))
	d := strconv.Itoa(ar.Time.Day())
	return y + "/" + m + "/" + d + "/"
}

//render custom pages
func (self *RenderFactory) RenderPages(root string, yamls map[string]interface{}) error {
	if !strings.HasSuffix(root, "/") {
		root += "/"
	}
	yCfg := yamls["config.yml"]
	var cfg = yCfg.(*yaml.File)
	generatePages(yamls)

	t := parseTemplate(root, PAGES_TPL, cfg)

	for _, p := range pages {
		p.Id = strings.TrimSuffix(p.Id, " ")
		filePath := root + "pages/" + p.Id + ".md"
		if !isExists(filePath) {
			log.Println(filePath + " is not found!")
			os.Exit(1)
		}
		f, err := os.Open(filePath)
		if err != nil {
			log.Println(err)

		}
		defer f.Close()
		rd := bufio.NewReader(f)
		var markdownStr string

		for {
			buf, _, err := rd.ReadLine()

			if err == io.EOF {
				break
			} else {
				content := string(buf)
				markdownStr += content + "\n"
			}

		}

		//deal markdown
		htmlByte := blackfriday.MarkdownCommon([]byte(markdownStr))
		//init other article infos
		htmlStr := html.UnescapeString(string(htmlByte))
		htmlStr = strings.Replace(htmlStr, "<pre><code", `<pre class="prettyprint linenums"`, -1)
		htmlStr = strings.Replace(htmlStr, `</code>`, "", -1)
		p.Content = htmlStr
		if !isExists(PUBLICSH_DIR + "/pages/") {
			os.MkdirAll(PUBLICSH_DIR+"/pages/", 0777)
		}
		targetFile := PUBLICSH_DIR + "/pages/" + p.Id + ".html"
		fout, err := os.Create(targetFile)
		if err != nil {
			log.Println("create file " + targetFile + " error!")
			os.Exit(1)
		}
		defer fout.Close()

		m := map[string]interface{}{"p": p, "nav": navBarList}
		t.Execute(fout, m)
	}

	return nil
}

//render archive
func (self *RenderFactory) RenderArchives(root string, yamls map[string]interface{}) error {
	if !strings.HasSuffix(root, "/") {
		root += "/"
	}
	yCfg := yamls["config.yml"]
	var cfg = yCfg.(*yaml.File)

	t := parseTemplate(root, ARCHIVE_TPL, cfg)
	targetFile := PUBLICSH_DIR + "/archive.html"
	fout, err := os.Create(targetFile)
	if err != nil {
		log.Println("create file " + targetFile + " error!")
		os.Exit(1)
	}
	defer fout.Close()

	generateArchive()
	//log.Println(allArchive)
	m := map[string]interface{}{"archives": allArchive, "nav": navBarList}
	exErr := t.Execute(fout, m)
	return exErr

}

func generateArchive() {
	archives = make(map[string]*YearArchive)
	for _, ar := range articles {
		y, m, _ := ar.Time.Date()
		year := fmt.Sprintf("%v", y)
		month := m.String()
		yArchive := archives[year]
		if yArchive == nil {
			yArchive = &YearArchive{year, make([]*MonthArchive, 0), make(map[string]*MonthArchive)}
			archives[year] = yArchive
		}
		mArchive := yArchive.months[month]
		if mArchive == nil {
			mArchive = &MonthArchive{month, m, make([]*ArticleBase, 0)}
			yArchive.months[month] = mArchive
		}
		mArchive.Articles = append(mArchive.Articles, &ArticleBase{ar.Link, ar.Title})

	}
	allArchive = make(YearArchives, 0)
	//sort by time
	for _, yArchive := range archives {
		monthCollect := make(MonthArchives, 0)
		for _, mArchive := range yArchive.months {
			monthCollect = append(monthCollect, mArchive)
		}
		sort.Sort(monthCollect)
		yArchive.months = nil
		yArchive.Months = monthCollect
		allArchive = append(allArchive, yArchive)
	}
	sort.Sort(allArchive)
}

func generateTags() {
	allTags = make(map[string]Tag)
	for _, ar := range articles {
		for _, tg := range ar.Tags {
			//log.Println(tg)
			t, ok := allTags[tg.Name]
			if ok {
				t.Articles = append(t.Articles, ArticleBase{ar.Link, ar.Title})
				t.Length = len(t.Articles)
				allTags[tg.Name] = t
			} else {
				art := ArticleBase{ar.Link, ar.Title}
				arts := make([]ArticleBase, 0)
				arts = append(arts, art)
				allTags[tg.Name] = Tag{tg.Name, arts, 1}
			}
		}
	}
}

func generateCategories() {
	categories = make(map[string]Category)
	for _, ar := range articles {
		c, ok := categories[ar.Category]
		if ok {
			c.Articles = append(c.Articles, ArticleBase{ar.Link, ar.Title})
			c.Length = len(c.Articles)
			categories[ar.Category] = c
		} else {
			art := ArticleBase{ar.Link, ar.Title}
			arts := make([]ArticleBase, 0)
			arts = append(arts, art)
			categories[ar.Category] = Category{ar.Category, arts, 1}
		}
	}
}

//process posts,get article title,post date
func processArticleFile(filePath, fileName string) (string, ArticleConfig, error) {
	f, err := os.Open(filePath)
	if err != nil {
		log.Println(err)

	}
	defer f.Close()
	rd := bufio.NewReader(f)
	var ct int = 0
	var yamlStr, markdownStr string
	for {
		buf, _, err := rd.ReadLine()
		if err == io.EOF {
			break
		} else {
			content := string(buf)
			if content == "---" {
				ct++
			}
			if ct == 2 {
				if content != "---" {
					markdownStr += content + "\n"
				}
			} else {
				yamlStr += content + "\n"
			}

		}

	}
	config := yaml.Config(strings.Replace(yamlStr, "---\n", "", -1))

	title, err := config.Get("title")
	date, err := config.Get("date")
	tagCount, err := config.Count("tags")
	if err != nil {
		log.Println(err)
	}

	var tags []TagConfig
	trName := strings.TrimSuffix(fileName, ".md")
	for i := 0; i < tagCount; i++ {
		tagName, err := config.Get("tags[" + strconv.Itoa(i) + "]")
		if err != nil {
			log.Println("generate Tags error " + err.Error())
		}
		tags = append(tags, TagConfig{tagName, title, trName + ".html"})
	}

	cat, err := config.Get("categories[0]")
	abstract, err := config.Get("abstract")
	author, err := config.Get("author")

	t, terr := time.Parse("2006-01-02 15:04:05", date)
	if terr != nil {
		log.Println(terr)
	}
	//log.Println(t)
	
	shortDate := t.UTC().Format("Jan 2, 2006")

	arInfo := ArticleConfig{title, date,shortDate, cat, tags, abstract, author, t, "", "", navBarList}

	//log.Println(markdownStr)
	return markdownStr, arInfo, nil

}

//sort articles by date
func addAndSortArticles(arInfo ArticleConfig) {
	//log.Println(len(articles))
	artLen := len(articles)
	if artLen < articleListSize {
		articles = append(articles, &arInfo)
	}
	sort.Sort(ByDate{articles})
	//log.Println(len(articles))
}

func unescaped(str string) interface{} { return template.HTML(str) }
func xmlHeader(blank string) string {
	return blank + `<?xml version="1.0" encoding="utf-8"?>`
}

func generateNavBar(yamls map[string]interface{}) {
	yCfg := yamls["nav.yml"]
	var cfg = yCfg.(*yaml.File)
	ct, err := cfg.Count("")
	if err != nil {
		log.Println(err)
	}
	for i := 0; i < ct; i++ {
		name, errn := cfg.Get("[" + strconv.Itoa(i) + "].label")
		if nil != errn {
			log.Println(errn)
		}
		href, errh := cfg.Get("[" + strconv.Itoa(i) + "].href")
		if nil != errh {
			log.Println(errh)
		}
		target, errt := cfg.Get("[" + strconv.Itoa(i) + "].target")
		if nil != errt {
			log.Println(errt)
		}

		nav := NavConfig{name, href, target}
		navBarList = append(navBarList, nav)

	}
	//log.Println(navBarList)
}

//generate custom pages
func generatePages(yamls map[string]interface{}) {
	yCfg := yamls["pages.yml"]
	var cfg = yCfg.(*yaml.File)
	ct, err := cfg.Count("")
	if err != nil {
		log.Println(err)
	}
	for i := 0; i < ct; i++ {
		id, erri := cfg.Get("[" + strconv.Itoa(i) + "].id")
		if nil != erri {
			log.Println(erri)
		}

		title, errt := cfg.Get("[" + strconv.Itoa(i) + "].title")
		if nil != errt {
			log.Println(errt)
		}

		page := CustomPage{id, title, ""}
		pages = append(pages, &page)
	}

}

func (self *RenderFactory) Render(root string) {
	yp := new(YamlParser)
	yamlData := yp.parse(root)
	self.RenderPosts(root, yamlData)
	self.RenderIndex(root, yamlData)
	self.RenderRss(root, yamlData)
	self.RenderTag(root, yamlData)
	self.RenderCategories(root, yamlData)
	self.RenderArchives(root, yamlData)
	self.RenderPages(root, yamlData)
}
