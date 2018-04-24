package alert_notifications

import (
	"github.com/grafana/grafana/pkg/models"
)
import "github.com/grafana/grafana/pkg/components/simplejson"

type configVersion struct {
	ApiVersion int64 `json:"apiVersion" yaml:"apiVersion"`
}

type notificationsAsConfig struct {
	ApiVersion int64

	Notifications       []*notificationFromConfig
	DeleteNotifications []*deleteNotificationConfig
}

type deleteNotificationConfig struct {
	Name  string
	OrgId int64
}

type notificationFromConfig struct {
	OrgId int64

	Name      string
	Type      string
	IsDefault bool
	Settings  map[string]interface{}
}

type notificationsAsConfigV0 struct {
	configVersion

	Notifications       []*notificationFromConfigV0   `json:"alert_notifications" yaml:"alert_notifications"`
	DeleteNotifications []*deleteNotificationConfigV0 `json:"delete_alert_notifications" yaml:"delete_alert_notifications"`
}

type deleteNotificationConfigV0 struct {
	Name  string `json:"name" yaml:"name"`
	OrgId int64  `json:"org_id" yaml:"org_id"`
}

type notificationFromConfigV0 struct {
	OrgId int64 `json:"org_id" yaml:"org_id"`

	Name      string                 `json:"name" yaml:"name"`
	Type      string                 `json:"type" yaml:"type"`
	IsDefault bool                   `json:"is_default" yaml:"is_default"`
	Settings  map[string]interface{} `json:"settings" yaml:"settings"`
}

func (cfg *notificationsAsConfigV0) mapToNotificationFromConfig(apiVersion int64) *notificationsAsConfig {
	r := &notificationsAsConfig{}

	r.ApiVersion = apiVersion

	if cfg == nil {
		return r
	}

	for _, notification := range cfg.Notifications {
		r.Notifications = append(r.Notifications, &notificationFromConfig{
			OrgId:     notification.OrgId,
			Name:      notification.Name,
			Type:      notification.Type,
			IsDefault: notification.IsDefault,
			Settings:  notification.Settings,
		})
	}

	for _, notification := range cfg.DeleteNotifications {
		r.DeleteNotifications = append(r.DeleteNotifications, &deleteNotificationConfig{
			OrgId: notification.OrgId,
			Name:  notification.Name,
		})
	}

	return r
}

func createInsertCommand(notification *notificationFromConfig) *models.CreateAlertNotificationCommand {
	settings := simplejson.New()
	if len(notification.Settings) > 0 {
		for k, v := range notification.Settings {
			settings.Set(k, v)
		}
	}

	return &models.CreateAlertNotificationCommand{
		Name:      notification.Name,
		Type:      notification.Type,
		IsDefault: notification.IsDefault,
		Settings:  settings,
		OrgId:     notification.OrgId,
	}
}

func createUpdateCommand(notification *notificationFromConfig, id int64) *models.UpdateAlertNotificationCommand {
	settings := simplejson.New()
	if len(notification.Settings) > 0 {
		for k, v := range notification.Settings {
			settings.Set(k, v)
		}
	}

	return &models.UpdateAlertNotificationCommand{
		Id:        id,
		Name:      notification.Name,
		Type:      notification.Type,
		IsDefault: notification.IsDefault,
		Settings:  settings,
		OrgId:     notification.OrgId,
	}
}
