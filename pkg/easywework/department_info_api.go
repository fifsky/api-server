package workwx

// deptListResp 部门列表响应
type deptListResp struct {
	CommonResp
	Department []*DeptInfo `json:"department"`
}

// deptResp 单个部门信息
type deptResp struct {
	CommonResp
	Department DeptInfo `json:"department"`
}

// execDeptList 获取部门列表
func (c *App) execDeptList(req deptListReq) (deptListResp, error) {

	dsr, err := c.execDeptSimpleList(deptSimpleListReq{
		ID: req.ID,
	})

	if err != nil {
		return deptListResp{}, err
	}
	resp := deptListResp{
		CommonResp: CommonResp{
			ErrCode: 0,
			ErrMsg:  "ok",
		},
		Department: make([]*DeptInfo, 0),
	}

	for _, v := range dsr.DepartmentId {
		sr, err := c.getDept(deptSimpleListReq{
			ID: v.ID,
		})
		if err == nil {
			resp.Department = append(resp.Department, &sr.Department)
		}
	}

	return resp, nil
}

// execDeptList 获取单个部门的信息
func (c *App) getDept(req deptSimpleListReq) (deptResp, error) {
	var resp deptResp
	err := c.executeWXApiGet("/cgi-bin/department/get", req, &resp, true)
	if err != nil {
		return deptResp{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return deptResp{}, bizErr
	}

	return resp, nil
}

// deptSimpleListResp 部门ID列表响应
type deptSimpleListResp struct {
	CommonResp
	DepartmentId []*DeptSimpleInfo `json:"department_id"`
}

// execDeptSimpleList 获取子部门ID列表
// https://developer.work.weixin.qq.com/document/path/95350
func (c *App) execDeptSimpleList(req deptSimpleListReq) (deptSimpleListResp, error) {
	var resp deptSimpleListResp
	err := c.executeWXApiGet("/cgi-bin/department/simplelist", req, &resp, true)
	if err != nil {
		return deptSimpleListResp{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return deptSimpleListResp{}, bizErr
	}

	return resp, nil
}
