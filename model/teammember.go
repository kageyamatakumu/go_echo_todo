package model

type TeamMember struct {
  ID     uint `json:"id" gorm:"primaryKey"`
  TeamID uint `json:"team_id"`
  UserID uint `json:"user_id"`
  DeleteFlg bool `json:"delete_flg" gorm:"default:false"`
}

type TeamMemberReponse struct {
  TeamID uint `json:"team_id"`
  UserID uint `json:"user_id"`
  DeleteFlg bool `json:"delete_flg" gorm:"default:false"`
}