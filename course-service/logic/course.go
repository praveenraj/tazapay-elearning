package logic

import (
	"github.com/jinzhu/gorm"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cast"
	"tazapay.com/elearning/common/models"

	"tazapay.com/elearning/svc/course/constants"

	commonconst "tazapay.com/elearning/common/constants"
	"tazapay.com/elearning/common/repository"
	"tazapay.com/elearning/common/request"
	"tazapay.com/elearning/common/responses"
	"tazapay.com/elearning/common/utils"
)

// GetAllCourses fetch all active courses which satisfy our business condition
func GetAllCourses(payload *request.Payload) *responses.Response {
	courses, err := repository.NewCourseDAO().GetAll()
	if err != nil {
		log.Error().Msgf("error fetching courses: %v", err)
		return responses.InternalError()
	}
	return responses.Success(courses)
}

// GetLessonContent fetch lesson content
func GetLessonContent(payload *request.Payload) *responses.Response {
	// get lesson id from path param
	lessonID, ok := payload.GetAPI().PathParameters[constants.PathParamLessonID]
	if !ok || lessonID == commonconst.Empty {
		log.Warn().Msgf("invalid lesson id")
		return responses.BadRequest()
	}

	// get lesson by lesson id
	lesson, err := repository.NewCourseSectionLessonDAO().GetByID(cast.ToInt(lessonID))
	if err != nil {
		log.Error().Msgf("error fetching lesson content: %v", err)
		return responses.InternalError()
	}
	return responses.Success(lesson)
}

// UpdateLessonContent draft/merge/save lesson content
func UpdateLessonContent(payload *request.Payload) *responses.Response {
	// get lesson id from path param
	lessonID, ok := payload.GetAPI().PathParameters[constants.PathParamLessonID]
	if !ok || lessonID == commonconst.Empty {
		log.Warn().Msgf("invalid lesson id")
		return responses.BadRequest()
	}

	var updateContent constants.UpdateContent
	err := utils.JSONMarshalAndUnmarshal(payload.GetAPI().Body, &updateContent)
	if err != nil {
		log.Error().Msgf("error marshaling update content request: %v", err)
		return responses.InternalError()
	}

	// get lesson by lesson id
	lesson, err := repository.NewCourseSectionLessonDAO().GetByID(cast.ToInt(lessonID))
	if err != nil {
		log.Error().Msgf("error fetching lesson content: %v", err)
		return responses.InternalError()
	}

	// switch through update action
	switch updateContent.Action {
	case constants.ActionDraft:
		return draftLessonContent(lesson, &updateContent)

	case constants.ActionMerge:
		return mergeLessonContent(lesson, &updateContent)

	case constants.ActionSave:
		return saveLessonContent(lesson, &updateContent)

	default:
		return responses.BadRequest()
	}
}

func draftLessonContent(lesson *models.CourseSectionLesson, updateContent *constants.UpdateContent) *responses.Response {
	contentDao := repository.NewCourseSectionLessonContentDAO()

	//  get content by parent id & version
	draftedContent, err := contentDao.GetByParentIDAndVersion(lesson.ContentID, updateContent.Version)
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Error().Msgf("error fetching content data by parentId & version: %v", err)
		return responses.InternalError()
	}

	if draftedContent == nil || (draftedContent != nil && draftedContent.State == models.StateEnumDraft) {
		// TODO: update the lesson content in the cloud storage location
	}

	// add new entry only if it is not exists
	if draftedContent == nil || draftedContent.ID == commonconst.IntZero {
		draft := models.CourseSectionLessonContent{
			Link:              "https://cloud.storage.update.branch." + updateContent.Version,
			Version:           updateContent.Version,
			State:             models.StateEnumDraft,
			TimeRequiredInSec: updateContent.TimeRequired,
			ParentID:          lesson.ContentID,
		}

		// add new lesson content as a drafted one
		err = contentDao.Create(&draft)
		if err != nil {
			log.Error().Msgf("error drafting lesson content: %v", err)
			return responses.InternalError()
		}
	}
	return responses.Success()
}

func mergeLessonContent(lesson *models.CourseSectionLesson, updateContent *constants.UpdateContent) *responses.Response {
	// TODO: needs to check whether the author has permission to merge the changes

	contentDao := repository.NewCourseSectionLessonContentDAO()

	//  get content by parent id & version
	draftedContent, err := contentDao.GetByParentIDAndVersion(lesson.ContentID, updateContent.Version)
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Error().Msgf("error fetching content data by parentId & version: %v", err)
		return responses.InternalError()
	}

	// check whether the content current state is permit to merge the changes
	if draftedContent == nil || (draftedContent != nil && draftedContent.State != models.StateEnumDraft) {
		log.Warn().Msgf("content current state is not permit to merge the changes")
		return responses.Conflict()
	}

	// move the state to merge
	draftedContent.State = models.StateEnumMerged
	err = contentDao.Update(draftedContent, []string{"state"})
	if err != nil {
		log.Error().Msgf("error moving lesson content to merge state: %v", err)
		return responses.InternalError()
	}
	return responses.Success()
}

func saveLessonContent(lesson *models.CourseSectionLesson, updateContent *constants.UpdateContent) *responses.Response {
	// TODO: needs to check whether the author has permission to merge the changes

	contentDao := repository.NewCourseSectionLessonContentDAO()

	//  get content by parent id & version
	draftedContent, err := contentDao.GetByParentIDAndVersion(lesson.ContentID, updateContent.Version)
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Error().Msgf("error fetching content data by parentId & version: %v", err)
		return responses.InternalError()
	}

	// check whether the content current state is permit to save the changes
	if draftedContent == nil || (draftedContent != nil && draftedContent.State != models.StateEnumMerged) {
		log.Warn().Msgf("content current state is not permit to save the changes")
		return responses.Conflict()
	}

	// move the state to save
	draftedContent.State = models.StateEnumSaved
	err = contentDao.Update(draftedContent, []string{"state"})
	if err != nil {
		log.Error().Msgf("error moving lesson content to save state: %v", err)
		return responses.InternalError()
	}

	// update the master branch to loads latest content
	newContent := lesson.Content
	newContent.Link = draftedContent.Link
	newContent.Version = draftedContent.Version
	newContent.TimeRequiredInSec = draftedContent.TimeRequiredInSec
	err = contentDao.Update(newContent, []string{"link", "version", "time_required_in_sec"})
	if err != nil {
		log.Error().Msgf("error updating master branch with new content: %v", err)
		return responses.InternalError()
	}
	return responses.Success()
}
