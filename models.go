package stocktwits

import (
	"encoding/json"
	"fmt"
	"time"
)

type Response struct {
	Status int
}

func (r *Response) UnmarshalJSON(b []byte) error {
	var err error
	var parts map[string]json.RawMessage
	if err = json.Unmarshal(b, &parts); err != nil {
		return err
	}
	var tmp Response
	var tmpInt int
	for k, v := range parts {
		switch k {
		case "status":
			// fmt.Printf("unmarshalling %v\n", v)
			if err = json.Unmarshal(v, &tmpInt); err != nil {
				return err
			}
			// fmt.Printf("unmarshalled %v\n", tmpInt)
			tmp.Status = tmpInt
		default:
			fmt.Printf("Response: unhandled case for key `%v` with value `%+v`\n", k, v)
		}
	}
	*r = tmp
	return nil
}

type Cursor struct {
	More  bool
	Since int
	Max   int
}

func (c *Cursor) UnmarshalJSON(b []byte) error {
	var err error
	var parts map[string]json.RawMessage
	if err = json.Unmarshal(b, &parts); err != nil {
		return err
	}
	var tmp Cursor
	var tmpBool bool
	var tmpInt int
	for k, v := range parts {
		switch k {
		case "more":
			if err = json.Unmarshal(v, &tmpBool); err != nil {
				return err
			}
			tmp.More = tmpBool
		case "since":
			if err = json.Unmarshal(v, &tmpInt); err != nil {
				return err
			}
			tmp.Since = tmpInt
		case "max":
			if err = json.Unmarshal(v, &tmpInt); err != nil {
				return err
			}
			tmp.Max = tmpInt
		default:
			fmt.Printf("Cursor: unhandled case for key `%v` with value `%+v`\n", k, v)
		}
	}
	*c = tmp
	return nil
}

type Symbol struct {
	Id             int
	Symbol         string
	Title          string
	Aliases        []string
	IsFollowing    bool
	WatchlistCount int
}

func (s *Symbol) UnmarshalJSON(b []byte) error {
	var err error
	var parts map[string]json.RawMessage
	if err = json.Unmarshal(b, &parts); err != nil {
		return err
	}
	var tmp Symbol
	var tmpInt int
	var tmpString string
	var tmpStrings []string
	var tmpBool bool
	for k, v := range parts {
		switch k {
		case "id":
			if err = json.Unmarshal(v, &tmpInt); err != nil {
				return err
			}
			tmp.Id = tmpInt
		case "symbol":
			if err = json.Unmarshal(v, &tmpString); err != nil {
				return err
			}
			tmp.Symbol = tmpString
		case "title":
			if err = json.Unmarshal(v, &tmpString); err != nil {
				return err
			}
			tmp.Title = tmpString
		case "aliases":
			if err = json.Unmarshal(v, &tmpStrings); err != nil {
				return err
			}
			tmp.Aliases = tmpStrings
		case "is_following":
			if err = json.Unmarshal(v, &tmpBool); err != nil {
				return err
			}
			tmp.IsFollowing = tmpBool
		case "watchlist_count":
			if err = json.Unmarshal(v, &tmpInt); err != nil {
				return err
			}
			tmp.WatchlistCount = tmpInt
		default:
			fmt.Printf("Symbol: unhandled case for key `%v` with value `%+v`\n", k, v)
		}
	}
	*s = tmp
	return nil
}

type Error struct {
	Message string
}

type Message struct {
	Id        int
	Body      string
	CreatedAt time.Time // timestamp
	User
	Source
	Symbols        []Symbol
	Links          []Link
	MentionedUsers []string
	Entities
	Conversation
	Likes
	Reshares
	Errors []Error
}

