package gosk

import (
	"testing"
)

func TestRenderPosts(t *testing.T) {
	var rf = new(RenderFactory)
	yp := new(YamlParser)
	yamlData := yp.parse("root")
	err := rf.RenderPosts("root", yamlData)
	if err != nil {
		t.Errorf("Render Index failed! -- " + err.Error())
	}
}

func TestRenderIndex(t *testing.T) {

	var rf = new(RenderFactory)
	yp := new(YamlParser)
	yamlData := yp.parse("root")
	err := rf.RenderIndex("root", yamlData)
	if err != nil {
		t.Errorf("Render Index failed! -- " + err.Error())
	}
}

func TestRenderRss(t *testing.T) {
	var rf = new(RenderFactory)
	yp := new(YamlParser)
	yamlData := yp.parse("root")
	err := rf.RenderRss("root", yamlData)
	if err != nil {
		t.Errorf("Render Rss failed! -- " + err.Error())
	}
}

func TestRenderTag(t *testing.T) {
	var rf = new(RenderFactory)
	yp := new(YamlParser)
	yamlData := yp.parse("root")
	err := rf.RenderTag("root", yamlData)
	if err != nil {
		t.Errorf("Render Tag failed! -- " + err.Error())
	}
}

func TestRenderCategories(t *testing.T) {
	var rf = new(RenderFactory)
	yp := new(YamlParser)
	yamlData := yp.parse("root")
	err := rf.RenderCategories("root", yamlData)
	if err != nil {
		t.Errorf("Render Tag failed! -- " + err.Error())
	}
}

func TestRenderArchives(t *testing.T) {
	var rf = new(RenderFactory)
	yp := new(YamlParser)
	yamlData := yp.parse("root")
	err := rf.RenderArchives("root", yamlData)
	if err != nil {
		t.Errorf("Render Pages failed! -- " + err.Error())
	}
}

func TestRenderPages(t *testing.T) {
	var rf = new(RenderFactory)
	yp := new(YamlParser)
	yamlData := yp.parse("root")
	err := rf.RenderPages("root", yamlData)
	if err != nil {
		t.Errorf("Render Pages failed! -- " + err.Error())
	}
}
