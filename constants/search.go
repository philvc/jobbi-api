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
	ErrorMissingAccess = "missing_access"
)


// Repository errors
const (
	ErrorGetMySearch = "repository_search_error_get_my_search"
	ErrorGetMySearchParticipants = "repository_search_error_get_my_search_participants"
	ErrorGetSharedSearches = "repository_search_error_get_shared_searches"
	ErrorGetFollowedSearches = "repository_search_error_get_followed_searches"
	ErrorAddSearch = "repository_search_error_add_search"
	ErrorModifySearch = "repository_search_error_modify_search"
	ErrorGetSearchById = "repository_search_error_get_search_by_id"
)