func (m *Message) UnmarshalJSON(b []byte) error {
	var err error
	var parts map[string]json.RawMessage
	if err = json.Unmarshal(b, &parts); err != nil {
		return err
	}
	var tmp Message
	var tmpInt int
	var tmpString string
	var tmpTime time.Time
	var tmpUser User
	var tmpSource Source
	var tmpSymbols []Symbol
	var tmpUsers []string
	var tmpEntities Entities
	var tmpLinks []Link
	var tmpConversation Conversation
	var tmpLikes Likes
	var tmpReshares Reshares
	for k, v := range parts {
		switch k {
		case "id":
			if err = json.Unmarshal(v, &tmpInt); err != nil {
				return err
			}
			tmp.Id = tmpInt
		case "body":
			if err = json.Unmarshal(v, &tmpString); err != nil {
				return err
			}
			tmp.Body = tmpString
		case "created_at":
			if err = json.Unmarshal(v, &tmpString); err != nil {
				return err
			}
			if tmpTime, err = time.Parse(time.RFC3339, tmpString); err != nil {
				return err
			}
			tmp.CreatedAt = tmpTime
		case "user":
			if err = json.Unmarshal(v, &tmpUser); err != nil {
				return err
			}
			tmp.User = tmpUser
		case "source":
			if err = json.Unmarshal(v, &tmpSource); err != nil {
				return err
			}
			tmp.Source = tmpSource
		case "symbols":
			if err = json.Unmarshal(v, &tmpSymbols); err != nil {
				return err
			}
			tmp.Symbols = tmpSymbols
		case "mentioned_users":
			if err = json.Unmarshal(v, &tmpUsers); err != nil {
				return err
			}
			tmp.MentionedUsers = tmpUsers
		case "entities":
			if err = json.Unmarshal(v, &tmpEntities); err != nil {
				return err
			}
			tmp.Entities = tmpEntities
		case "links":
			if err = json.Unmarshal(v, &tmpLinks); err != nil {
				return err
			}
			tmp.Links = tmpLinks
		case "conversation":
			if err = json.Unmarshal(v, &tmpConversation); err != nil {
				return err
			}
			tmp.Conversation = tmpConversation
		case "likes":
			if err = json.Unmarshal(v, &tmpLikes); err != nil {
				return err
			}
			tmp.Likes = tmpLikes
		case "reshares":
			if err = json.Unmarshal(v, &tmpReshares); err != nil {
				return err
			}
			tmp.Reshares = tmpReshares
		default:
			fmt.Printf("Message: unhandled case for key `%v` with value `%+v`\n", k, v)
		}
	}
	*m = tmp
	return nil
}

type User struct {
	Id                  int
	Username            string
	Name                string
	AvatarUrl           string
	AvatarUrlSsl        string
	JoinDate            time.Time // will just be date
	Official            bool
	Identity            string
	Classification      []string
	Followers           int
	Following           int
	Ideas               int
	WatchlistStockCount int
	LikeCount           int
}

func (u *User) UnmarshalJSON(b []byte) error {
	var err error
	var parts map[string]json.RawMessage
	if err = json.Unmarshal(b, &parts); err != nil {
		return err
	}
	var tmp User
	var tmpInt int
	var tmpString string
	var tmpTime time.Time
	var tmpBool bool
	var tmpStrings []string
	for k, v := range parts {
		switch k {
		case "id":
			if err = json.Unmarshal(v, &tmpInt); err != nil {
				return err
			}
			tmp.Id = tmpInt
		case "username":
			if err = json.Unmarshal(v, &tmpString); err != nil {
				return err
			}
			tmp.Username = tmpString
		case "name":
			if err = json.Unmarshal(v, &tmpString); err != nil {
				return err
			}
			tmp.Name = tmpString
		case "avatar_url":
			if err = json.Unmarshal(v, &tmpString); err != nil {
				return err
			}
			tmp.AvatarUrl = tmpString
		case "avatar_url_ssl":
			if err = json.Unmarshal(v, &tmpString); err != nil {
				return err
			}
			tmp.AvatarUrlSsl = tmpString
		case "join_date":
			if err = json.Unmarshal(v, &tmpString); err != nil {
				return err
			}
			if tmpTime, err = time.Parse("2006-01-02", tmpString); err != nil {
				return err
			}
			tmp.JoinDate = tmpTime
		case "official":
			if err = json.Unmarshal(v, &tmpBool); err != nil {
				return err
			}
			tmp.Official = tmpBool
		case "identity":
			if err = json.Unmarshal(v, &tmpString); err != nil {
				return err
			}
			tmp.Identity = tmpString
		case "followers":
			if err = json.Unmarshal(v, &tmpInt); err != nil {
				return err
			}
			tmp.Followers = tmpInt
		case "following":
			if err = json.Unmarshal(v, &tmpInt); err != nil {
				return err
			}
			tmp.Following = tmpInt
		case "ideas":
			if err = json.Unmarshal(v, &tmpInt); err != nil {
				return err
			}
			tmp.Ideas = tmpInt
		case "watchlist_stocks_count":
			if err = json.Unmarshal(v, &tmpInt); err != nil {
				return err
			}
			tmp.WatchlistStockCount = tmpInt
		case "like_count":
			if err = json.Unmarshal(v, &tmpInt); err != nil {
				return err
			}
			tmp.LikeCount = tmpInt
		case "classification":
			if err = json.Unmarshal(v, &tmpStrings); err != nil {
				return err
			}
			tmp.Classification = tmpStrings
		default:
			fmt.Printf("User: unhandled case for key `%v` with value `%+v`\n", k, v)
		}
	}
	*u = tmp
	return nil
}

