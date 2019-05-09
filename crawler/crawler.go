package main

import (
	"database/sql"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gocolly/colly"
)

var db *sql.DB

func getDBIns() *sql.DB {
	var err error
	if db == nil {
		db, err = sql.Open("mysql", "fengbo:FengBoPwd1.@tcp(127.0.0.1:3306)/practices")
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}
	}
	return db
}

func main() {

	maxPlayTimes := selectLatest("https://ximalaya.com/ertong/20262183/")
	fmt.Println("max_play_times: ", maxPlayTimes)
	return
	c := colly.NewCollector()
	var cnt int

	//喜马拉雅
	//<span class="count _t4_">
	siteName := "喜马拉雅"
	xmly := map[string]string{
		"八路叔叔张福远：少年特战队":        "https://www.ximalaya.com/ertong/20262183/",
		"八路叔叔张福远：特种兵学校":        "https://www.ximalaya.com/ertong/20262172/",
		"八路叔叔张福远：海军陆战队":        "https://www.ximalaya.com/ertong/21751406/",
		"八路叔叔张福远：铁血战鹰队":        "https://www.ximalaya.com/ertong/21751426/",
		"超级战舰合集【1-4部】":         "https://www.ximalaya.com/ertong/16682540/",
		"超级战舰1：绝境逃生":           "https://www.ximalaya.com/ertong/7919167/",
		"超级战舰2：生化战士":           "http://www.ximalaya.com/album/7919214",
		"超级战舰3：决战黑海湾":          "http://www.ximalaya.com/album/16681512",
		"超级战舰4：北极之战":           "http://www.ximalaya.com/album/16682018",
		"特种兵学校1 【限时免费】":        "https://www.ximalaya.com/ertong/21642975/",
		"海军陆战队1【限时免费】":         "https://www.ximalaya.com/ertong/22492367/",
		"特种兵学校":                "https://www.ximalaya.com/ertong/23169908/",
		"海军陆战队":                "https://www.ximalaya.com/ertong/23301112/",
		"少年特战队":                "https://www.ximalaya.com/ertong/23170266/",
		"铁血战鹰队":                "https://www.ximalaya.com/ertong/23335543/",
		"特种兵大冒险":               "https://www.ximalaya.com/ertong/22886040/",
		"【合集】小屁孩日记：一年级男生版":     "http://www.ximalaya.com/album/16856376",
		"【合集】超级战舰（1-4）":        "http://www.ximalaya.com/album/16682540",
		"【合集】女巫来了（1-3）":        "http://www.ximalaya.com/album/18452110",
		"【合集】小糊涂日记（1-4）":       "http://www.ximalaya.com/album/16851586",
		"【合集】幽默西游（1-4）":        "http://www.ximalaya.com/album/18449593",
		"【合集】章鱼哥派出所（1-5）":      "http://www.ximalaya.com/album/19598035",
		"阿丽和她的恐龙朋友":            "http://www.ximalaya.com/album/20099181",
		"儿童相声：酷哥酷发明1":          "http://www.ximalaya.com/album/7746080",
		"儿童相声：酷哥酷发明2":          "http://www.ximalaya.com/album/7746116",
		"儿童相声：酷哥酷发明3":          "http://www.ximalaya.com/album/19457498",
		"儿童相声：酷哥酷发明4":          "http://www.ximalaya.com/album/19581394",
		"儿童相声：酷哥酷发明5":          "http://www.ximalaya.com/album/19581514",
		"儿童相声：酷哥酷发明6":          "http://www.ximalaya.com/album/19581704",
		"疯狂动物学校1：胡萝卜怪兽":        "http://www.ximalaya.com/album/7303013",
		"疯狂动物学校2：梦里的贪吃怪":       "http://www.ximalaya.com/album/7303055",
		"疯狂蔬菜学校1":              "http://www.ximalaya.com/album/20115761",
		"疯狂蔬菜学校2":              "http://www.ximalaya.com/album/20115834",
		"疯狂蔬菜学校3":              "http://www.ximalaya.com/album/21639406",
		"疯狂蔬菜学校4":              "http://www.ximalaya.com/album/21639528",
		"疯狂蔬菜学校5":              "http://www.ximalaya.com/album/21639548",
		"疯狂蔬菜学校6":              "http://www.ximalaya.com/album/21639575",
		"疯狂蔬菜学校7":              "http://www.ximalaya.com/album/21639593",
		"疯狂蔬菜学校8":              "http://www.ximalaya.com/album/21639622",
		"和恐龙一起玩：我们不是恐龙":        "http://www.ximalaya.com/album/21567821",
		"和恐龙一起玩：我们的名字叫恐龙":      "http://www.ximalaya.com/album/21568349",
		"和恐龙一起玩：我们喜欢吃树叶":       "http://www.ximalaya.com/album/21568495",
		"和恐龙一起玩：我们用两条腿走路":      "http://www.ximalaya.com/album/21568593",
		"和恐龙一起玩：我们长得有点怪":       "http://www.ximalaya.com/album/21568639",
		"解读西方文明之源：《少儿版希腊神话》":   "http://www.ximalaya.com/album/6417149",
		"恐龙鲁鲁的故事":              "http://www.ximalaya.com/album/19366831",
		"口袋故事：晚安宝贝童话1":         "http://www.ximalaya.com/album/16863533",
		"罗琳：哈利波特之母":            "http://www.ximalaya.com/album/7919290",
		"女巫来了1：猫饼干传奇":          "http://www.ximalaya.com/album/7746515",
		"女巫来了2：水银表传奇":          "http://www.ximalaya.com/album/7746640",
		"女巫来了3：鹦鹉螺传奇":          "http://www.ximalaya.com/album/18450866",
		"森林里的恐龙朋友":             "http://www.ximalaya.com/album/7747178",
		"少儿神话广播剧《哪吒》":          "http://www.ximalaya.com/album/6416974",
		"小糊涂日记：第二部全集":          "http://www.ximalaya.com/album/7287951",
		"小糊涂日记：第三部全集":          "http://www.ximalaya.com/album/7287981",
		"小糊涂日记：第四部全集":          "http://www.ximalaya.com/album/7288007",
		"小糊涂日记：第一部全集":          "http://www.ximalaya.com/album/7236362",
		"小屁孩日记(男生版)：一年级屁事多(上)": "http://www.ximalaya.com/album/6923458",
		"小屁孩日记(男生版)：一年级屁事多(下)": "http://www.ximalaya.com/album/7100977",
		"小屁孩日记(女生版)：一年级快乐多":    "http://www.ximalaya.com/album/7180514",
		"小嘴鳄鱼和大嘴猴子1":           "http://www.ximalaya.com/album/7524367",
		"小嘴鳄鱼和大嘴猴子2":           "http://www.ximalaya.com/album/7524403",
		"小嘴鳄鱼和大嘴猴子3":           "http://www.ximalaya.com/album/20100653",
		"幽默西游1":                "http://www.ximalaya.com/album/7748747",
		"幽默西游2":                "http://www.ximalaya.com/album/7748867",
		"幽默西游3":                "http://www.ximalaya.com/album/18449503",
		"幽默西游4":                "http://www.ximalaya.com/album/18449529",
		"幼儿园的大鼻子先生":            "http://www.ximalaya.com/album/7747042",
		"幼儿园的男老师":              "http://www.ximalaya.com/album/7746968",
		"章鱼哥派出所1：镇长失踪事件":       "http://www.ximalaya.com/album/7919234",
		"章鱼哥派出所2：沉船里的秘密":       "http://www.ximalaya.com/album/7919249",
		"章鱼哥派出所3：能源谷危机":        "http://www.ximalaya.com/album/19458228",
		"章鱼哥派出所4：真假鱼博士":        "http://www.ximalaya.com/album/19581900",
		"章鱼哥派出所5：谎言俱乐部":        "http://www.ximalaya.com/album/19582058",
		"章鱼国小时代1：学霸归来":         "http://www.ximalaya.com/album/20116006",
		"章鱼国小时代2：学校大变样":        "http://www.ximalaya.com/album/20192002",
		"章鱼国小时代3：馋猫症候群":        "http://www.ximalaya.com/album/20192214",
		"章鱼国小时代4:校园黑客风波":       "http://www.ximalaya.com/album/20963607",
		"章鱼国小时代5：法定玩闹日":        "http://www.ximalaya.com/album/20192180",
		"章鱼国小时代6：名侦探向日葵 ":      "http://www.ximalaya.com/album/20965142",
		"章鱼国小时代7：课堂大爆炸":        "http://www.ximalaya.com/album/20965299",
		"章鱼国小时代8：藏书谜案":         "http://www.ximalaya.com/album/20965934",
		"章鱼国小时代9：超级小学声":        "http://www.ximalaya.com/album/20966484",
		"章鱼国小时代10：花炮大明星":       "http://www.ximalaya.com/album/20966540",
		"章鱼国小时代11：女王驾到":        "http://www.ximalaya.com/album/20966623",
		"章鱼国小时代12：考试向前冲":       "http://www.ximalaya.com/album/20966857",
	}
	c.OnHTML("span[class]", func(element *colly.HTMLElement) {
		if element.Attr("class") == "count _t4_" {
			cnt = int(parse(deleteBlankSpace(element.Text)))
		}
	})
	fmt.Println(siteName)
	fmt.Printf("\n")
	success := 0
	for audioName, url := range xmly {
		cnt = 0
		err := c.Visit(url)
		if err != nil {
			if err.Error() == "Forbidden" ||
				err.Error() == "URL already visited" {
				fmt.Println(url + ": " + err.Error())
				cnt = selectLatest(url)
			}
		}

		fmt.Printf("%20s\t%30s\t%d\n", audioName, url, cnt)
		saveDB(siteName, audioName, url, cnt)
		success++
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Printf("xmly total %d success %d\n", len(xmly), success)
	fmt.Printf("\n\n\n")
	//蜻蜓
	//<span class="_8-O6" data-reactid="186">
	siteName = "蜻蜓"
	qt := map[string]string{
		"八路叔叔张福远：少年特战队":           "http://www.qingting.fm/channels/279506",
		"八路叔叔张福远：特种兵学校":           "http://www.qingting.fm/channels/279466",
		"【合集】超级战舰（1-4部）":          "http://www.qingting.fm/channels/263991",
		"超级战舰1：绝境逃生":              "http://www.qingting.fm/channels/221502",
		"超级战舰2：生化战士":              "http://www.qingting.fm/channels/221503",
		"超级战舰3：决战黑海湾":             "http://www.qingting.fm/channels/221506",
		"超级战舰4：北极之战":              "http://www.qingting.fm/channels/263992",
		"儿童相声：酷哥酷发明1":             "https://www.qingting.fm/channels/214949/",
		"儿童相声：酷哥酷发明2":             "https://www.qingting.fm/channels/217409/",
		"儿童相声：酷哥酷发明3":             "https://www.qingting.fm/channels/278565/",
		"儿童相声：酷哥酷发明4":             "https://www.qingting.fm/channels/278566/",
		"儿童相声：酷哥酷发明5":             "https://www.qingting.fm/channels/278569/",
		"儿童相声：酷哥酷发明6":             "https://www.qingting.fm/channels/278571/",
		"解读西方文明之源：《少儿版希腊神话》":      "https://www.qingting.fm/channels/214937/",
		"口袋故事：晚安宝贝童话1（对应口袋故事3、4）": "https://www.qingting.fm/channels/278614",
		"口袋故事：晚安宝贝童话2（对应口袋故事1、2）": "https://www.qingting.fm/channels/278615/",
		"口袋故事：晚安宝贝童话3（对应口袋故事5、6）": "https://www.qingting.fm/channels/278616/",
		"口袋故事：晚安宝贝童话4（对应幼幼小故事上下）": "https://www.qingting.fm/channels/278618/",
		"罗琳：哈利波特之母":               "https://www.qingting.fm/channels/217565/",
		"女巫来了1：猫饼干传奇":             "https://www.qingting.fm/channels/217545/",
		"女巫来了2：水银表传奇":             "https://www.qingting.fm/channels/221310",
		"女巫来了3：鹦鹉螺传奇":             "https://www.qingting.fm/channels/269035",
		"【合集】女巫来了（1-3）":           "https://www.qingting.fm/channels/269778/",
		"森林里的恐龙朋友":                "https://www.qingting.fm/channels/214942/",
		"小糊涂日记：第一部全集":             "https://www.qingting.fm/channels/216045/",
		"小糊涂日记：第二部全集":             "https://www.qingting.fm/channels/221297",
		"小糊涂日记：第三部全集":             "https://www.qingting.fm/channels/221491/",
		"小糊涂日记：第四部全集":             "https://www.qingting.fm/channels/221493",
		"【合集】小糊涂日记（1-4）":          "https://www.qingting.fm/channels/269023/",
		"小屁孩日记(男生版)：一年级屁事多（上）":    "https://www.qingting.fm/channels/210040/",
		"小屁孩日记(男生版)：一年级屁事多（下）":    "https://www.qingting.fm/channels/217404/",
		"小屁孩日记(女生版)：一年级快乐多":       "https://www.qingting.fm/channels/216335/",
		"小嘴鳄鱼和大嘴猴子1":              "https://www.qingting.fm/channels/214939/",
		"小嘴鳄鱼和大嘴猴子2":              "https://www.qingting.fm/channels/214941/",
		"小嘴鳄鱼和大嘴猴子3":              "https://www.qingting.fm/channels/278563/",
		"幽默西游1":                   "https://www.qingting.fm/channels/214929/",
		"幽默西游2":                   "https://www.qingting.fm/channels/217353/",
		"幽默西游3":                   "https://www.qingting.fm/channels/269024/",
		"幽默西游4":                   "https://www.qingting.fm/channels/269025/",
		"幼儿园的大鼻子先生":               "https://www.qingting.fm/channels/217342/",
		"幼儿园的男老师":                 "https://www.qingting.fm/channels/217348",
		"章鱼哥派出所1":                 "https://www.qingting.fm/channels/217554/",
		"章鱼哥派出所2":                 "https://www.qingting.fm/channels/221322/",
		"章鱼哥派出所3":                 "https://www.qingting.fm/channels/221500/",
		"章鱼哥派出所4":                 "https://www.qingting.fm/channels/278572",
		"章鱼哥派出所5":                 "https://www.qingting.fm/channels/278573",
		"最豪华的机器人":                 "https://www.qingting.fm/channels/221325",
	}

	//c1 := colly.NewCollector()
	c.OnHTML("span[class]", func(element *colly.HTMLElement) {
		if element.Attr("class") == "playcount" {
			cnt = int(parse(deleteBlankSpace(element.Text)))
		}
	})

	fmt.Println(siteName)
	fmt.Printf("\n")
	success = 0
	for audioName, url := range qt {
		cnt = 0
		fmt.Println("visting: ", url)
		err := c.Visit(url)
		if err != nil {
			// panic(url + ": " + err.Error())
			fmt.Println(url + ": " + err.Error())
			cnt = selectLatest(url)
		}
		fmt.Printf("%20s\t%30s\t%d\n", audioName, url, cnt)
		saveDB(siteName, audioName, url, cnt)
		success++
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Printf("qt total %d success %d\n", len(qt), success)
	fmt.Printf("\n\n\n")
	//蜻蜓
	//<span class="_8-O6" data-reactid="186">
	siteName = "懒人听书"
	lrts := map[string]string{
		"超级战舰1：绝境逃生":              "http://www.lrts.me/book/34257",
		"超级战舰2：生化战士":              "http://www.lrts.me/book/34259",
		"超级战舰3：决战黑海湾":             "http://www.lrts.me/book/34416",
		"超级战舰4：北极之战":              "http://www.lrts.me/book/35462",
		"超级战舰合集":                  "http://www.lrts.me/book/35915",
		"儿童相声：酷哥酷发明1":             "http://www.lrts.me/book/34021",
		"儿童相声：酷哥酷发明2":             "http://www.lrts.me/book/34255",
		"儿童相声：酷哥酷发明3":             "http://www.lrts.me/book/42242",
		"儿童相声：酷哥酷发明4":             "http://www.lrts.me/book/42243",
		"儿童相声：酷哥酷发明5":             "http://www.lrts.me/book/42244",
		"儿童相声：酷哥酷发明6":             "http://www.lrts.me/book/42245",
		"疯狂蔬菜学校1":                 "http://www.lrts.me/book/42261",
		"疯狂蔬菜学校2":                 "http://www.lrts.me/book/42262",
		"疯狂蔬菜学校3":                 "http://www.lrts.me/book/42263",
		"疯狂蔬菜学校4":                 "http://www.lrts.me/book/42264",
		"解读西方文明之源：《少儿版希腊神话》":      "http://www.lrts.me/book/34418",
		"口袋故事：晚安宝贝童话1（对应口袋故事3、4）": "http://www.lrts.me/book/39827",
		"口袋故事：晚安宝贝童话2（对应口袋故事1、2）": "http://www.lrts.me/book/42258",
		"口袋故事：晚安宝贝童话3（对应口袋故事5、6）": "http://www.lrts.me/book/42259",
		"口袋故事：晚安宝贝童话4（对应幼幼小故事上下）": "http://www.lrts.me/book/42260",
		"罗琳：哈利波特之母":               "http://www.lrts.me/book/34423",
		"女巫来了1：猫饼干传奇":             "http://www.lrts.me/book/34222",
		"女巫来了2：水银表传奇":             "http://www.lrts.me/book/34419",
		"女巫来了3：鹦鹉螺传奇":             "http://www.lrts.me/book/41884",
		"森林里的恐龙朋友":                "http://www.lrts.me/book/34020",
		"少儿神话广播剧《哪吒》":             "http://www.lrts.me/book/34417",
		"小糊涂日记：我从哪个星球来":           "http://www.lrts.me/book/34251",
		"小糊涂日记：可恶的米多":             "http://www.lrts.me/book/34252",
		"小糊涂日记：我是胡小涂，不是小糊涂":       "http://www.lrts.me/book/34249",
		"小糊涂日记：今天不是说话日":           "http://www.lrts.me/book/34250",
		"小糊涂日记全集（1-4册）":           "http://www.lrts.me/book/34277",
		"小屁孩日记(男生版)：一年级屁事多(上)":    "http://www.lrts.me/book/33660",
		"小屁孩日记(男生版)：一年级屁事多(下)":    "http://www.lrts.me/book/34220",
		"小屁孩日记(女生版)：一年级快乐多":       "http://www.lrts.me/book/33661",
		"小嘴鳄鱼和大嘴猴子1":              "http://www.lrts.me/book/34253",
		"小嘴鳄鱼和大嘴猴子2":              "http://www.lrts.me/book/34254",
		"小嘴鳄鱼和大嘴猴子全集（共2册）":        "http://www.lrts.me/book/34278",
		"幽默西游1":                   "http://www.lrts.me/book/34022",
		"幽默西游2":                   "http://www.lrts.me/book/34256",
		"幽默西游3":                   "http://www.lrts.me/book/41885",
		"幽默西游4":                   "http://www.lrts.me/book/41886",
		"幼儿园的大鼻子先生":               "http://www.lrts.me/book/34421",
		"幼儿园的男老师":                 "http://www.lrts.me/book/34420",
		"章鱼哥派出所1":                 "http://www.lrts.me/book/34260",
		"章鱼哥派出所2":                 "http://www.lrts.me/book/34261",
		"章鱼哥派出所3":                 "http://www.lrts.me/book/34262",
		"章鱼哥派出所4":                 "http://www.lrts.me/book/42968",
		"章鱼哥派出所5":                 "http://www.lrts.me/book/42969",
		"章鱼国小时代1：学霸归来":            "http://www.lrts.me/book/42246",
		"章鱼国小时代10：花炮大明星":          "http://www.lrts.me/book/42255",
		"章鱼国小时代11：女王驾到":           "http://www.lrts.me/book/42256",
		"章鱼国小时代12：考试向前冲":          "http://www.lrts.me/book/42257",
		"章鱼国小时代2：学校大变样":           "http://www.lrts.me/book/42247",
		"章鱼国小时代3：馋猫症候群":           "http://www.lrts.me/book/42248",
		"章鱼国小时代4:校园黑客风波":          "http://www.lrts.me/book/42249",
		"章鱼国小时代5：法定玩闹日":           "http://www.lrts.me/book/42250",
		"章鱼国小时代6：名侦探向日葵":          "http://www.lrts.me/book/42251",
		"章鱼国小时代7：课堂大爆炸":           "http://www.lrts.me/book/42252",
		"章鱼国小时代8：藏书谜案":            "http://www.lrts.me/book/42253",
		"章鱼国小时代9：超级小学声":           "http://www.lrts.me/book/42254",
		"最豪华的机器人":                 "http://www.lrts.me/book/34422",
	}
	c.OnHTML("em", func(element *colly.HTMLElement) {
		if element.Attr("id") != "commentCount" {
			cnt = int(parse(deleteBlankSpace(element.Text)))
		}
	})
	fmt.Println(siteName)
	fmt.Printf("\n")
	success = 0
	for audioName, url := range lrts {
		cnt = 0
		err := c.Visit(url)
		if err != nil {
			// panic(url + ": " + err.Error())
			fmt.Println(url + ": " + err.Error())
			cnt = selectLatest(url)
		}
		fmt.Printf("%20s\t%30s\t%d\n", audioName, url, cnt)
		saveDB(siteName, audioName, url, cnt)
		success++
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Printf("lrts total %d success %d\n", len(lrts), success)
	c.Wait()
}

func deleteBlankSpace(str string) string {
	reg := regexp.MustCompile(`\s+`)
	return reg.ReplaceAllString(str, "")
}

func parse(str string) float64 {
	fmt.Println("parsing ", str)
	reg := regexp.MustCompile(`[0-9\.]+`)
	a, err := strconv.ParseFloat(reg.FindString(str), 64)
	if err != nil {
		panic(err.Error())
	}

	if strings.Contains(str, "万") || strings.Contains(str, "w") {
		a *= 10000
	}

	return a
}

func selectLatest(url string) int {
	sqlDB := getDBIns()
	q := `select max(play_times) from practices.audio_play_times_outside where url=?`
	rows, err := sqlDB.Query(q, url)
	if err != nil {
		fmt.Println("query err: ", err.Error())
		return 0
	}
	defer rows.Close()

	var (
		cnt int
	)
	if rows.Next() {
		if err = rows.Scan(&cnt); err != nil {
			fmt.Println("scan err: ", err.Error())
			return 0
		}
	}
	return int(cnt)
}

func saveDB(siteName string, audioName string, url string, cnt int) {
	db := getDBIns()
	var err error

	_, err = db.Exec("insert into practices.audio_play_times_outside (site_name,audio_name,play_times,url) values (?,?,?,?)",
		siteName, audioName, cnt, url)
	if err != nil {
		panic(err.Error())
	}
}
