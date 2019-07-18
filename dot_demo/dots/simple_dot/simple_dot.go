package simple_dot

import "github.com/scryinfo/dot/dot"

const (
    SimpleDotTypeId = "ea383705-9e3e-44c9-b64d-bac82b732c08"
)

type SimpleDot struct {
    description string
}

//construct dot
func newSimpleDot() (dot.Dot, error) {
    var err error
    d := &SimpleDot{}

    return d, err
}

//Data structure needed when generating newer component
func SimpleDotTypeLive() *dot.TypeLives {
    return &dot.TypeLives{
        Meta: dot.Metadata{TypeId: SimpleDotTypeId,
            NewDoter: func(_ interface{}) (dot dot.Dot, err error) {
                return newSimpleDot()
            }},
    }
}

func (s *SimpleDot) Create(l dot.Line) error {
    s.description = "test simple dot description. "
    return nil
}

func (s *SimpleDot) ShowDescription(extra string) string {
    return s.description + extra
}
