package infra

import (
	"fmt"
	"github.com/kznrluk/alchemist/app/domain"
)

// TODO: インメモリだとスケールしないのでRedisとかに移設する
var mixPairRepoSingleton domain.MixPairRepository

type inMemoryMixPairDB struct {
	mixPair map[domain.CardMixCaseID]domain.CardMixCase
}

func (i *inMemoryMixPairDB) Get(id domain.CardMixCaseID) (domain.CardMixCase, error) {
	mixPair, ok := i.mixPair[id]

	if !ok {
		return domain.CardMixCase{}, fmt.Errorf("infra: ミックスペアがメモリ上に存在しません")
	}

	return mixPair, nil
}

func (i *inMemoryMixPairDB) Save(mixPair domain.CardMixCase) {
	i.mixPair[mixPair.CardMixCaseID] = mixPair
}

func (i *inMemoryMixPairDB) Delete(id domain.CardMixCaseID) {
	delete(i.mixPair, id)
}

func NewInMemoryBattleDB() domain.MixPairRepository {
	if mixPairRepoSingleton == nil {
		mixPairRepoSingleton = &inMemoryMixPairDB{
			mixPair: map[domain.CardMixCaseID]domain.CardMixCase{},
		}
	}
	return mixPairRepoSingleton
}
