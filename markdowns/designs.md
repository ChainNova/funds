## 设计实现

funds基于Hypperledger实现的基金管理，Hyperledger为我们提供了如下的功能：

* 用户管理

Hyperledger的membersrvc模块提供了基本的用户管理功能，基于PKI体系的用户系统保证了交易的安全性。用户管理本身采用配置文件进行初始化，我们会进行一些扩展。

* 共识算法

共识算法提供了在分布式环境下解决数据一致性问题的方法。

* 区块链存储

区块链存储把所有的交易结果都存储在区块链上，称为ledger，任何人都可以查询ledger上的信息。

### 架构设计

架构设计包含三大部分：web client、App、Hyperledger。如下图

![fund架构图](./images/architecture.jpg)

web client：提供对外操作UI，实现user的输入输出简单处理后向App发送http request并接收response。

App：连接client与Hyperledger的中间层，负责接收client的httprequest，将request数据整理打包后通过Hyperledger提供API发送给Hyperledger处理；Hyperledger处理完成后返回处理结果给App，并有App包装后返回给client。

Hyperledger：基金管理系统底层区块链技术实现，提供memberSrv服务、peer共识服务、chaincode服务。负责执行交易并将交易相关信息保存于Ledger中。

###数据结构及流程

####数据结构

1. 基金基本信息：基金序号、基金名称、管理员
2. 基金净值：基金序号、净值
3. 系统全局量：基金序号、基金池容量、基金池中剩余基金数、系统资金量
4. 基金参与限制：基金序号、参与者资金量、参与者注册时间、认购起点
5. 基金认购限制：基金序号、认购单量、认购总量
6. 账户信息：账户证书、资金量、注册时间
7. 用户基金信息：账户证书、基金序号、所购基金份额
8. 排队信息：交易者证书、基金序号、交易额（认购或赎回）、申请时间

以上数据结构对应的数据都通过chaincode操作并保存在block的worldstate里。另外，系统账户的注册由Hyperledger的membersrv服务实现。

####系统流程

如下图流程图内所展示的逻辑是在chaincode实现。

![流程](./images/flowchart.jpg)

###APP接口设计

App模块为web client提供REST API。

#### 登陆

Request：

```
POST http://203.12.202.133:9900/login
{
	"enrollId":"lukas",//用户名
	"enrollSecret": "xoao",//密码
}
```

Response:

```
{
	"status": "OK",//或者"Err"
	"msg": "xxx"//错误信息
}
```

#### 创建基金

Request：

```
POST http://203.12.202.133:9900/create
{
	"name": "fundName",//基金名称
	"funds": 100,//初始基金数
	"assets": 100,//初始资金数
	"partnerAssets": 100,//注册资金
	"partnerTime": 100,//注册时间
	"buyStart": 100,//入购起点
	"buyPer": 100,//限购单量
	"buyAll": 100,//限购总量
	"net": 100//基金净值
}
```

Response:

```
{
	"status": "OK",//或者"Err"
	"msg": "xxx"//错误信息
}
```

#### 设置基金净值

Request：

```
POST http://203.12.202.133:9900/setnet
{
	"name": "fundName",//基金名称
	"net": 100//基金净值
}
```

Response:

```
{
	"status": "OK",//或者"Err"
	"msg": "xxx"//错误信息
}
```


#### 设置基金限制

Request：

```
POST http://203.12.202.133:9900/setlimit
{
	"name": "fundName",//基金名称
	"partnerAssets": 100,//注册资金
	"partnerTime": 100,//注册时间
	"buyStart": 100,//入购起点
	"buyPer": 100,//限购单量
	"buyAll": 100,//限购总量
}
```

Response:

```
{
	"status": "OK",//或者"Err"
	"msg": "xxx"//错误信息
}
```


#### 扩股回购

Request：

```
POST http://203.12.202.133:9900/setpool
{
	"name": "fundName",//基金名称
	"funds": 100,//扩股回购数，>0扩股  <0为回购
}
```

Response:

```
{
	"status": "OK",//或者"Err"
	"msg": "xxx"//错误信息
}
```


#### 认购赎回

Request：

```
POST http://203.12.202.133:9900/transfer
{
	"enrollID":"lukas",//用户ID
	"name": "fundName",//基金名称
	"funds": 100,//认购赎回数，>0认购  <0为赎回
}
```

Response:

```
{
	"status": "OK",//或者"Err"
	"msg": "xxx"//错误信息
}
```

#### 根据基金名称查询基金信息

Request：

```
GET http://203.12.202.133:9900/fund/:name

```

Response:

