package service

import (
	"goapp-gin/entity"
	"goapp-gin/repository"
)

type VideoService interface {
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
	Update(video entity.Video)
	Delete(video entity.Video)
}

type videoService struct {
	// videos []entity.Video
	videoRepository repository.VideoRepository
}

func New(repository repository.VideoRepository) VideoService {
	return &videoService{
		videoRepository: repository,
	}
}
func (service *videoService) Save(video entity.Video) entity.Video {
	service.videoRepository.Save(video)
	return video
}

func (service *videoService) Update(video entity.Video) {
	service.videoRepository.Update(video)
}
func (service *videoService) Delete(video entity.Video) {
	service.videoRepository.Delete(video)
}

func (service *videoService) FindAll() []entity.Video {
	return service.videoRepository.FindAll()
}
