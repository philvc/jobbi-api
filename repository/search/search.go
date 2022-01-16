package search_repository

import (
	"errors"
	"time"

	"github.com/google/uuid"
	constant "github.com/philvc/jobbi-api/constants"
	contract "github.com/philvc/jobbi-api/contract"
	"github.com/philvc/jobbi-api/database/model"
	"gorm.io/gorm"
)

type SearchRepository struct {
	database *gorm.DB
}

func Default(db *gorm.DB) SearchRepository {
	return SearchRepository{
		database: db,
	}
}

func (repository SearchRepository) GetMySearch(userId string) (*contract.MySearchDTO, error) {

	var result contract.MySearchDTO

	// Get search by user id
	if err := repository.database.
		Model(&model.Search{}).
		Select("title, id, sector, tags, type").
		Where("user_id = ? ", userId).
		Where("deleted_at IS NULL").
		Scan(&result).
		Error; err != nil {
		print(err)
		return nil, errors.New(constant.ErrorGetMySearch)
	}

	// Return nil if no search
	if result.Id == "" {
		return nil, nil
	}

	// Get friends
	if err := repository.database.Model(&model.Friendship{}).
		Where("search_id = ? ", result.Id).
		Where("friendships.deleted_at IS NULL").
		Joins("JOIN users ON users.id = friendships.user_id").
		Select("users.id, users.first_name, users.last_name, users.email, users.avatar_url").
		Find(&result.Friends).Error; err != nil {
		return nil, errors.New(constant.ErrorGetMySearchParticipants)
	}

	return &result, nil
}

func (repository SearchRepository) GetSharedSearches(userId string) (*[]contract.SharedSearchDTO, error) {

	var results []contract.SharedSearchDTO

	if err := repository.database.
		Model(&model.Friendship{}).
		Where("friendships.user_id = ?", userId).
		Where("friendships.deleted_at IS NULL").

		// GET Searches details
		Joins("JOIN searches ON searches.id = friendships.search_id").

		// Get search owner details
		Joins("JOIN users ON users.id = searches.user_id").
		Select("searches.id, searches.title, searches.type, searches.sector, searches.tags, searches.description, searches.user_id, users.first_name, users.last_name, users.avatar_url").
		Find(&results).
		Error; err != nil {
		return nil, errors.New(constant.ErrorGetSharedSearches)
	}

	return &results, nil
}

func (repository SearchRepository) GetFollowedSearches(userId string) (*[]contract.FollowedSearchDTO, error) {

	var results []contract.FollowedSearchDTO

	if err := repository.database.
		Model(&model.Follower{}).
		Where("followers.user_id = ?", userId).
		Where("followers.deleted_at IS NULL").

		// GET Searches details
		Joins("JOIN searches ON searches.id = followers.search_id").

		// Get search owner details
		Joins("JOIN users ON users.id = searches.user_id").
		Select("searches.id, searches.sector, searches.type, searches.title, searches.tags, searches.description, searches.user_id, users.first_name, users.last_name, users.avatar_url").
		Find(&results).
		Error; err != nil {
		return nil, errors.New(constant.ErrorGetFollowedSearches)
	}

	return &results, nil
}

func (repository SearchRepository) GetSearchById(searchId string) (*contract.SearchDTOById, error) {
	var result contract.SearchDTOById

	if err := repository.database.
		Model(&model.Search{}).
		Where("searches.id = ?", searchId).
		Where("searches.deleted_at IS NULL").
		Select("title, searches.id, sector, tags, description, searches.type, users.email, users.id as user_id, users.first_name, users.last_name, users.avatar_url").
		Joins("JOIN users ON users.id = searches.user_id").
		Scan(&result).
		Error; err != nil {
		return nil, errors.New(constant.ErrorGetSearchById)
	}

	// Return nil is no search
	if result.Id == "" {
		return nil, nil
	}

	// Get friends
	if err := repository.database.
		Model(&model.Friendship{}).
		Where("search_id = ? ", result.Id).
		Where("friendships.deleted_at IS NULL").
		Joins("JOIN users ON users.id = friendships.user_id").
		Select("users.id, users.first_name, users.last_name, users.email, users.avatar_url").
		Find(&result.Friends).Error; err != nil {
		return nil, errors.New(constant.ErrorGetMySearchParticipants)
	}

	return &result, nil
}

func (repository SearchRepository) GetFriendsBySearchId(searchId string) (*[]contract.UserDTO, error) {

	var friends []model.User
	if err := repository.database.Model(model.Friendship{}).Where("friendships.search_id = ? AND friendships.deleted_at IS NULL", searchId).Joins("JOIN users ON users.id = friendships.user_id").Select("users.id as id, users.first_name, users.last_name, users.email, users.avatar_url").Find(&friends).Error; err != nil {
		return nil, errors.New(constant.ErrorGetFriendsBySearchId)
	}

	friendDtos := model.ToUserDTOs(friends)

	return &friendDtos, nil
}

