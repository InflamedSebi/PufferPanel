package models

import (
	"github.com/pufferpanel/pufferpanel/errors"
	"gopkg.in/oauth2.v3"
	"strconv"
	"time"
)

type TokenInfo struct {
	ID uint `json:"-"`

	ClientInfoID uint       `gorm:"NOT NULL" json:"-"`
	ClientInfo   ClientInfo `gorm:"save_associations:false" json:"-"`

	//Scope            string
	//Code             string
	//CodeCreateAt     time.Time
	//CodeExpiresIn    time.Duration
	Access          string        `gorm:"UNIQUE_INDEX" json:"-"`
	AccessCreateAt  time.Time     `json:"-"`
	AccessExpiresIn time.Duration `json:"-"`
	//Refresh          string
	//RefreshCreateAt  time.Time
	//RefreshExpiresIn time.Duration
}

func (ti *TokenInfo) New() oauth2.TokenInfo {
	return &TokenInfo{}
}

func (ti *TokenInfo) GetClientID() string {
	return ti.ClientInfo.ClientID
}

func (ti *TokenInfo) SetClientID(clientId string) {
	ti.ClientInfo.ClientID = clientId
}

func (ti *TokenInfo) GetUserID() string {
	return strconv.Itoa(int(ti.ClientInfo.UserID))
}

func (ti *TokenInfo) SetUserID(id string) {
	result, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}
	if result < 0 {
		panic(errors.ErrFieldTooSmall("node id", 1))
	}
	ti.ClientInfo.UserID = uint(result)
}

func (ti *TokenInfo) GetRedirectURI() string {
	return ""
}

func (ti *TokenInfo) SetRedirectURI(string) {
	//NO-OP
}

func (ti *TokenInfo) GetScope() string {
	//return ti.Scope
	return ""
}

func (ti *TokenInfo) SetScope(scope string) {
	//ti.Scope = scope
}

func (ti *TokenInfo) GetCode() string {
	//return ti.Code
	return ""
}

func (ti *TokenInfo) SetCode(code string) {
	//ti.Code = code
}

func (ti *TokenInfo) GetCodeCreateAt() time.Time {
	//return ti.CodeCreateAt
	return time.Now()
}

func (ti *TokenInfo) SetCodeCreateAt(time time.Time) {
	//ti.CodeCreateAt = time
}

func (ti *TokenInfo) GetCodeExpiresIn() time.Duration {
	//return ti.CodeExpiresIn
	return 0
}

func (ti *TokenInfo) SetCodeExpiresIn(dur time.Duration) {
	//ti.CodeExpiresIn = dur
}

func (ti *TokenInfo) GetAccess() string {
	return ti.Access
}

func (ti *TokenInfo) SetAccess(access string) {
	ti.Access = access
}

func (ti *TokenInfo) GetAccessCreateAt() time.Time {
	return ti.AccessCreateAt
}

func (ti *TokenInfo) SetAccessCreateAt(t time.Time) {
	ti.AccessCreateAt = t
}

func (ti *TokenInfo) GetAccessExpiresIn() time.Duration {
	return ti.AccessExpiresIn
}

func (ti *TokenInfo) SetAccessExpiresIn(dur time.Duration) {
	ti.AccessExpiresIn = dur
}

func (ti *TokenInfo) GetRefresh() string {
	//return ti.Refresh
	return ""
}

func (ti *TokenInfo) SetRefresh(ref string) {
	//ti.Refresh = ref
}

func (ti *TokenInfo) GetRefreshCreateAt() time.Time {
	//return ti.RefreshCreateAt
	return time.Now()
}

func (ti *TokenInfo) SetRefreshCreateAt(t time.Time) {
	//ti.RefreshCreateAt = t
}

func (ti *TokenInfo) GetRefreshExpiresIn() time.Duration {
	//return ti.RefreshExpiresIn
	return 0
}

func (ti *TokenInfo) SetRefreshExpiresIn(dur time.Duration) {
	//ti.RefreshExpiresIn = dur
}

func Copy(info oauth2.TokenInfo) *TokenInfo {
	userId, _ := strconv.Atoi(info.GetUserID())
	return &TokenInfo{
		ClientInfo: ClientInfo{
			ClientID: info.GetClientID(),
			UserID:   uint(userId),
		},
		//Scope: info.GetScope(),
		//Code: info.GetCode(),
		//CodeCreateAt: info.GetCodeCreateAt(),
		//CodeExpiresIn: info.GetCodeExpiresIn(),
		Access:          info.GetAccess(),
		AccessCreateAt:  info.GetAccessCreateAt(),
		AccessExpiresIn: info.GetAccessExpiresIn(),
		//Refresh: info.GetRefresh(),
		//RefreshCreateAt: info.GetRefreshCreateAt(),
		//RefreshExpiresIn: info.GetRefreshExpiresIn(),
	}
}
