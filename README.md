# API\_SEUXW\_BACKEND

Backend system code for [seuxw.cc](http://seuxw.cc), based on [Golang](https://golang.org/) and [Python3](https://www.python.org/).

seuxw.cc 项目的 HTTP API 服务后端系统代码，基于 Golang 和 Python3 开发。

seuxw.cc 是一个以 “东大小微” 为虚拟载体，为东南大学本科生提供线上服务的平台。主要提供的服务包括了学生信息查询、基础娱乐功能、个人定制功能等，服务的平台囊括 QQ Robot、 Qzone、 Web 等。项目始于 2014 年，目前项目的用户量约 6000 人次。

## 0. 写在前面

由于本项目使用了 git 的子模块功能，请使用

```
git clone --recursive git@github.com:seuxw/api_seuxw_backend.git

```

命令克隆本项目的完整内容。


## 1. 项目简介
| 项目名称 | 项目中文名 | 项目团队 | 作者 | 创建日期 |
| :-: | :-: | :-: | :-: | :-: |
| **API\_SEUXW\_BACKEND** | **东大小微 HTTP API 服务后端系统** *基于项目 [api\_seuxw\_template](https://github.com/seuxw/api_seuxw_template) 拓展开发* | **[SEUXW](https://github.com/seuxw)** | **[@TauWu](https://github.com/TauWu)** | 2018-06-29 |

## 2. 项目规范

- 命名规范

    - 1. 属性命名
    
    中文属性 | English Property | Python | Golang |
    :-: | :-: | :-: | :-: |
    函数 | Function | `sample_func()` | `SampleFunc` |
    变量 | Variable | `sample_var` | `sampleVar` |
    文件名 | File | `sample_file` | `sample_file` |
    类 | Class | `SampleClass` | - |
    类公共函数 | Class Public Function | `sample_func()` | `SampleFunc()` |
    类公共属性 | Class Public Property | `sample_property` | - |
    类保护函数 | Class Protected Function | `_sample_func()_` | - |
    类保护属性 | Class Protected Property | `_sample_property` | - |
    类私有函数 | Class Private Function | `__sample_func()__` | `sampleFunc()` |
    类私有属性 | Class Private Property | `__sample_property` | - |
    结构体 | Struct | - | `sampleStruct` |
    结构体变量 | Struct Variable | - | `SampleStructVariable` |
    结构体接口 | Struct Interface | - | `SampleInterface` |
    标签 | Label | - | `SampleLabel` |

    - 2. 函数块与注释位置

        - Python
        
        ```py
        class SampleClass():
            '''SampleClass
            示例类
            '''
            def sample_func():
                '''sample_func
                示例函数
                params:
                    xxx                    
                returns:
                    xxx
                '''
                # Comments 这里写流程控制中比较重要的注释
                if True:
                    user_id = 0  # 用户 ID
                return
        ```
        - Golang

        ```go

        // Pagination
        // 接口返回分页信息
        type Pagination struct {
            Page     int `json:"page" db:"page"`            // 页数
            PageSize int `json:"page_size" db:"page_size"`  // 页面容量
            Total    int `json:"total" db:"total"`          // 总数
        }

        // SampleFunction xxx
        // 示例功能函数
        func (s *sampleStruct) SampleFunction(sampleVar sampleStruct)(sampleVar sampleStruct, err error) {
            var (
                int64 userID  // 用户 ID
            )
            
            // Comments 这里写流程控制中比较重要的注释
            if true {
                goto END
            }
            return

        // Label 注释
        END:
            return
        }
        ```

- 错误处理

    - Python:
        - 代码示例

        ```py
        try:
            sample_func()

        # 错误情况 1
        except ValueError as e1:
            err_func1(e1)

        # 未知错误
        except Exception as e:
            err_func(e)
        
        # 无错误抛出
        else:
            else_func()
            
        finally:
            final_func()
        ```
        - 说明：
            1. 将 **可预见的错误** 和 **警告** 信息放在 Exception 之前，最后只处理未知错误。
            2. err_func() 不一定是一个函数，可能只是一个日志的记录或者是个 raise。
            3. 可以使用断言（`assert`）进行简单的抛错处理。

    - Golang:
        - 写法 1：

        ```go
        result, err := SampleFunction()
        if err != nil {
            // 错误处理
            return
        }
        ```

        - 写法 2：

        ```go
        if result, err := SampleFunction(); err !=nil {
            // 错误处理
            return
        }
        ```

- 参数传递
    - Golang：
        1. 对于少量数据，不要传递指针
        2. 对于很多字段的 struct，建议使用指针操作
        3. 参数是 map, slice, chan 的 **请勿** 传递指针
        4. 参数是 Mutex 类型的 **必须** 指针

- 代码格式
    1. 函数中可以化零为整的代码段尽量定义子函数，如果实在不方便可以将一个函数一些可以整合在一起的代码段放在一起，相对独立的部分用 <ENTER> 隔开。
    2. 一行代码最好不要超过 80 个字符。
    3. Golang 的 IDE 中开启 gofmt 自动格式化代码。
    4. 一个 TAB 是 4 个空格。
    
- 日志处理
    1. API 部分需要携带 **所有的参数** 和 **所有的返回（data 部分为列表的看具体情况返回）**
    2. 日志格式：
    ```
    [15722] 2012/03/04 14:00:00.123123 [TRC] [GetUserInfoV2:230] 获取用户信息V2流程成功～ => TraceID: 7ee9bdda-ee49-4830-9546-8e3f7b5a23d0, Data: {"user_id":0}, Message:success., Pagination: <nil>
    ```
    3. 接受到请求的日志打印 **传递参数**，响应返回的日志打印 **返回数据信息**，中间有关键流程也建议一并打印。

- 其他规范
    1. **使用** cfg 文件设置配置文件 **而不是** 使用类似于 config.go 文件。
    2. sqlx 的使用
        - 查询多行数据

        ```go
        people := []Person{}
        db.Select(&people, "select * from person order by id asc")        
        ```
        - 查询单行数据

        ```go
        tau := Person{}
        err = db.Get(&tau, "select * from person where name = ?", "tau")
        ```

        - 执行 INSERT, UPDATE, DELETE 语句

        ```go
        result := db.MustExec(insertSQL)
        ```

        - 执行数据库事务

        ```go
        var (
            resultSet   []sql.Result
            result      sql.Result
        )
        tx := db.MustBegin()

        result = tx.MustExec("insert into (param1, param2) values ($1, $2)", "1", "2")
        resultSet = append(resultSet, result)

        result = tx.MustExec("insert into (param1, param2) values ($1, $2)", "1", "2")
        resultSet = append(resultSet, result)

        err := tx.Commit()
        if err != nil {
            db.log.Error("事务提交错误 %s", err)
            err = tx.RollBack()
            if err != nil {
                db.log.Error("事务回滚错误 %s", err)
            }
        }

        for idx := range resultSet {
            rowaffected, err := resultSet[idx].RowsAffected()
            if err != nil {
                // 错误处理
            }
            self.log.Trace("执行成功: lastInsertId: [%d]", lastInsertId)
            self.log.Trace("执行成功: rowaffected: [%d]", rowaffected)
        }
        ```

## 3. 项目开始前的工作

- 可以通过直接执行项目根目录的 make.sh 程序进行编译
- 进入 seuxw 目录后，执行 make 可以在 seuxw/_output/local/bin 目录生成可执行程序
- 执行完 make 后可以在 VSCode 中按 F5 进行程序调试
- 需要在 GOPATH(/data/code/com/go) 中添加部分 VSCode 插件依赖 [下载地址](https://share.weiyun.com/5IzgLKh)

## 4. 目录结构树

- `.vscode` VSCode 配置文件
- `database` 数据库建表语句
    - Initialize_database.sql 建表语句
- `apidoc` [API 文档项目](https://github.com/seuxw/apidoc)
- `seuxw` Golang 项目代码
    - `_output` 在 ./seuxw 根目录下执行 make 将会在此处生成软链接
    - `bash` shell
    - `filter` API 项目路径
        - `user` 项目 user
    - `vendor` github 等网络资源
    - `x` 外部模块
    - `embrice` 内部模块
        - `api` 调用外部 API
        - `constant` 常量
        - `entity` 结构体
        - `extension` 扩展
        - `middleware` 中间件
        - `rdb` 数据库操作
- `py_seuxw` [Python 脚本项目](https://github.com/seuxw/py_seuxw)
- .gitignore Git 忽略文件
- make.sh 快捷编译脚本
- README.md

## 5. Q&A

Q: 能编译成功，但是编译后不能在本地执行怎么办？

A: [build.sh](./seuxw/bash/build.sh) 的第 49 行

```sh
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $GO_BUILD_FLAGS \
```

指定了编译系统为 Linux。如果您是设备为 Mac，可以将 GOOS=linux 改为 GOOS=darwin。不推荐使用 Windows 系统编译这一套代码。关于 Win10 内部的 ubuntu 子系统，可以 [参考这里](https://tutorials.ubuntu.com/tutorial/tutorial-ubuntu-on-windows)。

Q: 为什么执行 ./make.sh 之后，编译的是 test 项目代码，并且执行的是 test.x 二进制文件？

A: 1. 项目编译问题

[build.sh](./seuxw/bash/build.sh) 的第 9-11 行

```sh
readonly SEUXW_TARGETS=(
	filter/test
)
```

指定了编译的项目路径，可以在后面添加如 filter/user 的内容来新增需要编译的项目，同样可以删除不需要编译的项目。

  2. 项目执行问题

[make.sh](./seuxw/make.sh) 中的最后一行指定了需要执行的文件名称，可以在这里修改成你需要执行的项目二进制文件。