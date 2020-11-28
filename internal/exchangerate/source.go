package exchangerate

type Source interface {
	GetRate(date string) float32
}
