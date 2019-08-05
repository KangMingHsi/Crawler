package view

import (
	"os"
	"spider/engine"
	"spider/frontend/model"
	datamodel "spider/model"
	"testing"
)

func TestSearchResultView(t *testing.T) {
	view := CreateSearchResultView("template.html")

	page := model.SearchResult{
		Hits:  1,
		Start: 0,
		Items: []engine.Item{
			engine.Item{
				URL:  "https://www.ptt.cc/bbs/LoL/M.1562223193.A.CA4.html",
				ID:   "[戰棋] 虛空生物增加三隻變成3/6BUFF有搞頭嗎",
				Type: "message",
				Payload: []datamodel.Message{
					datamodel.Message{
						IsRecommended: "→",
						AccountName:   "timcida",
						Msg:           "拿弓弩的才會叫遊俠吧",
						Time:          " 07/04 14:54",
					},

					datamodel.Message{
						IsRecommended: "噓",
						AccountName:   "RaysMoon",
						Msg:           "唉",
						Time:          " 07/04 14:54",
					},
				},
			},

			engine.Item{
				URL:  "https://www.ptt.cc/bbs/LoL/M.1562222306.A.DD1.html",
				ID:   "[電競] 2019 Rift Rivals:KR/CN/LMS/VN Day1",
				Type: "message",
				Payload: []datamodel.Message{
					datamodel.Message{
						IsRecommended: "推",
						AccountName:   "FlandreMiku",
						Msg:           "死人棋邀請賽?",
						Time:          " 07/04 14:38",
					},

					datamodel.Message{
						IsRecommended: "推",
						AccountName:   "homeqq520",
						Msg:           "沒空 下棋",
						Time:          " 07/04 14:39",
					},
				},
			},
		},
	}

	out, err := os.Create("template.test.html")

	err = view.Render(out, page)

	if err != nil {
		panic(err)
	}
}
