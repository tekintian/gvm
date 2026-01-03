package official

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/assert"
)

func getCollector() (*Collector, error) {
	// 使用 NewCollector 函数来创建 Collector 实例，确保所有字段都被正确初始化
	c, err := NewCollector(DefaultDownloadPageURL)
	if err != nil {
		return nil, err
	}

	// 然后读取测试数据文件并设置 doc 字段
	b, err := os.ReadFile("./testdata/golang_dl.html")
	if err != nil {
		return nil, err
	}
	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	c.doc = doc

	return c, nil
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

		// 查找所有表格元素
		tables := c.doc.Find("table.downloadtable")
		assert.NotEmpty(t, tables.Text(), "应该找到至少一个表格元素")

		// 使用第一个表格元素
		table := tables.First()
		assert.NotEmpty(t, table.Text(), "第一个表格元素应该不为空")

		// 将表格元素传递给findPackages函数
		pkgs := c.findPackages(table)
		assert.Greater(t, len(pkgs), 0, "应该找到至少一个安装包")
	})
}

func TestUnstableVersions(t *testing.T) {
	t.Run("查询unstable状态的go版本列表", func(t *testing.T) {
		c, err := getCollector()
		assert.Nil(t, err)
		assert.NotNil(t, c)

		// items, err := c.UnstableVersions()
		// assert.Nil(t, err)
		// assert.Equal(t, 62, len(items))
		// assert.Equal(t, "1.19rc2", items[0].Name)
		// assert.Equal(t, 19, len(items[0].Packages))
	})
}

func TestArchivedVersions(t *testing.T) {
	t.Run("查询archived状态的go版本列表", func(t *testing.T) {
		c, err := getCollector()
		assert.Nil(t, err)
		assert.NotNil(t, c)

		// items, err := c.ArchivedVersions()
		// assert.Nil(t, err)
		// assert.Equal(t, 70, len(items))
		// 第一个版本
		// assert.Equal(t, "1.19.2", items[0].Name)
		// assert.Equal(t, 15, len(items[0].Packages))
	})
}
func TestAllVersions(t *testing.T) {
	t.Run("查询所有go版本列表", func(t *testing.T) {
		c, err := getCollector()
		assert.Nil(t, err)
		assert.NotNil(t, c)
		// 这里的测试在不同时间结果不同,到此就OK

	})
}
