package consts

//主人公の名前
const HERO_NAME = "たかし"

//Xの名前（宝物の名前）
const X_NAME = "たい焼き器"

//地図の横のマス数
const MAP_COL_NUMBER = 3

//地図の縦のマス数
const MAP_ROW_NUMBER = 3

//地図のマスの合計数
const MAP_SIZE = MAP_COL_NUMBER * MAP_ROW_NUMBER

//北へ移動コマンド
const MOVE_NORTH = 1

//東へ移動コマンド
const MOVE_EAST = 2

//南へ移動コマンド
const MOVE_SOUTH = 3

//西へ移動コマンド
const MOVE_WEST = 4

//探すコマンド
const SEARCH_COMMAND = 9

//1文字
const LENGTH_ONE_CHAR = 1

//主人公の1回の移動マスス
const MOVE_LENGTH = 1

//地図のマスごとの街名
var CITY_NAME [9]string = [9]string{"新宿", "品川", "渋谷", "田町", "新橋", "丸の内", "日暮里", "駒込", "大塚"}
