package system

import (
	"errors"
	"new-go-edas-admin/common/global"
	mod "new-go-edas-admin/models/system"
)

type InterfaceCasbin interface {
	AddPolicy(policy []*mod.CasbinPolicy) error
	DelPolicy(policy []*mod.CasbinPolicy) error
	ListPolicy(limit, page int) *mod.CasbinPolicyList
}
type casbinInfo struct {
}

func NewCasbinInterface() InterfaceCasbin {
	return &casbinInfo{}
}

// 添加权限
func (c *casbinInfo) AddPolicy(policy []*mod.CasbinPolicy) error {
	if len(policy) > 0 {
		for _, v := range policy {
			if ok, _ := global.CasbinEnforcer.AddPolicy(v.RoleID, v.Path, v.Method, v.Desc); !ok {
				global.TPLogger.Error("权限已经存在")
				continue
			}

		}
		global.TPLogger.Info("权限添加成功")
		return nil
	}
	global.TPLogger.Error("权限不能为空")
	return errors.New("权限不能为空")
}

// 删除权限
func (c *casbinInfo) DelPolicy(policy []*mod.CasbinPolicy) error {
	if len(policy) > 0 {
		for _, v := range policy {
			if ok, _ := global.CasbinEnforcer.RemovePolicy(v.RoleID, v.Path, v.Method, v.Desc); !ok {
				global.TPLogger.Error("权限不存在")
				continue
			}
		}
		return nil
	}
	global.TPLogger.Error("权限不能为空")
	return errors.New("权限不能为空")
}

// 查看授权
func (c *casbinInfo) ListPolicy(limit, page int) *mod.CasbinPolicyList {
	var (
		policy  mod.CasbinPolicy
		policys []mod.CasbinPolicy
		total   int
	)
	casbinData, _ := global.CasbinEnforcer.GetPolicy()
	total = len(casbinData)
	//组装数据
	for _, v := range casbinData {
		policy.RoleID = v[0]
		policy.Path = v[1]
		policy.Method = v[2]
		policy.Desc = v[3]
		policys = append(policys, policy)
	}
	// 自定义处理分页
	if limit <= 0 || page <= 0 {
		return &mod.CasbinPolicyList{
			Items: policys,
			Total: total,
		}
	}
	/*
		举例1：
		limit 2  page 1  也就是 一页2条数据
		startIndex 0
		endIndex 1
		policys[0:1]
		举例2：
		limit 1  page  1 也即是 一页一条数据
		startIndex 0
		endIndex 1
		policys[0:1]
	*/
	startIndex := (page - 1) * limit
	endIndex := page * limit

	if endIndex > total {
		endIndex = total
	}
	policys = policys[startIndex:endIndex]
	return &mod.CasbinPolicyList{
		Items: policys,
		Total: total,
	}
}
