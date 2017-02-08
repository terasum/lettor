# Documentation

## `public HyperchainAPI(String providerURL, HttpProvider httpProvider,boolean isOutput,String ecertPath,String ecertPrivPath,String privateKeyPath,String publicKeyPath) throws Exception`

指定网络请求的provider的来源

 * **Parameters:**
   * `providerURL` — 请求的地址,如果是本地文本请求,请直接指向JSON文件所在文件夹,如果是网络请求,请按照 http://localhost:8081 格式填写
   * `httpProvider` — 请求类型,包括INTERNET 和 PLANTEXT

## `public SingleValueReturn getTransactions(String start, String end, int id)`

2.1 取得区块中的所有交易信息 tx_getTransactions

 * **Parameters:**
   * `start` — 查询开始区块号
   * `end` — 查询结束区块号
   * `id` — 查询批次ss
 * **Returns:** 标准格式返回值

## `public SingleValueReturn getDiscardTransactions(int id)`

2.2 查询所有非法交易

 * **Parameters:** `id` — 查询批次
 * **Returns:** 返回交易json列表

## `public SingleValueReturn getTransactionByHash(String txHash, int id)`

2.3 查询交易 by TxHash

 * **Parameters:**
   * `txHash` — -
   * `id` — 查询批次
 * **Returns:** 单值json string返回值

## `public SingleValueReturn getTransactionByBlockHashAndIndex(String blkHash, int index, int id)`

2.4 查询交易 by block hash and TxIndex

 * **Parameters:**
   * `blkHash` — 块hash
   * `index` — 交易序号
   * `id` — 查询批次
 * **Returns:** 返回交易信息数组

## `public SingleValueReturn getTransactionByBlockNumberAndIndex(int blkNumber, int index, int id)`

2.5 查询交易 by block number and TxIndex

 * **Parameters:**
   * `blkNumber` — 块Num
   * `index` — 交易index
   * `id` — -
 * **Returns:** 单值字符串返回

## `public SingleValueReturn sendTransaction(String from, String to, long value, String accountJson, String passwd, int id) throws Exception`

2.6 发送交易(带签名)

 * **Parameters:**
   * `from` — 发送账户地址
   * `to` — 接受账户地址
   * `value` — 发送的交易值
   * `accountJson` — 账户存储json字符串
   * `passwd` — 账户解锁密码
   * `id` — -
 * **Returns:** 返回交易hash
 * **Exceptions:** `Exception` — 自定义exception 一般是json解析错

## `public SingleValueReturn sendTransaction(String from, String to, long value, String accountJson, String passwd,boolean simulate, int id) throws Exception`

 * **Parameters:**
   * `from` — 发送账户地址
   * `to` — 接受账户地址
   * `value` — 发送的交易值
   * `accountJson` — 账户存储json字符串
   * `passwd` — 账户解锁密码
   * `simulate` — 模拟flag
   * `id` — -
 * **Returns:** 标准单值返回
 * **Exceptions:** `Exception` — -

## `public HmReturn sendTransaction(String contractAddr,String from , String to, long transferamount, long allamount, BigInteger illegalamount, PublicKey reciverpublicKey, String accountJson, String passwd, int id) throws Exception`

 * **Parameters:**
   * `from` — 发起交易人地址
   * `to` — 收款人地址
   * `transferamount` — 要转的金额
   * `allamount` — 所有的金额
   * `illegalamount` — 非法同态值
   * `reciverpublicKey` — 收款人公钥
   * `accountJson` — 发起人私钥文件的内容
   * `passwd` — 发起人密码
   * `id` — 交易id
 * **Returns:** 返回

## `public HmVerifyReturn ReciverVerify(String contractAddr,String from,String abi,String accountJson,String passwd,int id) throws Exception`

 * **Parameters:**
   * `contractAddr` — 专门针对同态转账的合约地址
   * `from` — 调用者
   * `abi` — 合约的abi，因为需要使用到abi进行解析合约方法的返回值
   * `accountJson` — 调用者私钥文件内容
   * `passwd` — 调用者密码
   * `id` — 调用号
 * **Returns:** 返回验证结果，包括里面调用合约的结果，以及平台验证结果
 * **Exceptions:** `Exception` — 

