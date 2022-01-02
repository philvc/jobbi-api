package search_usecase

import (
	"errors"

	constant "github.com/philvc/jobbi-api/constants"
	"github.com/philvc/jobbi-api/contract"
	"github.com/philvc/jobbi-api/repository"
)

type SearchUseCase struct {
	repository repository.Repository
}

// Returns an instance of a search use-case
func Default(repository repository.Repository) SearchUseCase {
	return SearchUseCase{
		repository: repository,
	}
}

// Get My search
func (usecase SearchUseCase) GetMySearch(sub string) (*contract.MySearchDTO, error) {

	// Check user exist
	user, err := usecase.repository.UserRepository.GetUserBySub(sub)

	if err != nil {
		return nil, err
	}

	// Get user search
	response, err := usecase.repository.SearchRepository.GetMySearch(user.Id)

	if err != nil {
		return nil, err
	}

	return response, nil
}

// Get posts by search id
func (usecase SearchUseCase) GetPostsBySearchId(sub string, searchId string) (*[]contract.PostDTOBySearchId, error) {

	// Check user
	userDto, err := usecase.repository.UserRepository.GetUserBySub(sub)
	if err != nil {
		return nil, err
	}

	// Check search exist
	_, err = usecase.IsSearchExist(searchId)
	if err != nil {
		return nil, err
	}

	// Check search access rights
	ok, err := usecase.hasSearchAccess(userDto.Id, searchId)

	if !ok || err != nil {

		return nil, err
	}

	// Get user searches
	response, err := usecase.repository.SearchRepository.GetPostsBySearchId(searchId)

	if err != nil {
		return nil, err
	}

	return response, nil
}

// Get participants by search id
func (usecase SearchUseCase) GetParticipantsBySearchId(sub string, searchId string) (*[]contract.ParticipantDTOForSearchById, error) {

	// Check user
	userDto, err := usecase.repository.UserRepository.GetUserBySub(sub)
	if err != nil {
		return nil, err
	}

	// Check search exist
	_, err = usecase.IsSearchExist(searchId)
	if err != nil {

		return nil, err
	}

	// Check search access rights
	ok, err := usecase.hasSearchAccess(userDto.Id, searchId)
	if !ok || err != nil {
		return nil, err
	}

	// Get search participants
	response, err := usecase.repository.SearchRepository.GetParticipantsBySearchId(searchId)

	if err != nil {
		return nil, err
	}

	return response, nil
}

// Get Shared searches
func (usecase SearchUseCase) GetSharedSearches(sub string) (*[]contract.SharedSearchDTO, error) {

	// Check user
	user, err := usecase.repository.UserRepository.GetUserBySub(sub)
	if err != nil {
		return nil, err
	}

	// Get user searches
	response, err := usecase.repository.SearchRepository.GetSharedSearches(user.Id)

	if err != nil {
		return nil, err
	}

	return response, nil
}

