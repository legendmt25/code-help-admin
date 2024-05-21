import { type CommunityRequest } from "../../generated/admin-api";
import { COMMUNITY_API, FORUM_CATEGORY_API } from "../api";

export const getAllCommunities = () => COMMUNITY_API.getAllCommunities();
export const getCommunityByName = (communityName: string) => COMMUNITY_API.getCommunityByUid({ name: communityName });
export const createCommunity = (body: CommunityRequest) => COMMUNITY_API.createCommunity({ communityRequest: body });
export const updateCommunity = (communityName: string, body: CommunityRequest) =>
  COMMUNITY_API.updateCommunity({ communityRequest: body, name: communityName });
export const deleteCommunity = (communityName: string) => COMMUNITY_API.deleteCommunity({ name: communityName });

export const getAllCategories = () => FORUM_CATEGORY_API.getAllForumCategories();
export const createCategory = (name: string) =>
  FORUM_CATEGORY_API.createForumCategory({
    categoryCreate: {
      name
    }
  });

export const deleteCategory = (uid: string) => new Promise((resolve) => resolve(undefined));
export const updateCategory = (uid: string, newCategoryName: string): Promise<void> =>
  new Promise((resolve) => {
    resolve(undefined);
  });