## `private SingleValueReturn sendSignedTransaction(SignedTxParams signedTxParams, int id) throws Exception`

发送已经签名的交易

 * **Parameters:**
   * `signedTxParams` — 已经签名的交易
   * `id` — 请求批次
 * **Returns:** 标准格式返回值

## `public SingleValueReturn getTxAvgTimeByBlockNumber(int from,int to, int id)`

2.8 查询指定区块中交易平均处理时间

 * **Parameters:**
   * `from` — 启始区块
   * `to` — 结束区块
   * `id` — 查询批次
 * **Returns:** 返回平均处理时间,16进制

## `public ReceiptReturn getTransactionReceipt(String TxHash, int id)`

2.9 查询指定交易中的收据信息

 * **Parameters:**
   * `TxHash` — 交易hash
   * `id` — 查询批次
 * **Returns:** 标准格式返回值

## `public SingleValueReturn getBlockTransactionCountByHash(String blkHash ,int id)`

2.10 查询区块交易数量

 * **Parameters:**
   * `blkHash` — 区块hash
   * `id` — -
 * **Returns:** 数目 16进制表示

## `private String getSignHash(TxParams txParams, int id)`

2.11获取交易签名哈希 取得交易的待签名 hash

 * **Returns:** string 签名hash

## `public SingleValueReturn getTransactionsCount(int id)`

查询链上所有交易量

 * **Parameters:** `id` — 批次
 * **Returns:** 返回相应的json数据

## `public SingleValueReturn getTransactionsByTime(BigInteger startTime,BigInteger endTime, int id)`

2.11 根据时间查询Transaction Hyperchain branch agriculture_addCheckingInterface_mxm 2016-12-08 tx_getTransactionsByTime

 * **Parameters:**
   * `startTime` — 开始时间戳 ns
   * `endTime` — 结束时间戳 ns
   * `id` — -
 * **Returns:** 返回txs 的json

## `public CompileReturn compileContract(String sourceCode, int id)`

3.1 编译源代码

 * **Parameters:**
   * `sourceCode` — 源代码
   * `id` — 请求批次
 * **Returns:** 标准格式返回值

## `public SingleValueReturn deployContract(String from, String bin, String accountJSON, String passwd, int id) throws Exception`

3.2 部署合约

 * **Parameters:**
   * `from` — 部署账户地址
   * `bin` — 需要部署的只能合约bin
   * `accountJSON` — 加密之后的账户秘钥信息
   * `passwd` — 解密密码
   * `id` — 查询批次
 * **Returns:** 交易hash
 * **Exceptions:** `Exception` — 部署错误

## `public SingleValueReturn deployContract(String from, String bin, ECPriv ecPriv, int id) throws Exception`

 * **Parameters:**
   * `from` — 调用地址
   * `bin` — 部署合约bin
   * `ecPriv` — 明文私钥
   * `id` — -
 * **Returns:** 返回部署 Txhash
 * **Exceptions:** `Exception` — -

## `public SingleValueReturn deployContract(String from, String bin, String accountJSON, String passwd,boolean simulate, int id) throws Exception`

3.2 部署合约 模拟

 * **Parameters:**
   * `from` — 部署账户地址
   * `bin` — 需要部署的只能合约bin
   * `accountJSON` — 加密之后的账户秘钥信息
   * `passwd` — 解密密码
   * `simulate` — 模拟flag
   * `id` — 查询批次
 * **Returns:** 交易hash
 * **Exceptions:** `Exception` — 异常

## `public SingleValueReturn deployContract(String from, String bin, String accountJSON, String passwd, int id, String constructMethodName, FuncParam... functionParams) throws Exception`

3.2 部署合约(构造函数,funcParam，弃用)

 * **Parameters:**
   * `from` — 部署账户地址
   * `bin` — 需要部署的只能合约bin
   * `accountJSON` — 加密之后的账户秘钥信息
   * `passwd` — 解密密码
   * `id` — 查询批次
   * `constructMethodName` — 构造函数方法名
   * `functionParams` — 构造函数参数
 * **Returns:** 交易hash
 * **Exceptions:** `Exception` — 部署错误

