package constant

// Constants
const (
	SearchTypePrivate = "private"
	SearchTypePublic = "public"
)

const (
	RequesterRoleOwner = "owner"
	RequesterRoleFriend = "friend"
	RequesterRoleFollower = "follower"
	RequesterRoleVisitor = "visitor"
)

// Errors

// Post Search usecase errors
const (
	ErrorAlreadyExistingSearch = "already_existing search"
	ErrorMissingDescription = "missing_description"
	ErrorMissingTitle = "missing_title"
	ErrorMissingType = "missing_type"
	ErrorMissingUserId = "missing_user_id"
	ErrorMissingSearchId = "missing_search_id"
	ErrorMissingPostId = "missing_post_id"
	ErrorMissingAccess = "missing_access"
	ErrorWrongBody = "fail_with_wrong_body"
	ErrorWrongParams = "fail_with_wrong_params"
	ErrorWrongParamsUsecase = "fail_with_wrong_params_usecase"
	ErrorGetRequesterRole = "get_requester_role_fail"
)

// Usecase errors
const (
	ErrorGetSearchForInvitationParams = "usecase_get_search_for_invitation_missing_params"
	ErrorUpsertFriendshipParams = "usecase_upsert_friendship_params"
	ErrorFriendshipAlreadyExist = "usecase_friendship_already_exist"
	ErrorFollowerAlreadyExist = "usecase_follower_already_exist"
	ErrorFriendshipFollowerAlreadyExist = "usecase_friendship_follower_already_exist"
	ErrorIsSearchOwner = "usecase_friendship_user_is_search_owner"
	ErrorFollowerNotAllowedOwner = "usecase_follower_not_allowed_owner"
	ErrorFollowerNotAllowedFriendship = "usecase_friendship_post_not_allowed_friendship"
)

// Repository errors
const (
	ErrorGetMySearch = "repository_search_error_get_my_search"
	ErrorSearchNotFound = "repository_search_error_exist"
	ErrorPostNotFound = "repository_search_error_post_not_found"
	ErrorGetMySearchParticipants = "repository_search_error_get_my_search_participants"
	ErrorGetSharedSearches = "repository_search_error_get_shared_searches"
	ErrorGetFollowedSearches = "repository_search_error_get_followed_searches"
	ErrorAddSearch = "repository_search_error_add_search"
	ErrorAddPost = "repository_search_error_add_post"
	ErrorUpdatePost = "repository_search_error_update_post"
	ErrorAddPostUserDetails = "repository_search_error_add_post_join_user_details"
	ErrorModifySearch = "repository_search_error_modify_search"
	ErrorGetSearchById = "repository_search_error_get_search_by_id"
	ErrorGetPostsBySearchId = "repository_search_error_get_posts_by_search_id"
	ErrorDeletePostById = "repository_search_error_delete_post_by_id"
	ErrorFriendshipNotFound = "repository_search_friendship_not_found"
	ErrorFriendshipDeletedNotFound = "repository_search_friendship_deleted_not_found"
	ErrorSaveFriendship = "repository_search_friendship_save_fail"
	ErrorDeleteFriendship = "repository_search_error_delete_friendship_by_id"
	ErrorSaveFollower = "save_follower_fail"
	ErrorFollowerNotFound = "repository_search_follower_by_id_not_found"
	ErrorDeleteFollower = "repository_search_error_delete_follower_by_id"
	ErrorGetFriendsBySearchId = "get_friends_by_search_id_failed"
	ErrorGetPublicSearches = "get_public_searches_failed"
	ErrorGetDeletedFriendship= "get_deleted_friendship_fail"
	ErrorGetDeletedFollower= "get_deleted_follower_fail"
	ErrorCreateCommentForPost = "create_comment_for_post_fail"
	ErrorUpdateCommentForPost = "update_comment_for_post_fail"
	ErrorGetAllCommentsForPost = "get_all_comment_for_post_fail"
)