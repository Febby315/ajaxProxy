# qq音乐接口

## 1. 各客户端下载地址

> https://y.qq.com/download/download.js?format=jsonp&jsonpCallback=MusicJsonCallback&loginUin=0&hostUin=0&format=jsonp&inCharset=utf8&outCharset=utf-8&notice=0&platform=yqq&needNewCode=0

## 热歌榜

> https://u.y.qq.com/cgi-bin/musicu.fcg?-=getUCGI6257840554017182&g_tk=5381&loginUin=0&hostUin=0&format=json&inCharset=utf8&outCharset=utf-8&notice=0&platform=yqq.json&needNewCode=0&data=%7B%22detail%22%3A%7B%22module%22%3A%22musicToplist.ToplistInfoServer%22%2C%22method%22%3A%22GetDetail%22%2C%22param%22%3A%7B%22topId%22%3A26%2C%22offset%22%3A0%2C%22num%22%3A20%2C%22period%22%3A%222019_34%22%7D%7D%2C%22comm%22%3A%7B%22ct%22%3A24%2C%22cv%22%3A0%7D%7D

```js
    var option = JSON.stringify({
        "detail":{
            "module":"musicToplist.ToplistInfoServer",
            "method":"GetDetail",
            "param":{
                "topId":26,"offset":0,"num":20,"period":"2019_34"
            }
        },
        "comm":{"ct":24,"cv":0}
    });
    var data = decodeURIComponent(option);
```

## 新歌榜

> https://u.y.qq.com/cgi-bin/musicu.fcg?-=getUCGI9511418813905428&g_tk=5381&loginUin=0&hostUin=0&format=json&inCharset=utf8&outCharset=utf-8&notice=0&platform=yqq.json&needNewCode=0&data=%7B%22detail%22%3A%7B%22module%22%3A%22musicToplist.ToplistInfoServer%22%2C%22method%22%3A%22GetDetail%22%2C%22param%22%3A%7B%22topId%22%3A27%2C%22offset%22%3A0%2C%22num%22%3A20%2C%22period%22%3A%222019-09-02%22%7D%7D%2C%22comm%22%3A%7B%22ct%22%3A24%2C%22cv%22%3A0%7D%7D

```js
    var option = JSON.stringify({
        "detail":{
            "module":"musicToplist.ToplistInfoServer",
            "method":"GetDetail",
            "param":{
                "topId":27,"offset":0,"num":20,"period":"2019-09-02"
            }
        },
        "comm":{"ct":24,"cv":0}
    });
    var data = decodeURIComponent(option);
```

## 歌词

> https://c.y.qq.com/lyric/fcgi-bin/fcg_query_lyric_new.fcg?-=MusicJsonCallback_lrc&pcachetime=1567416245009&songmid=0020VnHM0U9uNh&g_tk=5381&loginUin=564822672&hostUin=0&format=json&inCharset=utf8&outCharset=utf-8&notice=0&platform=yqq.json&needNewCode=0

## 作者信息

> https://c.y.qq.com/v8/fcg-bin/fcg_v8_album_info_cp.fcg?ct=24&albummid=001QvT9r1lhLUr&g_tk=5381&loginUin=564822672&hostUin=0&format=json&inCharset=utf8&outCharset=utf-8&notice=0&platform=yqq.json&needNewCode=0

## 在线播放地址获取

> https://u.y.qq.com/cgi-bin/musicu.fcg?-=getplaysongvkey6580556423248451&g_tk=5381&loginUin=564822672&hostUin=0&format=json&inCharset=utf8&outCharset=utf-8&notice=0&platform=yqq.json&needNewCode=0&data=%7B%22req%22%3A%7B%22module%22%3A%22CDN.SrfCdnDispatchServer%22%2C%22method%22%3A%22GetCdnDispatch%22%2C%22param%22%3A%7B%22guid%22%3A%227947292832%22%2C%22calltype%22%3A0%2C%22userip%22%3A%22%22%7D%7D%2C%22req_0%22%3A%7B%22module%22%3A%22vkey.GetVkeyServer%22%2C%22method%22%3A%22CgiGetVkey%22%2C%22param%22%3A%7B%22guid%22%3A%227947292832%22%2C%22songmid%22%3A%5B%220020VnHM0U9uNh%22%5D%2C%22songtype%22%3A%5B0%5D%2C%22uin%22%3A%22564822672%22%2C%22loginflag%22%3A1%2C%22platform%22%3A%2220%22%7D%7D%2C%22comm%22%3A%7B%22uin%22%3A%22564822672%22%2C%22format%22%3A%22json%22%2C%22ct%22%3A24%2C%22cv%22%3A0%7D%7D

```js
    var option = JSON.stringify({
        "req":{
            "module":"CDN.SrfCdnDispatchServer",
            "method":"GetCdnDispatch",
            "param":{"guid":"7947292832","calltype":0,"userip":""}
        },
        "req_0":{
            "module":"vkey.GetVkeyServer",
            "method":"CgiGetVkey",
            "param":{
                "guid":"7947292832",
                "songmid":["0020VnHM0U9uNh"],
                "songtype":[0],
                "uin":"564822672",
                "loginflag":1,
                "platform":"20"
            }
        },
        "comm":{"uin":"564822672","format":"json","ct":24,"cv":0}
    });
    var data = decodeURIComponent(option);
```

## web在线音频

> http://123.138.114.15/amobile.music.tc.qq.com/C4000020VnHM0U9uNh.m4a?guid=7947292832&vkey=76EA28BBB6C0EC721C5660A65ACE802921F43AD16048B81D1A0D69573CDE57A7914289087B2674AB1A6671BEA023CDF8CD361BFB36396444&uin=656&fromtag=66