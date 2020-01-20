## 简介
这个项目为go练手项目。功能简单，里面有些功能是为了做而做，实际上并没有太大意义。
这里分为client端跟server端。client端处理server启停动作，日志级别打印。
client其实就是为了做而做的一个地方。去掉client，直接使用命令行启动server端即可。
server 端每隔一段时间读取系统的资源信息，后发给指定的接收端，使用http请求，
当然也可以udp，换种方式实现。

### 使用的一些技术或者依赖包
- 类库管理使用go新增的mod 管理功能。用惯maven 这个mod挺难用的，java的生态还是很厉害的
- 使用[shirou/gopsutil](github.com/shirou/gopsutil) 类库，读取系统的信息。使用简单
- 使用[sirupsen](github.com/sirupsen/logrus)类库，docker 使用的日志模块，简单使用。
其中hook机制很好用，项目也简单写了一个

## 主要结构为
- clinet端
   + 命令行启停server功能
- server端
   + banner... 就是为了好看，没什么用...
   + loggerManager 日志处理类，里面有hook实现（每天一个新日志文件）。还有一个是日志文件管理线程超过15天删掉日志，坑肯定有的。
   + setting 读取go-monitor-config.json配置文件，给系统监听交互使用。
   + monitorManager 系统监听并将数据发给接收方
   + build.sh

## 开发过程中遇到的问题
### 第三方包管理
### 系统异常管理



##后续会增加的功能
 - 监听docker信息
 - client端命令更加丰富，比如添加查看日志功能（鸡肋，其实tial 就行了）
 - 其他...
