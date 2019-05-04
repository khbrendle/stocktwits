package stocktwits

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestFormatUrl(t *testing.T) {
	symbol := "F"
	var o string
	o = FormatUrl(symbol, map[string]interface{}{"max": 13})
	if o != "https://api.stocktwits.com/api/2/streams/symbol/F.json?max=13" {
		t.Errorf("TestFormatUrl run 1 failed")
	}
	o = FormatUrl(symbol, map[string]interface{}{"max": 13, "limit": 5})
	if o != "https://api.stocktwits.com/api/2/streams/symbol/F.json?max=13&limit=5" {
		fmt.Println(o)
		// t.Errorf("TestFormatUrl run 1 failed")
	}
}

func TestStreamSymbolUnmarshal(t *testing.T) {
	x := []byte(`
    {
        "response": {
            "status": 200
        },
        "symbol": {
            "id": 5281,
            "symbol": "F",
            "title": "Ford Motor Co.",
            "aliases": [],
            "is_following": false,
            "watchlist_count": 58268
        },
        "cursor": {
            "more": true,
            "since": 161970974,
            "max": 161961124
        },
        "messages": [
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
            },
            {
                "id": 161970555,
                "body": "$F DTCC probably knows nothing 💁🏻‍♂️",
                "created_at": "2019-04-26T02:37:20Z",
                "user": {
                    "id": 1036165,
                    "username": "CADLEV",
                    "name": "Erik",
                    "avatar_url": "https://avatars.stocktwits.com/production/1036165/thumb-1532653894.png",
                    "avatar_url_ssl": "https://avatars.stocktwits.com/production/1036165/thumb-1532653894.png",
                    "join_date": "2017-04-23",
                    "official": false,
                    "identity": "User",
                    "classification": [],
                    "followers": 52,
                    "following": 188,
                    "ideas": 4413,
                    "watchlist_stocks_count": 135,
                    "like_count": 4652
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
                        "thumb": "https://charts.stocktwits.com/production/small_161970555.png",
                        "large": "https://charts.stocktwits.com/production/large_161970555.png",
                        "original": "https://charts.stocktwits.com/production/original_161970555.png",
                        "url": "https://charts.stocktwits.com/production/original_161970555.png"
                    },
                    "sentiment": null
                }
            },
            {
                "id": 161970151,
                "body": "Hope all are doing good here 🙏. \nLet&#39;s begin the fun 👏\nTomorrow&#39;s watch list $RBZ $TLSA ( Tiziana Life) $AMZN $HUSA $DTEA $F $YUMA",
                "created_at": "2019-04-26T02:31:41Z",
                "user": {
                    "id": 1019959,
                    "username": "todmacx",
                    "name": "Todmacx",
                    "avatar_url": "https://avatars.stocktwits.com/production/1019959/thumb-1496308904.png",
                    "avatar_url_ssl": "https://avatars.stocktwits.com/production/1019959/thumb-1496308904.png",
                    "join_date": "2017-04-05",
                    "official": false,
                    "identity": "User",
                    "classification": [],
                    "followers": 804,
                    "following": 66,
                    "ideas": 3687,
                    "watchlist_stocks_count": 300,
                    "like_count": 258
                },
                "source": {
                    "id": 2095,
                    "title": "StockTwits For Android ",
                    "url": "http://www.stocktwits.com/mobile"
                },
                "symbols": [
                    {
                        "id": 864,
                        "symbol": "AMZN",
                        "title": "Amazon.com Inc.",
                        "aliases": [],
                        "is_following": false,
                        "watchlist_count": 203545
                    },
                    {
                        "id": 2196,
                        "symbol": "HUSA",
                        "title": "Houston American Energy Corp.",
                        "aliases": [],
                        "is_following": false,
                        "watchlist_count": 12198
                    },
                    {
                        "id": 5281,
                        "symbol": "F",
                        "title": "Ford Motor Co.",
                        "aliases": [],
                        "is_following": false,
                        "watchlist_count": 58268
                    },
                    {
                        "id": 12573,
                        "symbol": "DTEA",
                        "title": "DAVIDsTEA",
                        "aliases": [],
                        "is_following": false,
                        "watchlist_count": 5980
                    },
                    {
                        "id": 13010,
                        "symbol": "YUMA",
                        "title": "Yuma Energy",
                        "aliases": [],
                        "is_following": false,
                        "watchlist_count": 9139
                    },
                    {
                        "id": 15058,
                        "symbol": "RBZ",
                        "title": "Reebonz Holding Limited Ordinary ",
                        "aliases": [],
                        "is_following": false,
                        "watchlist_count": 6528
                    }
                ],
                "likes": {
                    "total": 2,
                    "user_ids": [
                        1638964,
                        996976
                    ]
                },
                "mentioned_users": [],
                "entities": {
                    "sentiment": {
                        "basic": "Bullish"
                    }
                }
            },
            {
                "id": 161969938,
                "body": "$F Huge Beat",
                "created_at": "2019-04-26T02:28:39Z",
                "user": {
                    "id": 363578,
                    "username": "mbrown2",
                    "name": "Mark Brown",
                    "avatar_url": "http://avatars.stocktwits.com/images/default_avatar_thumb.jpg",
                    "avatar_url_ssl": "https://s3.amazonaws.com/st-avatars/images/default_avatar_thumb.jpg",
                    "join_date": "2014-06-11",
                    "official": false,
                    "identity": "User",
                    "classification": [],
                    "followers": 6,
                    "following": 42,
                    "ideas": 269,
                    "watchlist_stocks_count": 44,
                    "like_count": 355
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
                "likes": {
                    "total": 2,
                    "user_ids": [
                        1414010,
                        2035539
                    ]
                },
                "mentioned_users": [],
                "entities": {
                    "sentiment": {
                        "basic": "Bullish"
                    }
                }
            },
            {
                "id": 161969917,
                "body": "$F Should see 11$ tomorrow folks. Congrats",
                "created_at": "2019-04-26T02:28:17Z",
                "user": {
                    "id": 363578,
                    "username": "mbrown2",
                    "name": "Mark Brown",
                    "avatar_url": "http://avatars.stocktwits.com/images/default_avatar_thumb.jpg",
                    "avatar_url_ssl": "https://s3.amazonaws.com/st-avatars/images/default_avatar_thumb.jpg",
                    "join_date": "2014-06-11",
                    "official": false,
                    "identity": "User",
                    "classification": [],
                    "followers": 6,
                    "following": 42,
                    "ideas": 269,
                    "watchlist_stocks_count": 44,
                    "like_count": 355
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
                "likes": {
                    "total": 1,
                    "user_ids": [
                        2035539
                    ]
                },
                "mentioned_users": [],
                "entities": {
                    "sentiment": {
                        "basic": "Bullish"
                    }
                }
            },
            {
                "id": 161969875,
                "body": "$F",
                "created_at": "2019-04-26T02:27:40Z",
                "user": {
                    "id": 948949,
                    "username": "Dalluge",
                    "name": "Aaron Dalluge",
                    "avatar_url": "https://avatars.stocktwits.com/production/948949/thumb-1551199285.png",
                    "avatar_url_ssl": "https://avatars.stocktwits.com/production/948949/thumb-1551199285.png",
                    "join_date": "2017-02-08",
                    "official": false,
                    "identity": "User",
                    "classification": [],
                    "followers": 62,
                    "following": 47,
                    "ideas": 4235,
                    "watchlist_stocks_count": 9,
                    "like_count": 25040
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
                "likes": {
                    "total": 1,
                    "user_ids": [
                        2035539
                    ]
                },
                "mentioned_users": [],
                "entities": {
                    "sentiment": {
                        "basic": "Bullish"
                    }
                }
            },
            {
                "id": 161969846,
                "body": "$GRUB will come up big tomorrow! Too many people focused on $INTC $AMZN $F etc!  Stoners will watch $NFLX and order Taco Bell! No brainer",
                "created_at": "2019-04-26T02:27:11Z",
                "user": {
                    "id": 1579017,
                    "username": "G_from_Mia",
                    "name": "Guillermo",
                    "avatar_url": "https://avatars.stocktwits.com/production/1579017/thumb-1534474010.png",
                    "avatar_url_ssl": "https://avatars.stocktwits.com/production/1579017/thumb-1534474010.png",
                    "join_date": "2018-08-17",
                    "official": false,
                    "identity": "User",
                    "classification": [],
                    "followers": 2,
                    "following": 9,
                    "ideas": 45,
                    "watchlist_stocks_count": 1,
                    "like_count": 23
                },
                "source": {
                    "id": 2269,
                    "title": "StockTwits Web",
                    "url": "https://stocktwits.com"
                },
                "symbols": [
                    {
                        "id": 864,
                        "symbol": "AMZN",
                        "title": "Amazon.com Inc.",
                        "aliases": [],
                        "is_following": false,
                        "watchlist_count": 203536
                    },
                    {
                        "id": 2310,
                        "symbol": "INTC",
                        "title": "Intel Corporation",
                        "aliases": [],
                        "is_following": false,
                        "watchlist_count": 54094
                    },
                    {
                        "id": 2839,
                        "symbol": "NFLX",
                        "title": "Netflix, Inc.",
                        "aliases": [],
                        "is_following": false,
                        "watchlist_count": 173279
                    },
                    {
                        "id": 5281,
                        "symbol": "F",
                        "title": "Ford Motor Co.",
                        "aliases": [],
                        "is_following": false,
                        "watchlist_count": 58268
                    },
                    {
                        "id": 11844,
                        "symbol": "GRUB",
                        "title": "GrubHub",
                        "aliases": [],
                        "is_following": false,
                        "watchlist_count": 6730
                    }
                ],
                "conversation": {
                    "parent_message_id": 161969846,
                    "in_reply_to_message_id": null,
                    "parent": true,
                    "replies": 1
                },
                "mentioned_users": [],
                "entities": {
                    "sentiment": null
                }
            },
            {
                "id": 161969569,
                "body": "$F What do you think about tomorrow, will this close higher? Or PB?",
                "created_at": "2019-04-26T02:23:11Z",
                "user": {
                    "id": 2016514,
                    "username": "Julianteachesfi",
                    "name": "Julian",
                    "avatar_url": "https://avatars.stocktwits.com/production/2016514/thumb-1555538341.png",
                    "avatar_url_ssl": "https://avatars.stocktwits.com/production/2016514/thumb-1555538341.png",
                    "join_date": "2019-04-17",
                    "official": false,
                    "identity": "User",
                    "classification": [],
                    "followers": 0,
                    "following": 19,
                    "ideas": 82,
                    "watchlist_stocks_count": 32,
                    "like_count": 91
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
                    "sentiment": null
                }
            },
            {
                "id": 161969535,
                "body": "$F self driving cars by end of year? $TSLA has head start but isn’t even profitable yet.  $F autonomous vehicles! https://www.cnbc.com/amp/2019/04/25/ford-aims-for-100-self-driving-cars-on-the-road-by-the-end-of-2019.html",
                "created_at": "2019-04-26T02:22:41Z",
                "user": {
                    "id": 1568163,
                    "username": "MoonWalk3000",
                    "name": "David",
                    "avatar_url": "https://avatars.stocktwits.com/production/1568163/thumb-1533658173.png",
                    "avatar_url_ssl": "https://avatars.stocktwits.com/production/1568163/thumb-1533658173.png",
                    "join_date": "2018-08-06",
                    "official": false,
                    "identity": "User",
                    "classification": [],
                    "followers": 18,
                    "following": 11,
                    "ideas": 1218,
                    "watchlist_stocks_count": 18,
                    "like_count": 1027
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
                    },
                    {
                        "id": 8660,
                        "symbol": "TSLA",
                        "title": "Tesla, Inc.",
                        "aliases": [],
                        "is_following": false,
                        "watchlist_count": 210996
                    }
                ],
                "conversation": {
                    "parent_message_id": 161969535,
                    "in_reply_to_message_id": null,
                    "parent": true,
                    "replies": 7
                },
                "links": [
                    {
                        "title": "Ford claims it will have 100 self-driving cars on the road by the end of the year",
                        "url": "https://www.cnbc.com/amp/2019/04/25/ford-aims-for-100-self-driving-cars-on-the-road-by-the-end-of-2019.html",
                        "shortened_url": "https://www.cnbc.com/amp/2019/04/25/ford-aims-for-100-self-driving-cars-on-the-road-by-the-end-of-2019.html",
                        "shortened_expanded_url": "cnbc.com/2019/04/25/ford-ai...",
                        "description": "Richard Levine | Getty Images Ford aims to have 100 self-driving vehicles on the road by the end of the year, its president of mobility said on Thursday. After reporting a top- and bottom-line earnings beat, Ford's Marcy Klevorn said on the company earnings call that the automaker's autonomous vehicle efforts are going well in Miami and Washington, D.C.",
                        "image": "https://fm.cnbc.com/applications/cnbc.com/resources/img/editorial/2018/10/09/105495697-1539085222459gettyimages-526666914.1910x1000.jpg",
                        "created_at": "2019-04-26T02:22:42Z",
                        "video_url": null,
                        "source": {
                            "name": "CNBC",
                            "website": "https://www.cnbc.com"
                        }
                    }
                ],
                "likes": {
                    "total": 2,
                    "user_ids": [
                        1678424,
                        1414010
                    ]
                },
                "mentioned_users": [],
                "entities": {
                    "sentiment": {
                        "basic": "Bullish"
                    }
                }
            },
            {
                "id": 161969294,
                "body": "$RBZ $NBEV $BTC.X $F $NVDA $INTC $SBUX 98% Professional traders recommended this room: urlzs.com/C1yU",
                "created_at": "2019-04-26T02:19:09Z",
                "user": {
                    "id": 2037129,
                    "username": "ralmibegno",
                    "name": "ralmibegno",
                    "avatar_url": "http://avatars.stocktwits.com/images/default_avatar_thumb.jpg",
                    "avatar_url_ssl": "https://s3.amazonaws.com/st-avatars/images/default_avatar_thumb.jpg",
                    "join_date": "2019-04-26",
                    "official": false,
                    "identity": "User",
                    "classification": [],
                    "followers": 0,
                    "following": 0,
                    "ideas": 0,
                    "watchlist_stocks_count": 0,
                    "like_count": 7
                },
                "source": {
                    "id": 2269,
                    "title": "StockTwits Web",
                    "url": "https://stocktwits.com"
                },
                "symbols": [
                    {
                        "id": 2310,
                        "symbol": "INTC",
                        "title": "Intel Corporation",
                        "aliases": [],
                        "is_following": false,
                        "watchlist_count": 54094
                    },
                    {
                        "id": 2925,
                        "symbol": "NVDA",
                        "title": "NVIDIA Corporation",
                        "aliases": [],
                        "is_following": false,
                        "watchlist_count": 133013
                    },
                    {
                        "id": 3465,
                        "symbol": "SBUX",
                        "title": "Starbucks Corporation",
                        "aliases": [],
                        "is_following": false,
                        "watchlist_count": 49200
                    },
                    {
                        "id": 5281,
                        "symbol": "F",
                        "title": "Ford Motor Co.",
                        "aliases": [],
                        "is_following": false,
                        "watchlist_count": 58268
                    },
                    {
                        "id": 11418,
                        "symbol": "BTC.X",
                        "title": "Bitcoin BTC/USD",
                        "aliases": [
                            "BTCUSD"
                        ],
                        "is_following": false,
                        "watchlist_count": 123630
                    },
                    {
                        "id": 13348,
                        "symbol": "NBEV",
                        "title": "New Age Beverages Corporation",
                        "aliases": [],
                        "is_following": false,
                        "watchlist_count": 33821
                    },
                    {
                        "id": 15058,
                        "symbol": "RBZ",
                        "title": "Reebonz Holding Limited Ordinary ",
                        "aliases": [],
                        "is_following": false,
                        "watchlist_count": 6528
                    }
                ],
                "conversation": {
                    "parent_message_id": 161969294,
                    "in_reply_to_message_id": null,
                    "parent": true,
                    "replies": 1
                },
                "mentioned_users": [],
                "entities": {
                    "sentiment": null
                }
            },
            {
                "id": 161969138,
                "body": "Potential #earnings trades (Whisper Reactors) for Fri am include $AMZN $EMN $F $SBUX $CVX $XOM, recent alerts here  http://www.whispernumber.com",
                "created_at": "2019-04-26T02:16:56Z",
                "user": {
                    "id": 11587,
                    "username": "WhisperNumber",
                    "name": "WhisperNumber.com",
                    "avatar_url": "https://avatars.stocktwits.com/production/11587/thumb-1462396294.png",
                    "avatar_url_ssl": "https://avatars.stocktwits.com/production/11587/thumb-1462396294.png",
                    "join_date": "2010-03-29",
                    "official": true,
                    "identity": "Official",
                    "classification": [
                        "official"
                    ],
                    "followers": 89896,
                    "following": 126,
                    "ideas": 4036,
                    "watchlist_stocks_count": 1,
                    "like_count": 66
                },
                "source": {
                    "id": 2269,
                    "title": "StockTwits Web",
                    "url": "https://stocktwits.com"
                },
                "symbols": [
                    {
                        "id": 864,
                        "symbol": "AMZN",
                        "title": "Amazon.com Inc.",
                        "aliases": [],
                        "is_following": false,
                        "watchlist_count": 203536
                    },
                    {
                        "id": 3465,
                        "symbol": "SBUX",
                        "title": "Starbucks Corporation",
                        "aliases": [],
                        "is_following": false,
                        "watchlist_count": 49200
                    },
                    {
                        "id": 4874,
                        "symbol": "CVX",
                        "title": "Chevron Corp.",
                        "aliases": [],
                        "is_following": false,
                        "watchlist_count": 11174
                    },
                    {
                        "id": 5169,
                        "symbol": "EMN",
                        "title": "Eastman Chemical Co.",
                        "aliases": [],
                        "is_following": false,
                        "watchlist_count": 478
                    },
                    {
                        "id": 5281,
                        "symbol": "F",
                        "title": "Ford Motor Co.",
                        "aliases": [],
                        "is_following": false,
                        "watchlist_count": 58268
                    },
                    {
                        "id": 7825,
                        "symbol": "XOM",
                        "title": "Exxon Mobil Corporation",
                        "aliases": [],
                        "is_following": false,
                        "watchlist_count": 29551
                    }
                ],
                "links": [
                    {
                        "title": "WhisperNumber.com",
                        "url": "http://www.whispernumber.com",
                        "shortened_url": "http://www.whispernumber.com",
                        "shortened_expanded_url": "whispernumber.com/",
                        "description": "The official site of Real Earnings Expectations, Whisper Numbers, Earnings Estimates, Earnings News, and Individual Investor Market Sentiment",
                        "image": "http://www.whispernumber.com/images/confidence2.jpg",
                        "created_at": "2019-04-26T02:16:57Z",
                        "video_url": null,
                        "source": {
                            "name": "Whispernumber",
                            "website": "http://www.whispernumber.com"
                        }
                    }
                ],
                "mentioned_users": [],
                "entities": {
                    "chart": {
                        "thumb": "https://charts.stocktwits.com/production/small_161969138.jpg",
                        "large": "https://charts.stocktwits.com/production/large_161969138.jpg",
                        "original": "https://charts.stocktwits.com/production/original_161969138.jpg",
                        "url": "https://charts.stocktwits.com/production/original_161969138.jpg"
                    },
                    "sentiment": null
                }
            },
            {
                "id": 161969124,
                "body": "$F my fave part about the Rivian deal is that JH wont have to beg that arrogant-ass VW guy for their EV platform",
                "created_at": "2019-04-26T02:16:42Z",
                "user": {
                    "id": 913225,
                    "username": "noobie_trader",
                    "name": "Jose",
                    "avatar_url": "https://avatars.stocktwits.com/production/913225/thumb-1487468635.png",
                    "avatar_url_ssl": "https://avatars.stocktwits.com/production/913225/thumb-1487468635.png",
                    "join_date": "2016-12-30",
                    "official": false,
                    "identity": "User",
                    "classification": [],
                    "followers": 5,
                    "following": 1,
                    "ideas": 598,
                    "watchlist_stocks_count": 16,
                    "like_count": 1349
                },
                "source": {
                    "id": 2269,
                    "title": "StockTwits Web",
                    "url": "https://stocktwits.com"
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
                "likes": {
                    "total": 2,
                    "user_ids": [
                        680800,
                        2035539
                    ]
                },
                "mentioned_users": [],
                "entities": {
                    "sentiment": {
                        "basic": "Bullish"
                    }
                }
            },
            {
                "id": 161968988,
                "body": "$F I’m very curious to see how this performs tomorrow",
                "created_at": "2019-04-26T02:15:05Z",
                "user": {
                    "id": 1382571,
                    "username": "bombbay",
                    "name": "Dylan",
                    "avatar_url": "https://avatars.stocktwits.com/production/1382571/thumb-1520423976.png",
                    "avatar_url_ssl": "https://avatars.stocktwits.com/production/1382571/thumb-1520423976.png",
                    "join_date": "2018-01-24",
                    "official": false,
                    "identity": "User",
                    "classification": [],
                    "followers": 0,
                    "following": 2,
                    "ideas": 129,
                    "watchlist_stocks_count": 71,
                    "like_count": 90
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
                "likes": {
                    "total": 3,
                    "user_ids": [
                        948949,
                        1559636,
                        2035539
                    ]
                },
                "mentioned_users": [],
                "entities": {
                    "sentiment": null
                }
            },
            {
                "id": 161968564,
                "body": "$F $TSLA F made more $ in Q1 w/ their financing arm, Ford Credit, than TSLA has EVER made in a qtr. F Credit alone can pay the div!",
                "created_at": "2019-04-26T02:09:02Z",
                "user": {
                    "id": 1414010,
                    "username": "jtcdub",
                    "name": "jtcdub",
                    "avatar_url": "https://avatars.stocktwits.com/production/1414010/thumb-1529587262.png",
                    "avatar_url_ssl": "https://avatars.stocktwits.com/production/1414010/thumb-1529587262.png",
                    "join_date": "2018-02-16",
                    "official": false,
                    "identity": "User",
                    "classification": [],
                    "followers": 14,
                    "following": 0,
                    "ideas": 4210,
                    "watchlist_stocks_count": 0,
                    "like_count": 2706
                },
                "source": {
                    "id": 2269,
                    "title": "StockTwits Web",
                    "url": "https://stocktwits.com"
                },
                "symbols": [
                    {
                        "id": 5281,
                        "symbol": "F",
                        "title": "Ford Motor Co.",
                        "aliases": [],
                        "is_following": false,
                        "watchlist_count": 58268
                    },
                    {
                        "id": 8660,
                        "symbol": "TSLA",
                        "title": "Tesla, Inc.",
                        "aliases": [],
                        "is_following": false,
                        "watchlist_count": 210989
                    }
                ],
                "conversation": {
                    "parent_message_id": 161968564,
                    "in_reply_to_message_id": null,
                    "parent": true,
                    "replies": 9
                },
                "likes": {
                    "total": 3,
                    "user_ids": [
                        1387046,
                        1651012,
                        948949
                    ]
                },
                "mentioned_users": [],
                "entities": {
                    "sentiment": null
                }
            },
            {
                "id": 161968162,
                "body": "$F Ford over $10=sign of the Apocalypse?",
                "created_at": "2019-04-26T02:03:03Z",
                "user": {
                    "id": 782057,
                    "username": "sndmn",
                    "name": "Sandman",
                    "avatar_url": "https://avatars.stocktwits.com/production/782057/thumb-1467319481.png",
                    "avatar_url_ssl": "https://avatars.stocktwits.com/production/782057/thumb-1467319481.png",
                    "join_date": "2016-06-26",
                    "official": false,
                    "identity": "User",
                    "classification": [],
                    "followers": 6,
                    "following": 74,
                    "ideas": 849,
                    "watchlist_stocks_count": 23,
                    "like_count": 1443
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
                "conversation": {
                    "parent_message_id": 161968162,
                    "in_reply_to_message_id": null,
                    "parent": true,
                    "replies": 1
                },
                "likes": {
                    "total": 1,
                    "user_ids": [
                        680800
                    ]
                },
                "mentioned_users": [],
                "entities": {
                    "sentiment": null
                }
            },
            {
                "id": 161967150,
                "body": "$F bought at 10.59 last July.  It’s good to see the 10’s in AH trading",
                "created_at": "2019-04-26T01:48:50Z",
                "user": {
                    "id": 1407029,
                    "username": "NY61789",
                    "name": "Neil Young",
                    "avatar_url": "http://avatars.stocktwits.com/images/default_avatar_thumb.jpg",
                    "avatar_url_ssl": "https://s3.amazonaws.com/st-avatars/images/default_avatar_thumb.jpg",
                    "join_date": "2018-02-10",
                    "official": false,
                    "identity": "User",
                    "classification": [],
                    "followers": 3,
                    "following": 0,
                    "ideas": 582,
                    "watchlist_stocks_count": 9,
                    "like_count": 352
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
                        "watchlist_count": 58250
                    }
                ],
                "likes": {
                    "total": 1,
                    "user_ids": [
                        1350495
                    ]
                },
                "mentioned_users": [],
                "entities": {
                    "sentiment": {
                        "basic": "Bullish"
                    }
                }
            },
            {
                "id": 161965058,
                "body": "$NBEV doesn&#39;t hit 10 bucks before $F ? That&#39;s how you know it&#39;s trash. LOL",
                "created_at": "2019-04-26T01:22:46Z",
                "user": {
                    "id": 1814602,
                    "username": "cashtradez",
                    "name": "Brandon Ca$h",
                    "avatar_url": "https://avatars.stocktwits.com/production/1814602/thumb-1555053002.png",
                    "avatar_url_ssl": "https://avatars.stocktwits.com/production/1814602/thumb-1555053002.png",
                    "join_date": "2019-01-18",
                    "official": false,
                    "identity": "User",
                    "classification": [],
                    "followers": 25,
                    "following": 63,
                    "ideas": 649,
                    "watchlist_stocks_count": 231,
                    "like_count": 1796
                },
                "source": {
                    "id": 2269,
                    "title": "StockTwits Web",
                    "url": "https://stocktwits.com"
                },
                "symbols": [
                    {
                        "id": 5281,
                        "symbol": "F",
                        "title": "Ford Motor Co.",
                        "aliases": [],
                        "is_following": false,
                        "watchlist_count": 58250
                    },
                    {
                        "id": 13348,
                        "symbol": "NBEV",
                        "title": "New Age Beverages Corporation",
                        "aliases": [],
                        "is_following": false,
                        "watchlist_count": 33812
                    }
                ],
                "conversation": {
                    "parent_message_id": 161965058,
                    "in_reply_to_message_id": null,
                    "parent": true,
                    "replies": 3
                },
                "likes": {
                    "total": 2,
                    "user_ids": [
                        1972198,
                        1143447
                    ]
                },
                "mentioned_users": [],
                "entities": {
                    "sentiment": {
                        "basic": "Bearish"
                    }
                }
            },
            {
                "id": 161964001,
                "body": "Today, some companies beat and had good earnings $AMZN $CMCSA $F $GRUB  , upside potential ! https://bullishdetective.com/earnings",
                "created_at": "2019-04-26T01:09:20Z",
                "user": {
                    "id": 2022444,
                    "username": "powereborn",
                    "name": "Nicolas Dei",
                    "avatar_url": "http://avatars.stocktwits.com/images/default_avatar_thumb.jpg",
                    "avatar_url_ssl": "https://s3.amazonaws.com/st-avatars/images/default_avatar_thumb.jpg",
                    "join_date": "2019-04-20",
                    "official": false,
                    "identity": "User",
                    "classification": [],
                    "followers": 3,
                    "following": 0,
                    "ideas": 23,
                    "watchlist_stocks_count": 0,
                    "like_count": 0
                },
                "source": {
                    "id": 2269,
                    "title": "StockTwits Web",
                    "url": "https://stocktwits.com"
                },
                "symbols": [
                    {
                        "id": 864,
                        "symbol": "AMZN",
                        "title": "Amazon.com Inc.",
                        "aliases": [],
                        "is_following": false,
                        "watchlist_count": 203508
                    },
                    {
                        "id": 1348,
                        "symbol": "CMCSA",
                        "title": "Comcast Corporation",
                        "aliases": [],
                        "is_following": false,
                        "watchlist_count": 6767
                    },
                    {
                        "id": 5281,
                        "symbol": "F",
                        "title": "Ford Motor Co.",
                        "aliases": [],
                        "is_following": false,
                        "watchlist_count": 58250
                    },
                    {
                        "id": 11844,
                        "symbol": "GRUB",
                        "title": "GrubHub",
                        "aliases": [],
                        "is_following": false,
                        "watchlist_count": 6723
                    }
                ],
                "links": [
                    {
                        "title": "Bullish Detective - Upcoming earnings",
                        "url": "https://bullishdetective.com/earnings",
                        "shortened_url": "https://bullishdetective.com/earnings",
                        "shortened_expanded_url": "bullishdetective.com/earnings",
                        "description": "Go through past and future earnings with our calendar. We filter results to only give you stocks that are talked about on the web.",
                        "image": "https://resources.bullishdetective.com/images/logos/nasdaq_sbux.png",
                        "created_at": "2019-04-26T01:09:22Z",
                        "video_url": null,
                        "source": {
                            "name": "Bullishdetective",
                            "website": "https://bullishdetective.com"
                        }
                    }
                ],
                "mentioned_users": [],
                "entities": {
                    "sentiment": null
                }
            },
            {
                "id": 161963761,
                "body": "Ford Motor announces earnings. $0.44 EPS. Beats estimates. $40.34b revenue.  https://www.marketbeat.com/s/405761 $F",
                "created_at": "2019-04-26T01:06:48Z",
                "user": {
                    "id": 284225,
                    "username": "AnalystRatingsNetwork",
                    "name": "MarketBeat.com",
                    "avatar_url": "https://avatars.stocktwits.com/production/284225/thumb-1484325936.png",
                    "avatar_url_ssl": "https://avatars.stocktwits.com/production/284225/thumb-1484325936.png",
                    "join_date": "2013-11-24",
                    "official": false,
                    "identity": "User",
                    "classification": [],
                    "followers": 6944,
                    "following": 9,
                    "ideas": 660181,
                    "watchlist_stocks_count": 0,
                    "like_count": 9
                },
                "source": {
                    "id": 1022,
                    "title": "Zapier",
                    "url": "https://zapier.com/"
                },
                "symbols": [
                    {
                        "id": 5281,
                        "symbol": "F",
                        "title": "Ford Motor Co.",
                        "aliases": [],
                        "is_following": false,
                        "watchlist_count": 58250
                    }
                ],
                "links": [
                    {
                        "title": "NYSE:F - Stock Price, News, & Analysis for Ford Motor",
                        "url": "https://www.marketbeat.com/s/405761",
                        "shortened_url": "https://www.marketbeat.com/s/405761",
                        "shortened_expanded_url": "marketbeat.com/stocks/NYSE/F/",
                        "description": "Researching Ford Motor (NYSE:F) stock? View F's stock price, price target, dividend, earnings, financials, insider trades, news and SEC filings at MarketBeat.",
                        "image": "https://www.marketbeat.com/logos/fordlogo2003.jpg",
                        "created_at": "2019-04-26T01:06:50Z",
                        "video_url": null,
                        "source": {
                            "name": "Marketbeat",
                            "website": "https://www.marketbeat.com"
                        }
                    }
                ],
                "likes": {
                    "total": 3,
                    "user_ids": [
                        1143447,
                        893205,
                        1950381
                    ]
                },
                "mentioned_users": [],
                "entities": {
                    "sentiment": null
                }
            },
            {
                "id": 161963755,
                "body": "Ford Motor announces earnings. $0.44 EPS. Beats estimates. $40.34b revenue.  https://www.marketbeat.com/s/405761 $F",
                "created_at": "2019-04-26T01:06:47Z",
                "user": {
                    "id": 284541,
                    "username": "EarningsInsider",
                    "name": "MarketBeat Earnings",
                    "avatar_url": "https://avatars.stocktwits.com/production/284541/thumb-1484326078.png",
                    "avatar_url_ssl": "https://avatars.stocktwits.com/production/284541/thumb-1484326078.png",
                    "join_date": "2013-11-25",
                    "official": false,
                    "identity": "User",
                    "classification": [],
                    "followers": 51251,
                    "following": 0,
                    "ideas": 131087,
                    "watchlist_stocks_count": 0,
                    "like_count": 0
                },
                "source": {
                    "id": 1022,
                    "title": "Zapier",
                    "url": "https://zapier.com/"
                },
                "symbols": [
                    {
                        "id": 5281,
                        "symbol": "F",
                        "title": "Ford Motor Co.",
                        "aliases": [],
                        "is_following": false,
                        "watchlist_count": 58235
                    }
                ],
                "links": [
                    {
                        "title": "NYSE:F - Stock Price, News, & Analysis for Ford Motor",
                        "url": "https://www.marketbeat.com/s/405761",
                        "shortened_url": "https://www.marketbeat.com/s/405761",
                        "shortened_expanded_url": "marketbeat.com/stocks/NYSE/F/",
                        "description": "Researching Ford Motor (NYSE:F) stock? View F's stock price, price target, dividend, earnings, financials, insider trades, news and SEC filings at MarketBeat.",
                        "image": "https://www.marketbeat.com/logos/fordlogo2003.jpg",
                        "created_at": "2019-04-26T01:06:48Z",
                        "video_url": null,
                        "source": {
                            "name": "Marketbeat",
                            "website": "https://www.marketbeat.com"
                        }
                    }
                ],
                "mentioned_users": [],
                "entities": {
                    "sentiment": null
                }
            },
            {
                "id": 161963319,
                "body": "$F Let&#39;s see if these follow through and  work out alright.",
                "created_at": "2019-04-26T01:02:50Z",
                "user": {
                    "id": 1181112,
                    "username": "RyanJett",
                    "name": "Jett",
                    "avatar_url": "https://avatars.stocktwits.com/production/1181112/thumb-1505743988.png",
                    "avatar_url_ssl": "https://avatars.stocktwits.com/production/1181112/thumb-1505743988.png",
                    "join_date": "2017-09-07",
                    "official": false,
                    "identity": "User",
                    "classification": [],
                    "followers": 6,
                    "following": 8,
                    "ideas": 1111,
                    "watchlist_stocks_count": 19,
                    "like_count": 3
                },
                "source": {
                    "id": 2095,
                    "title": "StockTwits For Android ",
                    "url": "http://www.stocktwits.com/mobile"
                },
                "symbols": [
                    {
                        "id": 5281,
                        "symbol": "F",
                        "title": "Ford Motor Co.",
                        "aliases": [],
                        "is_following": false,
                        "watchlist_count": 58250
                    }
                ],
                "conversation": {
                    "parent_message_id": 161963319,
                    "in_reply_to_message_id": null,
                    "parent": true,
                    "replies": 2
                },
                "likes": {
                    "total": 1,
                    "user_ids": [
                        1540590
                    ]
                },
                "mentioned_users": [],
                "entities": {
                    "chart": {
                        "thumb": "https://charts.stocktwits.com/production/small_161963319.jpg",
                        "large": "https://charts.stocktwits.com/production/large_161963319.jpg",
                        "original": "https://charts.stocktwits.com/production/original_161963319.jpg",
                        "url": "https://charts.stocktwits.com/production/original_161963319.jpg"
                    },
                    "sentiment": null
                }
            },
            {
                "id": 161963133,
                "body": "Find the latest earning report of $F analyzed by machine learning ,  Ford was a beautiful beat ! https://bullishdetective.com/stocks/NYSE/F",
                "created_at": "2019-04-26T01:00:48Z",
                "user": {
                    "id": 2022444,
                    "username": "powereborn",
                    "name": "Nicolas Dei",
                    "avatar_url": "http://avatars.stocktwits.com/images/default_avatar_thumb.jpg",
                    "avatar_url_ssl": "https://s3.amazonaws.com/st-avatars/images/default_avatar_thumb.jpg",
                    "join_date": "2019-04-20",
                    "official": false,
                    "identity": "User",
                    "classification": [],
                    "followers": 3,
                    "following": 0,
                    "ideas": 23,
                    "watchlist_stocks_count": 0,
                    "like_count": 0
                },
                "source": {
                    "id": 2269,
                    "title": "StockTwits Web",
                    "url": "https://stocktwits.com"
                },
                "symbols": [
                    {
                        "id": 5281,
                        "symbol": "F",
                        "title": "Ford Motor Co.",
                        "aliases": [],
                        "is_following": false,
                        "watchlist_count": 58235
                    }
                ],
                "links": [
                    {
                        "title": "Bullish Detective - Ford Motor Company",
                        "url": "https://bullishdetective.com/stocks/NYSE/F",
                        "shortened_url": "https://bullishdetective.com/stocks/NYSE/F",
                        "shortened_expanded_url": "bullishdetective.com/stocks...",
                        "description": "Ford Motor Company (NYSE:F). Ford Motor Company is an American multinational automaker that has its main headquarter in Dearborn, Michigan, a suburb of Detroit. It was founded by Henry Ford and incorporated on June 16, 1903. The company sells automobiles and commercial vehicles under the Ford brand and most luxury cars under the Lincoln brand.",
                        "image": "https://resources.bullishdetective.com/images/logos/nyse_f.png",
                        "created_at": "2019-04-26T01:00:49Z",
                        "video_url": null,
                        "source": {
                            "name": "Bullishdetective",
                            "website": "https://bullishdetective.com"
                        }
                    }
                ],
                "mentioned_users": [],
                "entities": {
                    "sentiment": null
                }
            },
            {
                "id": 161962937,
                "body": "$F love my $F",
                "created_at": "2019-04-26T00:58:17Z",
                "user": {
                    "id": 709160,
                    "username": "fernsl514",
                    "name": "Fernando Silva",
                    "avatar_url": "https://avatars.stocktwits.com/production/709160/thumb-1466610455.png",
                    "avatar_url_ssl": "https://avatars.stocktwits.com/production/709160/thumb-1466610455.png",
                    "join_date": "2016-03-17",
                    "official": false,
                    "identity": "User",
                    "classification": [],
                    "followers": 1,
                    "following": 63,
                    "ideas": 203,
                    "watchlist_stocks_count": 39,
                    "like_count": 165
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
                        "watchlist_count": 58250
                    }
                ],
                "conversation": {
                    "parent_message_id": 161962937,
                    "in_reply_to_message_id": null,
                    "parent": true,
                    "replies": 3
                },
                "likes": {
                    "total": 6,
                    "user_ids": [
                        948949,
                        1143447,
                        999704,
                        1194580,
                        1571704,
                        1370321
                    ]
                },
                "mentioned_users": [],
                "entities": {
                    "chart": {
                        "thumb": "https://charts.stocktwits.com/production/small_161962937.png",
                        "large": "https://charts.stocktwits.com/production/large_161962937.png",
                        "original": "https://charts.stocktwits.com/production/original_161962937.png",
                        "url": "https://charts.stocktwits.com/production/original_161962937.png"
                    },
                    "sentiment": null
                }
            },
            {
                "id": 161962577,
                "body": "$F is a sell tomorrow",
                "created_at": "2019-04-26T00:53:59Z",
                "user": {
                    "id": 356099,
                    "username": "ZigguratAI",
                    "name": "Ziggurat Technologies, Inc.",
                    "avatar_url": "https://avatars.stocktwits.com/production/356099/thumb-1552859256.png",
                    "avatar_url_ssl": "https://avatars.stocktwits.com/production/356099/thumb-1552859256.png",
                    "join_date": "2014-05-22",
                    "official": false,
                    "identity": "User",
                    "classification": [],
                    "followers": 114,
                    "following": 0,
                    "ideas": 386,
                    "watchlist_stocks_count": 5,
                    "like_count": 8
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
                        "watchlist_count": 58220
                    }
                ],
                "likes": {
                    "total": 2,
                    "user_ids": [
                        905371,
                        2001469
                    ]
                },
                "mentioned_users": [],
                "entities": {
                    "chart": {
                        "thumb": "https://charts.stocktwits.com/production/small_161962577.png",
                        "large": "https://charts.stocktwits.com/production/large_161962577.png",
                        "original": "https://charts.stocktwits.com/production/original_161962577.png",
                        "url": "https://charts.stocktwits.com/production/original_161962577.png"
                    },
                    "sentiment": {
                        "basic": "Bearish"
                    }
                }
            },
            {
                "id": 161961826,
                "body": "$F",
                "created_at": "2019-04-26T00:44:09Z",
                "user": {
                    "id": 948949,
                    "username": "Dalluge",
                    "name": "Aaron Dalluge",
                    "avatar_url": "https://avatars.stocktwits.com/production/948949/thumb-1551199285.png",
                    "avatar_url_ssl": "https://avatars.stocktwits.com/production/948949/thumb-1551199285.png",
                    "join_date": "2017-02-08",
                    "official": false,
                    "identity": "User",
                    "classification": [],
                    "followers": 62,
                    "following": 47,
                    "ideas": 4234,
                    "watchlist_stocks_count": 9,
                    "like_count": 25036
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
                        "watchlist_count": 58220
                    }
                ],
                "mentioned_users": [],
                "entities": {
                    "sentiment": {
                        "basic": "Bullish"
                    }
                }
            },
            {
                "id": 161961539,
                "body": "$F reported 2.33% YoY EPS growth for Q1, up from -23.08% in Q4 and higher than the Estimize community expected.\nhttp://www.estimize.com/intro/f?chart=historical&amp;metric_name=eps&amp;utm_content=F&amp;utm_medium=eps_growth&amp;utm_source=stocktwits",
                "created_at": "2019-04-26T00:40:43Z",
                "user": {
                    "id": 727510,
                    "username": "EstimizeAlerts",
                    "name": "Estimize",
                    "avatar_url": "https://avatars.stocktwits.com/production/727510/thumb-1461073766.png",
                    "avatar_url_ssl": "https://avatars.stocktwits.com/production/727510/thumb-1461073766.png",
                    "join_date": "2016-04-11",
                    "official": false,
                    "identity": "User",
                    "classification": [],
                    "followers": 6141,
                    "following": 56,
                    "ideas": 1623911,
                    "watchlist_stocks_count": 0,
                    "like_count": 0
                },
                "source": {
                    "id": 1021,
                    "title": "Estimize",
                    "url": "http://estimize.com"
                },
                "symbols": [
                    {
                        "id": 5281,
                        "symbol": "F",
                        "title": "Ford Motor Co.",
                        "aliases": [],
                        "is_following": false,
                        "watchlist_count": 58220
                    }
                ],
                "links": [
                    {
                        "title": "F - Ford Motor Co. Crowdsourced Earnings Estimates - Estimize",
                        "url": "http://www.estimize.com/intro/f?chart=historical&metric_name=eps&utm_content=F&utm_medium=eps_growth&utm_source=stocktwits",
                        "shortened_url": "http://www.estimize.com/intro/f?chart=historical&metric_name=eps&utm_content=F&utm_medium=eps_growth&utm_source=stocktwits",
                        "shortened_expanded_url": "estimize.com/intro/f?chart=...",
                        "description": "Earnings estimates for F from thousands of professional and amateur analysts who contribute to a open, crowdsourced estimate data set.",
                        "image": "https://www.estimize.com/assets/home_access_contributors/graph-efb02d56d858b4923bf00ffcef83249a.png",
                        "created_at": "2019-04-26T00:40:44Z",
                        "video_url": null,
                        "source": {
                            "name": "Estimize",
                            "website": "https://www.estimize.com"
                        }
                    }
                ],
                "mentioned_users": [],
                "entities": {
                    "chart": {
                        "thumb": "https://charts.stocktwits.com/production/small_161961539.png",
                        "large": "https://charts.stocktwits.com/production/large_161961539.png",
                        "original": "https://charts.stocktwits.com/production/original_161961539.png",
                        "url": "https://charts.stocktwits.com/production/original_161961539.png"
                    },
                    "sentiment": null
                }
            },
            {
                "id": 161961396,
                "body": "$F IMO $TSLA Folk Should bring their Money to $F Sounds like $F has a better head on its shoulders Than what Elon Muskateer has on his.",
                "created_at": "2019-04-26T00:39:19Z",
                "user": {
                    "id": 1454594,
                    "username": "Wranglerx",
                    "name": "Ricky Greene",
                    "avatar_url": "http://avatars.stocktwits.com/images/default_avatar_thumb.jpg",
                    "avatar_url_ssl": "https://s3.amazonaws.com/st-avatars/images/default_avatar_thumb.jpg",
                    "join_date": "2018-03-24",
                    "official": false,
                    "identity": "User",
                    "classification": [],
                    "followers": 1,
                    "following": 5,
                    "ideas": 44,
                    "watchlist_stocks_count": 44,
                    "like_count": 27
                },
                "source": {
                    "id": 2095,
                    "title": "StockTwits For Android ",
                    "url": "http://www.stocktwits.com/mobile"
                },
                "symbols": [
                    {
                        "id": 5281,
                        "symbol": "F",
                        "title": "Ford Motor Co.",
                        "aliases": [],
                        "is_following": false,
                        "watchlist_count": 58250
                    },
                    {
                        "id": 8660,
                        "symbol": "TSLA",
                        "title": "Tesla, Inc.",
                        "aliases": [],
                        "is_following": false,
                        "watchlist_count": 210979
                    }
                ],
                "likes": {
                    "total": 4,
                    "user_ids": [
                        363511,
                        1875796,
                        1143447,
                        1507193
                    ]
                },
                "mentioned_users": [],
                "entities": {
                    "sentiment": {
                        "basic": "Bullish"
                    }
                }
            },
            {
                "id": 161961386,
                "body": "$F quit with the ugly cars and boom!!!",
                "created_at": "2019-04-26T00:39:12Z",
                "user": {
                    "id": 780321,
                    "username": "Jason1974",
                    "name": "jason",
                    "avatar_url": "https://avatars.stocktwits.com/production/780321/thumb-1492535492.png",
                    "avatar_url_ssl": "https://avatars.stocktwits.com/production/780321/thumb-1492535492.png",
                    "join_date": "2016-06-23",
                    "official": false,
                    "identity": "User",
                    "classification": [],
                    "followers": 115,
                    "following": 74,
                    "ideas": 8256,
                    "watchlist_stocks_count": 21,
                    "like_count": 6825
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
                        "watchlist_count": 58220
                    }
                ],
                "likes": {
                    "total": 1,
                    "user_ids": [
                        1143447
                    ]
                },
                "mentioned_users": [],
                "entities": {
                    "sentiment": {
                        "basic": "Bullish"
                    }
                }
            },
            {
                "id": 161961290,
                "body": "$F 4K shares at 9.84 been holding for what seems like forever!!! Let’s roll",
                "created_at": "2019-04-26T00:38:08Z",
                "user": {
                    "id": 780321,
                    "username": "Jason1974",
                    "name": "jason",
                    "avatar_url": "https://avatars.stocktwits.com/production/780321/thumb-1492535492.png",
                    "avatar_url_ssl": "https://avatars.stocktwits.com/production/780321/thumb-1492535492.png",
                    "join_date": "2016-06-23",
                    "official": false,
                    "identity": "User",
                    "classification": [],
                    "followers": 115,
                    "following": 74,
                    "ideas": 8256,
                    "watchlist_stocks_count": 21,
                    "like_count": 6825
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
                        "watchlist_count": 58235
                    }
                ],
                "likes": {
                    "total": 2,
                    "user_ids": [
                        1143447,
                        999056
                    ]
                },
                "mentioned_users": [],
                "entities": {
                    "sentiment": {
                        "basic": "Bullish"
                    }
                }
            },
            {
                "id": 161961124,
                "body": "$F Only fuck with the best .",
                "created_at": "2019-04-26T00:36:27Z",
                "user": {
                    "id": 347686,
                    "username": "JDubbbb",
                    "name": "Jon",
                    "avatar_url": "https://avatars.stocktwits.com/production/347686/thumb-1498190222.png",
                    "avatar_url_ssl": "https://avatars.stocktwits.com/production/347686/thumb-1498190222.png",
                    "join_date": "2014-04-28",
                    "official": false,
                    "identity": "User",
                    "classification": [],
                    "followers": 78,
                    "following": 44,
                    "ideas": 11299,
                    "watchlist_stocks_count": 56,
                    "like_count": 3229
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
                "conversation": {
                    "parent_message_id": 161961124,
                    "in_reply_to_message_id": null,
                    "parent": true,
                    "replies": 3
                },
                "likes": {
                    "total": 5,
                    "user_ids": [
                        893205,
                        1806574,
                        1143447,
                        999704,
                        1114590
                    ]
                },
                "mentioned_users": [],
                "entities": {
                    "chart": {
                        "thumb": "https://charts.stocktwits.com/production/small_161961124.png",
                        "large": "https://charts.stocktwits.com/production/large_161961124.png",
                        "original": "https://charts.stocktwits.com/production/original_161961124.png",
                        "url": "https://charts.stocktwits.com/production/original_161961124.png"
                    },
                    "sentiment": {
                        "basic": "Bullish"
                    }
                }
            }
        ]
    }
  `)
	var o StreamSymbol
	var err error
	if err = json.Unmarshal(x, &o); err != nil {
		t.Error(err)
	}

	if o.GetMaxId() != 161970974 {
		t.Error("error in GetMinId")
	}

	if o.GetMinId() != 161961124 {
		t.Error("error in GetMinId")
	}
}
