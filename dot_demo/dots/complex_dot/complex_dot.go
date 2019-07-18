package complex_dot

import (
    "github.com/scryinfo/dot/dot"
)

const (
    ComplexDotTypeId = "8ff771db-14db-46a2-bea0-14f812733834"
)

type ComplexDot struct {
    description string
    config complexDotConfig
}

type complexDotConfig struct {
    ConfigOne string `json:"config_one"`
    ConfigTwo string `json:"config_two"`
}

//construct dot
func newComplexDot(conf interface{}) (dot.Dot, error) {
    var err error
    var bs []byte
    if bt, ok := conf.([]byte); ok {
        bs = bt
    } else {
        return nil, dot.SError.Parameter
    }

    dConf := &complexDotConfig{}
    err = dot.UnMarshalConfig(bs, dConf)
    if err != nil {
        return nil, err
    }

    d := &ComplexDot{config: *dConf}
    return d, nil
}

//Data structure needed when generating newer component
func ComplexDotTypeLive() *dot.TypeLives {
    return &dot.TypeLives{
        Meta: dot.Metadata{TypeId: ComplexDotTypeId,
            NewDoter: func(conf interface{}) (dot dot.Dot, err error) {
                return newComplexDot(conf)
            }},
    }
}

func (c *ComplexDot) Create(l dot.Line) error {
    c.description = "test complex dot description. "
    return nil
}

func (c *ComplexDot) ShowDescription(extra string) string {
    return c.description + c.config.ConfigOne + c.config.ConfigTwo + extra
}
