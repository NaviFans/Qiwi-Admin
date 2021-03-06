package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

// ===== BEGIN of all query sets

// ===== BEGIN of query set AppQuerySet

// AppQuerySet is an queryset type for App
type AppQuerySet struct {
	db *gorm.DB
}

// NewAppQuerySet constructs new AppQuerySet
func NewAppQuerySet(db *gorm.DB) AppQuerySet {
	return AppQuerySet{
		db: db.Model(&App{}),
	}
}

func (qs AppQuerySet) w(db *gorm.DB) AppQuerySet {
	return NewAppQuerySet(db)
}

// AccountIDEq is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) AccountIDEq(accountID uint) AppQuerySet {
	return qs.w(qs.db.Where("account_id = ?", accountID))
}

// AccountIDGt is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) AccountIDGt(accountID uint) AppQuerySet {
	return qs.w(qs.db.Where("account_id > ?", accountID))
}

// AccountIDGte is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) AccountIDGte(accountID uint) AppQuerySet {
	return qs.w(qs.db.Where("account_id >= ?", accountID))
}

// AccountIDIn is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) AccountIDIn(accountID uint, accountIDRest ...uint) AppQuerySet {
	iArgs := []interface{}{accountID}
	for _, arg := range accountIDRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("account_id IN (?)", iArgs))
}

// AccountIDLt is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) AccountIDLt(accountID uint) AppQuerySet {
	return qs.w(qs.db.Where("account_id < ?", accountID))
}

// AccountIDLte is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) AccountIDLte(accountID uint) AppQuerySet {
	return qs.w(qs.db.Where("account_id <= ?", accountID))
}

// AccountIDNe is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) AccountIDNe(accountID uint) AppQuerySet {
	return qs.w(qs.db.Where("account_id != ?", accountID))
}

// AccountIDNotIn is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) AccountIDNotIn(accountID uint, accountIDRest ...uint) AppQuerySet {
	iArgs := []interface{}{accountID}
	for _, arg := range accountIDRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("account_id NOT IN (?)", iArgs))
}

// All is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) All(ret *[]App) error {
	return qs.db.Find(ret).Error
}

// Count is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) Count() (int, error) {
	var count int
	err := qs.db.Count(&count).Error
	return count, err
}

// Create is an autogenerated method
// nolint: dupl
func (o *App) Create(db *gorm.DB) error {
	return db.Create(o).Error
}

// CreatedAtEq is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) CreatedAtEq(createdAt time.Time) AppQuerySet {
	return qs.w(qs.db.Where("created_at = ?", createdAt))
}

// CreatedAtGt is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) CreatedAtGt(createdAt time.Time) AppQuerySet {
	return qs.w(qs.db.Where("created_at > ?", createdAt))
}

// CreatedAtGte is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) CreatedAtGte(createdAt time.Time) AppQuerySet {
	return qs.w(qs.db.Where("created_at >= ?", createdAt))
}

// CreatedAtLt is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) CreatedAtLt(createdAt time.Time) AppQuerySet {
	return qs.w(qs.db.Where("created_at < ?", createdAt))
}

// CreatedAtLte is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) CreatedAtLte(createdAt time.Time) AppQuerySet {
	return qs.w(qs.db.Where("created_at <= ?", createdAt))
}

// CreatedAtNe is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) CreatedAtNe(createdAt time.Time) AppQuerySet {
	return qs.w(qs.db.Where("created_at != ?", createdAt))
}

// Delete is an autogenerated method
// nolint: dupl
func (o *App) Delete(db *gorm.DB) error {
	return db.Delete(o).Error
}

// Delete is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) Delete() error {
	return qs.db.Delete(App{}).Error
}

// DeletedAtEq is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) DeletedAtEq(deletedAt time.Time) AppQuerySet {
	return qs.w(qs.db.Where("deleted_at = ?", deletedAt))
}

// DeletedAtGt is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) DeletedAtGt(deletedAt time.Time) AppQuerySet {
	return qs.w(qs.db.Where("deleted_at > ?", deletedAt))
}

// DeletedAtGte is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) DeletedAtGte(deletedAt time.Time) AppQuerySet {
	return qs.w(qs.db.Where("deleted_at >= ?", deletedAt))
}

// DeletedAtIsNotNull is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) DeletedAtIsNotNull() AppQuerySet {
	return qs.w(qs.db.Where("deleted_at IS NOT NULL"))
}

