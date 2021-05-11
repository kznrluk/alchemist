package domain

type CardMixCaseID string
type CardMixCaseGenerationID string

type CardMixCase struct {
	CardMixCaseID CardMixCaseID
	From          []Card
	To            CardDefinition
	SuccessRate   int
	// successRate(成功度合い)を入れれば、クライアントがレスポンス無しでミックスすることができる
	// これをレスポンスに入れると成功度合いを見てミックスするか決める半チートみたいなことができてしまう
	// ただ、入れないとミックス時にAPI通信が発生する... それは微妙なので仕方なく返している
}

type Mix interface {
	FindMix(cardList []Card) []CardMixCase
	ExecMix(cardList []Card) (CardMixCase, error)
}

type MixPairRepository interface {
	Get(id CardMixCaseID) (CardMixCase, error)
	Save(mixPair CardMixCase)
	Delete(id CardMixCaseID)
}
