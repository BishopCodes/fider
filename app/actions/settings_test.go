package actions_test

import (
	"testing"

	"github.com/getfider/fider/app/actions"
	"github.com/getfider/fider/app/models"
	. "github.com/getfider/fider/app/pkg/assert"
)

func TestInvalidUserNames(t *testing.T) {
	RegisterT(t)

	for _, name := range []string{
		"",
		"123456789012345678901234567890123456789012345678901", // 51 chars
	} {

		action := actions.UpdateUserSettings{}
		action.Initialize()
		action.Model.Name = name
		action.Model.AvatarType = models.AvatarTypeGravatar
		result := action.Validate(&models.User{}, services)
		ExpectFailed(result, "name")
	}
}

func TestValidUserNames(t *testing.T) {
	RegisterT(t)

	for _, name := range []string{
		"Jon Snow",
		"Arya",
	} {
		action := actions.UpdateUserSettings{}
		action.Initialize()
		action.Model.Name = name
		action.Model.AvatarType = models.AvatarTypeGravatar
		result := action.Validate(&models.User{}, services)
		ExpectSuccess(result)
	}
}

func TestInvalidSettings(t *testing.T) {
	RegisterT(t)

	for _, settings := range []map[string]string{
		map[string]string{
			"bad_name": "3",
		},
		map[string]string{
			models.NotificationEventNewComment.UserSettingsKeyName: "4",
		},
	} {
		action := actions.UpdateUserSettings{}
		action.Initialize()
		action.Model.Name = "John Snow"
		action.Model.Settings = settings
		result := action.Validate(&models.User{}, services)
		ExpectFailed(result, "settings", "avatarType")
	}
}

func TestValidSettings(t *testing.T) {
	RegisterT(t)

	for _, settings := range []map[string]string{
		nil,
		map[string]string{
			models.NotificationEventNewPost.UserSettingsKeyName:      models.NotificationEventNewPost.DefaultSettingValue,
			models.NotificationEventNewComment.UserSettingsKeyName:   models.NotificationEventNewComment.DefaultSettingValue,
			models.NotificationEventChangeStatus.UserSettingsKeyName: models.NotificationEventChangeStatus.DefaultSettingValue,
		},
		map[string]string{
			models.NotificationEventNewComment.UserSettingsKeyName: models.NotificationEventNewComment.DefaultSettingValue,
		},
	} {
		action := actions.UpdateUserSettings{}
		action.Initialize()
		action.Model.Name = "John Snow"
		action.Model.Settings = settings
		action.Model.AvatarType = models.AvatarTypeGravatar

		result := action.Validate(&models.User{
			AvatarBlobKey: "jon.png",
		}, services)

		ExpectSuccess(result)
		Expect(action.Model.Avatar.BlobKey).Equals("jon.png")
	}
}
