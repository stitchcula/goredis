//简单的使用配置文件的启动工具，配置文件默认为当前目录下的config.ini
//
// ### 示例配置
//     [ssdb]
//     #ssdb的主机IP
//     host=127.0.0.1
//     #ssdb的端口
//     port=8888
//     #连接池检查时间间隔
//     health_second=5
//     #连接密码，默认为空
//     password=
//     #最大等待数目，当连接池满后，新建连接将等待池中连接释放后才可以继续，本值限制最大等待的数量，超过本值后将抛出异常。默认值: 1000
//     max_wait_size=1000
//     #当连接池中的连接耗尽的时候一次同时获取的连接数。默认值: 5
//     acquire_increment=5
//     #最小连接池数。默认值: 5
//     min_pool_size=5
//     #最大连接池个数。默认值: 20
//     max_pool_size=20
//     #获取连接超时时间，单位为秒。默认值: 5
//     get_client_timeout=5
//
// ### 示例
//     if err := ssdb.Start(); err != nil {
//     println("无法连接到ssdb")
//     os.Exit(1)
//     }
//     defer ssdb.Close()
//     //获取client
//     client, err := ssdb.Client()
//     if err != nil {
//     println("无法获取连接")
//     os.Exit(1)
//     }
//     defer client.Close()
//     client.Set("a", 1)
//     client.Get("a")
package ssdb
