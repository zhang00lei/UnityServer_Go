local string={	
	id="1001"	desc="服务器返回未知错误！"	module="EorrorCode EC_UNKNOWN"	
	id="1002"	desc="重复登录"	module="EorrorCode  EC_Login_1"	
	id="1003"	desc="用户被冻结"	module="EorrorCode EC_Login_2"	
	id="1004"	desc="图片索引错误"	module="EorrorCode EC_ImageIndex_ERROR"	
	id="1005"	desc="昵称错误"	module="EorrorCode EC_NickName_ERROR"	
	id="1006"	desc="性别错误"	module="EorrorCode EC_SEX_ERROR"	
	id="1007"	desc="手机号码错误"	module="EorrorCode EC_Phonenum_ERROR"	
	id="1008"	desc="集结号已经被绑定，此次无法绑定"	module="EorrorCode EC_JJHBind_Have"	
	id="1009"	desc="用户名或密码不正确"	module="EorrorCode EC_checkAndGetUserInfo_1"	
	id="1010"	desc="用户名不合法"	module="EorrorCode EC_registerUserInfo_1"	
	id="1011"	desc="用户名已存在"	module="EorrorCode EC_registerUserInfo_2"	
	id="1012"	desc="密码不合法"	module="EorrorCode EC_registerUserInfo_3"	
	id="1013"	desc="手机号不合法"	module="EorrorCode EC_registerUserInfo_4"	
	id="1014"	desc="用户不存在---用户编号不合法"	module="EorrorCode EC_queryCoin_1"	
	id="1015"	desc="用户不存在----编号不合法"	module="EorrorCode EC_drawCoin_1"	
	id="1016"	desc="密码不合法"	module="EorrorCode EC_drawCoin_2"	
	id="1017"	desc="余额不足"	module="EorrorCode EC_drawCoin_3"	
	id="1018"	desc="用户不存在----编号不合法"	module="EorrorCode EC_depositCoin_1"	
	id="1019"	desc="数值不合法"	module="EorrorCode EC_depositCoin_2"	
	id="1020"	desc="无此用户"	module="EorrorCode EC_NoUser"	
	id="1021"	desc="NULL MachineUUID"	module="EorrorCode EC_MachineUUID_IsNull"	
	id="1022"	desc="获得用户信息失败"	module="EorrorCode EC_GetUserInfo_Faild"	
	id="1023"	desc="参数个数错误"	module="EorrorCode EC_Param_Cont_ERROR"	
	id="1024"	desc="桌号错误"	module="EorrorCode EC_Zhuohao_ERROR"	
	id="1025"	desc="座位号错误"	module="EorrorCode EC_Zuoweihao_ERROR"	
	id="1026"	desc="场景类型错误"	module="EorrorCode EC_Scenetype_ERROR"	
	id="1027"	desc="武器类型错误"	module="EorrorCode EC_Weapontype_ERROR"	
	id="1028"	desc="用户等级不够错误"	module="EorrorCode EC_Userlevel_NotEnough"	
	id="1029"	desc="炮值错误"	module="EorrorCode EC_Paovalue_ERROR"	
	id="1030"	desc="无此炮错误"	module="EorrorCode EC_Paodan_NO"	
	id="1031"	desc="商品类型错误"	module="EorrorCode EC_Commoditytype_ERROR"	
	id="1032"	desc="强化等级错误"	module="EorrorCode EC_Weaponforce_Level_ERROR"	
	id="1033"	desc="玩家无此武器错误"	module="EorrorCode EC_UserNoWeapon_ERROR"	
	id="1034"	desc="道具类型错误"	module="EorrorCode EC_Daojutype_ERROR"	
	id="1035"	desc="MachineUUID不匹配错误"	module="EorrorCode EC_MachineUUID_NotSame"	
	id="1036"	desc="桌号，座位号不匹配错误"	module="EorrorCode EC_ZhuohaoZuoweihao_NotSame"	
	id="1037"	desc="金币不足"	module="EorrorCode EC_Jinbi_NotEnough"	
	id="1038"	desc="钻石不足"	module="EorrorCode EC_Zuanshi_NotEnough"	
	id="1039"	desc="场景类型错误"	module="EorrorCode EC_AwardData_ERROR"	
	id="1040"	desc="无此礼包错误"	module="EorrorCode EC_Libao_NO"	
	id="1041"	desc="已有此武器类型错误"	module="EorrorCode EC_Have_Weapon"	
	id="1042"	desc="已有相同的UUID"	module="EorrorCode EC_MachineUUID_SAME"	
	id="1043"	desc="无此成功支付记录"	module="EorrorCode EC_Paymoney_NOPay"	
	id="1044"	desc="此支付中数据错误"	module="EorrorCode EC_Paymoney_DataError"	
	id="1045"	desc="此集结号已在PC端登录"	module="EorrorCode EC_PC_JJHLOGIN_ERROR"	
	id="1046"	desc="银行功能关闭"	module="EorrorCode EC_PC_JJH_CLOSEBANK"	
	id="1047"	desc="查询失败"	module="EorrorCode EC_queryCharm_1"	
	id="1048"	desc="兑换失败，您今日兑换已达上限请明日再来。"	module="EorrorCode EC_ExchangeCharm_TooManyTimes"	
	id="1049"	desc="兑换失败"	module="EorrorCode EC_ExchangeCharm_1"	
	id="1050"	desc="兑换失败，您输入的银行密码有误，请重新输入。"	module="EorrorCode EC_ExchangeCharm_2"	
	id="1051"	desc="兑换失败，您的魅力值不足"	module="EorrorCode EC_ExchangeCharm_3"	
	id="1052"	desc="兑换码无效!"	module="database  Exchange_WrongNumber Exchange_WrongSuccess"	
	id="1053"	desc="兑换码已经被使用!"	module="database Exchange_AlreadyUse"	
	id="1054"	desc="该集结号账号已绑定过，请直接登录游戏。"	module="errorcode"	
	id="1055"	desc="该功能正在维护中，暂时无法使用"	module="errorcode"	
	id="1056"	desc="验证码错误，请填写正确验证码。"	module="errorcode register"	
	id="1057"	desc="该手机号已注册，请更换其它手机注册，或者直接登录！"	module="errorcode register"	
	id="1058"	desc="验证码申请次数到达上限！"	module="errorcode register"	
	id="1059"	desc="请在60秒后重新尝试申请！"	module="errorcode register"	
	id="1060"	desc="验证信息过期，请重新登录！"	module="errorcode w88 token invalid"	
	id="1061"	desc="注册失败，请重试！"	module="errorcode regiseter"	
	id="1062"	desc="注册过于频繁，请稍后重试！"	module="errorcode regiseter"	
	id="1063"	desc="该账号已绑定集结号微信公众号，为了您的账号安全，请在集结号公众号内点击 用户中心 获取  手游授权。"	module="errorcode weixin "	
	id="1064"	desc="您今天游戏币消耗已达上限，请适度游戏！"	module="errorcode win_enough"	
	id="1065"	desc="珍珠不足，无法抽奖。
快去打鱼收集更多珍珠吧！"	module="errorcode pearlnot enouth"	
	id="1066"	desc="暂时无法绑定，请退出后注册或登录。"	module="errorcode EC_JJHBind_NO"	
	id="1067"	desc="您还未绑定集结号帐号，无法使用该功能。"	module="errorcode"	
	id="1068"	desc="领取奖励还有剩余，无法转出。"	module="errorcode EC_NO_TRANSFER1"	
	id="1069"	desc="兑换奖励还有剩余，无法转出。"	module="errorcode EC_NO_TRANSFER2"	
	id="1070"	desc="兑换失败，此账号不满足兑换条件。"	module="errorcode EC_ExchangeCharm_4"	
	id="1071"	desc="您当日奖励领取次数已超限！"	module="errorcode EC_LoginAwardLimit"	
	id="1072"	desc="您还有很多金币可以使用，先去玩游戏吧！"	module="errorcode EC_OnlineAwardLimit"	
	id="1073"	desc="要绑定的集结号账号正在游戏登陆中，请先退出集结号账号登录再进行绑定！"	module="errorcode EC_BIND_ONLINE"	
	id="1500"	desc="数据加载中"	module="网络模块"	
	id="1501"	desc="服务器连接失败"	module="网络模块输入有误，请输入4-8为的字母、数字组合。"	
	id="1502"	desc="服务器返回未知错误！"	module="网络模块"	
	id="1503"	desc="网络不稳定，请检查网络连接"	module="网络模块 System_CheckNet"	
	id="1504"	desc="网络不稳定，正在重连服务器"	module="网络模块 System_ReConnect"	
	id="1505"	desc="服务器连接失败，是否重连？"	module="网络模块 Net_Down"	
	id="1506"	desc="服务器维护中，请稍后再试!"	module="网络模块 Net_Maintenance"	
	id="1507"	desc="服务器人数已满，请稍后再试!"	module="网络模块 Net_Full"	
	id="1508"	desc="版本升级了！请更新到最新版本，点击确定直接下载！"	module="loadsens界面文本"	
	id="1509"	desc="网络异常，请稍后再试"	module="抽奖模块"	
	id="1510"	desc="服务器断开连接，是否重连？"	module="网络模块 Net_Down"	
	id="1551"	desc="秒"	module="database   second"	
	id="1552"	desc="倍"	module="database   Multify"	
	id="1553"	desc="钻石"	module="database   Diamond"	
	id="1554"	desc="金币"	module="database   Coin"	
	id="1555"	desc="明日再来"	module="database   Tomorrow"	
	id="1556"	desc="点击领取"	module="database   ClickToGet"	
	id="1557"	desc="用户名错误或不存在"	module="database JJH_Login_WrongName"	
	id="1558"	desc="我"	module="database me"	
	id="1559"	desc="昵称"	module="公用标题文本"	
	id="1560"	desc="手机"	module="公用标题文本"	
	id="1561"	desc="性别"	module="公用标题文本"	
	id="1562"	desc="男"	module="公用标题文本"	
	id="1563"	desc="女"	module="公用标题文本"	
	id="1564"	desc="手机端每天都有登录奖励哦！亲！记着来领取哦！"	module="登录奖励随机提示（集结号账号）"	
	id="1565"	desc="在银行中可以查看自己的银行金币和魅力值哦！"	module="登录奖励随机提示（集结号账号）"	
	id="1566"	desc="高等级的场景中有更高倍率的鱼哦！"	module="登录奖励随机提示（集结号账号）"	
	id="1567"	desc="更多好玩游戏，尽在集结号游戏中心！"	module="登录奖励随机提示（集结号账号）"	
	id="1568"	desc="更多精品棋牌游戏都在集结号游戏中心等你来玩哦！"	module="登录奖励随机提示（游客账号）"	
	id="1569"	desc="注册集结号帐号，让帐号更安全！"	module="登录奖励随机提示（游客账号）"	
	id="1570"	desc="请认准我们的官网，谨防山寨钓鱼网站！"	module="登录奖励随机提示（游客账号）"	
	id="1571"	desc="您还是游客玩家，为了您的帐号安全建议您请先注册集结号帐号。"	module="登录奖励随机提示(游客账号)"	
	id="1572"	desc="获得VIP{0}加成{1}%"	module="登录奖励加成提示"	
	id="1573"	desc="明天奖励{0}金币
记着来拿哦！！"	module="登录奖励加成提示"	
	id="1574"	desc="是否进入[06ef00]新手引导[-]？"	module="新手引导显示文本"	
	id="1575"	desc="欢迎来到海洋世界，
箭头指的就是您的[06ef00]炮台[-]位置"	module="新手引导显示文本"	
	id="1576"	desc="箭头位置就是您拥有的
[06ef00]金币[-]和[06ef00]钻石[-]数量"	module="新手引导显示文本"	
	id="1577"	desc="现在，让我们[06ef00]升级武器[-]吧！"	module="新手引导显示文本"	
	id="1578"	desc="[06ef00]炮台等级[-]越高威力越大
消耗的[06ef00]金币数量[-]也越多"	module="新手引导显示文本"	
	id="1579"	desc="让我们试试[06ef00]道具[-]，点击箭头所指的
[06ef00]锁定道具[-]"	module="新手引导显示文本"	
	id="1580"	desc="目标已经锁定，现在尝试一下
[06ef00]切换[-]，试试点击[06ef00]其他鱼群[-]"	module="新手引导显示文本"	
	id="1581"	desc="恭喜您完成全部引导，现在
让我们去[06ef00]领取奖励[-]吧！"	module="新手引导显示文本"	
	id="2001"	desc="账号不能为空"	module="集结号登录模块（客户端校验）"	
	id="2002"	desc="密码不能为空"	module="集结号登录模块（客户端校验）"	
	id="2003"	desc="该账号已被注册或输入有误，请重新输入4-8位的字母、数字组合。"	module="集结号登录模块（客户端校验）（服务器返回）JJH_Reg_WrongName JJH_Login_CheckWrongName"	
	id="2004"	desc="您的密码输入有误，请输入6-12位的字母、数字组合。"	module="集结号登录模块（客户端校验）（服务器返回）JJH_Reg_WrongPass JJH_Login_WrongPass"	
	id="2005"	desc="您的账号已在其他设备登录，请勿重复登录。"	module="集结号登录模块（服务器返回）"	
	id="2006"	desc="您的账号存在异常，已被系统冻结，请联系客服查询。"	module="集结号登录模块（服务器返回）"	
	id="2007"	desc="账号或密码不正确，请重新输入。"	module="集结号登录模块（服务器返回）"	
	id="2008"	desc="该集结号账号已绑定用户，可以直接使用该账号登录。"	module=""	
	id="2009"	desc="该账号已被注册，请重新输入。"	module="集结号注册模块"	
	id="2010"	desc="确认密码输入不同"	module="集结号注册模块 JJH_Reg_WrongPassEqual"	
	id="2011"	desc="昵称不能为空"	module="用户信息模块（客户端校验）"	
	id="2012"	desc="选择了未知头像！"	module="用户信息模块（服务器返回）"	
	id="2013"	desc="输入的昵称非法！"	module="用户信息模块（服务器返回）（客户端校验）"	
	id="2014"	desc="请输入正确手机号。"	module="用户信息模块（服务器返回）（客户端校验）JJH_Reg_WrongPhone"	
	id="2015"	desc="填写真实信息有助于账号安全哟！"	module="用户信息模块（客户端提示）"	
	id="2016"	desc="用户名不足4位"	module="登录模块（服务器端）（客户端提示）  JJH_Login_CheckWrongName"	
	id="2017"	desc="输入错误，昵称为4-12位字符或汉字。"	module="JJH_WrongNickName"	
	id="2018"	desc="请输入要绑定的集结号账号"	module="用户绑定信息提示"	
	id="2019"	desc="卡号"	module="用户绑定信息提示"	
	id="2020"	desc="密码"	module="用户绑定信息文本"	
	id="2021"	desc="如果您还没有集结号账号
请登录我们的官网注册！"	module="用户绑定信息文本"	
	id="2022"	desc="使用中"	module="用户绑定信息文本"	
	id="2023"	desc="包含敏感字符，请重新输入！"	module="用户绑定信息文本"	
	id="2024"	desc="亲，您还是游客哦！
建议您绑定集结账号更多游戏等您来玩！"	module="用户绑定信息文本"	
	id="2025"	desc="账号绑定成功"	module="用户绑定信息文本"	
	id="2026"	desc="已携带金币："	module="财富银行标题"	
	id="2027"	desc="银行内金币："	module="财富银行标题"	
	id="2028"	desc="存取金币："	module="财富银行标题"	
	id="2029"	desc="银行密码："	module="财富银行标题"	
	id="2030"	desc="我的魅力："	module="魅力兑换标题"	
	id="2031"	desc="今日剩余兑换次数："	module="魅力兑换标题"	
	id="2032"	desc="兑换魅力"	module="魅力兑换标题"	
	id="2033"	desc="银行密码"	module="魅力兑换标题"	
	id="2034"	desc="点击输入存取金额"	module="财富银行输入框提示文本"	
	id="2035"	desc="点击输入银行密码"	module="财富银行输入框提示文本"	
	id="2036"	desc="点击输入兑换数量"	module="财富银行输入框提示文本"	
	id="2037"	desc="为了你的账号安全，请先设置银行密码！"	module="财富银行输入框提示文本"	
	id="2038"	desc="确认密码:"	module="财富银行输入框提示文本"	
	id="2039"	desc="请输入6-14位英文数字"	module="财富银行输入框提示文本"	
	id="2040"	desc="请确认要设置的银行密码"	module="财富银行输入框提示文本"	
	id="2041"	desc="[55e7f8]小提示：银行密码请不同于登录密码，银行密码非常重要请牢记！[-]"	module="财富银行输入框提示文本"	
	id="2042"	desc="两次密码输入不一致，请重新确认密码。"	module="财富银行输入框提示文本"	
	id="2043"	desc="请输入需要设置的银行密码。"	module="财富银行输入框提示文本"	
	id="2044"	desc="设置密码请输入6-14位英文数字"	module="财富银行输入框提示文本"	
	id="2045"	desc="密码设置成功，请牢记哦！"	module="财富银行输入框提示文本"	
	id="2046"	desc="密码设置失败，请重试"	module=""	
	id="2051"	desc="您的帐号已登录PC端游戏，请退出后再使用银行功能！"	module="财富模块（服务器返回）"	
	id="2052"	desc="您的帐号存在异常无法使用该功能，请咨询在线客服：800090166！"	module="财富模块（服务器返回）"	
	id="2053"	desc="抱歉，该功能正在维护。请稍后再试。"	module="财富模块（服务器返回）"	
	id="2054"	desc="每次存入金币数量需大于5万金币哦！"	module="财富模块（服务器返回）（客户端校验）GoldLess50000"	
	id="2055"	desc="您的金币不足，存入失败。"	module="财富模块（服务器返回）（客户端校验）PackageNotEnough"	
	id="2056"	desc="存入失败，您输入金币数量有误请重新输入。"	module="财富模块（服务器返回）（客户端校验）"	
	id="2057"	desc="存入金币成功。"	module="财富模块（服务器返回）Depositok"	
	id="2058"	desc="取出失败，您银行内金币不足。"	module="财富模块（服务器返回）（客户端校验）FetchGold_Wrong BankNotEnough"	
	id="2059"	desc="取出失败，请输入正确银行密码。"	module="财富模块（服务器返回）（客户端校验）passWordEorro"	
	id="2060"	desc="取出金币成功。"	module="财富模块（服务器返回）FetchMoneyOk"	
	id="2061"	desc="取出金币的数量不能为空"	module="（客户端校验）GoldNotNull"	
	id="2062"	desc="密码不能为空"	module="（客户端校验）Charm_MissPass PasswordNotNull"	
	id="2063"	desc="兑换失败，您的魅力值不足。"	module="财富模块（服务器返回）（客户端校验）Charm_ValueNotEnough CharmNotEnough"	
	id="2064"	desc="兑换失败，您今日魅力兑换次数已用完。"	module="财富模块（服务器返回）（客户端校验）"	
	id="2065"	desc="每次魅力兑换最大为100魅力哦！"	module="财富模块（服务器返回）（客户端校验）Charm_ValueLimit CharmThen100"	
	id="2066"	desc="兑换失败，您输入金币数量有误请重新输入"	module="财富模块（服务器返回）（客户端校验）"	
	id="2067"	desc="兑换成功，请在财富银行中查看哦！"	module="财富模块（服务器返回）Charm_Success Exchangeok"	
	id="2068"	desc="小提示：存入金币无需密码，每次存入金币数量最低为5万哦！（7*24小时客服QQ：800090166）"	module="财富银行Tips"	
	id="2069"	desc="小提示：魅力兑换成功后，金币将保存在银行中哦！（7 * 24小时客服QQ：800090166）"	module="魅力兑换tips"	
	id="2070"	desc="您当前为游客账号无法使用魅力兑换。"	module="财富模块（服务器返回）Charm_NotBind"	
	id="2071"	desc="您还未绑定集结号帐号，无法使用该功能。"	module="财富模块(客户端)"	
	id="2072"	desc="魅力兑换值不能为空"	module="（客户端校验）CharmValueNotNull"	
	id="2073"	desc="魅力值输入有误请重新输入。"	module="财富模块Charm_ValueError"	
	id="2074"	desc="取出失败！"	module="取金币通用提示"	
	id="2075"	desc="存入失败！"	module="存金币不特指失败"	
	id="2076"	desc="无人上榜"	module="排行榜模块"	
	id="2077"	desc="第{0}名"	module="排行榜模块"	
	id="2078"	desc="未上榜"	module="排行榜模块"	
	id="2079"	desc="玩家昵称:  {0}"	module="排行榜模块"	
	id="2080"	desc="一，二，三，四，五，六，七，八，九，十，十一，十二，十三，十四，十五，十六，十七，十八，十九，二十"	module="排行榜模块"	
	id="2081"	desc="金币排行榜每日中午12点刷新！
金币数量排名前20的玩家将荣登金币榜！奖励直接发放到玩家银行哦！"	module="排行榜模块"	
	id="2082"	desc="实力榜展示每日获得金币最多的前20位玩家！
榜单每天中午12点刷新！奖励直接发放到玩家银行哦！"	module="排行榜模块"	
	id="2083"	desc="我的排名："	module="排行榜模块"	
	id="3001"	desc="购买金币成功，获得@"	module="商城模块 PaymoneyOK_coin"	
	id="3002"	desc="购买钻石成功，获得@"	module="商城模块 PaymoneyOK_diamond"	
	id="3003"	desc="购买成功，获得@"	module="商城模块 PaymoneyOK_package"	
	id="3004"	desc="购买成功，获得畅玩礼包。"	module="商城模块 PaymoneyOK_package1"	
	id="3005"	desc="购买成功，获得土豪礼包。"	module="商城模块 PaymoneyOK_package2"	
	id="3006"	desc="您目前还是游客账号，绑定后购买更安全！"	module="商城模块 GuestBuyItem"	
	id="3007"	desc="您的每日充值额度已达上限！"	module="商城模块 GuestBuyItem"	
	id="3008"	desc="商城数据已更新，请重新购买！"	module="商城模块 GuestBuyItem"	
	id="3050"	desc="这几条鱼看起来有点特殊，赶紧用锁定技能将其捕获吧!"	module="新手引导 NewGuide_1"	
	id="3051"	desc="太棒了，赚了好多钻石，我们赶紧去升级炮台吧!"	module="新手引导 NewGuide_2"	
	id="3052"	desc="您已经掌握捕鱼技能了，赶紧开始你的捕鱼旅程吧!"	module="新手引导 NewGuide_3"	
	id="3053"	desc="欢迎来到海洋世界！
我们为您准备了50钻石！请收下吧！"	module="新手引导 "	
	id="4001"	desc="充值任意金额就能开启VIP1哦！"	module="vip模块 Vip0"	
	id="4002"	desc="充值@元马上享受更高特权！"	module="vip模块 Vip1"	
	id="4003"	desc="充值@元马上享有顶级特权！"	module="vip模块 Vip2"	
	id="4004"	desc="VIP"	module="显示文本"	
	id="4005"	desc="非VIP无加成"	module="显示文本"	
	id="4006"	desc="ID:{0}"	module="个人信息UI显示文本"	
	id="4007"	desc="ID:YK{0}"	module="个人信息UI显示文本"	
	id="4008"	desc="级"	module="个人信息UI显示文本"	
	id="4009"	desc="LV"	module="用户信息面板"	
	id="4010"	desc="YK"	module="middlescens 更新用户信息"	
	id="4011"	desc="提示"	module="转盘显示文本"	
	id="4012"	desc="真遗憾
与1888888金币插肩而过。"	module="转盘显示文本"	
	id="4014"	desc="恭喜!获得"	module="转盘显示文本后面带物品名称"	
	id="4015"	desc="恭喜玩家"	module="转盘显示文本后面带物品名称"	
	id="4016"	desc="抽奖获得"	module="转盘显示文本后面带物品名称"	
	id="4017"	desc="鱼潮进行中，请稍后"	module="鱼潮 tideFishManger提示文本"	
	id="4018"	desc="礼包"	module="显示文本"	
	id="4500"	desc="已达到最高等级！"	module="武器模块 Weapon_LevelMax"	
	id="4501"	desc="最大炮值增强@经验加成时间增强@加速时间增强"	module="武器模块Weapon_LevelUp"	
	id="4502"	desc="确定花费s% 钻石购买w%吗？"	module="武器模块（购买提示）Comfirm_BuyWeapon"	
	id="4503"	desc="钻石不足！"	module="武器模块"	
	id="4504"	desc="最大炮值："	module="武器模块"	
	id="4505"	desc="加速："	module="武器模块"	
	id="4506"	desc="经验："	module="武器模块"	
	id="4507"	desc="自动："	module="武器模块"	
	id="4508"	desc="未购买"	module="武器模块"	
	id="4509"	desc="炮值提升："	module="武器模块"	
	id="4510"	desc="加速升级："	module="武器模块"	
	id="4511"	desc="经验升级："	module="武器模块"	
	id="4512"	desc="自动升级："	module="武器模块"	
	id="4600"	desc="经验加成"	module="道具名称"	
	id="4601"	desc="自动发炮"	module="道具名称"	
	id="4602"	desc="锁定"	module="道具名称"	
	id="4603"	desc="加速"	module="道具名称"	
	id="4604"	desc="我们的客服会24小时为您提供帮助，
充值订单可能有一些延迟请您稍后重新登陆查看。"	module="联系我们显示文本"	
	id="4605"	desc="自动"	module="道具名称"	
	id="4606"	desc="经验"	module="道具名称"	
	id="4607"	desc="道具"	module="道具分类"	
	id="4614"	desc="金币"	module=""	
	id="4615"	desc="财富的象征！用于鱼炮发射！"	module=""	
	id="4616"	desc="钻石"	module=""	
	id="4617"	desc="闪闪的钻石！可以购买和升级更高的鱼炮！"	module=""	
	id="4618"	desc="珍珠"	module=""	
	id="4619"	desc="稀有的海底珍珠，可以兑换和参与更多活动哦！"	module=""	
	id="4608"	desc="直接锁定大鱼！让炮弹百发百中！"	module="道具描述"	
	id="4609"	desc="一段时间内自动发射炮弹！释放你双手！"	module="道具描述"	
	id="4610"	desc="获得更高的经验加成！升级快起飞！"	module="道具描述"	
	id="4611"	desc="获得更快的炮弹发射速度！大鱼往哪跑！"	module="道具描述"	
	id="4612"	desc="请期待更多新道具!"	module="道具描述"	
	id="4613"	desc="神秘道具"	module="道具描述"	
	id="4620"	desc="黄金宝箱"	module="道具名称"	
	id="4621"	desc="罕见的黄金宝箱！打爆后可以获得超能炮！"	module="道具描述"	
	id="4622"	desc="白银宝箱"	module="道具名称"	
	id="4623"	desc="稀有的白银宝箱！打爆后可以获得钻头炮！"	module="道具描述"	
	id="4624"	desc="超级攻击，可以打出爆炸伤害。也可赠送给其他玩家。"	module=""	
	id="4625"	desc="攻击可以穿透各种鱼群。也可赠送给其他玩家。"	module=""	
	id="4626"	desc="黄铜鱼雷"	module="道具名称"	
	id="4627"	desc="黄铜打造的普通鱼雷；
游戏中使用可以获得少量金币。"	module="道具描述"	
	id="4628"	desc="白银鱼雷"	module="道具名称"	
	id="4629"	desc="白银制作的中级鱼雷；
游戏中使用可以获得一些金币。"	module="道具描述"	
	id="4630"	desc="黄金鱼雷"	module="道具名称"	
	id="4631"	desc="黄金打造价值昂贵的高级鱼雷；
游戏中使用可以获得很多金币。"	module="道具描述"	
	id="4632"	desc="紫金鱼雷"	module="道具名称"	
	id="4633"	desc="紫金铸造非常罕见的稀有鱼雷；
游戏中使用可以获得大量金币。"	module="道具描述"	
	id="4634"	desc="七彩鱼雷"	module="道具名称"	
	id="4635"	desc="七彩陨石铸造的传说级鱼雷；
游戏中使用可以获得超多金币。"	module="道具描述"	
	id="4650"	desc="火炮"	module="武器名称"	
	id="4651"	desc="穿透炮"	module="武器名称"	
	id="4652"	desc="速射炮"	module="武器名称"	
	id="4653"	desc="散射炮"	module="武器名称"	
	id="4654"	desc="神秘武器"	module="武器名称"	
	id="4655"	desc="超能炮"	module="武器名称"	
	id="4656"	desc="钻头炮"	module="武器名称"	
	id="4700"	desc="最大炮值增强"	module="属性名"	
	id="4701"	desc="经验加成时间增强"	module="属性名"	
	id="4702"	desc="自动发炮时间增强"	module="属性名"	
	id="4703"	desc="锁定时间增强"	module="属性名"	
	id="4704"	desc="加速时间增强"	module="属性名"	
	id="4710"	desc="一种普通的新手武器。
最高炮值：19999
价格:免费"	module="炮的长描述"	
	id="4711"	desc="炮弹可以穿透攻击到下一条鱼
最高炮值：19999
价格：300红钻"	module="炮的长描述"	
	id="4712"	desc="炮弹的发射速度超快！
最高炮值：19999
价格：600红钻"	module="炮的长描述"	
	id="4713"	desc="能同时发射三枚炮弹！
最高炮值：19999
价格：900红钻"	module="炮的长描述"	
	id="4714"	desc="最大炮值"	module="升级信息"	
	id="4715"	desc="经验加成"	module="升级信息"	
	id="4716"	desc="加速"	module="升级信息"	
	id="4717"	desc="自动发炮"	module="升级信息"	
	id="4718"	desc="请期待更多的武器！"	module="炮的长描述"	
	id="4719"	desc="一种普通的新手武器。"	module="炮的短描述"	
	id="4720"	desc="炮弹可以穿透攻击到下一条鱼"	module="炮的短描述"	
	id="4721"	desc="炮弹的发射速度超快！"	module="炮的短描述"	
	id="4722"	desc="能同时发射三枚炮弹！"	module="炮的短描述"	
	id="4723"	desc="不同的武器拥有不同的功能。每种武器都可以升级哦！最大炮值都是19999！"	module="我的武器的tips"	
	id="4724"	desc="升级后"	module="升级信息"	
	id="4725"	desc="当前炮值"	module="升级信息"	
	id="4726"	desc="升级炮值："	module="升级信息"	
	id="4727"	desc="赠送"	module="升级信息"	
	id="4800"	desc="珍珠"	module="珍珠兑换中心模块"	
	id="4801"	desc="兑换成功"	module="珍珠兑换中心模块"	
	id="4802"	desc="兑换失败"	module="珍珠兑换中心模块"	
	id="4803"	desc="兑换物品"	module="珍珠兑换中心模块"	
	id="4804"	desc="消耗珍珠"	module="珍珠兑换中心模块"	
	id="4805"	desc="兑换时间"	module="珍珠兑换中心模块"	
	id="4806"	desc="兑换状态"	module="珍珠兑换中心模块"	
	id="4807"	desc="暂无兑换记录"	module="珍珠兑换中心模块"	
	id="4808"	desc="请填写需要充值的手机号"	module="珍珠兑换中心模块"	
	id="4809"	desc="两次输入的手机号码不一致，请重新输入！"	module="珍珠兑换中心模块"	
	id="4810"	desc="实物兑换成功后，将由我们客服人员依次在1-3个工作日为您发放。兑换状态请在兑换记录中查看。
若有疑问请咨询客服QQ：800-090-166"	module="珍珠兑换中心模块"	
	id="4811"	desc="系统将为您保存两周内的记录，若想查看更多请联系在线客服。
若有疑问请咨询客服QQ：800-090-166"	module="珍珠兑换中心模块"	
	id="4812"	desc="兑换中"	module="珍珠兑换中心模块"	
	id="4813"	desc="20"	module="珍珠兑换中心模块"	
	id="4814"	desc="30"	module="珍珠兑换中心模块"	
	id="4815"	desc="50"	module="珍珠兑换中心模块"	
	id="4816"	desc="100"	module="珍珠兑换中心模块"	
	id="4817"	desc="20元充值卡，兑换后直接充值到账哦！"	module="珍珠兑换中心模块"	
	id="4818"	desc="金币加油包"	module="珍珠兑换中心模块"	
	id="4819"	desc="红钻加油包"	module="珍珠兑换中心模块"	
	id="4820"	desc="金币加油箱"	module="珍珠兑换中心模块"	
	id="4821"	desc="红钻加油箱"	module="珍珠兑换中心模块"	
	id="4822"	desc="兑换成功获得16万金币哦！"	module="珍珠兑换中心模块"	
	id="4823"	desc="兑换成功获得800红钻哦！"	module="珍珠兑换中心模块"	
	id="4824"	desc="手机号"	module="珍珠兑换中心模块"	
	id="4825"	desc="点击输入手机号"	module="珍珠兑换中心模块"	
	id="4826"	desc="确认手机号"	module="珍珠兑换中心模块"	
	id="4827"	desc="银行密码："	module="珍珠兑换中心模块"	
	id="4828"	desc="点击输入银行密码"	module="珍珠兑换中心模块"	
	id="4829"	desc="兑换异常，请重新尝试兑换！"	module="珍珠兑换中心模块"	
	id="4830"	desc="兑换成功"	module="珍珠兑换中心模块"	
	id="4831"	desc="商品已兑换完，请兑换其它物品！"	module="珍珠兑换中心模块"	
	id="4832"	desc="珍珠不足，无法兑换！"	module="珍珠兑换中心模块"	
	id="4833"	desc="商品已经下架，请兑换其它商品！"	module="珍珠兑换中心模块"	
	id="4834"	desc="银行密码错误，请重新输入"	module="珍珠兑换中心模块"	
	id="4835"	desc="今日剩余：{0}"	module="珍珠兑换中心模块"	
	id="4836"	desc="30元充值卡，兑换后直接充值到账哦！"	module="珍珠兑换中心模块"	
	id="4837"	desc="50元充值卡，兑换后直接充值到账哦！"	module="珍珠兑换中心模块"	
	id="4839"	desc="100元充值卡，兑换后直接充值到账哦！"	module="珍珠兑换中心模块"	
	id="4840"	desc="兑换成功获得85万金币哦！"	module="珍珠兑换中心模块"	
	id="4841"	desc="兑换成功获得4000红钻哦！"	module="珍珠兑换中心模块"	
	id="4842"	desc="游戏中使用可以获得一些金币。"	module="珍珠兑换中心模块"	
	id="4843"	desc="游戏中使用可以获得很多金币"	module="珍珠兑换中心模块"	
	id="4844"	desc="登录密码"	module="珍珠兑换中心模块"	
	id="4900"	desc="蝴蝶鱼"	module="鱼的名称"	
	id="4901"	desc="鳕鱼"	module="鱼的名称"	
	id="4902"	desc="金枪鱼"	module="鱼的名称"	
	id="4903"	desc="黄金鱼"	module="鱼的名称"	
	id="4904"	desc="小丑鱼"	module="鱼的名称"	
	id="4905"	desc="凤尾鱼"	module="鱼的名称"	
	id="4906"	desc="尖嘴鱼"	module="鱼的名称"	
	id="4907"	desc="旗鱼"	module="鱼的名称"	
	id="4908"	desc="海龟"	module="鱼的名称"	
	id="4909"	desc="狮子鱼"	module="鱼的名称"	
	id="4910"	desc="河豚"	module="鱼的名称"	
	id="4911"	desc="灯笼鱼"	module="鱼的名称"	
	id="4912"	desc="蝙蝠鱼"	module="鱼的名称"	
	id="4913"	desc="鲨鱼"	module="鱼的名称"	
	id="4914"	desc="独角鲸"	module="鱼的名称"	
	id="4915"	desc="黄金魔鬼鱼"	module="鱼的名称"	
	id="4916"	desc="黄金鲨鱼"	module="鱼的名称"	
	id="4917"	desc="李逵"	module="鱼的名称"	
	id="4918"	desc="局部炸弹"	module="特殊鱼的名称"	
	id="4919"	desc="全屏炸弹"	module="特殊鱼的名称"	
	id="4920"	desc="一网打尽"	module="特殊鱼的名称"	
	id="4921"	desc="铁索连环"	module="特殊鱼的名称"	
	id="4922"	desc="神秘鱼"	module="特殊鱼的名称"	
	id="4923"	desc="爆炸后对自身周围的其他鱼造成致命伤害!"	module="特殊鱼的属性"	
	id="4924"	desc="爆炸后对屏幕内的所有鱼造成致命伤害!"	module="特殊鱼的属性"	
	id="4925"	desc="对屏幕中其他同种类的鱼造成致命伤害!"	module="特殊鱼的属性"	
	id="4926"	desc="三条鱼紧紧链接在一起，打中一条获得3倍奖励！"	module="特殊鱼的属性"	
	id="4927"	desc="请期待更多新鱼!"	module="特殊鱼属性"	
	id="4928"	desc="鱼群"	module="鱼分类"	
	id="4929"	desc="中鱼"	module="鱼分类"	
	id="4930"	desc="大鱼"	module="鱼分类"	
	id="4931"	desc="特殊鱼"	module="鱼分类"	
	id="4932"	desc="鲶鱼"	module="鱼的名称"	
	id="4933"	desc="海豚"	module="鱼的名称"	
	id="4934"	desc="锤头鲨"	module="鱼的名称"	
	id="4935"	desc="龙"	module="鱼的名称"	
	id="4936"	desc="鮟鱇"	module="鱼的名称"	
	id="4937"	desc="鳄鱼"	module="鱼的名称"	
	id="4938"	desc="螃蟹"	module="鱼的名称"	
	id="4939"	desc="章鱼"	module="鱼的名称"	
	id="4940"	desc="因其凶残的长相与习性而被成为“海鬼”"	module="鱼的描述"	
	id="4941"	desc="拥有悠久历史的“水中一霸”"	module="鱼的描述"	
	id="4942"	desc="大部分时间都在寻找各类食物的“吃货蟹王”"	module="鱼的描述"	
	id="4943"	desc="块头很大，性子胆小“深海巨怪”"	module="鱼的描述"	
	id="4999"	desc="剩余珍珠："	module="抽奖模块"	
	id="5000"	desc="一小袋金币"	module="抽奖模块"	
	id="5001"	desc="一中袋金币"	module="抽奖模块"	
	id="5002"	desc="一大袋金币"	module="抽奖模块"	
	id="5003"	desc="一小箱金币"	module="抽奖模块"	
	id="5004"	desc="一中箱金币"	module="抽奖模块"	
	id="5005"	desc="一大箱金币"	module="抽奖模块"	
	id="5006"	desc="一小袋钻石"	module="抽奖模块"	
	id="5007"	desc="100"	module="抽奖模块"	
	id="5008"	desc="一中袋钻石"	module=""	
	id="5009"	desc="一大袋钻石"	module=""	
	id="5010"	desc="一小箱钻石"	module=""	
	id="5011"	desc="一中箱钻石"	module=""	
	id="5012"	desc="一大箱钻石"	module=""	
	id="5013"	desc="畅玩礼包"	module="商城礼包名称"	
	id="5014"	desc="土豪礼包"	module="商城礼包名称"	
	id="5500"	desc="经验额外加成{0}"	module="vip模块"	
	id="5501"	desc="登录奖励金币额外加成{0}"	module="vip模块"	
	id="5502"	desc="在线时长奖励额外加成{0}"	module="vip模块"	
	id="5503"	desc="VIP专属标识，vip模块"	module="vip模块"	
	id="5504"	desc="10%"	module=""	
	id="5505"	desc="20%"	module=""	
	id="5510"	desc="经验额外加成10%;登录奖励金币额外加成20%;VIP专属标识"	module=""	
	id="5511"	desc="经验额外加成20%;登录奖励金币额外加成50%;在线时长奖励额外加成50%;VIP专属标识"	module=""	
	id="5512"	desc="经验额外加成40%;登录奖励金币额外加成100%;在线时长奖励额外加成50%;VIP专属标识"	module=""	
	id="5513"	desc="220000;1000;x2"	module="礼品内容"	
	id="5514"	desc="450000;2000;x1"	module="礼品内容"	
	id="5515"	desc="金币;钻石;自动发炮"	module="礼品标题"	
	id="5516"	desc="金币;钻石;速射炮"	module="礼品标题"	
	id="5517"	desc="200000;1000;x2"	module=""	
	id="5518"	desc="350000;2000;x1"	module=""	
	id="5601"	desc="我的邮箱ID:m{0}"	module="邮箱模块"	
	id="5602"	desc="{0}/20 邮件最多保存7天"	module="邮箱模块"	
	id="5603"	desc="没有新的邮件…"	module="邮箱模块"	
	id="5604"	desc="个"	module="邮箱模块"	
	id="5605"	desc="请在左侧点击要查看的邮件"	module="邮箱模块"	
	id="5606"	desc="领取附件失败"	module="邮箱模块"	
	id="5607"	desc="领取成功"	module="邮箱模块"	
	id="5608"	desc="道具不存在"	module="邮箱模块"	
	id="5609"	desc="道具不足"	module="邮箱模块"	
	id="5610"	desc="附件已领取"	module="邮箱模块"	
	id="5650"	desc="[248bdb]请确认，是否将[-][ff7a19]{0}[-][248bdb]个[-][ff7a19]{1}[-] [248bdb]赠送给下面的玩家[-]
[248bdb]玩家昵称：[-][ff7a19]{2}[-]
[248bdb]玩家邮箱编号：[-][ff7a19]{3}[-]
[248bdb]玩家集结号ID：[-][ff7a19]{4}[-] "	module="superWeapon 模块"	
	id="5651"	desc="请输入正确的邮箱ID"	module="superWeapon 模块"	
	id="5652"	desc="请输入赠送数量"	module="superWeapon 模块"	
	id="5653"	desc="出售一个{0}可以获得{1}万金币哦！"	module="superWeapon 模块"	
	id="5654"	desc="请输入正确的邮箱编号，邮箱编号请在邮箱界面内查看！"	module="superWeapon 模块"	
	id="5655"	desc="请输入出售数量"	module="superWeapon 模块"	
	id="5656"	desc="道具不足"	module="superWeapon 模块"	
	id="5657"	desc="道具不存在"	module="superWeapon 模块"	
	id="5658"	desc="失败"	module="superWeapon 模块"	
	id="5659"	desc="出售成功"	module="superWeapon 模块"	
	id="5660"	desc="赠送成功"	module="superWeapon 模块"	
	id="5661"	desc="交易失败"	module="superWeapon 模块"	
	id="5700"	desc="抵制不良游戏，拒绝盗版游戏。注意自我保护，谨防上当受骗。适度游戏益脑，沉迷游戏伤身。合理安排时间，享受健康生活。"	module="登陆模块和注册模块"	
	id="5701"	desc="集结号账号可以直接登录哦！每天都有大量免费金币，记得来领取哦！"	module="登陆模块和注册模块"	
	id="5702"	desc="还没有帐号？马上注册！"	module="登陆模块和注册模块"	
	id="5703"	desc="点击输入账号"	module="登陆模块和注册模块"	
	id="5704"	desc="点击输入密码"	module="登陆模块和注册模块"	
	id="5705"	desc="点击输入电话号码"	module="登陆模块和注册模块"	
	id="5706"	desc="点击输入验证码"	module="登陆模块和注册模块"	
	id="5707"	desc="注册账号可以使账号更安全哦！"	module="登陆模块和注册模块"	
	id="5708"	desc="获取验证码失败，稍后请重新获取！"	module="登陆模块和注册模块"	
	id="5709"	desc="注册码获取成功，请注意在手机端查收！"	module="登陆模块和注册模块"	
	id="5710"	desc="该手机号已经注册！"	module="登陆模块和注册模块"	
	id="5711"	desc="请输入6位短信验证码"	module="登陆模块和注册模块"	
	id="5712"	desc="手机号"	module="登陆模块和注册模块"	
	id="5713"	desc="密码"	module="登陆模块和注册模块"	
	id="5714"	desc="验证码"	module="登陆模块和注册模块"	
	id="5715"	desc="请填写6-24个字符，并区分大小写"	module="登陆模块和注册模块"	
	id="5716"	desc="您的手机号就是登录账号哦"	module="登陆模块和注册模块"	
	id="5717"	desc="新广出审[2016]1745号    浙新广[2016]395号   出版物号ISBN978-7-7979-0542-8"	module="登陆模块和注册模块"	
	id="5750"	desc="我是不是看起来萌萌哒！"	module="聊天模块常用语"	
	id="5751"	desc="我有钱，我任性！"	module="聊天模块常用语"	
	id="5752"	desc="我才是真正的捕鱼达人！"	module="聊天模块常用语"	
	id="5753"	desc="有缘千里来相会，交个朋友吧！"	module="聊天模块常用语"	
	id="5754"	desc="放开那鲨鱼，让我来！"	module="聊天模块常用语"	
	id="5755"	desc="哈哈，这下赚大了！"	module="聊天模块常用语"	
	id="5756"	desc="嘿嘿，总算被我逮着条大的！"	module="聊天模块常用语"	
	id="5757"	desc="这么小的鱼，回家炖汤都不够啊！"	module="聊天模块常用语"	
	id="5758"	desc="为何大鱼总是与我插肩而过？"	module="聊天模块常用语"	
	id="5759"	desc="屠龙宝刀在我手，小样看你往哪游！"	module="聊天模块常用语"	
	id="5760"	desc="青山不改绿水长流，大家后会有期！"	module="聊天模块常用语"	
	id="5761"	desc="非VIP用户不能发言"	module="聊天模块提醒"	
	id="5762"	desc="还有{0}秒可以发言，VIP3用户无限制"	module="聊天模块提醒"	
	id="5763"	desc="[d65420]每发送一条消息消耗[-][ff0000]1000[-][d65420]金币[-]"	module="聊天模块提醒"	
	id="5764"	desc="必须大于等于10级才能收发信息"	module=""	
	id="5800"	desc="元"	module="支付选择模块"	
	id="5900"	desc="手机未安装微信，请安装后重试!"	module="sdk模块"	
	id="5901"	desc="服务器繁忙，稍后再试!"	module="sdk模块"	
	id="5902"	desc="QQ登录失败，请稍后重试！"	module="sdk模块"	
	id="5903"	desc="QQ登录网路异常，请稍后重试！"	module="sdk模块"	
	id="5904"	desc="未安装QQ，请先安装QQ！"	module="sdk模块"	
	id="5905"	desc="QQ版本太低，请先升级！"	module="sdk模块"	
	id="5906"	desc="微信登录失败，请稍后重试！"	module="sdk模块"	
	id="5907"	desc="本地数据异常，请稍后重试！"	module="sdk模块"	
	id="5908"	desc="用户未实名认证，请先实名认证！"	module="sdk模块"	
	id="5909"	desc="登录信息过期，请重新授权登录！"	module="sdk模块"	
	id="5910"	desc="支付参数异常，请稍后重试！"	module="sdk模块"	
	id="5911"	desc="微信版本太低，请先升级！"	module="sdk模块"	
	id="6000"	desc="点击批量赠送，可以一次赠送多个道具！"	module="背包模块"	
	id="6001"	desc="道具名称：{0}
道具数量：{1}
道具说明：{2}"	module="背包模块"	
	id="6002"	desc="点需要赠送的道具，
然后在右侧修改赠送数量。"	module="背包模块"	
	id="6003"	desc="批量赠送每次最多赠送4种道具哦！"	module="背包模块"	
	id="6004"	desc="赠送玩家邮箱：m{0}
赠送玩家昵称：{1}
赠送道具内容：{2}"	module="背包模块"	
	id="6005"	desc="赠送成功，已发送到玩家邮箱！"	module="背包模块"	
	id="6006"	desc="赠送失败，道具不足！"	module="背包模块"	
	id="6007"	desc="点击输入邮箱ID"	module="背包模块"	
	id="6008"	desc="点击输入数量"	module="背包模块"	
	id="6009"	desc="背包里没有道具哦！"	module="背包模块"	
	id="6010"	desc="等级{0}以上的玩家才可以使用赠送功能哦！"	module="背包模块"	
	id="6011"	desc="VIP{0}以上的玩家才可以使用赠送功能哦！"	module="背包模块"	
	id="6012"	desc="邮箱ID"	module="背包模块"	
	id="6013"	desc="数量"	module="背包模块"	
	id="6100"	desc="兑换成功后，请在邮件内领取道具！"	module="兑换码模块"	
	id="6101"	desc="卡号不能为空"	module="兑换码模块"	
	id="6102"	desc="卡号无效"	module="兑换码模块"	
	id="6103"	desc="兑换成功，请稍后在邮件中查阅！"	module="兑换码模块"	
	id="6104"	desc="兑换服务器繁忙，请稍后再试！"	module="兑换码模块"	
	id="6105"	desc="兑换已达上限，无法兑换！"	module="兑换码模块"	
	id="6106"	desc="请输入卡号"	module="兑换码模块"	
	id="6107"	desc="请输入密码"	module="兑换码模块"	
	id="6108"	desc="卡号或者密码错误"	module=""	
	id="6150"	desc="双倍爆率模式下金币消耗双倍；打鱼威力也更大哦！双倍爆率模式状态下不可使用其他道具哦！
是否消耗10000钻石解锁双倍模式？
"	module="双倍模块"	
	id="6151"	desc="是否开启双倍模式？"	module="双倍模块"	
	id="6152"	desc="双倍金币模式下金币消耗双倍；但是也可以获得双倍的金币返还哦！
是否消耗10000钻石解锁双倍模式？"	module=""	
	id="6153"	desc="消耗{0}钻石解锁双倍爆率功能。
双倍爆率模式下金币消耗双倍；打鱼威力也更大哦！双倍爆率模式状态下不可使用其他道具哦！
限时折扣：{1}
是否消耗{2}钻石解锁双倍模式？"	module=""	
	id="6154"	desc="消耗{0}钻石解锁双倍金币功能。
双倍金币开启后金币消耗翻倍，同时获得金币翻倍。
限时折扣：{1}
是否消耗{2}钻石解锁双倍模式？"	module=""	
	id="6160"	desc="客服QQ："	module="联系我们模块"	
	id="6161"	desc="800 090 166"	module="联系我们模块"	
	id="6162"	desc="热线电话："	module="联系我们模块"	
	id="6163"	desc="0579-8291 7777"	module="联系我们模块"	
	id="6170"	desc="音量"	module="设置模块"	
	id="6171"	desc="音效"	module="设置模块"	
	id="6172"	desc="全部静音"	module="设置模块"	
	id="6173"	desc="屏蔽聊天"	module="设置模块"	
	id="6300"	desc="登录QQ授权支付"	module="SDK模块"	
	id="6301"	desc="登录微信授权支付"	module="SDK模块"	
	id="6302"	desc="-----请选择登录方式-----"	module="SDK模块"	
	id="6303"	desc="集结号账号绑定成功，需要重新登录游戏!"	module="SDK微信和QQ登录后的绑定提示"	
	id="6400"	desc="鱼 雷"	module="排行榜模块"	
	id="6401"	desc="实 力"	module="排行榜模块"	
	id="6402"	desc="金币奖励"	module="排行榜模块"	
	id="6403"	desc="一键领取"	module="邮件模块"	
	id="6404"	desc="领取"	module="邮件模块"	
	id="6405"	desc="删除"	module="邮件模块"	
	id="6406"	desc="存入"	module="银行模块"	
	id="6407"	desc="取出"	module="银行模块"	
	id="6408"	desc="每次最大兑换100魅力"	module="银行模块"	
	id="6409"	desc="兑换"	module="兑换模块"	
	id="6410"	desc="道具名称:"	module="背包模块"	
	id="6411"	desc="道具数量:"	module="背包模块"	
	id="6412"	desc="接收人ID:"	module="背包模块"	
	id="6413"	desc="赠送:"	module="背包模块"	
	id="6414"	desc="玩家昵称:"	module="背包模块"	
	id="6415"	desc="不可输入自己邮箱"	module="背包模块"	
	id="6416"	desc="谢谢
参与"	module="欢乐转盘"	
	id="6417"	desc="黄铜
鱼雷"	module="欢乐转盘"	
	id="6418"	desc="金币
8888
"	module="欢乐转盘"	
	id="6419"	desc="金币
18888
"	module="欢乐转盘"	
	id="6420"	desc="黄金
鱼雷"	module="欢乐转盘"	
	id="6421"	desc="100
钻石"	module="欢乐转盘"	
	id="6422"	desc="白银
鱼雷"	module="欢乐转盘"	
	id="6423"	desc="金币
1888888"	module="欢乐转盘"	
	id="6424"	desc="获奖名单"	module="欢乐转盘"	
	id="6425"	desc="抽奖记录"	module="欢乐转盘"	
	id="6426"	desc="取消"	module="对话框"	
	id="6427"	desc="确定"	module="对话框"	
	id="6428"	desc="继续抽奖"	module="欢乐转盘"	
	id="6429"	desc="500金币"	module="签到奖励"	
	id="6430"	desc="700金币"	module="签到奖励"	
	id="6431"	desc="900金币"	module="签到奖励"	
	id="6432"	desc="1200金币"	module="签到奖励"	
	id="6433"	desc="1500金币"	module="签到奖励"	
	id="6434"	desc="1800金币"	module="签到奖励"	
	id="6435"	desc="2500金币"	module="签到奖励"	
	id="6436"	desc="发送"	module="聊天"	
	id="6437"	desc="聊天记录"	module="聊天"	
	id="6438"	desc="常用语"	module="聊天"	
	id="6439"	desc="支付金额"	module="支付方式"	
	id="6440"	desc="是否进行兑换"	module="银行模块"	
	id="6441"	desc="注册"	module="登录模块"	
	id="6442"	desc="登录"	module="登录模块"	
	id="6443"	desc="信息填写不完整"	module="登录模块"	
	id="6444"	desc="金 币"	module="排行榜模块"	
	id="6445"	desc="银行密码"	module="兑换"	
	id="6446"	desc="开"	module="设置模块"	
	id="6447"	desc="关"	module="设置模块"	
	id="6448"	desc="{0}邮件最多保存7天"	module="邮件模块"	
	id="6449"	desc="系统邮件"	module="邮件模块"	
	id="6450"	desc="不可低于赠送最低数量"	module="背包模块"	
	id="6451"	desc="不可高于赠送最高数量"	module="背包模块"	
	id="6452"	desc="使用"	module="背包模块"	
	id="6453"	desc="赠送"	module="背包模块"	
	id="6454"	desc="赠送成功"	module="背包模块"	
}
 return string