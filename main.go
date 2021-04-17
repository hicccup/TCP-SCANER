package main



import (
	"fmt"
	"net"
	"sync"
	"time"
)
//并发TCP-SCAN
func main()  {
	//记录当前时间
	time1 := time.Now()
	//创建一个计数器
	var wg sync.WaitGroup
	//创建端口的for循环
	for i := 1; i < 65536; i++ {
		//计数器+1
		wg.Add(1)
		//创建匿名函数并直接运行go func (){}()
		go func(j int) {
			//计数器-1
			defer wg.Done()
			//新建一个ip地址
			ip := fmt.Sprintf("x.x.x.x:%d",j)
			//conn连接 err错误  Dial(传输类型，地址	)
			conn, err := net.Dial("tcp",ip)
			//判断一下是否有错误发生 如果！=nil 说明关闭或者防火墙阻断
			if err != nil {
				//输出信息
				fmt.Printf("%v close\n",ip)
				return
			}
			//关闭连接
			conn.Close()
			//打印信息
			fmt.Printf("%v open\n",ip)
		}(i)
	}
	//等待 直到计数器=0
	wg.Wait()
	//算时间差
	time2 := time.Since(time1) / 1e9
	println(time2,"s")
}