## `public SingleValueReturn deployContract(String from, String bin, String accountJSON, String passwd, int id, String constructMethodName, FuncParamReal... functionParams) throws Exception`

 * **Parameters:**
   * `from` — 部署账户地址
   * `bin` — 需要部署的只能合约bin
   * `accountJSON` — 加密之后的账户秘钥信息
   * `passwd` — 解密密码
   * `id` — 查询批次
   * `constructMethodName` — 构造函数方法名
   * `functionParams` — 构造函数参数
 * **Returns:** 交易hash
 * **Exceptions:** `Exception` — 部署错误

## `public SingleValueReturn deployContract(String from, String bin, ECPriv ecPriv, int id, String constructMethodName, FuncParam... functionParams) throws Exception`

 * **Parameters:**
   * `from` — 调用地址
   * `bin` — 合约bin
   * `ecPriv` — 明文私钥
   * `id` — 批次
   * `constructMethodName` — 构造方法名
   * `functionParams` — 构造方法参数
 * **Returns:** 交易hash
 * **Exceptions:** `Exception` — -

## `public SingleValueReturn deployContract(String from, String bin, ECPriv ecPriv, int id, String constructMethodName, FuncParamReal... functionParams) throws Exception`

 * **Parameters:**
   * `from` — 部署账户地址
   * `bin` — 需要部署的只能合约bin
   * `ecPriv` — 明文私钥文件
   * `id` — 查询批次
   * `constructMethodName` — 构造方法名
   * `functionParams` — 构造方法参数，FuncParamReal类型
 * **Returns:** 交易hash
 * **Exceptions:** `Exception` — 

## `public SingleValueReturn deployContract(String from, String bin, String accountJSON, String passwd,boolean simulate,int id, String constructMethodName, FuncParam... functionParams) throws Exception`

3.2 部署合约(构造函数) 模拟

 * **Parameters:**
   * `from` — 部署账户地址
   * `bin` — 需要部署的只能合约bin
   * `accountJSON` — 加密之后的账户秘钥信息
   * `passwd` — 解密密码
   * `simulate` — 模拟
   * `id` — 查询批次
   * `constructMethodName` — 构造函数方法名
   * `functionParams` — 构造函数参数
 * **Returns:** 交易hash
 * **Exceptions:** `Exception` — 部署错误

## `public SingleValueReturn deployContract(String from, String bin, ECPriv ecPriv,boolean simulate,int id, String constructMethodName, FuncParam... functionParams) throws Exception`

 * **Parameters:**
   * `from` — 部署账户地址
   * `bin` — 需要部署的只能合约bin
   * `ecPriv` — 明文私钥
   * `simulate` — 模拟
   * `id` — 查询批次
   * `constructMethodName` — 构造函数方法名
   * `functionParams` — 构造函数参数
 * **Returns:** 部署交易hash
 * **Exceptions:** `Exception` — -

## `private SingleValueReturn getSignHashForContract(DeployParams deployParams, int id)`

取得部署合约的待签名 hash

 * **Returns:** string 签名hash

## `public SingleValueReturn invokeContract(String from, String contractAddr, String input, String accountJson, String passwd, int id) throws Exception`

3.3 调用合约

 * **Parameters:**
   * `from` — 合约from
   * `contractAddr` — 合约地址
   * `input` — 调用编码
   * `accountJson` — 账户加密json
   * `passwd` — 账目密码
   * `id` — 查询批次
 * **Returns:** 返回单值结果
 * **Exceptions:** `Exception` — 抛出签名错误

## `public SingleValueReturn invokeContract(String from, String contractAddr, String input, ECPriv ecPriv, int id) throws Exception`

3.3 调用合约 明文私钥

 * **Parameters:**
   * `from` — 合约from
   * `contractAddr` — 合约地址
   * `input` — 调用编码
   * `ecPriv` — 明文私钥
   * `id` — 查询批次
 * **Returns:** 返回单值结果
 * **Exceptions:** `Exception` — 抛出签名错误

## `public String invokeContract(String from,String contractAddr,String abi,String methodName,String input,String accountJson,String passwd,int id) throws Exception`

