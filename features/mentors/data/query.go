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
	txSelect := mq.db.Preload("MentorInstrument.Instrument").Limit(4).Order("avg_rating DESC").Find(&dataModel)
	if txSelect.Error != nil {
		return nil, errors.New(consts.QUERY_ErrorReadData)
	}
	return ListModelToCore(dataModel), nil
}

// SelectAll implements mentors.MentorData
func (mq *mentorQuery) SelectAll(limit int, offset int, filter mentors.MentorFilter) ([]mentors.Core, error) {
	var dataModel []Mentor

	// if all filter empty
	if filter.Name == "" && filter.Instrument == "" && filter.Genre == "" && filter.Rating == 0 && filter.Qualification == "" {
		txSelect := mq.db.Preload("MentorInstrument.Instrument").Limit(limit).Offset(offset).Find(&dataModel)
		if txSelect.Error != nil {
			return nil, errors.New(consts.QUERY_ErrorReadData)
		}
	}

	// if filter not empty

	// if filter.Name != "" {
	// 	nameSearch := "%" + filter.Name + "%"
	// 	name := fmt.Sprintf("name LIKE %s", nameSearch)
	// }

	// if filter.Instrument != "" {
	// 	instrument := fmt.Sprintf("MentorInstrument.Instrument.Name = %s", filter.Instrument)
	// }

	// if filter.Genre != "" {
	// 	genre := fmt.Sprintf("MentorGenre.Genre.Name = %s", filter.Genre)
	// }

	// if filter.Rating != 0 {
	// 	var rating string
	// 	min := filter.Rating
	// 	max := filter.Rating + 0.9
	// 	if filter.Rating == 5 {
	// 		rating = fmt.Sprintf("avg_rating = %f", filter.Rating)
	// 	} else {
	// 		rating = fmt.Sprintf("avg_rating BETWEEN %f AND  %f", min, max)
	// 	}
	// }

	// if filter.Qualification != "" {
	// 	qualification := fmt.Sprintf("Mentor.Credential.Type")
	// }

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
