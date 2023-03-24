package data

import (
	"errors"
	"musiclab-be/features/mentors"
	"musiclab-be/features/reviews/data"
	"musiclab-be/utils/consts"

	"gorm.io/gorm"
)

type mentorQuery struct {
	db *gorm.DB
}

// UpdateData implements mentors.MentorData
func (*mentorQuery) UpdateData(idMentor uint, input mentors.Core) error {
	panic("unimplemented")
}

// SelectProfile implements mentors.MentorData
func (mq *mentorQuery) SelectProfile(idMentor uint) (mentors.Core, error) {
	var row int64
	txRow := mq.db.Model(&data.Review{}).Where("mentor_id", idMentor).Count(&row)
	if txRow.Error != nil {
		return mentors.Core{}, errors.New(consts.QUERY_NotFound)
	}

	dataModel := Mentor{}
	txSelect := mq.db.First(&dataModel, idMentor)
	if txSelect.Error != nil {
		return mentors.Core{}, errors.New(consts.QUERY_NotFound)
	}

	dataModel.CountReviews = row

	return ModelToCore(dataModel), nil
}

func New(db *gorm.DB) mentors.MentorData {
	return &mentorQuery{
		db: db,
	}
}