type Source struct {
	Id           int
	Title        string
	Url          string
	Name         string
	Website      string
	VideoUrl     string
	ShortenedUrl string
}

func (s *Source) UnmarshalJSON(b []byte) error {
	var err error
	var parts map[string]json.RawMessage
	if err = json.Unmarshal(b, &parts); err != nil {
		return err
	}
	var tmp Source
	var tmpInt int
	var tmpString string
	for k, v := range parts {
		switch k {
		case "id":
			if err = json.Unmarshal(v, &tmpInt); err != nil {
				return err
			}
			tmp.Id = tmpInt
		case "title":
			if err = json.Unmarshal(v, &tmpString); err != nil {
				return err
			}
			tmp.Title = tmpString
		case "url":
			if err = json.Unmarshal(v, &tmpString); err != nil {
				return err
			}
			tmp.Url = tmpString
		case "name":
			if err = json.Unmarshal(v, &tmpString); err != nil {
				return err
			}
			tmp.Name = tmpString
		case "website":
			if err = json.Unmarshal(v, &tmpString); err != nil {
				return err
			}
			tmp.Website = tmpString
		case "video_url":
			if err = json.Unmarshal(v, &tmpString); err != nil {
				return err
			}
			tmp.VideoUrl = tmpString
		case "shortened_url":
			if err = json.Unmarshal(v, &tmpString); err != nil {
				return err
			}
			tmp.ShortenedUrl = tmpString
		default:
			fmt.Printf("Source: unhandled case for key `%v` with value `%+v`\n", k, v)
		}
	}
	*s = tmp
	return nil
}

type Entities struct {
	Chart     Chart
	Sentiment Sentiment
}

func (e *Entities) UnmarshalJSON(b []byte) error {
	var err error
	var parts map[string]json.RawMessage
	if err = json.Unmarshal(b, &parts); err != nil {
		return err
	}
	var tmp Entities
	var tmpChart Chart
	var tmpSentiment Sentiment
	for k, v := range parts {
		switch k {
		case "chart":
			if err = json.Unmarshal(v, &tmpChart); err != nil {
				return err
			}
			tmp.Chart = tmpChart
		case "sentiment":
			if err = json.Unmarshal(v, &tmpSentiment); err != nil {
				return err
			}
			tmp.Sentiment = tmpSentiment
		default:
			fmt.Printf("Entities: unhandled case for key `%v` with value `%+v`\n", k, v)
		}
	}
	*e = tmp
	return nil
}

type Chart struct {
	Thumb    string
	Large    string
	Original string
	Url      string
}

func (c *Chart) UnmarshalJSON(b []byte) error {
	var err error
	var parts map[string]json.RawMessage
	if err = json.Unmarshal(b, &parts); err != nil {
		return err
	}
	var tmp Chart
	var tmpString string
	for k, v := range parts {
		switch k {
		case "thumb":
			if err = json.Unmarshal(v, &tmpString); err != nil {
				return err
			}
			tmp.Thumb = tmpString
		case "large":
			if err = json.Unmarshal(v, &tmpString); err != nil {
				return err
			}
			tmp.Large = tmpString
		case "original":
			if err = json.Unmarshal(v, &tmpString); err != nil {
				return err
			}
			tmp.Original = tmpString
		case "url":
			if err = json.Unmarshal(v, &tmpString); err != nil {
				return err
			}
			tmp.Url = tmpString
		default:
			fmt.Printf("Chart: unhandled case for key `%v` with value `%+v`\n", k, v)
		}
	}
	*c = tmp
	return nil
}

type Sentiment struct {
	Basic string
}

func (s *Sentiment) UnmarshalJSON(b []byte) error {
	var err error
	var parts map[string]json.RawMessage
	if err = json.Unmarshal(b, &parts); err != nil {
		return err
	}
	var tmp Sentiment
	var tmpString string
	for k, v := range parts {
		switch k {
		case "basic":
			if err = json.Unmarshal(v, &tmpString); err != nil {
				return err
			}
			tmp.Basic = tmpString
		default:
			fmt.Printf("Sentiment: unhandled case for key `%v` with value `%+v`\n", k, v)
		}
	}
	*s = tmp
	return nil
}

type Link struct {
	Title                string
	Url                  string
	ShortenedUrl         string
	ShortenedExpandedUrl string
	Description          string
	Image                string
	CreatedAt            string
	VideoUrl             string
	Source
}

