package domain

type CardID string
type CardInstanceID string
type CardType string

const (
	CARD_TYPE_MATERIAL = "MATERIAL"
	CARD_TYPE_ACTION   = "ACTION"
)

type CardDefinition struct {
	ID         int
	Name       string
	AssetLink  string
	Attribute  Attribute
	CardType   CardType
	ActionType ActionType
	Count      int
	BasePower  float32
	// AttributePower float32
	Description string
}

type CardState struct {
	CardDefinition

	CardInstanceID CardInstanceID // デッキにある他のカードと識別するためのID
	CurrentPower   float32        // 錬金大成功 成功 失敗 等で決まる
}

type Card interface {
	GetCardInstanceId() CardInstanceID
	GetCurrentPower()
	GetState() CardState
}
