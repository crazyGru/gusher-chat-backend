package main

import (
	"backend/access"
	protocolModule "backend/protocol"
	"backend/protocol/protocol"
	"backend/shared"
	"backend/startup"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProtocolResponse struct {
	Id          int              `json:"id" example:"1"`
	Name        string           `json:"name" example:"FinTech"`
	SampleTasks []*protocol.Task `json:"sample_tasks"`
	Total       int              `json:"total" example:"30"`
}

type ProtocolResponseAdmin struct {
	Id    int              `json:"id" example:"1"`
	Name  string           `json:"name" example:"FinTech"`
	Tasks []*protocol.Task `json:"tasks"`
}

type TaskResponse struct {
	Id          int                `json:"id" example:"123"`
	StartupRole shared.StartupRole `json:"startupRole" example:"developer"`
	Title       string             `json:"title" example:"Develop company website"`
	Content     string             `json:"content" example:"html content..."`
}

type ProtocolProgress struct {
	Id       int            `json:"id" example:"1"`
	Name     string         `json:"name" example:"FinTech"`
	Total    int            `json:"total"`
	Complete []TaskResponse `json:"complete"`
	Pending  []TaskResponse `json:"pending"`
}

type assignProtocolRequest struct {
	ProtocolId int `json:"protocol_id"`
}

type assignProtocolRolesRequest struct {
	ProtocolRoles []string `json:"protocol_roles"`
}

type completeTaskRequest struct {
	Complete bool `json:"complete"`
}

func (w *webserver) routesProtocol(r *gin.Engine) {
	r.GET("/protocols", ginHandler(w.protocolsList))
	r.GET("/protocols/:id/roles", ginHandler(w.protocolRoles))
	r.GET("/protocols/:id", ginHandler(w.protocolGet))
	//	r.POST("/protocols", ginHandler(w.requireUser), ginHandler(w.protocolAdminSave))
	//r.GET("/protocols/startup-protocol", ginHandler(w.protocolProgress))
	r.GET("/protocols/progress", ginHandler(w.requireUser), ginHandler(w.protocolProgress))
	r.PATCH("/startup/:startup_id/protocol-task/:task_id", ginHandler(w.requireUser), ginHandler(w.protocolCompleteTask))
	r.POST("/startup/:startup_id/assign-protocol", ginHandler(w.requireUser), ginHandler(w.protocolAssignToStartup))
	r.POST("/startup/:startup_id/role/:role_id/protocol-roles", ginHandler(w.requireUser), ginHandler(w.protocolAssignToStartupRoles))

	r.GET("/admin/protocols", ginHandler(w.requireUser), ginHandler(w.protocolsAdminList))
	r.POST("/admin/protocols", ginHandler(w.requireUser), ginHandler(w.protocolAdminSave))
	r.GET("/admin/protocols/:id", ginHandler(w.requireUser), ginHandler(w.protocolAdminGet))

}

// @Summary Shows list of available protocols
// @Description This will list all available protocols in system including their corresponding tasks and their dependencies.
// @Description Method should not be used by regular users, because we do not want to allow them to see full list of tasks
// @Tags protocol
// @Produce  json
// @Success 200 {object} []ProtocolResponse "Full list of protocols"
// @Failure 500 {object} string "Error message"
// @Router /protocols  [get]
func (w *webserver) protocolsList(c *gin.Context) error {
	var resp []ProtocolResponse
	params := protocolModule.ProtocolSearch{
		Id: 0,
	}
	protocols, err := w.protocol.Find(params)
	for _, p := range protocols {
		resp = append(resp, ProtocolResponse{
			Id:          p.Id(),
			Name:        p.Name,
			SampleTasks: p.SampleTasks(),
			Total:       p.Total(),
		})
	}

	if err != nil {
		return err
	}
	c.JSON(200, resp)
	return nil
}

// @Summary Shows list of available protocols
// @Security ApiKeyAuth
// @Tags protocol, admin
// @Description This will list all available protocols in system including their corresponding tasks and their dependencies.
// @Description Method should not be used by regular users, because we do not want to allow them to see full list of tasks
// @Produce  json
// @Success 200 {object} ListResponse[ProtocolResponseAdmin] "Full list of protocols"
// @Failure 500 {object} string "Error message"
// @Router /admin/protocols  [get]
func (w *webserver) protocolsAdminList(c *gin.Context) error {
	var resp ListResponse[ProtocolResponseAdmin]
	if !w.access.CanManageProtocols(w.identity.UserId) {
		return access.ErrAccessDenied
	}
	params := protocolModule.ProtocolSearch{
		Id: 0,
	}
	protocols, err := w.protocol.Find(params)
	resp.Total = len(protocols)
	for _, p := range protocols {
		resp.Items = append(resp.Items, ProtocolResponseAdmin{
			Id:    p.Id(),
			Name:  p.Name,
			Tasks: p.Tasks,
		})
	}

	if err != nil {
		return err
	}
	c.JSON(200, resp)
	return nil
}

// @Summary Shows protocols details
// @Description This will list all available protocols in system including their corresponding tasks and their dependencies.
// @Description Method should not be used by regular users, because we do not want to allow them to see full list of tasks
// @Tags protocol
// @Param id path int true "Protocols ID"
// @Produce  json
// @Success 200 {object} ProtocolResponse "Protocol details"
// @Failure 500 {object} string "Error message"
// @Router /protocols/{id} [get]
func (w *webserver) protocolGet(c *gin.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	params := protocolModule.ProtocolSearch{
		Id: id,
	}
	protocols, err := w.protocol.Find(params)
	if err != nil {
		return err
	}
	c.JSON(200, ProtocolResponse{
		Id:          protocols[0].Id(),
		Name:        protocols[0].Name,
		SampleTasks: protocols[0].SampleTasks(),
		Total:       protocols[0].Total(),
	})
	return nil
}

// @Summary Shows protocol details with tasks
// @Security ApiKeyAuth
// @Tags protocol, admin
// @Description This will list all available protocols in system including their corresponding tasks and their dependencies.
// @Description Method should not be used by regular users, because we do not want to allow them to see full list of tasks
// @Param id path int true "Protocols ID"
// @Produce  json
// @Success 200 {object} ProtocolResponseAdmin "Protocol details
// "
// @Failure 500 {object} string "Error message"
// @Router /admin/protocols/{id}  [get]
func (w *webserver) protocolAdminGet(c *gin.Context) error {
	if !w.access.CanManageProtocols(w.identity.UserId) {
		return access.ErrAccessDenied
	}
	id, err := strconv.Atoi(c.Param("id"))
	params := protocolModule.ProtocolSearch{
		Id: id,
	}
	protocols, err := w.protocol.Find(params)
	if err != nil {
		return err
	}
	c.JSON(200, ProtocolResponseAdmin{
		Id:    protocols[0].Id(),
		Name:  protocols[0].Name,
		Tasks: protocols[0].Tasks,
	})
	return nil
}

// @Summary Shows list of available roles for specified protocol
// @Param id path int true "Protocol ID"
// @Tags protocol
// @Produce  json
// @Success 200 {object} []shared.StartupRole "All roles available in protocol"
// @Failure 500 {object} string "Error message"
// @Router /protocols/{id}/roles  [get]
func (w *webserver) protocolRoles(c *gin.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return errors.New("incorrect backend id")
	}
	params := protocolModule.ProtocolSearch{Id: id}
	protocols, err := w.protocol.Find(params)
	if err != nil {
		return err
	}
	c.JSON(200, protocols[0].Roles())
	return nil
}

