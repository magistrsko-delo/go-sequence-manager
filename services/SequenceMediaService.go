package services

type SequenceMediaService struct {

}

func (sequenceMediaService *SequenceMediaService) GetSequenceMedias(sequenceId int32) (int32, error)  {

	return sequenceId, nil
}

func InitSequenceMediaService() *SequenceMediaService  {
	return &SequenceMediaService{}
}