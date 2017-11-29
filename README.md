# golang 构建数据服务

本项目的主体框架fork自潘老师的github，entiies2部分是我自己编写的，使用xorm来代替database/sql,使用了xorm就不需要写dao服务，非常方便，但使用反射技术，牺牲了性能。