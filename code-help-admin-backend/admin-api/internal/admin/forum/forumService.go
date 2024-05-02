package forum

import (
	"admin-api/internal/admin/forum/category"
	"admin-api/internal/admin/forum/comment"
	"admin-api/internal/admin/forum/community"
	"admin-api/internal/admin/forum/post"
	codeHelpForumGen "api-spec/generated/code-help-forum"
)

type Service interface {
	Community() community.ForumCommunityService
	Post() post.ForumPostService
	Comment() comment.ForumCommentService
	Category() category.ForumCategoryService
}

type forumServiceImpl struct {
	client           codeHelpForumGen.ClientInterface
	communityService community.ForumCommunityService
	postService      post.ForumPostService
	categoryService  category.ForumCategoryService
	commentService   comment.ForumCommentService
}

func NewForumService(client codeHelpForumGen.ClientInterface) Service {
	return &forumServiceImpl{
		client:           client,
		communityService: community.NewCommunityServiceImpl(client),
		postService:      post.NewPostServiceImpl(client),
		categoryService:  category.NewCategoryServiceImpl(client),
		commentService:   comment.NewCommentServiceImpl(client),
	}
}

func (it *forumServiceImpl) Community() community.ForumCommunityService {
	return it.communityService
}

func (it *forumServiceImpl) Post() post.ForumPostService {
	return it.postService
}

func (it *forumServiceImpl) Comment() comment.ForumCommentService {
	return it.commentService
}

func (it *forumServiceImpl) Category() category.ForumCategoryService {
	return it.categoryService
}