// @Summary Creates or updates protocol in database
// @Param protocol body protocol.Protocol true "Protocol body"
// @Tags protocol, admin
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 500 {object} string "Error message"
// @Router /admin/protocols  [post]
func (w *webserver) protocolAdminSave(c *gin.Context) error {
	var err error
	if !w.access.CanManageProtocols(w.identity.UserId) {
		return access.ErrAccessDenied
	}
	dto := &protocol.ProtocolDto{
		ID:    0,
		Name:  "",
		Tasks: nil,
	}
	err = c.BindJSON(dto)
	p := protocol.FromDto(dto)

	if err != nil {
		return err
	}
	err = w.protocol.Save(p)
	if err != nil {
		return err
	}
	c.JSON(200, gin.H{"id": p.Id()})
	return nil
}

// @Summary Get protocol for startup by its id or vanity
// @Security ApiKeyAuth
// @Tags protocol
// @Param startup_id query int false "Startup ID"
// @Param startup_vanity query string false "Startup Vanity"
// @Param skip_content query string false "1|0 show tasks without content"
// @Produce  json
// @Success 200 {object} main.ProtocolProgress
// @Failure 500 {object} string "Error message"
// @Router /protocols/progress  [get]
func (w *webserver) protocolProgress(c *gin.Context) error {
	// todo total tasks count, allow to get tasks without description(content)
	var err error
	var startupId int
	if c.Query("startup_id") != "" {
		startupId, err = strconv.Atoi(c.Query("startup_id"))
		if err != nil {
			return err
		}
	}
	if c.Query("startup_vanity") != "" {
		s, err := w.startup.FindStartups(startup.Search{Vanity: c.Query("startup_vanity")})
		if err != nil {
			return err
		}
		if len(s) == 0 {
			return errors.New("startup with specified vanity does not exist")
		}
		startupId = s[0].Id()
	}
	protocolId, err := w.startup.StartupProtocolId(startupId)
	if err != nil {
		return err
	}

	pr, err := w.protocol.GetProgress(startupId, protocolId)
	if err != nil {
		return err
	}
	var emptyWhen = func(str string, cond bool) string {
		if cond {
			return ""
		}
		return str
	}
	var resp ProtocolProgress
	for _, t := range pr.PendingTasks(w.startup.UserRolesInStartup(startupId, w.identity.UserId)) {
		resp.Pending = append(resp.Pending, TaskResponse{
			Id:          t.Id,
			StartupRole: t.StartupRole,
			Title:       t.Title,
			Content:     emptyWhen(t.Content, c.Query("skip_content") == "1"),
		})
	}

	for _, t := range pr.CompleteTasks() {
		resp.Complete = append(resp.Complete, TaskResponse{
			Id:          t.Id,
			StartupRole: t.StartupRole,
			Title:       t.Title,
			Content:     emptyWhen(t.Content, c.Query("skip_content") == "1"),
		})
	}
	resp.Total = pr.TotalTasks()
	resp.Id = pr.ProtocolId()
	resp.Name = pr.ProtocolName()

	c.JSON(200, resp)
	return nil
}