// Get Followed searches
func (usecase SearchUseCase) GetFollowedSearches(sub string) (*[]contract.FollowedSearchDTO, error) {

	// Check user
	user, err := usecase.repository.UserRepository.GetUserBySub(sub)
	if err != nil {
		return nil, err
	}

	// Get user searches
	response, err := usecase.repository.SearchRepository.GetFollowedSearches(user.Id)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (usecase SearchUseCase) GetSearchById(searchId string, sub string) (*contract.SearchDTOById, error) {

	// Check user exist
	userDto, err := usecase.repository.UserRepository.GetUserBySub(sub)
	if err != nil {
		return nil, err
	}

	// Check search exist
	_, err = usecase.IsSearchExist(searchId)
	if err != nil {

		return nil, err
	}

	// Check access rights: user is owner or friend or follower or search i public
	_, err = usecase.hasSearchAccess(userDto.Id, searchId)
	if err != nil {
		return nil, err
	}

	search, err := usecase.repository.SearchRepository.GetSearchById(searchId)
	return search, err
}

func (usecase SearchUseCase) AddSearch(sub string, searchDTO contract.SearchDTO) (*contract.SearchDTO, error) {

	if searchDTO.Title == "" {
		return nil, errors.New(constant.ErrorMissingTitle)
	}

	if searchDTO.Description == "" {
		return nil, errors.New(constant.ErrorMissingDescription)
	}

	if searchDTO.Type == "" {
		return nil, errors.New(constant.ErrorMissingType)
	}

	if searchDTO.UserID == "" {
		return nil, errors.New(constant.ErrorMissingUserId)
	}

	// Check user exist
	_, err := usecase.repository.UserRepository.GetUserBySub(sub)
	if err != nil {
		return nil, err
	}

	// Check if user has already an existing search
	existingSearch, _ := usecase.repository.SearchRepository.GetMySearch(searchDTO.UserID)
	if existingSearch.Id != "" {
		return nil, errors.New(constant.ErrorAlreadyExistingSearch)
	}

	// Call repository
	newSearch, err := usecase.repository.SearchRepository.AddSearch(searchDTO)

	return newSearch, err
}

func (usecase SearchUseCase) ModifySearch(searchDTO contract.SearchDTO) (*contract.SearchDTO, error) {

	// Check mandatory fields
	if searchDTO.Id == "" {
		return nil, errors.New(constant.ErrorMissingSearchId)
	}

	if searchDTO.UserID == "" {
		return nil, errors.New(constant.ErrorMissingUserId)
	}

	if searchDTO.Type == "" {
		return nil, errors.New(constant.ErrorMissingType)
	}

	// Check search exit
	_, err := usecase.IsSearchExist(searchDTO.Id)
	if err != nil {
		return nil, err
	}

	// Check access rights
	isOwner := usecase.IsOwner(searchDTO.UserID, searchDTO.Id)
	if !isOwner {
		return nil, errors.New(constant.ErrorMissingAccess)
	}

	// Call repository
	search, err := usecase.repository.SearchRepository.ModifySearch(searchDTO)

	return search, err
}

func (usecase SearchUseCase) AddPost(postDTO *contract.PostDTO) (*contract.AddPostResponseDTO, error) {

	// Check mandatory fields
	if postDTO.SearchID == "" {
		return nil, errors.New(constant.ErrorMissingSearchId)
	}

	if postDTO.UserID == "" {
		return nil, errors.New(constant.ErrorMissingUserId)
	}
	if postDTO.Title == "" {
		return nil, errors.New(constant.ErrorMissingTitle)
	}
	if postDTO.Description == "" {
		return nil, errors.New(constant.ErrorMissingDescription)
	}

	if postDTO.Type == "" {
		return nil, errors.New(constant.ErrorMissingType)
	}

	// Check search exist
	_, err := usecase.IsSearchExist(postDTO.SearchID)
	if err != nil {
		return nil, err
	}

	// Check user access rights to search
	ok, err := usecase.hasSearchAccess(postDTO.UserID, postDTO.SearchID)
	if !ok || err != nil {

		return nil, err
	}

	// Call repository
	postResponseDto, err := usecase.repository.SearchRepository.AddPost(postDTO)
	if err != nil {
		return nil, err
	}

	return postResponseDto, nil
}

func (usecase SearchUseCase) UpdatePostById(postDTO *contract.PostDTO) (*contract.UpdatePostResponseDTO, error) {

	// Check mandatory fields
	if postDTO.Id == "" {
		return nil, errors.New(constant.ErrorMissingPostId)
	}
	if postDTO.SearchID == "" {
		return nil, errors.New(constant.ErrorMissingSearchId)
	}

	if postDTO.UserID == "" {
		return nil, errors.New(constant.ErrorMissingUserId)
	}
	if postDTO.Title == "" {
		return nil, errors.New(constant.ErrorMissingTitle)
	}

	if postDTO.Type == "" {
		return nil, errors.New(constant.ErrorMissingType)
	}

	// Check search exist
	_, err := usecase.IsSearchExist(postDTO.SearchID)
	if err != nil {
		return nil, err
	}
	
	// Check post exist
	_, err = usecase.IsPostExist(postDTO.Id)
	if err != nil {
		return nil, err
	}

	// Check user is post owner
	ok := usecase.IsPostOwner(postDTO.UserID, postDTO.Id)
	if !ok  {
		return nil, errors.New(constant.ErrorMissingAccess)
	}

	// Update post
	updatePostResponse, err := usecase.repository.SearchRepository.UpdatePost(postDTO)
	if err != nil {

		return nil, err
	}

	return updatePostResponse, nil
}

func (usecase SearchUseCase) DeletePostById(sub string, searchId string, postId string) (bool, error) {

	// Check mandatory fields
	if postId == "" {
		return false, errors.New(constant.ErrorMissingPostId)
	}
	if searchId == "" {
		return false, errors.New(constant.ErrorMissingSearchId)
	}

	if sub == "" {
		return false, errors.New(constant.ErrorMissingUserId)
	}

	// Check search exist
	_, err := usecase.IsSearchExist(searchId)
	if err != nil {
		return false, err
	}

	// Check user exist
	userDto, err := usecase.repository.UserRepository.GetUserBySub(sub)
	if err != nil {
		return false, err
	}

	// Check user is owner of post
	ok := usecase.IsPostOwner(userDto.Id, postId)
	if !ok {
		return false, errors.New(constant.ErrorMissingAccess)
	}

	// Call repository
	ok, err = usecase.repository.SearchRepository.DeletePostById(postId)
	if err != nil {
		return false, err
	}

	return ok, nil

}

func (usecase SearchUseCase) IsPostOwner(userId string, postId string) bool {

	ok := usecase.repository.SearchRepository.IsPostOwner(userId, postId)

	return ok
}

func (usecase SearchUseCase) IsOwner(userId string, searchId string) bool {

	ok := usecase.repository.SearchRepository.IsSearchOwner(userId, searchId)

	return ok
}

func (usecase SearchUseCase) IsPublic(searchId string) bool {

	ok := usecase.repository.SearchRepository.IsPublic(searchId)

	return ok
}

func (usecase SearchUseCase) IsFriend(userId string, searchId string) bool {

	ok := usecase.repository.SearchRepository.IsFriend(userId, searchId)

	return ok
}

func (usecase SearchUseCase) IsSearchExist(searchId string) (*contract.SearchDTO, error) {
	search, err := usecase.repository.SearchRepository.IsSearchExist(searchId)
	if err != nil {
		return nil, err
	}

	return search, nil
}

func (usecase SearchUseCase) IsPostExist(postId string) (*contract.PostDTO, error) {
	post, err := usecase.repository.SearchRepository.IsPostExist(postId)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (usecase SearchUseCase) hasSearchAccess(userId string, searchId string) (bool, error) {

	// Check if search is public
	isPublic := usecase.IsPublic(searchId)
	if !isPublic {

		// Check if user is owner
		isOwner := usecase.IsOwner(userId, searchId)
		if !isOwner {

			// Check if user is friend or follower
			isFriend := usecase.IsFriend(userId, searchId)

			if !isFriend {
				return false, errors.New(constant.ErrorMissingAccess)
			}
		}
	}
	return true, nil
}
