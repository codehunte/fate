package model_test

var ttList = [][]string{
	{"木", "木", "木", "大吉"},
	{"木", "木", "火", "大吉"},
	{"木", "木", "土", "大吉"},
	{"木", "木", "金", "凶多吉少"},
	{"木", "木", "水", "吉多于凶"},
	{"木", "火", "木", "大吉"},
	{"木", "火", "火", "中吉"},
	{"木", "火", "土", "大吉"},
	{"木", "火", "金", "凶多于吉"},
	{"木", "火", "水", "大凶"},
	{"木", "土", "木", "大凶"},
	{"木", "土", "火", "中吉"},
	{"木", "土", "土", "吉"},
	{"木", "土", "金", "吉多于凶"},
	{"木", "土", "水", "大凶"},
	{"木", "金", "木", "大凶"},
	{"木", "金", "火", "大凶"},
	{"木", "金", "土", "凶多于吉"},
	{"木", "金", "金", "大凶"},
	{"木", "金", "水", "大凶"},
	{"木", "水", "木", "大吉"},
	{"木", "水", "火", "凶多于吉"},
	{"木", "水", "土", "凶多于吉"},
	{"木", "水", "金", "大吉"},
	{"木", "水", "水", "大吉"},
	{"火", "木", "木", "大吉"},
	{"火", "木", "火", "大吉"},
	{"火", "木", "土", "大吉"},
	{"火", "木", "金", "凶多于吉"},
	{"火", "木", "水", "中吉"},
	{"火", "火", "木", "大吉"},
	{"火", "火", "火", "中吉"},
	{"火", "火", "土", "大吉"},
	{"火", "火", "金", "大凶"},
	{"火", "火", "水", "大凶"},
	{"火", "土", "木", "吉多于凶"},
	{"火", "土", "火", "大吉"},
	{"火", "土", "土", "大吉"},
	{"火", "土", "金", "大吉"},
	{"火", "土", "水", "吉多于凶"},
	{"火", "金", "木", "大凶"},
	{"火", "金", "火", "大凶"},
	{"火", "金", "土", "吉凶参半"},
	{"火", "金", "金", "大凶"},
	{"火", "金", "水", "大凶"},
	{"火", "水", "木", "凶多于吉"},
	{"火", "水", "火", "大凶"},
	{"火", "水", "土", "大凶"},
	{"火", "水", "金", "大凶"},
	{"火", "水", "水", "大凶"},
	{"土", "木", "木", "中吉"},
	{"土", "木", "火", "中吉"},
	{"土", "木", "土", "凶多于吉"},
	{"土", "木", "金", "大凶"},
	{"土", "木", "水", "凶多于吉"},
	{"土", "火", "木", "大吉"},
	{"土", "火", "火", "大吉"},
	{"土", "火", "土", "大吉"},
	{"土", "火", "金", "吉多于凶"},
	{"土", "火", "水", "大凶"},
	{"土", "土", "木", "中吉"},
	{"土", "土", "火", "大吉"},
	{"土", "土", "土", "大吉"},
	{"土", "土", "金", "大吉"},
	{"土", "土", "水", "凶多于吉"},
	{"土", "金", "木", "凶多于吉"},
	{"土", "金", "火", "凶多于吉"},
	{"土", "金", "土", "大吉"},
	{"土", "金", "金", "大吉"},
	{"土", "金", "水", "大吉"},
	{"土", "水", "木", "凶多于吉"},
	{"土", "水", "火", "大凶"},
	{"土", "水", "土", "大凶"},
	{"土", "水", "金", "吉凶参半"},
	{"土", "水", "水", "大凶"},
	{"金", "木", "木", "凶多于吉"},
	{"金", "木", "火", "凶多于吉"},
	{"金", "木", "土", "凶多于吉"},
	{"金", "木", "金", "大凶"},
	{"金", "木", "水", "凶多于吉"},
	{"金", "火", "木", "凶多于吉"},
	{"金", "火", "火", "吉凶参半"},
	{"金", "火", "土", "吉凶参半"},
	{"金", "火", "金", "大凶"},
	{"金", "火", "水", "大凶"},
	{"金", "土", "木", "中吉"},
	{"金", "土", "火", "大吉"},
	{"金", "土", "土", "大吉"},
	{"金", "土", "金", "大吉"},
	{"金", "土", "水", "吉多于凶"},
	{"金", "金", "木", "大凶"},
	{"金", "金", "木", "大凶"},
	{"金", "金", "土", "大吉"},
	{"金", "金", "金", "中吉"},
	{"金", "金", "水", "中吉"},
	{"金", "水", "木", "大吉"},
	{"金", "水", "火", "凶多于吉"},
	{"金", "水", "土", "吉"},
	{"金", "水", "金", "大吉"},
	{"金", "水", "水", "中吉"},
	{"水", "木", "木", "大吉"},
	{"水", "木", "火", "大吉"},
	{"水", "木", "土", "大吉"},
	{"水", "木", "金", "凶多于吉"},
	{"水", "木", "水", "大吉"},
	{"水", "火", "木", "中吉"},
	{"水", "火", "火", "大凶"},
	{"水", "火", "土", "凶多于吉"},
	{"水", "火", "金", "大凶"},
	{"水", "火", "水", "大凶"},
	{"水", "土", "木", "大凶"},
	{"水", "土", "火", "中吉"},
	{"水", "土", "土", "中吉"},
	{"水", "土", "金", "中吉"},
	{"水", "土", "水", "大凶"},
	{"水", "金", "木", "凶多于吉"},
	{"水", "金", "火", "凶多于"},
	{"水", "金", "土", "大吉"},
	{"水", "金", "金", "中吉"},
	{"水", "金", "水", "大吉"},
	{"水", "水", "木", "大吉"},
	{"水", "水", "火", "大凶"},
	{"水", "水", "土", "大凶"},
	{"水", "水", "金", "大吉"},
	{"水", "水", "水", "中吉"},
}

//
//func TestThreeFive_InitSave(t *testing.T) {
//	for _, v := range name {
//		five := model.ThreeFive{
//			SurStrokes:    v[0],
//			SecondStrokes: v[1],
//			ThirdStrokes:  v[2],
//		}
//		five.InitSave()
//	}
//
//}
//func TestFindSecondStrokes(t *testing.T) {
//	s := model.FindSecondStrokes(17)
//	log.Println(s)
//}
//
//func TestFindThirdStrokes(t *testing.T) {
//	s := model.FindThirdStrokes(17, 8)
//	log.Println(s)
//}