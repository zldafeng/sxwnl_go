package src

type Vector3 struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

type Vector2 struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type GZ struct {
	tg int
	dz int
}

type Month struct {
	y int //公历年份
	m int //公历月分

	d0 int // 月首儒略日数
	dn int // 月天数

	Ly  string //干支纪年
	ShX int //年生肖

	nianhao string

	days    []string
	yearGan int
	yearZhi int
}
type Day struct {
	d0    float64 //儒略日
	di    float64 //公历月内日序数
	y     int //公历年
	m     int //公历月
	d     float64
	dn    float64 //所在公历月的天数
	week0 float64 //月首的星期
	week  float64 //当前日的星期数
	weeki float64 //本日所在的周序号
	weekN float64 //本月的总周数

	Ldi    float64 //距农历月首的编移量,0对应初一
	Ldc string // 农历日名称

	cur_dz float64 //距冬至的天数
	cur_xz float64 //距夏至的天数
	cur_lq float64 //距立秋的天数
	cur_mz float64 //距芒种的天数
	cur_xs float64 //距小暑的天数

	cur_jq []float64 //节气时间索引 （这里可以使用JD::JD2DD转换成准确时间）
	cur_cn []int     //节气名称索引

	Lmc   int  //阴历月的月份
	Ldn   float64  //阴历月的天数
	Lleap bool //是不是阴历的润月
	Lmc2  int  //下个月的月份，不存在则为-1

	qk int //节气，不存在则为-1

	Lyear  float64 //农历纪年(10进制,1984年起算)
	Lyear0 float64 //农历纪年(10进制,1984年起算)
	Lyear2 GZ  //干支纪年(立春)
	Lyear3 GZ  //干支纪年(正月)
	Lyear4 float64

	Lmonth  int
	Lmonth2 GZ //月天干地支

	Lday2 GZ

	XiZ  int //星座
	jqmc float64 //节气类型
	jqjd float64
	jqsj string //节气时间

	yxmc float64 //月像类型
	yxjd float64
	yxsj string //月像时间
}
