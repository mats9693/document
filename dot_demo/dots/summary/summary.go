package summary

import (
    "fmt"
    "github.com/scryinfo/dot/dot"
    "github.com/scryinfo/dp_app_dot_demo/dots/complex_dot"
    "github.com/scryinfo/dp_app_dot_demo/dots/simple_dot"
    "github.com/scryinfo/dp_app_dot_demo/dots/summary/sw"
)

const (
    SumTypeId = "f74c6b0f-a6b0-4e01-9e5a-231325749463"
)

type Summary struct {
    sumWrapper sw.SumWrapper
    description string
    config summaryConfig
    One *simple_dot.SimpleDot `dot:""`
    Two *complex_dot.ComplexDot `dot:""`
}

type summaryConfig struct {
    SummaryConfigNums int `json:"summary_num"`
}

//construct dot
func newSummaryDot(conf interface{}) (dot.Dot, error) {
    var err error
    var bs []byte
    if bt, ok := conf.([]byte); ok {
        bs = bt
    } else {
        return nil, dot.SError.Parameter
    }

    dConf := &summaryConfig{}
    err = dot.UnMarshalConfig(bs, dConf)
    if err != nil {
        return nil, err
    }

    d := &Summary{config: *dConf}

    return d, err
}

//Data structure needed when generating newer component
func SumTypeLive() []*dot.TypeLives {
    return []*dot.TypeLives{
        {
            Meta: dot.Metadata{TypeId: SumTypeId,
                NewDoter: func(conf interface{}) (dot.Dot, error) {
                    return newSummaryDot(conf)
            }},
        },
        simple_dot.SimpleDotTypeLive(),
        complex_dot.ComplexDotTypeLive(),
    }
}

func (s *Summary) Create(l dot.Line) error {
    s.description = "a summary dot template, use to initialize in start step. "
    return nil
}

func (s *Summary) Start(ignore bool) error {
    s.sumWrapper, _ = sw.NewSumWrapper(s.description)

    fmt.Println("> Summary started. ")

    {
        fmt.Println("> -------test start:-------")
        fmt.Println("> ", instance.ShowDescription(" [test in summary]"))
        fmt.Println("> ", instance.ShowSimpleDotDesc())
        fmt.Println("> ", instance.ShowComplexDotDesc())
        fmt.Println("> -------test end.  -------")
    }

    return nil
}

func (s *Summary) SumWrapper() sw.SumWrapper {
    return s.sumWrapper
}

func (s *Summary) Stop(ignore bool) error {
    return nil
}

func (s *Summary) Destroy(ignore bool) error {
    return nil
}