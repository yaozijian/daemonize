
简单封装了[daemonigo](https://github.com/tyranron/daemonigo)和[seelog](https://github.com/cihub/seelog)，方便日志处理和将程序启动为守护进程。

### 使用方法

* 在适当地方调用daemonize.Daemonize()，将程序启动为守护进程。一般在main()函数开始的时候调用daemonize.Daemonize()。
* 注意：如果main()中有启动TCP监听的代码，请将daemonize.Daemonize()放在相关代码前面，否则会使得子进程（后台进程）再次启动监听失败。

* 启动程序的时候，用参数start表示启动为守护进程；用参数stop表示停止守护进程；用参数fg表示启动为前台进程。
* 启动为守护进程时日志默认输出到<可执行文件名>.log中；前台运行时日志默认输出到控制台。
* 注意：仅在Linux中支持start/stop/fg这三个参数，在Windows中程序总是前台运行。

* 可以在可执行程序目录下放置<可执行文件名>.file.xml文件，用start参数启动程序时daemonize.Daemonize()会读取并使用其中的seelog配置。
* 可以在可执行程序目录下放置<可执行文件名>.console.xml文件，用fg参数启动程序，或者在Windows中运行程序时daemonize.Daemonize()会读取并使用其中的seelog配置。
* 上述两个配置文件是可选的，如果没有这两个文件，daemonize.Daemonize()会使用默认配置。
* seelog配置文件的写法请参考[seelog文档](https://github.com/cihub/seelog/wiki)。