// DeletedAtIsNull is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) DeletedAtIsNull() AppQuerySet {
	return qs.w(qs.db.Where("deleted_at IS NULL"))
}

// DeletedAtLt is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) DeletedAtLt(deletedAt time.Time) AppQuerySet {
	return qs.w(qs.db.Where("deleted_at < ?", deletedAt))
}

// DeletedAtLte is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) DeletedAtLte(deletedAt time.Time) AppQuerySet {
	return qs.w(qs.db.Where("deleted_at <= ?", deletedAt))
}

// DeletedAtNe is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) DeletedAtNe(deletedAt time.Time) AppQuerySet {
	return qs.w(qs.db.Where("deleted_at != ?", deletedAt))
}

// GetUpdater is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) GetUpdater() AppUpdater {
	return NewAppUpdater(qs.db)
}

// GroupIDEq is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) GroupIDEq(groupID uint) AppQuerySet {
	return qs.w(qs.db.Where("group_id = ?", groupID))
}

// GroupIDGt is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) GroupIDGt(groupID uint) AppQuerySet {
	return qs.w(qs.db.Where("group_id > ?", groupID))
}

// GroupIDGte is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) GroupIDGte(groupID uint) AppQuerySet {
	return qs.w(qs.db.Where("group_id >= ?", groupID))
}

// GroupIDIn is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) GroupIDIn(groupID uint, groupIDRest ...uint) AppQuerySet {
	iArgs := []interface{}{groupID}
	for _, arg := range groupIDRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("group_id IN (?)", iArgs))
}

// GroupIDLt is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) GroupIDLt(groupID uint) AppQuerySet {
	return qs.w(qs.db.Where("group_id < ?", groupID))
}

// GroupIDLte is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) GroupIDLte(groupID uint) AppQuerySet {
	return qs.w(qs.db.Where("group_id <= ?", groupID))
}

// GroupIDNe is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) GroupIDNe(groupID uint) AppQuerySet {
	return qs.w(qs.db.Where("group_id != ?", groupID))
}

// GroupIDNotIn is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) GroupIDNotIn(groupID uint, groupIDRest ...uint) AppQuerySet {
	iArgs := []interface{}{groupID}
	for _, arg := range groupIDRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("group_id NOT IN (?)", iArgs))
}

// IDEq is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) IDEq(ID uint) AppQuerySet {
	return qs.w(qs.db.Where("id = ?", ID))
}

// IDGt is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) IDGt(ID uint) AppQuerySet {
	return qs.w(qs.db.Where("id > ?", ID))
}

// IDGte is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) IDGte(ID uint) AppQuerySet {
	return qs.w(qs.db.Where("id >= ?", ID))
}

// IDIn is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) IDIn(ID uint, IDRest ...uint) AppQuerySet {
	iArgs := []interface{}{ID}
	for _, arg := range IDRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("id IN (?)", iArgs))
}

// IDLt is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) IDLt(ID uint) AppQuerySet {
	return qs.w(qs.db.Where("id < ?", ID))
}

// IDLte is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) IDLte(ID uint) AppQuerySet {
	return qs.w(qs.db.Where("id <= ?", ID))
}

// IDNe is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) IDNe(ID uint) AppQuerySet {
	return qs.w(qs.db.Where("id != ?", ID))
}

// IDNotIn is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) IDNotIn(ID uint, IDRest ...uint) AppQuerySet {
	iArgs := []interface{}{ID}
	for _, arg := range IDRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("id NOT IN (?)", iArgs))
}

// LevelEq is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) LevelEq(level AppLevel) AppQuerySet {
	return qs.w(qs.db.Where("level = ?", level))
}

// LevelGt is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) LevelGt(level AppLevel) AppQuerySet {
	return qs.w(qs.db.Where("level > ?", level))
}

// LevelGte is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) LevelGte(level AppLevel) AppQuerySet {
	return qs.w(qs.db.Where("level >= ?", level))
}

// LevelIn is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) LevelIn(level AppLevel, levelRest ...AppLevel) AppQuerySet {
	iArgs := []interface{}{level}
	for _, arg := range levelRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("level IN (?)", iArgs))
}

// LevelLt is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) LevelLt(level AppLevel) AppQuerySet {
	return qs.w(qs.db.Where("level < ?", level))
}

// LevelLte is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) LevelLte(level AppLevel) AppQuerySet {
	return qs.w(qs.db.Where("level <= ?", level))
}

// LevelNe is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) LevelNe(level AppLevel) AppQuerySet {
	return qs.w(qs.db.Where("level != ?", level))
}

