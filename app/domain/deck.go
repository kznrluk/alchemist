package domain

type CardDeckState struct {
	Hand  []CardState // 手札
	Stack []CardState // 山札
}

type CardDeck interface {
	Draw(count int)

	FindFromHand(id CardInstanceID) Card

	GetHand() []Card
	AddCard(card Card)
	CardUsed(id CardInstanceID) error
	GetState() CardDeckState
}