func (repository SearchRepository) GetPostsBySearchId(searchId string) (*[]contract.PostDTOBySearchId, error) {

	var results []contract.PostDTOBySearchId

	if err := repository.database.
		Model(&model.Post{}).
		Where("posts.search_id = ?", searchId).
		Where("posts.deleted_at IS NULL").
		Order("posts.created_at DESC").
		Joins("JOIN users ON users.id = posts.user_id").
		Select("users.id as user_id, posts.updated_at, posts.id, users.email as user_email, users.first_name as user_first_name, posts.search_id, users.last_name as user_last_name, title, description, posts.id as id, posts.type, tags, contact_first_name, contact_last_name, contact_email, company_name, company_address, company_email, company_phone_number, company_url, url, contact_phone_number").
		Find(&results).
		Error; err != nil {
		return nil, errors.New(constant.ErrorGetPostsBySearchId)
	}

	return &results, nil

}

func (repository SearchRepository) GetParticipantsBySearchId(searchId string) (*[]contract.ParticipantDTOForSearchById, error) {

	var results []contract.ParticipantDTOForSearchById
	var posts []model.Post

	// Get Friends
	var friends []contract.ParticipantDTOForSearchById
	if err := repository.database.
		Model(&model.Friendship{}).
		Where("friendships.search_id = ?", searchId).
		Where("friendships.deleted_at IS NULL").
		Joins("JOIN users ON users.id = friendships.user_id").
		Select("users.id, users.first_name, users.last_name, users.email, users.avatar_url, friendships.type, friendships.id as friendship_id").
		Find(&friends).
		Joins("JOIN posts ON posts.user_id = users.id AND posts.search_id = ?", searchId).
		Select("posts.id, posts.user_id").
		Find(&posts).
		Error; err != nil {
		return nil, errors.New(constant.ErrorGetPostsBySearchId)
	}

	// Count total of post for each participant
	if len(friends) != 0 {
		for index, item := range friends {
			var count int64 = 0

			for _, post := range posts {
				if post.UserID == item.Id {
					count = count +1
				}
			}

			friends[index].NumberOfPosts = count
		}
	}

	// Get Followers
	var followers []contract.ParticipantDTOForSearchById
	if err := repository.database.
		Model(&model.Follower{}).
		Where("followers.search_id = ?", searchId).
		Where("followers.deleted_at IS NULL").
		Joins("JOIN users ON users.id = followers.user_id").
		Select("users.id, users.first_name, users.last_name, users.email, users.avatar_url, followers.id as follower_id").
		Find(&followers).
		Joins("JOIN posts ON posts.user_id = users.id AND posts.search_id = ?", searchId).
		Select("posts.id, posts.user_id").
		Find(&posts).
		Error; err != nil {
		return nil, errors.New(constant.ErrorGetPostsBySearchId)
	}

	// Count total of post for each participant
	if len(followers) != 0 {
		for index, item := range followers {
			var count int64 = 0

			for _, post := range posts {
				if post.UserID == item.Id {
					count = count +1
				}
			}

			followers[index].NumberOfPosts = count
		}
	}

	results = append(friends, followers...)

	return &results, nil

}

func (repository SearchRepository) AddSearch(SearchDTO contract.SearchDTO) (*contract.SearchDTO, error) {

	search := model.ToSearch(SearchDTO)

	// Add new search uuid
	id := uuid.New()

	search.ID = id.String()

	if err := repository.database.Create(&search).Error; err != nil {
		return nil, errors.New(constant.ErrorAddSearch)
	}

	searchDTO := model.ToSearchDTO(search)

	return &searchDTO, nil
}

func (repository SearchRepository) ModifySearch(SearchDTO contract.SearchDTO) (*contract.SearchDTO, error) {

	search := model.ToSearch(SearchDTO)

	if err := repository.database.Model(&search).Where("id = ?", search.ID).Where("deleted_at IS NULL").Updates(map[string]interface{}{"title": search.Title,
		"description": search.Description, "type": search.Type, "tags": search.Tags, "sector": search.Sector}).Error; err != nil {
		return nil, errors.New(constant.ErrorModifySearch)
	}

	searchDTO := model.ToSearchDTO(search)

	return &searchDTO, nil
}