// LevelNotIn is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) LevelNotIn(level AppLevel, levelRest ...AppLevel) AppQuerySet {
	iArgs := []interface{}{level}
	for _, arg := range levelRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("level NOT IN (?)", iArgs))
}

// Limit is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) Limit(limit int) AppQuerySet {
	return qs.w(qs.db.Limit(limit))
}

// NameEq is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) NameEq(name string) AppQuerySet {
	return qs.w(qs.db.Where("name = ?", name))
}

// NameIn is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) NameIn(name string, nameRest ...string) AppQuerySet {
	iArgs := []interface{}{name}
	for _, arg := range nameRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("name IN (?)", iArgs))
}

// NameNe is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) NameNe(name string) AppQuerySet {
	return qs.w(qs.db.Where("name != ?", name))
}

// NameNotIn is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) NameNotIn(name string, nameRest ...string) AppQuerySet {
	iArgs := []interface{}{name}
	for _, arg := range nameRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("name NOT IN (?)", iArgs))
}

// One is used to retrieve one result. It returns gorm.ErrRecordNotFound
// if nothing was fetched
func (qs AppQuerySet) One(ret *App) error {
	return qs.db.First(ret).Error
}

// OrderAscByAccountID is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) OrderAscByAccountID() AppQuerySet {
	return qs.w(qs.db.Order("account_id ASC"))
}

// OrderAscByCreatedAt is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) OrderAscByCreatedAt() AppQuerySet {
	return qs.w(qs.db.Order("created_at ASC"))
}

// OrderAscByDeletedAt is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) OrderAscByDeletedAt() AppQuerySet {
	return qs.w(qs.db.Order("deleted_at ASC"))
}

// OrderAscByGroupID is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) OrderAscByGroupID() AppQuerySet {
	return qs.w(qs.db.Order("group_id ASC"))
}

// OrderAscByID is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) OrderAscByID() AppQuerySet {
	return qs.w(qs.db.Order("id ASC"))
}

// OrderAscByLevel is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) OrderAscByLevel() AppQuerySet {
	return qs.w(qs.db.Order("level ASC"))
}

// OrderAscByUpdatedAt is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) OrderAscByUpdatedAt() AppQuerySet {
	return qs.w(qs.db.Order("updated_at ASC"))
}

// OrderAscByWalletID is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) OrderAscByWalletID() AppQuerySet {
	return qs.w(qs.db.Order("wallet_id ASC"))
}

// OrderDescByAccountID is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) OrderDescByAccountID() AppQuerySet {
	return qs.w(qs.db.Order("account_id DESC"))
}

// OrderDescByCreatedAt is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) OrderDescByCreatedAt() AppQuerySet {
	return qs.w(qs.db.Order("created_at DESC"))
}

// OrderDescByDeletedAt is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) OrderDescByDeletedAt() AppQuerySet {
	return qs.w(qs.db.Order("deleted_at DESC"))
}

// OrderDescByGroupID is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) OrderDescByGroupID() AppQuerySet {
	return qs.w(qs.db.Order("group_id DESC"))
}

// OrderDescByID is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) OrderDescByID() AppQuerySet {
	return qs.w(qs.db.Order("id DESC"))
}

// OrderDescByLevel is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) OrderDescByLevel() AppQuerySet {
	return qs.w(qs.db.Order("level DESC"))
}

// OrderDescByUpdatedAt is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) OrderDescByUpdatedAt() AppQuerySet {
	return qs.w(qs.db.Order("updated_at DESC"))
}

// OrderDescByWalletID is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) OrderDescByWalletID() AppQuerySet {
	return qs.w(qs.db.Order("wallet_id DESC"))
}

// SetAccountID is an autogenerated method
// nolint: dupl
func (u AppUpdater) SetAccountID(accountID uint) AppUpdater {
	u.fields[string(AppDBSchema.AccountID)] = accountID
	return u
}

// SetCreatedAt is an autogenerated method
// nolint: dupl
func (u AppUpdater) SetCreatedAt(createdAt time.Time) AppUpdater {
	u.fields[string(AppDBSchema.CreatedAt)] = createdAt
	return u
}

// SetDeletedAt is an autogenerated method
// nolint: dupl
func (u AppUpdater) SetDeletedAt(deletedAt *time.Time) AppUpdater {
	u.fields[string(AppDBSchema.DeletedAt)] = deletedAt
	return u
}

