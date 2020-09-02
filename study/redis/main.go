package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

/**
------------------------------
//sorted set应用场景

//统计用户登录排行榜

//zadd login_times uid score
zadd login_times 004 1
zadd login_times 002 2
zadd login_times 003 4

//当用户登录时，对该用户登录次数自增1
zincrby("login_times", 1, "001")
//获取登录次数最多的用户(排名前N的用户)
zrevrange("login_times", 0, N-1)

------------------------------
RDB持久化

* 当redis需要做持久化时，会fork一个子进程
* 子进程会将数据写到磁盘上一个RDB文件上
* 子进程写完之后，会把原来的RDB文件换掉
* 应用到的特性是copy-on-write机制(子进程备份时候，父进程新写会触发拷贝)

------------------------------
Redis主从同步

* 防止出现单点
* 缓解主节点的读压力，可水平扩展
* 数据冗余，保证数据库安全性

------------------------------
主从复制的流程
* slave和master建立TCP连接
* slave向master发起数据同步请求
* slave接收master发送过来的同步数据RDB
* slave基于数据构建内存数据库

*/

func main() {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Printf("conn redis failed, err: %v\n", err)
		return
	}

	defer c.Close()

	_, err = c.Do("Set", "aaa", 100)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	res, err := redis.String(c.Do("Get", "aaa"))
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Printf("res :%v\n", res)
}
