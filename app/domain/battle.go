package domain

type BattleID string

type BattleState struct {
	Alpha []CharacterState // 味方側チーム
	Beta  []CharacterState // 敵側チーム

	CurrentControllableCharacterID CharacterID
	CurrentTurn                    int
}

type Battle interface {
	GetBattleID() BattleID
	CurrentControllableCharacterID() CharacterID

	Mix(invoker CharacterID, cardIDList []CardInstanceID) (BattleState, error)
	UseCard(invoker CharacterID, cardID CardInstanceID, target CharacterID) (BattleState, error)
	TurnEnd(invoker CharacterID) error

	GetState() BattleState
}

type BattleRepository interface {
	GetFromBattleID(id BattleID) (Battle, error)
	Save(battle Battle) error
}