func (l *Link) UnmarshalJSON(b []byte) error {
	var err error
	var parts map[string]json.RawMessage
	if err = json.Unmarshal(b, &parts); err != nil {
		return err
	}
	var tmp Link
	var tmpString string
	var tmpSource Source
	for k, v := range parts {
		switch k {
		case "title":
			if err = json.Unmarshal(v, &tmpString); err != nil {
				return err
			}
			tmp.Title = tmpString
		case "url":
			if err = json.Unmarshal(v, &tmpString); err != nil {
				return err
			}
			tmp.Url = tmpString
		case "shortened_url":
			if err = json.Unmarshal(v, &tmpString); err != nil {
				return err
			}
			tmp.ShortenedUrl = tmpString
		case "shortened_expanded_url":
			if err = json.Unmarshal(v, &tmpString); err != nil {
				return err
			}
			tmp.ShortenedExpandedUrl = tmpString
		case "description":
			if err = json.Unmarshal(v, &tmpString); err != nil {
				return err
			}
			tmp.Description = tmpString
		case "image":
			if err = json.Unmarshal(v, &tmpString); err != nil {
				return err
			}
			tmp.Image = tmpString
		case "created_at":
			if err = json.Unmarshal(v, &tmpString); err != nil {
				return err
			}
			tmp.CreatedAt = tmpString
		case "video_url":
			if err = json.Unmarshal(v, &tmpString); err != nil {
				return err
			}
			tmp.VideoUrl = tmpString
		case "source":
			if err = json.Unmarshal(v, &tmpSource); err != nil {
				return err
			}
			tmp.Source = tmpSource
		default:
			fmt.Printf("Link: unhandled case for key `%v` with value `%+v`\n", k, v)
		}
	}
	*l = tmp
	return nil
}

type Conversation struct {
	ParentMessageId    int
	InReplyToMessageId int
	Parent             bool
	Replies            int
}

func (c *Conversation) UnmarshalJSON(b []byte) error {
	var err error
	var parts map[string]json.RawMessage
	if err = json.Unmarshal(b, &parts); err != nil {
		return err
	}
	var tmp Conversation
	var tmpInt int
	var tmpBool bool
	for k, v := range parts {
		switch k {
		case "parent_message_id":
			if err = json.Unmarshal(v, &tmpInt); err != nil {
				return err
			}
			tmp.ParentMessageId = tmpInt
		case "in_reply_to_message_id":
			if err = json.Unmarshal(v, &tmpInt); err != nil {
				return err
			}
			tmp.InReplyToMessageId = tmpInt
		case "parent":
			if err = json.Unmarshal(v, &tmpBool); err != nil {
				return err
			}
			tmp.Parent = tmpBool
		case "replies":
			if err = json.Unmarshal(v, &tmpInt); err != nil {
				return err
			}
			tmp.Replies = tmpInt
		default:
			fmt.Printf("Conversation: unhandled case for key `%v` with value `%+v`\n", k, v)
		}
	}
	*c = tmp
	return nil
}

type Likes struct {
	Total   int
	UserIds []int
}

func (l *Likes) UnmarshalJSON(b []byte) error {
	var err error
	var parts map[string]json.RawMessage
	if err = json.Unmarshal(b, &parts); err != nil {
		return err
	}
	var tmp Likes
	var tmpInt int
	var tmpInts []int
	for k, v := range parts {
		switch k {
		case "total":
			if err = json.Unmarshal(v, &tmpInt); err != nil {
				return err
			}
			tmp.Total = tmpInt
		case "user_ids":
			if err = json.Unmarshal(v, &tmpInts); err != nil {
				return err
			}
			tmp.UserIds = tmpInts
		default:
			fmt.Printf("Likes: unhandled case for key `%v` with value `%+v`\n", k, v)
		}
	}
	*l = tmp
	return nil
}

type Reshares struct {
	ResharedCount int
	UserIds       []int
}

func (r *Reshares) UnmarshalJSON(b []byte) error {
	var err error
	var parts map[string]json.RawMessage
	if err = json.Unmarshal(b, &parts); err != nil {
		return err
	}
	var tmp Reshares
	var tmpInt int
	var tmpInts []int
	for k, v := range parts {
		switch k {
		case "reshared_count":
			if err = json.Unmarshal(v, &tmpInt); err != nil {
				return err
			}
			tmp.ResharedCount = tmpInt
		case "user_ids":
			if err = json.Unmarshal(v, &tmpInts); err != nil {
				return err
			}
			tmp.UserIds = tmpInts
		default:
			fmt.Printf("Likes: unhandled case for key `%v` with value `%+v`\n", k, v)
		}
	}
	*r = tmp
	return nil
}
