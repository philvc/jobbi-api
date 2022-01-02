package constant

// Constants
const (
	SearchTypePrivate = "private"
	SearchTypePublic = "public"

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
)