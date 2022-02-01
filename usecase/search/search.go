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

// Get friends by search id
func (usecase *SearchUseCase) GetFriendsBySearchId(sub string, searchId string) (*[]contract.UserDTO, error) {

	// Verify params
	if sub == "" || searchId == "" {
		return nil, errors.New(constant.ErrorWrongParamsUsecase)
	}

	// Check user
	userDto, err := usecase.repository.UserRepository.GetUserBySub(sub)
	if err != nil {
		return nil, err
	}

	// Search exist
	_, err = usecase.IsSearchExist(searchId)
	if err != nil {

		return nil, err
	}

	// Check search access rights
	ok, err := usecase.hasSearchAccess(userDto.Id, searchId)
	if !ok || err != nil {
		return nil, err
	}

	// Call repository
	friends, err := usecase.repository.SearchRepository.GetFriendsBySearchId(searchId)
	if err != nil {
		return nil, err
	}
	return friends, nil

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
	if existingSearch != nil && existingSearch.Id != "" {
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
	IsSearchOwner := usecase.IsSearchOwner(searchDTO.UserID, searchDTO.Id)
	if !IsSearchOwner {
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

	// Check user is search owner OR friend OR follower
	ok := usecase.IsSearchOwner(postDTO.UserID, postDTO.SearchID)
	if !ok {

		_, err := usecase.IsFriendshipExist(postDTO.SearchID, postDTO.UserID)
		if err != nil {
			_, err := usecase.IsFollowerExist(postDTO.SearchID, postDTO.UserID)
			if err != nil {
				return nil, errors.New(constant.ErrorMissingAccess)
			}
		}

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
	if !ok {
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

func (usecase SearchUseCase) GetSearchByIdForInvitation(sub string, searchId string) (*contract.SearchDTOById, error) {

	// Check params
	if sub == "" || searchId == "" {
		return nil, errors.New(constant.ErrorGetSearchForInvitationParams)
	}

	// Context: get search for invitation fetch search by id without any access checks
	// Check user exist
	_, err := usecase.repository.UserRepository.GetUserBySub(sub)
	if err != nil {
		return nil, err
	}

	// Check search exist
	_, err = usecase.IsSearchExist(searchId)
	if err != nil {
		return nil, err
	}

	// Call repository
	search, err := usecase.repository.SearchRepository.GetSearchById(searchId)
	if err != nil {
		return nil, err
	}

	return search, nil
}

func (usecase SearchUseCase) UpsertFriendship(friendshipDto *contract.FriendshipDTO) (*contract.FriendshipDTO, error) {

	// Check mandatory fields
	if friendshipDto.SearchId == "" || friendshipDto.UserId == "" || friendshipDto.Type == "" {
		return nil, errors.New(constant.ErrorUpsertFriendshipParams)
	}

	// Check user exist
	_, err := usecase.repository.UserRepository.GetUserById(friendshipDto.UserId)
	if err != nil {
		return nil, err
	}

	// Check search exist
	_, err = usecase.IsSearchExist(friendshipDto.SearchId)
	if err != nil {
		return nil, err
	}

	// Check user is not owner of search
	ok := usecase.IsSearchOwner(friendshipDto.UserId, friendshipDto.SearchId)
	if ok {
		return nil, errors.New(constant.ErrorIsSearchOwner)
	}

	// Check friendship exist
	friendship, err := usecase.repository.SearchRepository.IsFriendshipExist(friendshipDto.SearchId, friendshipDto.UserId)

	// If friendship exist return error to avoid duplicate friendships
	if err == nil && friendship.Id != "" {
		return nil, errors.New(constant.ErrorFriendshipAlreadyExist)
	}

	// Check if user is already a follower
	follower, _ := usecase.IsFollowerExist(friendshipDto.SearchId, friendshipDto.UserId)
	if follower != nil && follower.Id != "" {
		return nil, errors.New(constant.ErrorFriendshipFollowerAlreadyExist)
	}

	// Check friendship has been deleted
	friendship, err = usecase.repository.SearchRepository.IsFriendshipDeleted(friendshipDto.SearchId, friendshipDto.UserId)
	if err != nil {
		return nil, err
	}

	// If it has been deleted, re-activate the friendships
	if friendship != nil && friendship.Id != "" {

		// Set id to dto
		friendshipDto.Id = friendship.Id

	}

	// Call repository
	friendshipResponseDto, err := usecase.repository.SearchRepository.SaveFriendship(friendshipDto)
	if err != nil {
		return nil, err
	}

	return friendshipResponseDto, nil
}

func (usecase SearchUseCase) DeleteFriendshipById(sub string, searchId string, friendshipId string) (bool, error) {
	// Check params
	if searchId == "" || friendshipId == "" || sub == "" {
		return false, errors.New(constant.ErrorWrongParamsUsecase)
	}

	// Check user exist
	userDto, err := usecase.repository.UserRepository.GetUserBySub(sub)
	if err != nil {
		return false, err
	}

	// Check search exist
	search, err := usecase.IsSearchExist(searchId)
	if err != nil {
		return false, err
	}

	// Check if user owner of search
	ok := usecase.IsSearchOwner(userDto.Id, searchId)

	// If not owner check if user is friend
	if !ok {

		// Check is user is friend
		_, err := usecase.IsFriendshipExist(search.Id, userDto.Id)
		if err != nil {
			return false, err
		}
	}

	// Call repository
	ok, err = usecase.repository.SearchRepository.DeleteFriendship(friendshipId)
	if err != nil {
		return false, err
	}

	return ok, nil
}

func (usecase SearchUseCase) PostFollower(sub string, searchId string) (*contract.FollowerDTO, error) {

	// Check Params
	if searchId == "" || sub == "" {
		return nil, errors.New(constant.ErrorWrongParamsUsecase)
	}

	// Check user exist
	userDto, err := usecase.repository.UserRepository.GetUserBySub(sub)
	if err != nil {
		return nil, err
	}

	// Check search exist
	search, err := usecase.IsSearchExist(searchId)
	if err != nil {
		return nil, err
	}

	// Check search is public
	if search.Type == constant.SearchTypePrivate {
		return nil, errors.New(constant.ErrorMissingAccess)
	}

	// Check requester is not search owner
	ok := usecase.IsSearchOwner(userDto.Id, searchId)
	if ok {
		return nil, errors.New(constant.ErrorFollowerNotAllowedOwner)
	}

	// Check if follower exist
	follower, _ := usecase.IsFollowerExist(searchId, userDto.Id)
	if follower != nil && follower.Id != "" {
		return nil, errors.New(constant.ErrorFollowerAlreadyExist)
	}

	// Check if friendship exist
	friendship, _ := usecase.IsFriendshipExist(searchId, userDto.Id)
	if friendship != nil && friendship.Id != "" {
		return nil, errors.New(constant.ErrorFollowerNotAllowedFriendship)
	}

	// Build follower dto
	postData := contract.FollowerDTO{
		Id:       "",
		SearchId: searchId,
		UserId:   userDto.Id,
	}

	// Check friendship has been deleted
	follower, err = usecase.repository.SearchRepository.IsFollowerDeleted(postData.SearchId, postData.UserId)
	if err != nil {
		return nil, err
	}

	// If it has been deleted, re-activate the friendships
	if follower != nil && follower.Id != "" {

		// Set id to dto
		postData.Id = follower.Id

	}

	// Call repo
	followerDto, err := usecase.repository.SearchRepository.SaveFollower(postData)
	if err != nil {
		return nil, err
	}

	return followerDto, nil
}

func (usecase SearchUseCase) DeleteFollowerById(sub string, searchId string, followerId string) (bool, error) {

	// Check params
	if sub == "" || searchId == "" || followerId == "" {
		return false, errors.New(constant.ErrorWrongParamsUsecase)
	}

	// Check requester exist
	userDto, err := usecase.repository.UserRepository.GetUserBySub(sub)
	if err != nil {
		return false, err
	}

	// Check search exist
	_, err = usecase.IsSearchExist(searchId)
	if err != nil {
		return false, nil
	}

	// Check follower exist
	follower, err := usecase.IsFollowerExistById(followerId)
	if err != nil {
		return false, err
	}

	// Check requesterId is search owner or follower
	if follower.UserId != userDto.Id {

		ok := usecase.IsSearchOwner(userDto.Id, searchId)
		if !ok {
			return false, errors.New(constant.ErrorMissingAccess)
		}
	}

	// Call repo
	ok, err := usecase.repository.SearchRepository.DeleteFollowerById(followerId)
	if err != nil {
		return false, err
	}

	return ok, nil
}

func (usecase SearchUseCase) GetPublicSearches(sub string) (*[]contract.PublicSearchDTO, error) {
	// Check user
	userDto, err := usecase.repository.UserRepository.GetUserBySub(sub)
	if err != nil {
		return nil, errors.New(constant.ErrorMissingAccess)
	}

	// Call repository
	searches, err := usecase.repository.SearchRepository.GetPublicSearches(userDto.Id)
	if err != nil {
		return nil, err
	}

	return searches, nil
}

func (usecase SearchUseCase) GetSearchRole(sub string, searchId string) (string, error) {

	// Check params
	if sub == "" || searchId == "" {
		return "", errors.New(constant.ErrorWrongParamsUsecase)
	}

	// Check user exist
	userDto, err := usecase.repository.UserRepository.GetUserBySub(sub)
	if err != nil {
		return "", err
	}

	// Check search exist
	search, err := usecase.IsSearchExist(searchId)
	if err != nil {
		return "", err
	}

	// Check requester role
	var requesterRole string

	// Check requester is search owner
	if ok := usecase.IsSearchOwner(userDto.Id, search.Id); ok {
		requesterRole = constant.RequesterRoleOwner
	}

	// Check requester is search follower
	if follower, _ := usecase.IsFollowerExist(searchId, userDto.Id); follower != nil && follower.Id != "" {
		requesterRole = constant.RequesterRoleFollower
	}

	// Check requester is search friend
	if friend, _ := usecase.IsFriendshipExist(search.Id, userDto.Id); friend != nil && friend.Id != "" {
		requesterRole = constant.RequesterRoleFriend
	}

	// Set role to visitor if search is public & requester role unknown
	if requesterRole == "" && search.Type == constant.SearchTypePublic {
		requesterRole = constant.RequesterRoleVisitor
	}

	// If requester role still undefined it means that he shoudn't be able to see the search
	if requesterRole == "" {
		return "", errors.New(constant.ErrorGetRequesterRole)
	}

	return requesterRole, nil
}

func (usecase SearchUseCase) CreateCommentForPost(sub string, request *contract.CommentDTO) (*contract.CreateCommentResponseDto, error) {

	// Check params
	if sub == "" || request == nil || request.Content == "" || request.SearchId == "" || request.PostId == "" {
		return nil, errors.New(constant.ErrorWrongParamsUsecase)
	}

	// Get user
	userDto, err := usecase.repository.UserRepository.GetUserBySub(sub)
	if err != nil {
		return nil, err
	}

	// Check Search exist
	searchDTO, err := usecase.IsSearchExist(request.SearchId)
	if err != nil {
		return nil, err
	}

	// Check Access: is search public or private ? isOwner ? isFriend ?
	ok, err := usecase.hasSearchAccess(userDto.Id, request.SearchId)
	if err != nil || !ok {
		return nil, err
	}

	// Check post exist
	_, err = usecase.IsPostExistForSearch(request.PostId, request.SearchId)
	if err != nil {
		return nil, err
	}

	// Build repository dto
	request.UserId = userDto.Id
	request.SearchId = searchDTO.Id

	// Call repository
	commentCreateDto, err := usecase.repository.SearchRepository.CreateCommentForPost(request)
	if err != nil {
		return nil, err
	}

	return commentCreateDto, nil

}

func (usecase SearchUseCase) UpdateCommentForPost(sub string, request *contract.CommentDTO) (*contract.CommentDTO, error) {

	// Check params
	if sub == "" || request == nil || request.Content == "" || request.SearchId == "" || request.PostId == "" || request.Id == "" {
		return nil, errors.New(constant.ErrorWrongParamsUsecase)
	}

	// Get user
	userDto, err := usecase.repository.UserRepository.GetUserBySub(sub)
	if err != nil {
		return nil, err
	}

	// Check post exist
	_, err = usecase.IsPostExist(request.PostId)
	if err != nil {
		return nil, err
	}

	// Check user is comment owner
	ok, err := usecase.IsCommentOwner(userDto.Id, request.Id)
	if err != nil || !ok {
		return nil, err
	}

	// Call repo
	response, err := usecase.repository.SearchRepository.UpdateCommentForPost(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (usecase SearchUseCase) GetAllCommentsForPost(sub string, searchId string, postId string) (*[]contract.CommentForPostDto, error) {

	// Check params
	if sub == "" || searchId == "" || postId == "" {
		return nil, errors.New(constant.ErrorWrongParamsUsecase)
	}

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

	// Check search exist
	ok, err := usecase.hasSearchAccess(userDto.Id, searchId)
	if err != nil || !ok {
		return nil, err
	}

	// Call repo
	response, err := usecase.repository.SearchRepository.GetAllCommentsForPost(postId)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (usecase SearchUseCase) DeleteCommentForPost(sub string, searchId string, postId string, commentId string )(bool, error){

	// Check Params
	if sub == "" || searchId == "" || postId == "" {
		return false, errors.New(constant.ErrorWrongParamsUsecase)
	}

	// Check user
	userDto, err := usecase.repository.UserRepository.GetUserBySub(sub)
	if err != nil {
		return false, err
	}

	// Check comment exist
	ok, err := usecase.IsCommentExist(commentId)
	if err != nil ||!ok {
		return false, err
	}

	// Check user is comment owner
	ok, err = usecase.IsCommentOwner(userDto.Id, commentId)
	if err != nil || !ok {
		return false, err
	}


	// Call repo
	ok, err = usecase.repository.SearchRepository.DeleteCommentForPost(commentId)
	if err != nil {
		return false, err
	}

	return ok, nil
}