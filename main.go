package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
	consts "travel_to_search_for_X/const"
)

//地図の座標
type Location struct {
	x int
	y int
}

//宝と爆弾の座標
type BombAndX struct {
	locationX     Location
	locationBomb1 Location
	//定義してあるが、未使用
	locationBomb2 Location
}

//主人公の座標
type HeroLocation struct {
	location Location
}

var bombAndX = new(BombAndX)
var hero = new(HeroLocation)

//地図。配列の各要素には街名が入る
var cityMap [consts.MAP_COL_NUMBER][consts.MAP_ROW_NUMBER]string

//初期化終了を待つチャンネル
var waitChannel = make(chan string)

//配列の要素に指定の値が含まれるかチェック
func arrayContains(arr []int, value int) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}
	return false
}

//CUIの画面を初期化
func clearScreen() {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	case "darwin":
		cmd = exec.Command("clear")
	case "linux":
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}

//宝と爆弾の座標を初期化
func initializeBombX() {
	rand.Seed(time.Now().Unix())
	bombAndX.locationX = Location{rand.Intn(consts.MAP_COL_NUMBER), rand.Intn(consts.MAP_ROW_NUMBER)}
	bombAndX.locationBomb1 = Location{rand.Intn(consts.MAP_COL_NUMBER), rand.Intn(consts.MAP_ROW_NUMBER)}
	bombAndX.locationBomb2 = Location{rand.Intn(consts.MAP_COL_NUMBER), rand.Intn(consts.MAP_ROW_NUMBER)}
}

//主人公の初期座標位置を初期化
func setHeroLocation() {
	var x = int(consts.MAP_COL_NUMBER / 2)
	var y = int(consts.MAP_ROW_NUMBER / 2)
	hero.location = Location{x, y}
}

//ゲーム全体の初期化
func initialize() {
	setHeroLocation()
	initializeBombX()
	rand.Seed(time.Now().Unix())
	var temp []int
	for i := 0; i < consts.MAP_COL_NUMBER; i++ {
		for k := 0; k < consts.MAP_ROW_NUMBER; k++ {
			r := rand.Intn(consts.MAP_SIZE)
			for arrayContains(temp, r) {
				r = rand.Intn(consts.MAP_SIZE)
			}
			cityMap[i][k] = consts.CITY_NAME[r]
			temp = append(temp, r)
		}
	}
	waitChannel <- ""
}

//主人公のコマンドを表示
func showCommands() []int {

	var acceptCommands []int

	if 0 <= hero.location.y-1 {
		fmt.Println("1 北へ行く")
		acceptCommands = append(acceptCommands, 1)
	} else {
		fmt.Println("北へは行けない")
	}

	if hero.location.x+1 < consts.MAP_ROW_NUMBER {
		fmt.Println("2 東へ行く")
		acceptCommands = append(acceptCommands, 2)
	} else {
		fmt.Println("東へは行けない")
	}

	if hero.location.y+1 < consts.MAP_COL_NUMBER {
		fmt.Println("3 南へ行く")
		acceptCommands = append(acceptCommands, 3)
	} else {
		fmt.Println("南へは行けない")
	}

	if 0 <= hero.location.x-1 {
		fmt.Println("4 西へ行く")
		acceptCommands = append(acceptCommands, 4)
	} else {
		fmt.Println("西へは行けない")
	}

	fmt.Println("9 探す")

	return acceptCommands
}

//宝物を探す。爆弾を引き当てると死亡。宝物、爆弾どちらも当てるとゲーム終了
func searchX() {
	clearScreen()
	fmt.Println(consts.HERO_NAME + "は" + cityMap[hero.location.x][hero.location.y] + "で" + consts.X_NAME + "を探した")
	waitAnyKeyPress()

	if hero.location.x == bombAndX.locationBomb1.x &&
		hero.location.y == bombAndX.locationBomb1.y {
		fmt.Println(consts.HERO_NAME + "は爆弾を引いてしまった！ちゅど～ん！！ ...死んだ")
		os.Exit(0)
	}

	if hero.location.x == bombAndX.locationX.x &&
		hero.location.y == bombAndX.locationX.y {
		fmt.Println(consts.HERO_NAME + "は" + consts.X_NAME + "を探し当てた！ おめでとう！！")
		os.Exit(0)
	} else {
		fmt.Println(cityMap[hero.location.x][hero.location.y] + "に" + consts.X_NAME + "は無かった...")
		waitAnyKeyPress()
	}
}

//タイトルの表示
func showTitle() {
	clearScreen()
	var title = fmt.Sprintf("%sの%sを求めて", consts.HERO_NAME, consts.X_NAME)
	fmt.Println(title)
	fmt.Println("Press return key...")
	waitAnyKeyPress()
}

//コマンドの入力
func inputCommands(acceptCommands []int) {
	var input string
	var commandNo int64
	reader := bufio.NewReader(os.Stdin)

	for {
		input, _ = reader.ReadString('\n')
		input = strings.TrimRight(input, "\r\n")
		input = strings.TrimRight(input, "\n")
		commandNo, _ = strconv.ParseInt(input, 10, 64)

		if (len(input) == consts.LENGTH_ONE_CHAR && arrayContains(acceptCommands, int(commandNo))) ||
			(len(input) == consts.LENGTH_ONE_CHAR && commandNo == consts.SEARCH_COMMAND) {
			break
		}
	}

	switch commandNo {
	case consts.MOVE_NORTH:
		hero.location.y -= consts.MOVE_LENGTH
	case consts.MOVE_EAST:
		hero.location.x += consts.MOVE_LENGTH
	case consts.MOVE_SOUTH:
		hero.location.y += consts.MOVE_LENGTH
	case consts.MOVE_WEST:
		hero.location.x -= consts.MOVE_LENGTH
	case consts.SEARCH_COMMAND:
		searchX()
	}
}

//キー入力待ち
func waitAnyKeyPress() {
	key := bufio.NewScanner(os.Stdin)
	key.Scan()
}

//メイン処理
func youAreInCity() {
	var msg string
	msg = fmt.Sprintf(consts.HERO_NAME+"は今、%sにいる", cityMap[hero.location.x][hero.location.y])
	clearScreen()
	fmt.Println(msg)
	acceptCommands := showCommands()
	inputCommands(acceptCommands)
}

//ここから処理が始まる
func main() {
	go initialize()
	showTitle()
	<-waitChannel
	for {
		youAreInCity()
	}
}
