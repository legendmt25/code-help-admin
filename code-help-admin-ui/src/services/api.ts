import env from "../env";
import {
  CategoryApi,
  CommentApi,
  CommunityApi,
  Configuration,
  ContestApi,
  ForumCategoryApi,
  PostApi,
  ProblemApi
} from "../generated/admin-api";
import { authMiddleware } from "./middlewares";

export const baseConfiguration: Configuration = new Configuration({
  basePath: env.ADMIN_API_URL,
  middleware: [
    {
      pre: authMiddleware
    }
  ]
});

export const PROBLEM_API = new ProblemApi(baseConfiguration);
export const CATEGORY_API = new CategoryApi(baseConfiguration);
export const FORUM_CATEGORY_API = new ForumCategoryApi(baseConfiguration);
export const CONTEST_API = new ContestApi(baseConfiguration);
export const COMMUNITY_API = new CommunityApi(baseConfiguration);
export const POST_API = new PostApi(baseConfiguration);
export const COMMENT_API = new CommentApi(baseConfiguration);
