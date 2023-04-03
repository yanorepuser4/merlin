// Copyright 2020 The Merlin Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package service

import (
	"fmt"

	"gopkg.in/yaml.v2"

	"github.com/caraml-dev/merlin/gitlab"
	"github.com/caraml-dev/merlin/models"
	"github.com/caraml-dev/merlin/storage"
	"github.com/caraml-dev/merlin/warden"
)

// ModelEndpointAlertService interface.
type ModelEndpointAlertService interface {
	ListTeams() ([]string, error)

	ListModelAlerts(modelID models.ID) ([]*models.ModelEndpointAlert, error)
	GetModelEndpointAlert(modelID models.ID, modelEndpointID models.ID) (*models.ModelEndpointAlert, error)
	CreateModelEndpointAlert(user string, alert *models.ModelEndpointAlert) (*models.ModelEndpointAlert, error)
	UpdateModelEndpointAlert(user string, alert *models.ModelEndpointAlert) (*models.ModelEndpointAlert, error)
}

type modelEndpointAlertService struct {
	alertStorage storage.AlertStorage
	gitlabClient gitlab.Client
	wardenClient warden.Client

	dashboardRepository string
	dashboardBranch     string
	alertRepository     string
	alertBranch         string

	dashboardBaseURL string
}

// NewModelEndpointAlertService initializes new alert service.
func NewModelEndpointAlertService(
	alertStorage storage.AlertStorage,
	gitlabClient gitlab.Client, wardenClient warden.Client,
	dashboardRepository, dashboardBranch,
	alertRepository, alertBranch, dashboardBaseURL string) ModelEndpointAlertService {
	return &modelEndpointAlertService{
		alertStorage: alertStorage,
		gitlabClient: gitlabClient,
		wardenClient: wardenClient,

		dashboardRepository: dashboardRepository,
		dashboardBranch:     dashboardBranch,
		alertRepository:     alertRepository,
		alertBranch:         alertBranch,

		dashboardBaseURL: dashboardBaseURL,
	}
}

func (s *modelEndpointAlertService) ListTeams() ([]string, error) {
	return s.wardenClient.GetAllTeams()
}

func (s *modelEndpointAlertService) ListModelAlerts(modelID models.ID) ([]*models.ModelEndpointAlert, error) {
	return s.alertStorage.ListModelEndpointAlerts(modelID)
}

func (s *modelEndpointAlertService) GetModelEndpointAlert(modelID models.ID, modelEndpointID models.ID) (*models.ModelEndpointAlert, error) {
	return s.alertStorage.GetModelEndpointAlert(modelID, modelEndpointID)
}

func (s *modelEndpointAlertService) CreateModelEndpointAlert(user string, alert *models.ModelEndpointAlert) (*models.ModelEndpointAlert, error) {
	commitMessage := fmt.Sprintf("Autogenerated by Merlin: Create alert for %s/%s in %s", alert.Model.Project.Name, alert.Model.Name, alert.EnvironmentName)

	alertSpec := alert.ToPromAlertSpec(s.dashboardBaseURL)
	alertFile, err := yaml.Marshal(&alertSpec)
	if err != nil {
		return nil, err
	}
	alertFilename := fmt.Sprintf("alerts/merlin/%s/%s_%s.yaml", alert.Model.Project.Name, alert.Model.Name, alert.EnvironmentName)

	createAlertOpt := gitlab.CreateFileOptions{
		Repository:    s.alertRepository,
		Branch:        s.alertBranch,
		FileName:      alertFilename,
		Content:       string(alertFile),
		CommitMessage: commitMessage,
		AuthorEmail:   user,
		AuthorName:    user,
	}
	if err := s.gitlabClient.CreateFile(createAlertOpt); err != nil {
		return nil, err
	}

	if err := s.alertStorage.CreateModelEndpointAlert(alert); err != nil {
		return nil, err
	}

	return alert, nil
}

func (s *modelEndpointAlertService) UpdateModelEndpointAlert(user string, alert *models.ModelEndpointAlert) (*models.ModelEndpointAlert, error) {
	commitMessage := fmt.Sprintf("Autogenerated by Merlin: Update alert for %s/%s in %s", alert.Model.Project.Name, alert.Model.Name, alert.EnvironmentName)

	alertSpec := alert.ToPromAlertSpec(s.dashboardBaseURL)
	alertFile, err := yaml.Marshal(&alertSpec)
	if err != nil {
		return nil, err
	}
	alertFilename := fmt.Sprintf("alerts/merlin/%s/%s_%s.yaml", alert.Model.Project.Name, alert.Model.Name, alert.EnvironmentName)

	updateAlertOpt := gitlab.UpdateFileOptions{
		Repository:    s.alertRepository,
		Branch:        s.alertBranch,
		FileName:      alertFilename,
		Content:       string(alertFile),
		CommitMessage: commitMessage,
		AuthorEmail:   user,
		AuthorName:    user,
	}
	if err := s.gitlabClient.UpdateFile(updateAlertOpt); err != nil {
		return nil, err
	}

	if err := s.alertStorage.UpdateModelEndpointAlert(alert); err != nil {
		return nil, err
	}

	return alert, nil
}
