package official

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/assert"
)

func getCollector() (*Collector, error) {
	b, err := ioutil.ReadFile("./testdata/golang_dl.html")
	if err != nil {
		return nil, err
	}
	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	return &Collector{
		url: DefaultDownloadPageURL,
		doc: doc,
	}, nil
}
func Test_stablePackages(t *testing.T) {
	c, err := NewCollector("https://go.dev/dl/")
	assert.Nil(t, err)
	assert.NotNil(t, c)
	items, err := c.StableVersions()
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, v := range items {
		fmt.Println(v.Name)
	}
}
func Test_findPackages(t *testing.T) {
	t.Run("查找目标go版本下的安装包列表", func(t *testing.T) {
		c, err := getCollector()
		assert.Nil(t, err)
		assert.NotNil(t, c)

		// 查询archive包中的第一个里面的所有安装包，应该是19个
		pkgs := c.findPackages(c.doc.Find("#archive").Find("div.toggle").First())
		assert.Equal(t, 19, len(pkgs))

		// for i, expected := range []*version.Package{
		// 	{
		// 		FileName:  "go1.13beta1.src.tar.gz",
		// 		URL:       "https://dl.google.com/go/go1.13beta1.src.tar.gz",
		// 		Kind:      version.SourceKind,
		// 		OS:        "",
		// 		Arch:      "",
		// 		Size:      "21MB",
		// 		Checksum:  "e8a7c504cd6775b8a6af101158b8871455918c9a61162f0180f7a9f118dc4102",
		// 		Algorithm: string(checksum.SHA256),
		// 	},
		// } {
		// 	assert.Equal(t, expected.Algorithm, pkgs[i].Algorithm)
		// 	assert.Equal(t, expected.FileName, pkgs[i].FileName)
		// 	assert.Equal(t, expected.Kind, pkgs[i].Kind)
		// 	assert.Equal(t, expected.OS, pkgs[i].OS)
		// 	assert.Equal(t, expected.Arch, pkgs[i].Arch)
		// 	assert.Equal(t, expected.Size, pkgs[i].Size)
		// 	assert.Equal(t, expected.Checksum, pkgs[i].Checksum)
		// }

	})
}

func TestUnstableVersions(t *testing.T) {
	t.Run("查询unstable状态的go版本列表", func(t *testing.T) {
		c, err := getCollector()
		assert.Nil(t, err)
		assert.NotNil(t, c)

		items, err := c.UnstableVersions()
		assert.Nil(t, err)
		assert.Equal(t, 62, len(items))
		assert.Equal(t, "1.19rc2", items[0].Name)
		assert.Equal(t, 19, len(items[0].Packages))
	})
}

func TestArchivedVersions(t *testing.T) {
	t.Run("查询archived状态的go版本列表", func(t *testing.T) {
		c, err := getCollector()
		assert.Nil(t, err)
		assert.NotNil(t, c)

		items, err := c.ArchivedVersions()
		assert.Nil(t, err)
		// assert.Equal(t, 70, len(items))
		// 第一个版本
		assert.Equal(t, "1.19.2", items[0].Name)
		// assert.Equal(t, 15, len(items[0].Packages))
	})
}

func TestAllVersions(t *testing.T) {
	t.Run("查询所有go版本列表", func(t *testing.T) {
		c, err := getCollector()
		assert.Nil(t, err)
		assert.NotNil(t, c)

		items, err := c.AllVersions()
		assert.Nil(t, err)
		// assert.Equal(t, 73, len(items))
		// 最后一个版本
		assert.Equal(t, "1.3rc1", items[len(items)-1].Name)
		// assert.Equal(t, 15, len(items[len(items)-1].Packages))
	})
}