```
{
	"status": "OK",//或者"Err"
	"msg": "{
				"name": "fundName",//基金名称
				"funds": 100,//初始基金数
				"assets": 100,//初始资金数
				"partnerAssets": 100,//注册资金
				"partnerTime": 100,//注册时间
				"buyStart": 100,//入购起点
				"buyPer": 100,//限购单量
				"buyAll": 100,//限购总量
				"net": 100//基金净值
			}"//或错误信息
}
```

#### 查询所有基金信息

Request：

```
GET http://203.12.202.133:9900/funds

```

Response:

```
{
	"status": "OK",//或者"Err"
	"msg": "[{
				"name": "fundName",//基金名称
				"funds": 100,//初始基金数
				"assets": 100,//初始资金数
				"partnerAssets": 100,//注册资金
				"partnerTime": 100,//注册时间
				"buyStart": 100,//入购起点
				"buyPer": 100,//限购单量
				"buyAll": 100,//限购总量
				"net": 100//基金净值
    }]"//或错误信息
}
```

#### 查询用户某一基金的信息
Request：

```
GET http://203.12.202.133:9900/user/:fundName/:enrollID

```

Response:

```
{
	"status": "OK",//或者"Err"
	"msg": "{
				"name": "fundName",//基金名称
				"owner": 100,//用户名
				"assets": 100,//资金数
				"fund": 100,//基金数
			}"//或错误信息
}
```


###Hyperledger API

#### 初始化

Deploy Request:

```
POST host:port/chaincode
{
    "jsonrpc": "2.0",
    "method": "deploy",
    "params": {
        "type": "GOLANG",
        "chaincodeID": {
            "path": "",
            "name": ""
        },
        "ctorMsg": {
            "args": "[][]byte{}"//参数 init
        },
        "timeout": 0,
        "secureContext": "string",
        "confidentialityLevel": 1,
        "metadata": "[]byte {}",
        "attributes": "[]string{}"
    },
    "id": {
        "StringValue": "*string",
        "IntValue": "*int64"
    }
}
```

Deploy Response:

```
{
    "jsonrpc": "2.0",
    "result": {
        "status": "ok",
        "message": "xxxx"
    },
    "id": {
        "StringValue": "*string",
        "IntValue": "*int64"
    }
}
```

#### 创建基金

Invoke Request:

```
POST host:port/chaincode
{
    "jsonrpc": "2.0",
    "method": "invoke",
    "params": {
        "type": "GOLANG",
        "chaincodeID": {
            "path": "",
            "name": ""
        },
        "ctorMsg": {
            "args": "[][]byte{}"//参数 1、“createFund“ 2、基金名称string  3、基金管理员  4、基金净值 5、基金池 6、系统资金 7、参与者资金量 8、参与者注册时间 9、认购起点 10、认购单量 11、认购总量 12、基金净值
        },
        "timeout": 0,
        "secureContext": "string",
        "confidentialityLevel": 1,
        "metadata": "[]byte {}",
        "attributes": "[]string{}"
    },
    "id": {
        "StringValue": "*string",
        "IntValue": "*int64"
    }
}
```

Invoke Response:

```
{
    "jsonrpc": "2.0",
    "result": {
        "status": "ok",
        "message": "xxxx"
    },
    "id": {
        "StringValue": "*string",
        "IntValue": "*int64"
    }
}
```

#### 设置基金净值

Invoke Request:

```
POST host:port/chaincode
{
    "jsonrpc": "2.0",
    "method": "invoke",
    "params": {
        "type": "GOLANG",
        "chaincodeID": {
            "path": "",
            "name": ""
        },
        "ctorMsg": {
            "args": "[][]byte{}"//参数 1、“setNet“ 2、基金名  3、净值int
        },
        "timeout": 0,
        "secureContext": "string",
        "confidentialityLevel": 1,
        "metadata": "[]byte {}",
        "attributes": "[]string{}"
    },
    "id": {
        "StringValue": "*string",
        "IntValue": "*int64"
    }
}
```

Invoke Response:

```
{
    "jsonrpc": "2.0",
    "result": {
        "status": "ok",
        "message": "xxxx"
    },
    "id": {
        "StringValue": "*string",
        "IntValue": "*int64"
    }
}
```

#### 设置基金池（扩股、回购）

Invoke Request:

```
POST host:port/chaincode
{
    "jsonrpc": "2.0",
    "method": "invoke",
    "params": {
        "type": "GOLANG",
        "chaincodeID": {
            "path": "",
            "name": ""
        },
        "ctorMsg": {
            "args": "[][]byte{}"//参数 1、“setFoundPool“ 2、基金名  3、扩股/回购数（>0为扩股 <0为回购）
        },
        "timeout": 0,
        "secureContext": "string",
        "confidentialityLevel": 1,
        "metadata": "[]byte {}",
        "attributes": "[]string{}"
    },
    "id": {
        "StringValue": "*string",
        "IntValue": "*int64"
    }
}
```

