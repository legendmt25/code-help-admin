import { type CommunityRequest, type PostRequest } from "../../generated/admin-forum-api";
import { COMMENT_API, COMMUNITY_API, FORUM_CATEGORY_API, POST_API } from "../api";

export const getPosts = (communityName: string) => POST_API.getPosts({ community: communityName });
export const getPost = (uid: string) => POST_API.getPost({ uid });
export const createPost = (communityName: string, body: PostRequest) =>
  POST_API.createCommunityPost({ community: communityName, postRequest: body });
export const updatePost = (uid: string, body: PostRequest) => POST_API.updatePost({ uid, postRequest: body });
export const deletePost = (uid: string) => POST_API.deletePost({ uid });

export const getComments = (postUid: string) => COMMENT_API.getCommentsForPost({ post: postUid });
export const getCommentReplies = (commentUid: string) => COMMENT_API.getCommentReplies({ uid: commentUid });
export const deleteComment = (uid: string) => COMMENT_API.deleteComment({ uid });
export const createComment = (postUid: string, content: string) =>
  COMMENT_API.commentOnPost({ post: postUid, commentRequest: { content } });
export const createReply = (uid: string, content: string) =>
  COMMENT_API.replyToComment({ uid, commentRequest: { content } });
export const deleteReply = deleteComment;

export const getAllCommunities = () => COMMUNITY_API.getAllCommunities();
export const getCommunityByName = (communityName: string) => COMMUNITY_API.getCommunityByUid({ name: communityName });
export const createCommunity = (body: CommunityRequest) => COMMUNITY_API.createCommunity({ communityRequest: body });
export const updateCommunity = (communityName: string, body: CommunityRequest) =>
  COMMUNITY_API.updateCommunity({ communityRequest: body, name: communityName });
export const deleteCommunity = (communityName: string) => COMMUNITY_API.deleteCommunity({ name: communityName });

export const getCommunityModerators = (communityName: string) =>
  COMMUNITY_API.getCommunityModerators({ name: communityName });
export const addCommunityModerator = (communityName: string, username: string) =>
  COMMUNITY_API.addModerator({ name: communityName, moderatorRequest: { username } });
export const removeCommunityModerator = (communityName: string, username: string) =>
  COMMUNITY_API.removeModerator({ name: communityName, username });

export const getAllCategories = () => FORUM_CATEGORY_API.getAllCategories();
export const createCategory = (name: string) => FORUM_CATEGORY_API.createCategory({ categoryCreate: { name } });
export const deleteCategory = (uid: string) => FORUM_CATEGORY_API.deleteCategory({ uid });
export const updateCategory = (uid: string, newCategoryName: string) =>
  FORUM_CATEGORY_API.updateCategory({ categoryCreate: { name: newCategoryName }, uid });
