package data

import (
	"errors"
	"fmt"
	"log"
	"musiclab-be/features/reviews"

	"gorm.io/gorm"
)

type reviewQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) reviews.ReviewData {
	return &reviewQuery{
		db: db,
	}
}

// PostMentorReview implements reviews.ReviewData
func (rq *reviewQuery) PostMentorReview(mentorID uint, newReview reviews.Core) error {
	// res := Mentor{}
	// err := rq.db.Where("id = ?", mentorID).First(&res).Error
	// if err != nil {
	// 	log.Println("query error", err.Error())
	// 	return errors.New("server error")
	// }
	cnv := CoreToData(newReview)
	// err = rq.db.Create(&cnv).Error
	// if err != nil {
	// 	log.Println("query error", err.Error())
	// 	return errors.New("server error")
	// }

	txTransaction := rq.db.Begin()
	if txTransaction.Error != nil {
		txTransaction.Rollback()
		return txTransaction.Error
	}

	tx := txTransaction.Create(&cnv)
	if tx.Error != nil || tx.RowsAffected == 0 {
		txTransaction.Rollback()
		return txTransaction.Error
	}

	var avgRating float32

	tx = txTransaction.Model(&cnv).Where("mentor_id = ?", cnv.MentorID).Select("AVG(rating)").First(&avgRating)
	if tx.Error != nil || tx.RowsAffected == 0 {
		txTransaction.Rollback()
		return txTransaction.Error
	}

	tx = txTransaction.Model(&Mentor{}).Where("id = ?", cnv.MentorID).Update("avg_rating", avgRating)
	if tx.Error != nil || tx.RowsAffected == 0 {
		txTransaction.Rollback()
		return txTransaction.Error
	}

	tx = txTransaction.Commit()
	if tx.Error != nil {
		tx.Rollback()
		return txTransaction.Error
	}

	return nil
}

// GetMentorReview implements reviews.ReviewData
func (rq *reviewQuery) GetMentorReview(limit, offset int, mentorID uint) ([]reviews.Core, error) {
	res := []Review{}
	err := rq.db.Preload("Student").Where("mentor_id = ?", mentorID).Limit(limit).Offset(offset).Order("created_at desc").Find(&res).Error
	if err != nil {
		log.Println("query error", err.Error())
		return []reviews.Core{}, errors.New("server error")
	}

	result := []reviews.Core{}
	i := 0
	for _, val := range res {
		result = append(result, ToCore(val))
		y := res[i].CreatedAt.Year()
		m := int(res[i].CreatedAt.Month())
		d := res[i].CreatedAt.Day()
		result[i].ReviewDate = fmt.Sprintf("%d-%d-%d", y, m, d)

		i++
	}

	return result, nil
}