func (repository SearchRepository) AddPost(postDto *contract.PostDTO) (*contract.AddPostResponseDTO, error) {

	post := model.ToPost(*postDto)

	// Add new post uuid
	id := uuid.New()
	post.ID = id.String()

	var postResponseDto contract.AddPostResponseDTO

	// Create post
	if err := repository.database.
		Create(&post).
		Joins("JOIN users ON users.id = ?", post.UserID).
		Select("users.id as user_id, users.first_name as user_first_name, users.last_name as user_last_name, users.email as user_email, posts.id as id, posts.title as title, posts.description as description, posts.type as type, posts.url as url, posts.search_id as search_id").
		Scan(&postResponseDto).
		Error; err != nil {
		return nil, errors.New(constant.ErrorAddPost)
	}

	return &postResponseDto, nil

}

func (repository SearchRepository) UpdatePost(postDto *contract.PostDTO) (*contract.UpdatePostResponseDTO, error) {

	post := model.ToPost(*postDto)

	var postResponseDto contract.UpdatePostResponseDTO

	// Update post
	if err := repository.database.
		Model(&post).
		Where("posts.id = ?", post.ID).
		Where("posts.deleted_at IS NULL").
		Updates(map[string]interface{}{"title": post.Title, "description": post.Description, "type": post.Type, "url": post.Url}).
		Scan(&postResponseDto).
		Error; err != nil {
		return nil, errors.New(constant.ErrorUpdatePost)
	}

	// Get user info
	if postResponseDto.Id != "" {
		err := repository.database.Model(&model.User{}).Where("id = ?", postResponseDto.UserID).Select("users.first_name as user_first_name, users.last_name as user_last_name, users.email as user_email").Scan(&postResponseDto).Error
		if err != nil {
			return nil, errors.New(constant.ErrorUpdatePost)
		}
	}

	return &postResponseDto, nil

}

func (repository SearchRepository) DeletePostById(postId string) (bool, error) {

	if err := repository.database.Model(&model.Post{}).Where("posts.id = ?", postId).Update("deleted_at", time.Now().UTC()).Error; err != nil {
		return false, errors.New(constant.ErrorDeletePostById)
	}

	return true, nil
}

func (repository SearchRepository) IsSearchOwner(userId string, searchId string) bool {

	if err := repository.database.Model(&model.Search{}).Where("id = ?", searchId).Where("user_id = ?", userId).First(&model.Search{}).Error; err != nil {
		return false
	}

	return true
}

func (repository SearchRepository) IsPostOwner(userId string, postId string) bool {

	if err := repository.database.Model(&model.Post{}).Where("id = ?", postId).Where("user_id = ?", userId).First(&model.Post{}).Error; err != nil {
		return false
	}

	return true
}

func (repository SearchRepository) IsPublic(searchId string) bool {

	var search model.Search

	// Get search by id
	if err := repository.database.Model(&model.Search{}).Where("id = ?", searchId).Where("type = ?", constant.SearchTypePublic).First(&search).Error; err != nil {
		return false
	}

	return true
}

func (repository SearchRepository) IsFriend(userId string, searchId string) bool {

	var friendship model.Friendship

	// Get search by id
	if err := repository.database.Model(&model.Friendship{}).Where("search_id = ?", searchId).Where("user_id = ? ", userId).First(&friendship).Error; err != nil {
		return false

	}

	return true
}

func (repository SearchRepository) IsSearchExist(searchId string) (*contract.SearchDTO, error) {

	var search model.Search

	// Get search by id
	if err := repository.database.Model(&model.Search{}).Where("id = ?", searchId).Where("deleted_at IS NULL ").First(&search).Error; err != nil {
		return nil, errors.New(constant.ErrorSearchNotFound)

	}

	searchDTO := model.ToSearchDTO(search)

	return &searchDTO, nil
}

func (repository SearchRepository) IsPostExist(postId string) (*contract.PostDTO, error) {

	var post model.Post

	// Get search by id
	if err := repository.database.Model(&model.Post{}).Where("id = ?", postId).Where("deleted_at IS NULL ").First(&post).Error; err != nil {
		return nil, errors.New(constant.ErrorPostNotFound)

	}

	postDto := model.ToPostDTO(post)

	return &postDto, nil
}

func (repository SearchRepository) IsFriendshipExist(searchId string, userId string) (*contract.FriendshipDTO, error) {
	var friendship model.Friendship

	if err := repository.database.
		Model(&friendship).
		Where("search_id = ?", searchId).
		Where("user_id = ?", userId).
		Where("deleted_at IS NULL").
		First(&friendship).
		Error; err != nil {

		return nil, errors.New(constant.ErrorFriendshipNotFound)
	}

	friendshipDto := model.ToFriendshipDTO(friendship)

	return &friendshipDto, nil
}