// @Summary Mark task as complete
// @Security ApiKeyAuth
// @Tags protocol
// @Param startup_id path int true "Startup ID"
// @Param task_id path int true "Task ID"
// @Param request body completeTaskRequest true "Request data"
// @Accepts  json
// @Produce  json
// @Success 200
// @Failure 500 {object} string "Error message"
// @Router /startup/{startup_id}/protocol-task/{task_id}  [patch]
func (w *webserver) protocolCompleteTask(c *gin.Context) error {
	var request struct {
		StartupId int `uri:"startup_id"`
		TaskId    int `uri:"task_id"`
	}
	var requestBody completeTaskRequest
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		return err
	}

	if err := c.ShouldBindUri(&request); err != nil {
		return err
	}

	if !w.access.CanAccessProtocolTask(w.identity.UserId, request.StartupId, request.TaskId) {
		return access.ErrAccessDenied
	}
	protocolId, err := w.startup.StartupProtocolId(request.StartupId)
	if err != nil {
		return err
	}
	sp, err := w.protocol.GetProgress(request.StartupId, protocolId)
	if err != nil {
		return err
	}
	if requestBody.Complete {
		err = sp.Complete(request.TaskId)
	} else {
		err = sp.NotComplete(request.TaskId)
	}
	if err != nil {
		return err
	}
	if err := w.protocol.SaveProgress(sp); err != nil {
		return err
	}
	c.JSON(200, "ok")
	return nil

}

// @Summary Assign protocol to startup
// @Security ApiKeyAuth
// @Tags protocol
// @Param startup_id path int true "Startup ID"
// @Param request body assignProtocolRequest true "Request body"
// @Accepts  json
// @Produce  json
// @Success 200
// @Failure 500 {object} string "Error message"
// @Router /startup/{startup_id}/assign-protocol  [post]
func (w *webserver) protocolAssignToStartup(c *gin.Context) error {
	var request struct {
		StartupId int `uri:"startup_id"`
	}
	if err := c.ShouldBindUri(&request); err != nil {
		return err
	}
	if !w.access.CanManageStartup(w.identity.UserId, request.StartupId) {
		return access.ErrAccessDenied
	}
	var reqBody assignProtocolRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		return err
	}

	items, err := w.startup.FindStartups(startup.Search{Id: request.StartupId})
	if err != nil {
		return err
	}
	items[0].SetProtocol(reqBody.ProtocolId)
	w.startup.SaveStartup(items[0])
	c.JSON(200, "ok")
	return nil
}

// @Summary Assign protocol-roles to startup-role
// @Security ApiKeyAuth
// @Tags protocol
// @Param startup_id path int true "Startup ID"
// @Param role_id path int true "Role ID"
// @Param request body assignProtocolRolesRequest true "Request body"
// @Accepts  json
// @Produce  json
// @Success 200
// @Failure 500 {object} string "Error message"
// @Router /startup/{startup_id}/role/{role_id}/protocol-roles  [post]
func (w *webserver) protocolAssignToStartupRoles(c *gin.Context) error {
	var request struct {
		StartupId int `uri:"startup_id"`
		RoleId    int `uri:"role_id"`
	}
	if err := c.ShouldBindUri(&request); err != nil {
		return err
	}
	var reqBody assignProtocolRolesRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		return err
	}

	if !w.access.CanManageStartup(w.identity.UserId, request.StartupId) {
		return access.ErrAccessDenied
	}

	r, err := w.startup.FindRoles(startup.RoleSearch{Id: request.RoleId})
	if err != nil {
		return err
	}
	r[0].SetProtocolRole(w.protoRoles(reqBody.ProtocolRoles))
	w.startup.SaveRole(r[0])
	c.JSON(200, "ok")
	return nil
}

func (w *webserver) protoRoles(items []string) []shared.StartupRole {
	var res []shared.StartupRole
	for i := range items {
		res = append(res, shared.StartupRole(items[i]))
	}
	return res
}
