package main

const template = `
{"p1_BaoSRQ":{"Text":"%v"},"p1_DangQSTZK":{"F_Items":[["良好","良好",1],["不适","不适",1]],"SelectedValue":"良好"},"p1_ZhengZhuang":{"Hidden":true,"F_Items":[["感冒","感冒",1],["咳嗽","咳嗽",1],["发热","发热",1]],"SelectedValueArray":[]},"p1_ZaiXiao":{"F_Items":[["不在校","不在校",1],["宝山","宝山校区",1],["延长","延长校区",1],["嘉定","嘉定校区",1],["新闸路","新闸路校区",1]],"SelectedValue":"%v"},"p1_GuoNei":{"F_Items":[["国内","国内",1],["国外","国外",1]],"SelectedValue":"%v"},"p1_ddlSheng":{"F_Items":[["-1","选择省份",1,"",""],["北京","北京",1,"",""],["天津","天津",1,"",""],["上海","上海",1,"",""],["重庆","重庆",1,"",""],["河北","河北",1,"",""],["山西","山西",1,"",""],["辽宁","辽宁",1,"",""],["吉林","吉林",1,"",""],["黑龙江","黑龙江",1,"",""],["江苏","江苏",1,"",""],["浙江","浙江",1,"",""],["安徽","安徽",1,"",""],["福建","福建",1,"",""],["江西","江西",1,"",""],["山东","山东",1,"",""],["河南","河南",1,"",""],["湖北","湖北",1,"",""],["湖南","湖南",1,"",""],["广东","广东",1,"",""],["海南","海南",1,"",""],["四川","四川",1,"",""],["贵州","贵州",1,"",""],["云南","云南",1,"",""],["陕西","陕西",1,"",""],["甘肃","甘肃",1,"",""],["青海","青海",1,"",""],["内蒙古","内蒙古",1,"",""],["广西","广西",1,"",""],["西藏","西藏",1,"",""],["宁夏","宁夏",1,"",""],["新疆","新疆",1,"",""],["香港","香港",1,"",""],["澳门","澳门",1,"",""],["台湾","台湾",1,"",""]],"SelectedValueArray":["%v"]},"p1_ddlShi":{"Enabled":true,"F_Items":%v,"SelectedValueArray":["%v"]},"p1_ddlXian":{"Enabled":true,"F_Items":%v,"SelectedValueArray":["%v"]},"p1_TongZWDLH":{"Required":true,"F_Items":[["是","是",1],["否","否",1]],"SelectedValue":"%v"},"p1_XiangXDZ":{"Label":"国内详细地址","Text":"%v"},"p1_QueZHZJC":{"F_Items":[["是","是",1,"",""],["否","否",1,"",""]],"SelectedValueArray":["%v"]},"p1_CengFWH":{"Label":"2020年1月10日后是否在湖北逗留过"},"p1_CengFWH_RiQi":{"Hidden":true},"p1_CengFWH_BeiZhu":{"Hidden":true},"p1_JieChu":{"Label":"01月26日至02月09日是否与来自湖北发热人员密切接触"},"p1_JieChu_RiQi":{"Hidden":true},"p1_JieChu_BeiZhu":{"Hidden":true},"p1_TuJWH":{"Label":"01月26日至02月09日是否乘坐公共交通途径湖北"},"p1_TuJWH_RiQi":{"Hidden":true},"p1_TuJWH_BeiZhu":{"Hidden":true},"p1_JiaRen":{"Label":"01月26日至02月09日家人是否有发热等症状"},"p1_JiaRen_BeiZhu":{"Hidden":true},"p1_SuiSM":{"Required":true,"F_Items":[["红色","红色",1],["黄色","黄色",1],["绿色","绿色",1]],"SelectedValue":"%v"},"p1_SuiSMSM":{"IFrameAttributes":{}},"p1":{"IFrameAttributes":{}}}`
