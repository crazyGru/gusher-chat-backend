package main

import (
	"backend/startup"
	entity "backend/startup/startup"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

type StartupResponse struct {
	Id              int             `json:"id" example:"1"`
	UserTitle       string          `json:"user_title" example:"Founder"`
	UserDescription string          `json:"user_description" example:"Hyperspace Pilot"`
	UserEquity      float64         `json:"user_equity" example:"10"`
	Name            string          `json:"name" example:"Gusher"`
	Title           string          `json:"title" example:"Gusher"`
	Description     string          `json:"description" example:"Greatest startup engine"`
	Avatar          entity.Image    `json:"avatar"`
	Logo            entity.Image    `json:"logo"`
	Icon            entity.Image    `json:"icon"`
	Cover           entity.Image    `json:"cover"`
	Video           entity.Video    `json:"video"`
	Pitch           string          `json:"pitch" example:"Once upon a time in a galaxy far-far away..."`
	EquityMin       float64         `json:"equity_min" example:"1"`
	EquityTotal     float64         `json:"equity_total" example:"25"`
	Deadline        int             `json:"deadline" example:"10"`
	Vanity          string          `json:"vanity" example:"gusher"`
	CompanyType     string          `json:"company_type" example:"fintech"`
	Followers       int             `json:"followers" example:"100"`
	Applicants      int             `json:"applicants" example:"333"`
	Views           int             `json:"views" example:"76633"`
	Topics          []*entity.Topic `json:"topics"`
}

func NewStartupResponse(s *entity.Startup) StartupResponse {
	return StartupResponse{
		Id:              s.Id(),
		UserTitle:       s.UserTitle(),
		UserDescription: s.UserDescription(),
		UserEquity:      s.UserEquity(),
		Name:            s.Name(),
		Title:           s.Title(),
		Description:     s.Description(),
		Avatar:          s.Avatar(),
		Logo:            s.Logo(),
		Icon:            s.Icon(),
		Cover:           s.Cover(),
		Video:           s.Video(),
		Pitch:           s.Pitch(),
		EquityMin:       s.EquityMin(),
		EquityTotal:     s.EquityTotal(),
		Deadline:        s.Deadline(),
		Vanity:          s.Vanity(),
		CompanyType:     s.CompanyType(),
		Followers:       s.Followers(),
		Applicants:      s.Applicants(),
		Views:           s.Views(),
		Topics:          s.Topics(),
	}
}

func (w *webserver) routesStartup(r *gin.Engine) {
	r.GET("/startups/details", ginHandler(w.startupDetails))
	r.GET("/startups", ginHandler(w.startupList))

}

// @Summary Startup details
// @Param id query int false "Startup ID"
// @Param vanity query string false "Startup Vanity"
// @Produce  json
// @Tags startup
// @Success 200 {object} main.StartupResponse
// @Failure 500 {object} string "Error message"
// @Router /startups/details [get]
func (w *webserver) startupDetails(c *gin.Context) error {
	var search startup.Search
	var err error
	if c.Query("id") != "" {
		search.Id, err = strconv.Atoi(c.Query("id"))
		if err != nil {
			return err
		}
	}
	search.Vanity = c.Query("vanity")
	items, err := w.startup.FindStartups(search)
	if err != nil {
		return err
	}
	if len(items) == 0 {
		return errors.New("Not found")
	}
	c.JSON(200, NewStartupResponse(items[0]))
	return nil
}

// @Summary Startup list
// @Param offset query int false "Offset"
// @Param limit query int false "Limit"
// @Produce  json
// @Tags startup
// @Success 200 {object} []main.ListResponse[StartupResponse]
// @Failure 500 {object} string "Error message"
// @Router /startups [get]
func (w *webserver) startupList(c *gin.Context) error {
	var search startup.Search
	var err error
	if c.Query("offset") != "" {
		search.Offset, err = strconv.Atoi(c.Query("offset"))
		if err != nil {
			return err
		}
	}
	if c.Query("limit") != "" {
		search.Limit, err = strconv.Atoi(c.Query("limit"))
		if err != nil {
			return err
		}
	}

	items, err := w.startup.FindStartups(search)
	if err != nil {
		return err
	}
	cnt, err := w.startup.CountStartups(search)
	if err != nil {
		return err
	}
	var resp ListResponse[StartupResponse]
	resp.Total = cnt

	for _, s := range items {
		resp.Items = append(resp.Items, NewStartupResponse(s))
	}
	c.JSON(200, resp)
	return nil
}