invoke the Smart contract and decode the result automatically. return string format is a predefined json string

 * **Parameters:**
   * `from` — the invoke user address
   * `contractAddr` — the address which the smart contract deployed before
   * `abi` — Smart contract abi string
   * `methodName` — the method name which you need to invoke
   * `input` — the encoded method params string
   * `accountJson` — the encrypted account Json string
   * `passwd` — the password which need to decrypt the encrypted json string
   * `id` — the batch id you just query
 * **Returns:** return a formatted json string
 * **Exceptions:** `Exception` — -

## `public String invokeContract(String from,String contractAddr,String abi,String methodName,String input,ECPriv ecPriv,int id) throws Exception`

便捷调用合约方法，将会自动返回解码之后的返回值

 * **Parameters:**
   * `from` — 调用账户地址
   * `contractAddr` — 合约部署地址
   * `abi` — 对应合约abi
   * `methodName` — 方法名
   * `input` — 调用参数input
   * `ecPriv` — 明文私钥
   * `id` — 调用批次
 * **Returns:** 返回明文解码返回值json
 * **Exceptions:** `Exception` — -

## `public SingleValueReturn invokeContract(String from, String contractAddr, String input, String accountJson, String passwd,boolean simulate, int id) throws Exception`

3.3 调用合约 模拟

 * **Parameters:**
   * `from` — 合约from
   * `contractAddr` — 合约地址
   * `input` — 调用编码
   * `accountJson` — 账户加密json
   * `passwd` — 账目密码
   * `simulate` — 模拟flag
   * `id` — 查询批次
 * **Returns:** 返回单值结果
 * **Exceptions:** `Exception` — 抛出签名错误

## `public SingleValueReturn invokeContract(String from, String contractAddr, String input, ECPriv ecPriv,boolean simulate, int id) throws Exception`

 * **Parameters:**
   * `from` — 调用账户地址
   * `contractAddr` — 合约地址
   * `input` — 调用输入
   * `ecPriv` — 明文私钥
   * `simulate` — 是否模拟请求
   * `id` — 调用批次
 * **Returns:** 返回SingleValueReturn 交易hash
 * **Exceptions:** `Exception` — -

## `public String newAccount(String passphrase) throws GeneralSecurityException`

生成账户keystore 文件, json字符串形式返回

 * **Parameters:** `passphrase` — 设置账户密码
 * **Returns:** json格式keystore文件
 * **Exceptions:** `GeneralSecurityException` — 加密算法异常

## `public AccountContent newHmAccount(String passphrase) throws GeneralSecurityException`

功能：当使用同态加密的时候，需要用这种方法产生公私钥，得到encode之后的pri和pub进行保存

 * **Parameters:** `passphrase` — 用户密码
 * **Returns:** 返回一个账户内容结构体
 * **Exceptions:** `GeneralSecurityException` — 

## `public ECPriv newAccount() throws GeneralSecurityException`

取得账户明文私钥以及地址

 * **Returns:** ECPriv 包括用户私钥以及账户地址
 * **Exceptions:** `GeneralSecurityException` — -

## `public String encryptAccount(ECPriv privateKey, String password) throws GeneralSecurityException`

加密明文私钥

 * **Parameters:**
   * `privateKey` — 账户明文私钥及地址
   * `password` — 设置账户密码
 * **Returns:** json格式加密私钥文件
 * **Exceptions:** `GeneralSecurityException` — 加密算法异常

## `public ECPriv decryptAccount(String encrypetedAccountJson, String pwd) throws Exception`

解密加密私钥文件

 * **Parameters:**
   * `encrypetedAccountJson` — 加密私钥文件
   * `pwd` — 用户密码
 * **Returns:** 明文私钥文件
 * **Exceptions:** `Exception` — 

## `public BlockReturn getBlockByHash(String blockHash, int id)`

查询区块信息 by blk Hash block_getBlocksByHash

 * **Parameters:**
   * `blockHash` — blk hash
   * `id` — -
 * **Returns:** 区块信息

## `public BlockNumReturn getBlocksByTime(BigInteger startTime,BigInteger endTime,int id)`

根据区块生成时间查询区块

 * **Parameters:**
   * `startTime` — 开始查询时间
   * `endTime` — 结束查询时间
   * `id` — -
 * **Returns:** 返回BlockNumReturn

## `public ArrayList<NodeInfoReturn> getNodes(int id)`

取得所有的节点信息

 * **Parameters:** `id` — -
 * **Returns:** 返回节点信息列表