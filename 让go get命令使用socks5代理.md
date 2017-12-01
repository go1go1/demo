
不知道是因为什么鬼打墙的原因，GFW把go的官网给禁了，也许因为Go出身Google吧，真是*了狗了。不仅是 golang.org 被墙了，现在Go的包管理站 http://gopkg.in 也被墙了。直接用 go get gopkg.in/package 半天都没反应，只好用代理了。

本地使用 shadowsocks做代理，因为ss是socks5代理，所以不能用上面的方法，go get又没有直接设置代理的地方，所以只好祭出代理利器：Proxifier 。
打开Proxifier后添加 proxies，添加 127.0.0.1:ss端口 类型 socket5，这样重新打开终端（我的git bash），直接

go get -u -v ...
