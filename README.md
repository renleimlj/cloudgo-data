# golang 构建数据服务

本项目的主体框架fork自潘老师的github，entiies2部分是我自己编写的，使用xorm来代替database/sql,使用了xorm就不需要写dao服务，非常方便，但使用反射技术，牺牲了性能。

## 安装mysql并创建数据库

![](https://s1.ax1x.com/2017/11/30/4SYct.png)

## 开启应用，添加两条数据，并进行查询

![](https://s1.ax1x.com/2017/11/30/4SaB8.png)

![](https://s1.ax1x.com/2017/11/30/4SdHS.png)

## 用ab进行压力测试

![](https://s1.ax1x.com/2017/11/30/4S0Ag.png)

![](https://s1.ax1x.com/2017/11/30/4SBNQ.png)

## 换用xorm进行压力测试

![](https://s1.ax1x.com/2017/11/30/4SDhj.png)

![](https://s1.ax1x.com/2017/11/30/4Ss9s.png)

## 对比结果

执行相同操作，使用xorm要比使用database/sql花费稍多一点的时间，但是xorm用起来更加方便，这一点时间代价还是可以接收的。