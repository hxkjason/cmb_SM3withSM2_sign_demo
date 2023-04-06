
### Golang 实现对接招行国密 SM3withSM2 的签名算法

招行对接文档: [__https://openapi.cmbchina.com/help/two/2-1-5.html__](https://openapi.cmbchina.com/help/two/2-1-5.html)

 |  鉴权方式   | 是否校验appid |签名算法  |对称/非对称签名 |是否校验时间戳 |是否签名secret|
 |:----:|:----:|:----:|:----:|:----:|:----:|
 | 非对称签名认证-带body摘要  | 是 |SM3withSM2 | 非对称 |是|是|

### 替换配置信息
将文件 [__cmb_openapi_service/const_var_struct.go__](https://github.com/hxkjason/cmb_openapi_sign_demo/cmb_openapi_service/const_var_struct.go) 中的 AppId、AppSecret、PrivateKey 替换为您在招行 Open API 中创建应用的真实信息。

### 依赖第三方包 **gmsm** v1.4.1

[__https://github.com/tjfoc/gmsm__](https://github.com/tjfoc/gmsm)

### SM3withSM2 签名代码

[__gmsm_service/gmsm_service.go__](https://github.com/hxkjason/cmb_SM3withSM2_sign_demo/blob/main/gmsm_service/gmsm_service.go#L18-L50)


