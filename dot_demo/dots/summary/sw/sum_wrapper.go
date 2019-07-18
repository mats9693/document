package sw

type SumWrapper interface {
    ShowDescription(extra string) string
    ShowSimpleDotDesc() string
    ShowComplexDotDesc() string
}