Invoke Response:

```
{
    "jsonrpc": "2.0",
    "result": {
        "status": "ok",
        "message": "xxxx"
    },
    "id": {
        "StringValue": "*string",
        "IntValue": "*int64"
    }
}
```

#### 基金交易（认购赎回）

Invoke Request:

```
POST host:port/chaincode
{
    "jsonrpc": "2.0",
    "method": "invoke",
    "params": {
        "type": "GOLANG",
        "chaincodeID": {
            "path": "",
            "name": ""
        },
        "ctorMsg": {
            "args": "[][]byte{}"//参数 1、“transferFound“ 2、基金ID  3、认购/赎回数（>0为认购 <0为赎回）
        },
        "timeout": 0,
        "secureContext": "string",
        "confidentialityLevel": 1,
        "metadata": "[]byte {}",
        "attributes": "[]string{}"
    },
    "id": {
        "StringValue": "*string",
        "IntValue": "*int64"
    }
}
```

Invoke Response:

```
{
    "jsonrpc": "2.0",
    "result": {
        "status": "ok",
        "message": "xxxx" //如果交易成功则为交易额（可能是部分交易完成），否则为错误信息
    },
    "id": {
        "StringValue": "*string",
        "IntValue": "*int64"
    }
}
```

#### 基金限制设置

Invoke Request:

```
POST host:port/chaincode
{
    "jsonrpc": "2.0",
    "method": "invoke",
    "params": {
        "type": "GOLANG",
        "chaincodeID": {
            "path": "",
            "name": ""
        },
        "ctorMsg": {
            "args": "[][]byte{}"//参数 1、“setFundLimit“ 2、基金名  3、参与者资金量 4、参与者注册时间 5、认购起点 6、认购单量 7、认购总量 
        },
        "timeout": 0,
        "secureContext": "string",
        "confidentialityLevel": 1,
        "metadata": "[]byte {}",
        "attributes": "[]string{}"
    },
    "id": {
        "StringValue": "*string",
        "IntValue": "*int64"
    }
}
```

Invoke Response:

```
{
    "jsonrpc": "2.0",
    "result": {
        "status": "ok",
        "message": "xxxx" //如果交易成功则为交易额（可能是部分交易完成），否则为错误信息
    },
    "id": {
        "StringValue": "*string",
        "IntValue": "*int64"
    }
}
```

#### 基金/列表信息查询

Query Request:

```
POST host:port/chaincode
{
    "jsonrpc": "2.0",
    "method": "query",
    "params": {
        "type": "GOLANG",
        "chaincodeID": {
            "path": "",
            "name": ""
        },
        "ctorMsg": {
            "args": "[][]byte{}"//参数 1、“getFund“ 2、one/list 3、基金ID（第二个参数为one时需要此参数） 
        },
        "timeout": 0,
        "secureContext": "string",
        "confidentialityLevel": 1,
        "metadata": "[]byte {}",
        "attributes": "[]string{}"
    },
    "id": {
        "StringValue": "*string",
        "IntValue": "*int64"
    }
}
```

Query Response:

```
{
    "jsonrpc": "2.0",
    "result": {
        "status": "ok",
        "message": "xxxx" //基金信息（包括所有基本信息的struct）
    },
    "id": {
        "StringValue": "*string",
        "IntValue": "*int64"
    }
}
```

#### 基金净值/列表查询

Query Request:

```
POST host:port/chaincode
{
    "jsonrpc": "2.0",
    "method": "query",
    "params": {
        "type": "GOLANG",
        "chaincodeID": {
            "path": "",
            "name": ""
        },
        "ctorMsg": {
            "args": "[][]byte{}"//参数 1、“getFundList“ 2、one/list  3、基金ID
        },
        "timeout": 0,
        "secureContext": "string",
        "confidentialityLevel": 1,
        "metadata": "[]byte {}",
        "attributes": "[]string{}"
    },
    "id": {
        "StringValue": "*string",
        "IntValue": "*int64"
    }
}
```

Query Response:

```
{
    "jsonrpc": "2.0",
    "result": {
        "status": "ok",
        "message": "xxxx" //基金净值信息
    },
    "id": {
        "StringValue": "*string",
        "IntValue": "*int64"
    }
}
```

#### 注册

Enrollment Request:

```
POST host:port/registrar

{
  "enrollId": "lukas",
  "enrollSecret": "NPKYL39uKbkj"
}

```

Enrollment Response:

```
{
    "OK": "Login successful for user 'lukas'."
}
```