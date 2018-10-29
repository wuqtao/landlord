package poker

/**
	定义扑克牌花色和数字
 */
 //定义扑克牌序号
const (
	PokerThree = iota
	PokerFour
	PokerFive
	PokerSix
	PokerSeven
	PokerEight
	PokerNine
	PokerTen
	PokerJack
	PokerQueen
	PokerKing
	PokerAce
	PokerTwo
	PokerBlackJoker
	PokerRedJoker
)
//定义扑克牌花色
const(
	PokerSuitDiamond = "Diamond" //方片
	PokerSuitHeart   = "Heart"   //红桃
	PokerSuitSpade   = "Spade"   //黑桃
	PokerSuitClub    = "Club"    //梅花
	PokerSuitJoker   = "Joker"   //大小王无花色
)
//定义扑克牌
type PokerCard struct {
	CardValue int			//card值用于排序比较
	CardSuit string			//card花色
	CardName string			//card显示的字符
}
//定义一副扑克牌
type PokerDeck struct {
	Cards [54]PokerCard
}

var originDeck PokerDeck

func init(){
	originDeck = CreateOriginDeck()
}

//创建原始扑克牌，后续只需要复制即可，不用再运算获得
func CreateOriginDeck() PokerDeck{
	deck := PokerDeck{}
	for i := 0;i<52;i++ {
		shang := i/4;
		yu := i%4;
		suit := ""
		switch(yu){
		case 0:
			suit = PokerSuitClub
		case 1:
			suit = PokerSuitDiamond
		case 2:
			suit = PokerSuitHeart
		case 3:
			suit = PokerSuitSpade
		}
		pokerValue := 0
		pokerName := ""
		switch(shang){
		case 0:
			pokerValue = PokerThree
			pokerName = "3"
		case 1:
			pokerValue = PokerFour
			pokerName = "4"
		case 2:
			pokerValue = PokerFive
			pokerName = "5"
		case 3:
			pokerValue = PokerSix
			pokerName = "6"
		case 4:
			pokerValue = PokerSeven
			pokerName = "7"
		case 5:
			pokerValue = PokerEight
			pokerName = "8"
		case 6:
			pokerValue = PokerNine
			pokerName = "9"
		case 7:
			pokerValue = PokerTen
			pokerName = "10"
		case 8:
			pokerValue = PokerJack
			pokerName = "J"
		case 9:
			pokerValue = PokerQueen
			pokerName = "Q"
		case 10:
			pokerValue = PokerKing
			pokerName = "K"
		case 11:
			pokerValue = PokerAce
			pokerName = "A"
		case 12:
			pokerValue = PokerTwo
			pokerName = "2"
		}
		deck.Cards[i] = PokerCard{
			pokerValue,
			suit,
			pokerName,
		}
	}
	deck.Cards[52] = PokerCard{
		PokerBlackJoker,
		PokerSuitJoker,
		"BlackJoker",
	}

	deck.Cards[53] = PokerCard{
		PokerRedJoker,
		PokerSuitJoker,
		"RedJoker",
	}
	return deck
}

//每个游戏桌子都有单独的扑克牌，防止洗牌等冲突
func CreateDeck() *PokerDeck{
	copyDeck := originDeck
	return &copyDeck
}






