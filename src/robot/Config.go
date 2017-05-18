package robot

//Config connect ip and port
type Config struct {
	Adds    []string
	AddIdx  int
	LogType int //0显示所有日志 1显示关键日志
}

var nameList = []string{
	"之桃", "慕青", "尔岚", "初夏", "沛菡", "傲珊", "曼文", "乐菱", "惜文", "香寒", "新柔", "语蓉", "海安", "夜蓉", "涵柏", "水桃", "醉蓝", "语琴", "从彤", "傲晴", "语兰", "又菱", "碧彤", "元霜", "怜梦", "紫寒", "妙彤", "曼易", "南莲", "紫翠", "雨寒", "易烟", "如萱", "若南", "寻真", "晓亦", "向珊", "慕灵", "以蕊", "映易", "雪柳", "海云", "凝天", "沛珊", "寒云", "冰旋", "宛儿", "绿真", "晓霜", "碧凡", "夏菡", "曼香", "若烟", "半梦", "雅绿", "冰蓝", "灵槐", "平安", "书翠", "翠风", "代云", "梦曼", "幼翠", "听寒", "梦柏", "醉易", "访旋", "亦玉", "凌萱", "访卉", "怀亦", "笑蓝", "靖柏", "夜蕾", "冰夏", "梦松", "书雪", "乐枫", "念薇", "靖雁", "从寒", "觅波", "静曼", "凡旋", "以亦", "念露", "芷蕾", "千兰", "新波", "代真", "新蕾", "雁玉", "冷卉", "紫山", "千琴", "傲芙", "盼山", "怀蝶", "冰兰", "山柏", "翠萱", "问旋", "白易", "问筠", "如霜", "半芹", "丹珍", "冰彤", "亦寒", "之瑶", "冰露", "尔珍", "谷雪", "乐萱", "涵菡", "海莲", "傲蕾", "青槐", "易梦", "惜雪", "宛海", "之柔", "夏青", "亦瑶", "妙菡", "紫蓝", "幻柏", "元风", "冰枫", "访蕊", "芷蕊", "凡蕾", "凡柔", "安蕾", "天荷", "含玉", "书兰", "雅琴", "书瑶", "从安", "夏槐", "念芹", "代曼", "幻珊", "谷丝", "秋翠", "白晴", "海露", "代荷", "含玉", "书蕾", "听白", "灵雁", "雪青", "乐瑶", "含烟", "涵双", "平蝶", "雅蕊", "傲之", "灵薇", "含蕾", "从梦", "从蓉", "初丹。听兰", "听蓉", "语芙", "夏彤", "凌瑶", "忆翠", "幻灵", "怜菡", "紫南", "依珊", "妙竹", "访烟", "怜蕾", "映寒", "友绿", "冰萍", "惜霜", "凌香", "芷蕾", "雁卉", "迎梦", "元柏", "代萱", "紫真", "千青", "凌寒", "紫安", "寒安", "怀蕊", "秋荷", "涵雁",
}

var headIconList = []string{
	"https://img.momocdn.com/album/A8/01/A80109FD-0976-2B07-0B19-74EBFCF6F96320170222_400x400.jpg",
	"https://img.momocdn.com/live/59/C2/59C2C96B-760D-1BFF-085A-2F1E61980D5220170331_400x400.jpg",
	"https://img.momocdn.com/album/FB/31/FB31F069-A2D5-7E78-06B8-F90EF0649BE320161105_400x400.jpg",
	"https://img.momocdn.com/album/21/13/21136F5E-9BA1-52FA-9763-B80D307F53EE20160126_400x400.jpg",
	"https://img.momocdn.com/album/83/29/8329BAF8-003C-1A0D-549C-6C891882FFAF20170215_400x400.jpg",
	"https://img.momocdn.com/live/C9/38/C9386294-7D96-B9AE-0011-F385A59FE1C420170408_400x400.jpg",
	"https://img.momocdn.com/live/C9/B0/C9B06F3B-943C-3133-5FF2-E04FF662A48220170410_400x400.jpg",
	"https://img.momocdn.com/live/DB/93/DB930AF5-2F58-72FC-B475-C3BCB76091D720161129_400x400.jpg",
	"https://img.momocdn.com/live/60/6F/606F8848-F972-9630-0E34-EE9425C96CB320170410_400x400.jpg",
	"https://img.momocdn.com/live/2B/44/2B44CB6E-FB8C-0A36-538F-6F2C73BE08A920170410_400x400.jpg",
	"https://img.momocdn.com/album/A4/6C/A46CAD08-AC03-61CC-B325-3746489DE89120161214_400x400.jpg",
	"https://img.momocdn.com/live/25/74/25741C6F-E2A0-1AF7-7B5C-950A4E45E43220170410_400x400.jpg",
	"https://img.momocdn.com/live/71/5B/715B99AE-BDF6-6BC3-A479-CB46268A770B20170327_400x400.jpg",
	"https://img.momocdn.com/live/AB/4D/AB4D649F-A0F7-9B3F-7912-D13B9E3D4C7F20170312_400x400.jpg",
	"https://img.momocdn.com/live/71/BC/71BC1FA3-E3FD-0740-176E-2B54F94571AF20170316_400x400.jpg",
	"https://img.momocdn.com/live/0C/B1/0CB1E5F1-9F84-2327-24E2-C399FF8A56C220170123_400x400.jpg",
	"https://img.momocdn.com/live/F7/F5/F7F5FF73-37C3-05B7-DBB3-16F4C20241F120170401_400x400.jpg",
	"https://img.momocdn.com/live/D1/8B/D18BD86A-06D5-681D-1A33-2254EBB99A3520161109_400x400.jpg",
	"https://img.momocdn.com/live/41/52/41527616-9AB4-C885-9EE3-45B68A07870720170309_400x400.jpg",
}

var publicKey = []byte(`
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCwtelan7tuLRT97nbBEKqqqb9d
LDW8wxIdrNbZIFm410GA3puhD0boI0bCzHz//PCZg7ZpFhJArLrnmv5EzBKwbp7/
QIyOHzZukSCGs9XFd3Mu94UfXY7G/3q9KwP1btPfmBJJFEbFBN6sj3j0+zob3hhw
EyY5hs0hYgVIFNpe/QIDAQAB
-----END PUBLIC KEY-----
`)
