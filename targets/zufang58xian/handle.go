package zufang58xian

type IV struct {
	Index int    `json:"I"`
	Value string `json:"V"`
}

type zufang58house struct {
	Name string
}

const (
	indexPrice         = 1016  // 价格
	indexRoomCount     = 1019  // 几室 (猜测)
	indexFloor         = 1020  // 楼层 (猜测)
	indexB             = 1021  //
	indexC             = 1022  //
	indexRoomSize      = 1025  // 房间大小
	indexD             = 1028  //
	indexE             = 1588  //
	indexF             = 1590  //
	indexG             = 1591  //
	indexH             = 1592  //
	indexI             = 1594  //
	indexFloorTotal    = 1596  // 楼层数
	indexJ             = 1597  //
	indexK             = 4902  //
	indexL             = 4903  //
	indexStreetAddress = 5100  // 街道地址
	indexRoomAdmin     = 5333  // 户主
	indexN             = 5379  //
	indexO             = 5410  //
	indexResInfo       = 5461  // 资源描述，用户域名解析跳转
	indexQ             = 5468  //
	indexR             = 5470  //
	indexLon           = 6691  // 经度
	indexLat           = 6692  // 维度
	indexS             = 7109  // 经度(未知)
	indexT             = 7110  // 维度(未知)
	indexU             = 8943  //
	indexV             = 8992  // 楼层 (猜测)
	indexReleaseTime   = 9184  // 发布时间
	indexW             = 10200 //
	indexX             = 10201 //
	indexY             = 10202 //
	indexZ             = 10203 //
	indexApartmentName = 10276 // 公寓名称
	indexA             = 10589 //
	indexM             = 10652 //
	indexP             = 10922 //
	indexAA            = 11123 //
	indexAB            = 12135 //
)

var (
	indexNameMap = map[int]string{}
)
