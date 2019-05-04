package stocktwits

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"
)

var err error

func TestResponseUnmarshal(t *testing.T) {
	x := []byte(`
    {
      "status": 200
    }
  `)
	e := Response{Status: 200}
	var o Response
	if err = json.Unmarshal(x, &o); err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(e, o) {
		t.Errorf("expected %+v, got %+v", e, o)
	}
}

func TestCursorUnmarshal(t *testing.T) {
	x := []byte(`
    {
        "more": true,
        "since": 161970974,
        "max": 161961124
    }
  `)
	e := Cursor{More: true, Since: 161970974, Max: 161961124}
	var o Cursor
	if err = json.Unmarshal(x, &o); err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(e, o) {
		t.Errorf("expected %+v, got %+v", e, o)
	}
}

func TestSymbolUnmarshal(t *testing.T) {
	x := []byte(`
    {
        "id": 5281,
        "symbol": "F",
        "title": "Ford Motor Co.",
        "aliases": [],
        "is_following": false,
        "watchlist_count": 58268
    }
  `)
	e := Symbol{
		Id:             5281,
		Symbol:         "F",
		Title:          "Ford Motor Co.",
		Aliases:        []string{},
		IsFollowing:    false,
		WatchlistCount: 58268}
	var o Symbol
	if err = json.Unmarshal(x, &o); err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(e, o) {
		t.Errorf("expected %+v, got %+v", e, o)
	}
}
func TestChartUnmarshal(t *testing.T) {
	x := []byte(`
    {
        "thumb": "https://charts.stocktwits.com/production/small_161970974.png",
        "large": "https://charts.stocktwits.com/production/large_161970974.png",
        "original": "https://charts.stocktwits.com/production/original_161970974.png",
        "url": "https://charts.stocktwits.com/production/original_161970974.png"
    }
  `)
	e := Chart{
		Thumb:    "https://charts.stocktwits.com/production/small_161970974.png",
		Large:    "https://charts.stocktwits.com/production/large_161970974.png",
		Original: "https://charts.stocktwits.com/production/original_161970974.png",
		Url:      "https://charts.stocktwits.com/production/original_161970974.png"}
	var o Chart
	if err = json.Unmarshal(x, &o); err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(e, o) {
		t.Errorf("expected %+v, got %+v", e, o)
	}
}

func TestUserUnmarshal(t *testing.T) {
	x := []byte(`
    {
        "id": 1014465,
        "username": "BigRolliePollie",
        "name": "Igor Vladir",
        "avatar_url": "https://avatars.stocktwits.com/production/1014465/thumb-1525661407.png",
        "avatar_url_ssl": "https://avatars.stocktwits.com/production/1014465/thumb-1525661407.png",
        "join_date": "2017-03-31",
        "official": false,
        "identity": "User",
        "classification": [],
        "followers": 25,
        "following": 8,
        "ideas": 1567,
        "watchlist_stocks_count": 42,
        "like_count": 646
    }
  `)
	e := User{
		Id:                  1014465,
		Username:            "BigRolliePollie",
		Name:                "Igor Vladir",
		AvatarUrl:           "https://avatars.stocktwits.com/production/1014465/thumb-1525661407.png",
		AvatarUrlSsl:        "https://avatars.stocktwits.com/production/1014465/thumb-1525661407.png",
		JoinDate:            time.Date(2017, 03, 31, 0, 0, 0, 0, time.UTC),
		Official:            false,
		Identity:            "User",
		Classification:      []string{},
		Followers:           25,
		Following:           8,
		Ideas:               1567,
		WatchlistStockCount: 42,
		LikeCount:           646}
	var o User
	if err = json.Unmarshal(x, &o); err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(e, o) {
		t.Errorf("expected %+v, got %+v", e, o)
	}
}

func TestSourceUnmarshal(t *testing.T) {
	x := []byte(`
    {
        "id": 1149,
        "title": "StockTwits for iOS",
        "url": "http://www.stocktwits.com/mobile"
    }
  `)
	e := Source{
		Id:    1149,
		Title: "StockTwits for iOS",
		Url:   "http://www.stocktwits.com/mobile"}
	var o Source
	if err = json.Unmarshal(x, &o); err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(e, o) {
		t.Errorf("expected %+v, got %+v", e, o)
	}
}