// SetGroupID is an autogenerated method
// nolint: dupl
func (u AppUpdater) SetGroupID(groupID uint) AppUpdater {
	u.fields[string(AppDBSchema.GroupID)] = groupID
	return u
}

// SetID is an autogenerated method
// nolint: dupl
func (u AppUpdater) SetID(ID uint) AppUpdater {
	u.fields[string(AppDBSchema.ID)] = ID
	return u
}

// SetLevel is an autogenerated method
// nolint: dupl
func (u AppUpdater) SetLevel(level AppLevel) AppUpdater {
	u.fields[string(AppDBSchema.Level)] = level
	return u
}

// SetName is an autogenerated method
// nolint: dupl
func (u AppUpdater) SetName(name string) AppUpdater {
	u.fields[string(AppDBSchema.Name)] = name
	return u
}

// SetToken is an autogenerated method
// nolint: dupl
func (u AppUpdater) SetToken(token string) AppUpdater {
	u.fields[string(AppDBSchema.Token)] = token
	return u
}

// SetUpdatedAt is an autogenerated method
// nolint: dupl
func (u AppUpdater) SetUpdatedAt(updatedAt time.Time) AppUpdater {
	u.fields[string(AppDBSchema.UpdatedAt)] = updatedAt
	return u
}

// SetWalletID is an autogenerated method
// nolint: dupl
func (u AppUpdater) SetWalletID(walletID uint) AppUpdater {
	u.fields[string(AppDBSchema.WalletID)] = walletID
	return u
}

// SetWebHookURL is an autogenerated method
// nolint: dupl
func (u AppUpdater) SetWebHookURL(webHookURL string) AppUpdater {
	u.fields[string(AppDBSchema.WebHookURL)] = webHookURL
	return u
}

// TokenEq is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) TokenEq(token string) AppQuerySet {
	return qs.w(qs.db.Where("token = ?", token))
}

// TokenIn is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) TokenIn(token string, tokenRest ...string) AppQuerySet {
	iArgs := []interface{}{token}
	for _, arg := range tokenRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("token IN (?)", iArgs))
}

// TokenNe is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) TokenNe(token string) AppQuerySet {
	return qs.w(qs.db.Where("token != ?", token))
}

// TokenNotIn is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) TokenNotIn(token string, tokenRest ...string) AppQuerySet {
	iArgs := []interface{}{token}
	for _, arg := range tokenRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("token NOT IN (?)", iArgs))
}

// Update is an autogenerated method
// nolint: dupl
func (u AppUpdater) Update() error {
	return u.db.Updates(u.fields).Error
}

// UpdateNum is an autogenerated method
// nolint: dupl
func (u AppUpdater) UpdateNum() (int64, error) {
	db := u.db.Updates(u.fields)
	return db.RowsAffected, db.Error
}

// UpdatedAtEq is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) UpdatedAtEq(updatedAt time.Time) AppQuerySet {
	return qs.w(qs.db.Where("updated_at = ?", updatedAt))
}

// UpdatedAtGt is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) UpdatedAtGt(updatedAt time.Time) AppQuerySet {
	return qs.w(qs.db.Where("updated_at > ?", updatedAt))
}

// UpdatedAtGte is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) UpdatedAtGte(updatedAt time.Time) AppQuerySet {
	return qs.w(qs.db.Where("updated_at >= ?", updatedAt))
}

// UpdatedAtLt is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) UpdatedAtLt(updatedAt time.Time) AppQuerySet {
	return qs.w(qs.db.Where("updated_at < ?", updatedAt))
}

// UpdatedAtLte is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) UpdatedAtLte(updatedAt time.Time) AppQuerySet {
	return qs.w(qs.db.Where("updated_at <= ?", updatedAt))
}

// UpdatedAtNe is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) UpdatedAtNe(updatedAt time.Time) AppQuerySet {
	return qs.w(qs.db.Where("updated_at != ?", updatedAt))
}

// WalletIDEq is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) WalletIDEq(walletID uint) AppQuerySet {
	return qs.w(qs.db.Where("wallet_id = ?", walletID))
}

// WalletIDGt is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) WalletIDGt(walletID uint) AppQuerySet {
	return qs.w(qs.db.Where("wallet_id > ?", walletID))
}

// WalletIDGte is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) WalletIDGte(walletID uint) AppQuerySet {
	return qs.w(qs.db.Where("wallet_id >= ?", walletID))
}

