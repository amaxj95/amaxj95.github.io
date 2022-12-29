package models

import "gorm.io/gorm"

type Crunchy_prereq struct {
    gorm.Model
    kubectl string `json:"kubectl version --short" gorm:"text;not null;default:null"`
    git   string    `json:"git --version" gorm:"text;not null;default:null"'
}
