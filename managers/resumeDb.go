package managers

import (
	"time"

	"github.com/laidingqing/feichong/helpers"
	"github.com/laidingqing/feichong/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	resumeCollectionName = "resumes"
)

// GetResumeByUser 查询用户的简历
func GetResumeByUser(userID string) (models.Resume, error) {
	var resume models.Resume
	query := func(c *mgo.Collection) error {
		return c.Find(bson.M{"userId": userID}).One(&resume)
	}

	executeQuery(resumeCollectionName, query)

	return resume, nil
}

// UpdateResumeByUser 更新用户的简历
func UpdateResumeByUser(userID string, resume models.Resume) (models.Resume, error) {
	rev, err := GetResumeByUser(userID)
	if err != nil {
		return models.Resume{}, err
	}
	log := helpers.NewLogger()
	var exErr error
	log.Log("ID", rev.ID)
	if rev.ID.Hex() == "" {
		rev = resume
		insertQuery := func(c *mgo.Collection) error {
			rev.ID = bson.NewObjectId()
			rev.UserID = userID
			rev.CreatedAt = time.Now()
			rev.IsAuth = false
			return c.Insert(rev)
		}
		exErr = executeQuery(resumeCollectionName, insertQuery)
	} else {
		updateQuery := func(c *mgo.Collection) error {
			rev.UpdateAt = time.Now()
			rev.Name = resume.Name
			rev.Bio = resume.Bio
			rev.Email = resume.Email
			rev.Phone = resume.Phone
			rev.Projects = resume.Projects
			rev.Educations = resume.Educations
			return c.UpdateId(rev.ID, rev)
		}
		exErr = executeQuery(resumeCollectionName, updateQuery)
	}

	return rev, exErr
}
