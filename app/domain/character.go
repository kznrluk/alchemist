package domain

type CharacterID string

type CharacterInitializeData struct {
	CharacterID     CharacterID
	PlayerID        PlayerID // キャラクターを操作できる人
	Name            string
	DeckCardIdList  []CardID
	DefaultStatuses map[Status]StatusValue
	IsCPU           bool
}

type CharacterState struct {
	CharacterID CharacterID
	PlayerID    PlayerID // キャラクターを操作できる人

	Status Statuses
	Deck   CardDeckState
}

type Character interface {
	GetCharacterID() CharacterID
	IsAlive() bool
	StatusChange(delta StatusDelta)

	IsCPU() bool

	DrawDeck() (CardDeckState, error)
	CreateMixCase() []CardMixCase
	DoMix(cardList []Card) error
	UseCard(cardInstanceId CardInstanceID, isDryRun bool) (Action, error)

	GetState() CharacterState
}
