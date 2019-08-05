package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseArticleList(t *testing.T) {
	contents, err := ioutil.ReadFile("articlelist_test_data.html")
	if err != nil {
		panic(err)
	}

	result := ParseArticleList(contents, "https://www.ptt.cc")

	expectedUrls := []string{
		"/bbs/LoL/M.1563039118.A.CD9.html", "/bbs/LoL/M.1563039515.A.41D.html", "/bbs/LoL/M.1554037664.A.684.html",
	}
	expectedArticles := []string{
		"Re: [閒聊] LEC 賽前閒聊 : Jiizuke 的推特", "[電競] 2019 LCS Summer W6D1", "[公告] 伺服器狀況詢問/聊天/揪團/抱怨/多功能區",
	}

	const resultSize = 6

	if len(result.Requests) != resultSize+1 {
		t.Errorf("result should have %d resquests; but had %d", resultSize+1, len(result.Requests))
	}

	for i, url := range expectedUrls {
		if result.Requests[i].URL != url {
			t.Errorf("expected url #%d: %s; but was %s", i, url, result.Requests[i].URL)
		}
	}

	if len(result.Items) != resultSize {
		t.Errorf("result should have %d resquests; but had %d", len(result.Items), resultSize)
	}

	for i, article := range expectedArticles {
		if result.Items[i].Payload.(string) != article {
			t.Errorf("expected article #%d: %s; but was %s", i, article, result.Items[i].Payload.(string))
		}
	}
}
