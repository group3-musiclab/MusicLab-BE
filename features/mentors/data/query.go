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

// SelectAllByRating implements mentors.MentorData
func (mq *mentorQuery) SelectAllByRating() ([]mentors.Core, error) {
	var dataModel []Mentor
	txSelect := mq.db.Raw("SELECT * FROM top_mentor_week ORDER BY avg_rating DESC LIMIT 6").Find(&dataModel)
	if txSelect.Error != nil {
		return nil, errors.New(consts.QUERY_ErrorReadData)
	}
	return ListModelToCore(dataModel), nil
}

// SelectAll implements mentors.MentorData
func (mq *mentorQuery) SelectAll(limit int, offset int, filter mentors.MentorFilter) ([]mentors.Core, error) {
	var dataModel []Mentor

	// if all filter empty
	if filter.Name == "" && filter.Instrument == 0 && filter.Genre == 0 && filter.Rating == 0 && filter.Qualification == "" {
		txSelect := mq.db.Select("id, avatar, name, about, instagram, avg_rating").Limit(limit).Offset(offset).Find(&dataModel)
		if txSelect.Error != nil {
			return nil, errors.New(consts.QUERY_ErrorReadData)
		}
	}

	return ListModelToCore(dataModel), nil
}

// Delete implements mentors.MentorData
func (mq *mentorQuery) Delete(mentorID uint) error {
	tx := mq.db.Delete(&Mentor{}, mentorID)
	if tx.Error != nil {
		return errors.New(consts.QUERY_NotFound)
	}
	if tx.RowsAffected == 0 {
		return errors.New(consts.QUERY_NoRowsAffected)
	}
	return nil
}

// InsertCredential implements mentors.MentorData
func (mq *mentorQuery) InsertCredential(input mentors.CredentialCore) error {
	dataModel := CredentialCoreToModel(input)
	txInsert := mq.db.Create(&dataModel)
	if txInsert.Error != nil {
		return errors.New(consts.QUERY_ErrorInsertData)
	}
	if txInsert.RowsAffected == 0 {
		return errors.New(consts.QUERY_NoRowsAffected)
	}
	return nil
}

// UpdateData implements mentors.MentorData
func (mq *mentorQuery) UpdateData(mentorID uint, input mentors.Core) error {
	dataModel := CoreToModel(input)
	tx := mq.db.Where("id = ?", mentorID).Updates(&dataModel)
	if tx.Error != nil {
		return errors.New(consts.QUERY_ErrorUpdateData)
	}
	if tx.RowsAffected == 0 {
		return errors.New(consts.QUERY_NoRowsAffected)
	}
	return nil
}

// SelectProfile implements mentors.MentorData
func (mq *mentorQuery) SelectProfile(mentorID uint) (mentors.Core, error) {
	var row int64
	txRow := mq.db.Model(&data.Review{}).Where("mentor_id", mentorID).Count(&row)
	if txRow.Error != nil {
		return mentors.Core{}, errors.New(consts.QUERY_ErrorSelect)
	}

	dataModel := Mentor{}
	txSelect := mq.db.First(&dataModel, mentorID)
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
