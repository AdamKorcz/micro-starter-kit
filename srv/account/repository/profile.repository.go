package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/util/log"

	"github.com/xmlking/micro-starter-kit/srv/account/model"
	pb "github.com/xmlking/micro-starter-kit/srv/account/proto/account"
)

// ProfileRepository interface
type ProfileRepository interface {
	List(req *pb.ProfileListQuery) (total uint32, users []*model.Profile, err error)
	Get(req *pb.ProfileRequest) (*model.Profile, error)
	Create(req *pb.ProfileRequest) error
}

// profileRepository struct
type profileRepository struct {
	db *gorm.DB
}

// NewProfileRepository returns an instance of `ProfileRepository`.
func NewProfileRepository(db *gorm.DB) ProfileRepository {
	return &profileRepository{
		db: db,
	}
}

// Exist
func (repo *profileRepository) Exist(req *pb.ProfileRequest) bool {
	var count int
	if req.UserId != 0 {
		repo.db.Model(&model.Profile{}).Where("user_id = ?", req.UserId).Count(&count)
		if count > 0 {
			return true
		}
	}
	return false
}

// List
func (repo *profileRepository) List(req *pb.ProfileListQuery) (total uint32, profiles []*model.Profile, err error) {
	db := repo.db

	var limit, offset uint32
	if req.Limit > 0 {
		limit = req.Limit
	} else {
		limit = 10
	}
	if req.Page > 1 {
		offset = (req.Page - 1) * limit
	} else {
		offset = 0
	}

	var sort string
	if req.Sort != "" {
		sort = req.Sort
	} else {
		sort = "created_at desc"
	}

	if req.UserId != 0 {
		db = db.Where("user_id = ?", req.UserId)
	}
	if req.Gender != "" {
		db = db.Where("gender = ?", req.Gender)
	}

	if err = db.Order(sort).Limit(limit).Offset(offset).Find(&profiles).Count(&total).Error; err != nil {
		log.Logf("Error in ProfileRepository: %v", err)
		return
	}
	return
}

// Find by ID
func (repo *profileRepository) Get(req *pb.ProfileRequest) (profile *model.Profile, err error) {
	// profile = &model.Profile{Model: gorm.Model{ID: uint(req.Id)}}
	profile = &model.Profile{}
	if err = repo.db.First(&profile, req.Id).Error; err != nil && err != gorm.ErrRecordNotFound {
		log.Logf("Error in ProfileRepository: %v", err)
	}
	return
}

// Create
func (repo *profileRepository) Create(req *pb.ProfileRequest) error {
	if exist := repo.Exist(req); exist == true {
		return fmt.Errorf("Profile already exist")
	}

	profile := &model.Profile{
		UserID: req.UserId,
		TZ:     req.Tz,
		Gender: req.Gender,
		Avatar: req.Avatar,
	}

	if err := repo.db.Create(profile).Error; err != nil {
		log.Logf("Error in ProfileRepository: %v", err)
		return err
	}
	return nil
}