package models

type Grade struct {
  ID        int    `json:"id"`
  StudentID int    `json:"student_id" validate:"required"`
  Course    string `json:"course" validate:"required"`
  Score     int    `json:"score" validate:"min=0,max=100"`
}
