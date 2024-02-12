package casuserResponse

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/casuser/model"
)

type LoginResponse struct {
	User      model.Casusers `json:"user"`
	Token     string         `json:"token"`
	ExpiresAt int64          `json:"expiresAt"`
}