func (repository SearchRepository) IsFriendshipDeleted(searchId string, userId string) (*contract.FriendshipDTO, error) {
	var friendship model.Friendship

	if err := repository.database.
		Model(&friendship).
		Where("search_id = ?", searchId).
		Where("user_id = ?", userId).
		First(&friendship).
		Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, errors.New(constant.ErrorGetDeletedFriendship)
		}
	}

	friendshipDto := model.ToFriendshipDTO(friendship)

	return &friendshipDto, nil
}

func (repository SearchRepository) SaveFriendship(friendshipDto *contract.FriendshipDTO) (*contract.FriendshipDTO, error) {

	friendship := model.ToFriendship(*friendshipDto)

	// Create uuid if none
	if friendship.ID == "" {

		id := uuid.New()
		friendship.ID = id.String()

	}

	// Save friendship
	if err := repository.database.Model(&friendship).Save(&friendship).Error; err != nil {
		return nil, errors.New(constant.ErrorSaveFriendship)
	}

	friendshipDtoResponse := model.ToFriendshipDTO(friendship)

	return &friendshipDtoResponse, nil
}

func (repository SearchRepository) DeleteFriendship(friendshipId string) (bool, error) {

	// Delete friendship
	if err := repository.database.Model(&model.Friendship{}).Where("id = ?", friendshipId).Where("deleted_at IS NULL").Update("deleted_at", time.Now().UTC()).Error; err != nil {
		return false, errors.New(constant.ErrorDeleteFriendship)
	}

	return true, nil
}

func (repository SearchRepository) SaveFollower(followerDto contract.FollowerDTO) (*contract.FollowerDTO, error) {

	// Format model
	follower := model.ToFollower(followerDto)

	// Create follower id
	if followerDto.Id == "" {

		// Add new search uuid
		id := uuid.New()

		follower.ID = id.String()
	}

	// Post follower
	if err := repository.database.Save(&follower).Error; err != nil {
		return nil, errors.New(constant.ErrorSaveFollower)
	}

	followerDTO := model.ToFollowerDto(follower)

	return &followerDTO, nil
}

func (repository SearchRepository) IsFollowerExist(searchId string, userId string) (*contract.FollowerDTO, error) {

	var follower model.Follower

	err := repository.database.Model(model.Follower{}).Where("search_id = ?", searchId).Where("deleted_at IS NULL").Where("user_id = ?", userId).First(&follower).Error
	if err != nil {
		return nil, err
	}

	followerDto := model.ToFollowerDto(follower)

	return &followerDto, nil

}

func (repository SearchRepository) GetFollowerById(followerId string) (*contract.FollowerDTO, error) {

	var follower model.Follower

	if err := repository.database.Model(&follower).Where("id = ?", followerId).Where("deleted_at IS NULL").First(&follower).Error; err != nil {
		return nil, errors.New(constant.ErrorFollowerNotFound)
	}

	followerDto := model.ToFollowerDto(follower)

	return &followerDto, nil
}

func (repository SearchRepository) DeleteFollowerById(followerId string) (bool, error) {

	if err := repository.database.Model(&model.Follower{}).Where("id = ?", followerId).Update("deleted_at", time.Now().UTC()).Error; err != nil {
		return false, errors.New(constant.ErrorDeleteFollower)
	}

	return true, nil
}

func (repository SearchRepository) GetPublicSearches(userId string) (*[]contract.PublicSearchDTO, error) {

	var results []contract.PublicSearchDTO

	if err := repository.database.
		Model(model.Search{}).
		Where("searches.type = ?", constant.SearchTypePublic).
		Not("searches.user_id = ?", userId).
		Joins("JOIN users ON users.id = searches.user_id").
		Joins("LEFT JOIN friendships ON friendships.search_id = searches.id AND friendships.deleted_at IS NULL AND friendships.user_id = ?", userId).
		Joins("LEFT JOIN followers ON followers.search_id = searches.id AND followers.deleted_at IS NULL AND followers.user_id = ?", userId).
		Select("friendships.id as friendship_id, followers.id as follower_id, users.id as user_id, users.first_name, users.last_name, users.email, users.avatar_url, searches.title, searches.id as id, searches.description, searches.tags, searches.sector, searches.type").
		Find(&results).
		Error; err != nil {
		return nil, errors.New(constant.ErrorGetPublicSearches)
	}

	return &results, nil

}

func (repository SearchRepository) IsFollowerDeleted(searchId string, userId string) (*contract.FollowerDTO, error) {
	var follower model.Follower

	if err := repository.database.
		Model(&follower).
		Where("search_id = ?", searchId).
		Where("user_id = ?", userId).
		First(&follower).
		Error; err != nil {
		if err != gorm.ErrRecordNotFound {

			return nil, errors.New(constant.ErrorGetDeletedFollower)
		}
	}

	followerDto := model.ToFollowerDto(follower)

	return &followerDto, nil
}
