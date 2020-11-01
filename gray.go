package grayhit

import "errors"

type GrayPoint interface {
	GetName() string
	GetValue() string
}

type DivMethod interface {
	// locate within [0,100)
	CalcIndicator() int
}

type ABTestPolicy struct {
	Layer     Layer
	GrayRules []GrayRule
	DivRule   DivRule
}

func (policy *ABTestPolicy) HitGray(grayPoint GrayPoint) bool {
	for _, grayRule := range policy.GrayRules {
		if grayRule.Enable && grayPoint.GetName() == grayRule.Name {

			return grayRule.IsHit(grayPoint.GetValue())
		}
	}

	return false
}

func (policy *ABTestPolicy) HitDiv(div DivMethod) (bool, error) {
	return policy.DivRule.IsHit(div.CalcIndicator())
}

type Layer struct {
	ID   string
	Data string
}

/**
 * 灰度规则规范
 * eg: {"name":"uid","enabled":true,"include":[],"exclude":[],"global":false}
 */
type GrayRule struct {
	Name    string
	Enable  bool
	Include []string
	Exclude []string
	Global  bool
}

/**
 * 校验是否命中灰度规则
 * 1、先看是否在黑名单中，若在，则永远不会命中灰度，即使开全量也不行
 * 2、再看是否在白名单中，若在，则必定命中灰度
 * 3、若未命中黑、白名单，则看是否已全量
 */
func (rule GrayRule) IsHit(value string) bool {
	if len(rule.Exclude) > 0 && contains(rule.Exclude, value) {
		return false
	}

	if len(rule.Include) > 0 && contains(rule.Include, value) {
		return true
	}

	return rule.Global
}

type DivRule struct {
	Percent int // [0,100)
}

func (rule DivRule) IsHit(value int) (bool, error) {
	if value < 0 || value > 100 {
		return false, errors.New("invalid div range")
	}
	return value <= rule.Percent, nil
}

func contains(arr []string, s string) bool {
	for _, item := range arr {
		if item == s {
			return true
		}
	}

	return false
}