// WalletIDIn is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) WalletIDIn(walletID uint, walletIDRest ...uint) AppQuerySet {
	iArgs := []interface{}{walletID}
	for _, arg := range walletIDRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("wallet_id IN (?)", iArgs))
}

// WalletIDLt is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) WalletIDLt(walletID uint) AppQuerySet {
	return qs.w(qs.db.Where("wallet_id < ?", walletID))
}

// WalletIDLte is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) WalletIDLte(walletID uint) AppQuerySet {
	return qs.w(qs.db.Where("wallet_id <= ?", walletID))
}

// WalletIDNe is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) WalletIDNe(walletID uint) AppQuerySet {
	return qs.w(qs.db.Where("wallet_id != ?", walletID))
}

// WalletIDNotIn is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) WalletIDNotIn(walletID uint, walletIDRest ...uint) AppQuerySet {
	iArgs := []interface{}{walletID}
	for _, arg := range walletIDRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("wallet_id NOT IN (?)", iArgs))
}

// WebHookURLEq is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) WebHookURLEq(webHookURL string) AppQuerySet {
	return qs.w(qs.db.Where("web_hook_url = ?", webHookURL))
}

// WebHookURLIn is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) WebHookURLIn(webHookURL string, webHookURLRest ...string) AppQuerySet {
	iArgs := []interface{}{webHookURL}
	for _, arg := range webHookURLRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("web_hook_url IN (?)", iArgs))
}

// WebHookURLNe is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) WebHookURLNe(webHookURL string) AppQuerySet {
	return qs.w(qs.db.Where("web_hook_url != ?", webHookURL))
}

// WebHookURLNotIn is an autogenerated method
// nolint: dupl
func (qs AppQuerySet) WebHookURLNotIn(webHookURL string, webHookURLRest ...string) AppQuerySet {
	iArgs := []interface{}{webHookURL}
	for _, arg := range webHookURLRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("web_hook_url NOT IN (?)", iArgs))
}

// ===== END of query set AppQuerySet

// ===== BEGIN of App modifiers

type appDBSchemaField string

func (f appDBSchemaField) String() string {
	return string(f)
}

// AppDBSchema stores db field names of App
var AppDBSchema = struct {
	ID         appDBSchemaField
	CreatedAt  appDBSchemaField
	UpdatedAt  appDBSchemaField
	DeletedAt  appDBSchemaField
	Name       appDBSchemaField
	Token      appDBSchemaField
	WebHookURL appDBSchemaField
	Level      appDBSchemaField
	WalletID   appDBSchemaField
	GroupID    appDBSchemaField
	AccountID  appDBSchemaField
}{

	ID:         appDBSchemaField("id"),
	CreatedAt:  appDBSchemaField("created_at"),
	UpdatedAt:  appDBSchemaField("updated_at"),
	DeletedAt:  appDBSchemaField("deleted_at"),
	Name:       appDBSchemaField("name"),
	Token:      appDBSchemaField("token"),
	WebHookURL: appDBSchemaField("web_hook_url"),
	Level:      appDBSchemaField("level"),
	WalletID:   appDBSchemaField("wallet_id"),
	GroupID:    appDBSchemaField("group_id"),
	AccountID:  appDBSchemaField("account_id"),
}

// Update updates App fields by primary key
func (o *App) Update(db *gorm.DB, fields ...appDBSchemaField) error {
	dbNameToFieldName := map[string]interface{}{
		"id":           o.ID,
		"created_at":   o.CreatedAt,
		"updated_at":   o.UpdatedAt,
		"deleted_at":   o.DeletedAt,
		"name":         o.Name,
		"token":        o.Token,
		"web_hook_url": o.WebHookURL,
		"level":        o.Level,
		"wallet_id":    o.WalletID,
		"group_id":     o.GroupID,
		"account_id":   o.AccountID,
	}
	u := map[string]interface{}{}
	for _, f := range fields {
		fs := f.String()
		u[fs] = dbNameToFieldName[fs]
	}
	if err := db.Model(o).Updates(u).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return err
		}

		return fmt.Errorf("can't update App %v fields %v: %s",
			o, fields, err)
	}

	return nil
}

// AppUpdater is an App updates manager
type AppUpdater struct {
	fields map[string]interface{}
	db     *gorm.DB
}

// NewAppUpdater creates new App updater
func NewAppUpdater(db *gorm.DB) AppUpdater {
	return AppUpdater{
		fields: map[string]interface{}{},
		db:     db.Model(&App{}),
	}
}

// ===== END of App modifiers

// ===== END of all query sets
