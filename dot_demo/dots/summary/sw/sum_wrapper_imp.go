package sw

import (
    "github.com/scryinfo/dot/dot"
    "github.com/scryinfo/dp_app_dot_demo/dots/complex_dot"
    "github.com/scryinfo/dp_app_dot_demo/dots/simple_dot"
)

type sumWrapperImp struct {
    description string
    One *simple_dot.SimpleDot `dot:"ea383705-9e3e-44c9-b64d-bac82b732c08"`
    Two *complex_dot.ComplexDot `dot:"8ff771db-14db-46a2-bea0-14f812733834"`
}

var _ SumWrapper = (*sumWrapperImp)(nil)

func NewSumWrapper(desc string) (SumWrapper, error) {
    c := &sumWrapperImp{}

    c.description = desc

    //load components
    dot.GetDefaultLine().ToInjecter().Inject(&c)

    return c, nil
}

func (s *sumWrapperImp) ShowDescription(extra string) string {
    return s.description + extra
}

func (s *sumWrapperImp) ShowSimpleDotDesc() string {
    return s.One.ShowDescription(" [call in sum wrapper]")
}

func (s *sumWrapperImp) ShowComplexDotDesc() string {
    return s.Two.ShowDescription(" [call in sum wrapper]")
}
