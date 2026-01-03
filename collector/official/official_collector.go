package official

import (
	"fmt"
	"net/http"
	stdurl "net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/tekintian/gvm/pkg/errs"
	"github.com/tekintian/gvm/version"
)

// var _ collector.Collector = (*Collector)(nil)

const (
	// DefaultDownloadPageURL 官方下载站点网址
	DefaultDownloadPageURL = "https://go.dev/dl/"
)

// Collector 官方站点版本采集器
type Collector struct {
	url  string
	pURL *stdurl.URL
	doc  *goquery.Document
}

// NewCollector 返回采集器实例
func NewCollector(url string) (*Collector, error) {
	if url == "" {
		url = DefaultDownloadPageURL
	}
	pURL, err := stdurl.Parse(url)
	if err != nil {
		return nil, err
	}

	c := Collector{
		url:  url,
		pURL: pURL,
	}
	if err = c.loadDocument(); err != nil {
		return nil, err
	}
	return &c, nil
}

func (c *Collector) loadDocument() (err error) {
	resp, err := http.Get(c.url)
	if err != nil {
		return errs.NewURLUnreachableError(c.url, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return errs.NewURLUnreachableError(c.url, nil)
	}
	c.doc, err = goquery.NewDocumentFromReader(resp.Body)
	return err
}

func (c *Collector) findPackages(table *goquery.Selection) (pkgs []*version.Package) {
	alg := strings.TrimSuffix(table.Find("thead").Find("th").Last().Text(), " Checksum") // 获取 table > thead > tr > 最后一个 th的文本，然后去除  Checksum
	table.Find("tbody").Find("tr").Each(func(i int, s *goquery.Selection) {
		td := s.Find("td")
		href := td.First().Find("a").AttrOr("href", "")
		if strings.HasPrefix(href, "/") {
			href = fmt.Sprintf("%s://%s%s", c.pURL.Scheme,c.pURL.Host, href)
		}
		pkgs = append(pkgs, &version.Package{
			FileName:  td.First().Find("a").Text(),
			URL:       href,
			Kind:      td.Eq(1).Text(),
			OS:        td.Eq(2).Text(),
			Arch:      td.Eq(3).Text(),
			Size:      td.Eq(4).Text(),
			Checksum:  td.Eq(5).Text(),
			Algorithm: alg,
		})
	})

	return pkgs
}

// StableVersions 返回所有稳定版本
func (c *Collector) StableVersions() (items []*version.Version, err error) {

	//先拿到最新稳定版的div
	// 查询id为stable的元素到 id为archive之间的所有元素
	divsStable := c.doc.Find("#stable").NextUntil("#archive")
	divsStable.Each(func(i int, s *goquery.Selection) {
		vname, ok := s.Attr("id")
		if !ok {
			return
		}
		items = append(items, &version.Version{
			Name:     strings.TrimPrefix(vname, "go"),
			Packages: c.findPackages(s.Find("table").First()),
		})
	})

	// 再从历史版本中拿stable版本

	// 这个doc选择的内容为id="archive"元素下面的样式为.expanded下的div
	// <div id="archive"><div class="expanded""><div xxx></div></div></div>
	// divsArchive := c.doc.Find("#archive").ChildrenFiltered(".expanded").Find("div")
	// 查询id为archive下面的所有div的class为toggle的div元素
	divsArchive := c.doc.Find("#archive").Find("div.toggle")
	divsArchive.Each(func(i int, div *goquery.Selection) {
		//这里的div为 <div class="toggle" id="go1.19.2">
		vname, ok := div.Attr("id")
		if !ok || !strings.HasPrefix(vname, "go") {
			return
		}
		if strings.Contains(vname, "rc") || strings.Contains(vname, "beta") {
			return
		}
		items = append(items, &version.Version{
			Name:     strings.TrimPrefix(vname, "go"),
			Packages: c.findPackages(div.Find("table").First()),
		})
	})

	return items, nil
}

// UnstableVersions 返回所有非稳定版本
func (c *Collector) UnstableVersions() (items []*version.Version, err error) {

	c.doc.Find("#archive").ChildrenFiltered(".expanded").Find("div").Each(func(i int, div *goquery.Selection) {
		vname, ok := div.Attr("id")
		if !ok {
			return
		}
		//仅返回非稳定版
		if strings.Contains(vname, "rc") || strings.Contains(vname, "beta") {
			items = append(items, &version.Version{
				Name:     strings.TrimPrefix(vname, "go"),
				Packages: c.findPackages(div.Find("table").First()),
			})
		}
	})
	return items, nil
}

// 过滤版本 仅返回符合条件的版本
func (c *Collector) FilterVersions(condition string) (items []*version.Version, err error) {
	c.doc.Find("#archive").Find("div.toggle").Each(func(i int, div *goquery.Selection) {
		vname, ok := div.Attr("id")
		if !ok {
			return
		}
		if strings.Contains(vname, condition) {
			items = append(items, &version.Version{
				Name:     strings.TrimPrefix(vname, "go"),
				Packages: c.findPackages(div.Find("table").First()),
			})
		}
	})
	return items, nil
}

// ArchivedVersions 返回已归档版本
func (c *Collector) ArchivedVersions() (items []*version.Version, err error) {
	c.doc.Find("#archive").Find("div.toggle").Each(func(i int, div *goquery.Selection) {
		vname, ok := div.Attr("id")
		if !ok {
			return
		}
		items = append(items, &version.Version{
			Name:     strings.TrimPrefix(vname, "go"),
			Packages: c.findPackages(div.Find("table").First()),
		})
	})
	return items, nil
}

// AllVersions 返回所有已知版本
func (c *Collector) AllVersions() (items []*version.Version, err error) {
	items, err = c.StableVersions()
	if err != nil {
		return nil, err
	}
	archives, err := c.ArchivedVersions()
	if err != nil {
		return nil, err
	}
	items = append(items, archives...)

	unstables, err := c.UnstableVersions()
	if err != nil {
		return nil, err
	}
	items = append(items, unstables...)
	return items, nil
}
