package domain

type UseCardCaseID string

type UseCardCase struct {
	UseCardCaseID UseCardCaseID
	UsedCard      CardInstanceID

	ActionResult ActionResult
}

type UseCardCaseRepo struct {
}