func TestSymbolsUnmarshal(t *testing.T) {
	x := []byte(`
    [
        {
            "id": 5281,
            "symbol": "F",
            "title": "Ford Motor Co.",
            "aliases": [],
            "is_following": false,
            "watchlist_count": 58268
        }
    ]
  `)
	e := []Symbol{
		Symbol{
			Id:             5281,
			Symbol:         "F",
			Title:          "Ford Motor Co.",
			Aliases:        []string{},
			IsFollowing:    false,
			WatchlistCount: 58268}}
	var o []Symbol
	if err = json.Unmarshal(x, &o); err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(e, o) {
		t.Errorf("expected %+v, got %+v", e, o)
	}
}

func TestSentimentUnmarshal(t *testing.T) {
	x := []byte(`
    {
        "basic": "Bullish"
    }
  `)
	e := Sentiment{
		Basic: "Bullish"}
	var o Sentiment
	if err = json.Unmarshal(x, &o); err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(e, o) {
		t.Errorf("expected %+v, got %+v", e, o)
	}
}

func TestMessageUnmarshal(t *testing.T) {
	x := []byte(`
    {
        "id": 161970974,
        "body": "$F 21yrs of bliss",
        "created_at": "2019-04-26T02:43:00Z",
        "user": {
            "id": 1014465,
            "username": "BigRolliePollie",
            "name": "Igor Vladir",
            "avatar_url": "https://avatars.stocktwits.com/production/1014465/thumb-1525661407.png",
            "avatar_url_ssl": "https://avatars.stocktwits.com/production/1014465/thumb-1525661407.png",
            "join_date": "2017-03-31",
            "official": false,
            "identity": "User",
            "classification": [],
            "followers": 25,
            "following": 8,
            "ideas": 1567,
            "watchlist_stocks_count": 42,
            "like_count": 646
        },
        "source": {
            "id": 1149,
            "title": "StockTwits for iOS",
            "url": "http://www.stocktwits.com/mobile"
        },
        "symbols": [
            {
                "id": 5281,
                "symbol": "F",
                "title": "Ford Motor Co.",
                "aliases": [],
                "is_following": false,
                "watchlist_count": 58268
            }
        ],
        "mentioned_users": [],
        "entities": {
            "chart": {
                "thumb": "https://charts.stocktwits.com/production/small_161970974.png",
                "large": "https://charts.stocktwits.com/production/large_161970974.png",
                "original": "https://charts.stocktwits.com/production/original_161970974.png",
                "url": "https://charts.stocktwits.com/production/original_161970974.png"
            },
            "sentiment": null
        }
    }
  `)
	e := Message{
		Id:        161970974,
		Body:      "$F 21yrs of bliss",
		CreatedAt: time.Date(2019, 04, 26, 2, 43, 0, 0, time.UTC),
		User: User{
			Id:                  1014465,
			Username:            "BigRolliePollie",
			Name:                "Igor Vladir",
			AvatarUrl:           "https://avatars.stocktwits.com/production/1014465/thumb-1525661407.png",
			AvatarUrlSsl:        "https://avatars.stocktwits.com/production/1014465/thumb-1525661407.png",
			JoinDate:            time.Date(2017, 03, 31, 0, 0, 0, 0, time.UTC),
			Official:            false,
			Identity:            "User",
			Classification:      []string{},
			Followers:           25,
			Following:           8,
			Ideas:               1567,
			WatchlistStockCount: 42,
			LikeCount:           646},
		Source: Source{
			Id:           1149,
			Title:        "StockTwits for iOS",
			Url:          "http://www.stocktwits.com/mobile",
			Name:         "",
			Website:      "",
			VideoUrl:     "",
			ShortenedUrl: ""},
		Symbols: []Symbol{
			Symbol{
				Id:             5281,
				Symbol:         "F",
				Title:          "Ford Motor Co.",
				Aliases:        []string{},
				IsFollowing:    false,
				WatchlistCount: 58268}},
		Links:          []Link{},
		MentionedUsers: []string{},
		Entities: Entities{
			Chart: Chart{
				Thumb:    "https://charts.stocktwits.com/production/small_161970974.png",
				Large:    "https://charts.stocktwits.com/production/large_161970974.png",
				Original: "https://charts.stocktwits.com/production/original_161970974.png",
				Url:      "https://charts.stocktwits.com/production/original_161970974.png"},
			Sentiment: Sentiment{}},
		Conversation: Conversation{
			ParentMessageId:    0,
			InReplyToMessageId: 0,
			Parent:             false,
			Replies:            0},
		Likes: Likes{
			Total:   0,
			UserIds: []int{}},
		Errors: []Error{}}

	var o Message
	if err = json.Unmarshal(x, &o); err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(e, o) {
		t.Errorf("expected %+v, got %+v", e, o)
	}
}
