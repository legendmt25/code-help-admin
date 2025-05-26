import env from "../env";
import {
  CategoryApi,
  CodeRunnerApi,
  Configuration as AdminCoreApiConfugiration,
  ContestApi,
  ProblemApi
} from "../generated/admin-core-api";
import {
  CommentApi,
  CommunityApi,
  CategoryApi as ForumCategoryApi,
  Configuration as AdminForumApiConfugiration,
  PostApi,
} from "../generated/admin-forum-api";
import { authMiddleware } from "./middlewares";

const createConfiguration = (basePath: string) => ({
  basePath: basePath,

  headers: {
    // "Access-Control-Allow-Origin": "http://localhost:30000",
    // "Access-Control-Allow-Methods": "POST,PUT,PATCH,OPTIONS,GET,DELETE",
  },
  middleware: [
    {
      pre: authMiddleware
    }
  ]
});

export const adminCoreApiConfiguration: AdminCoreApiConfugiration = new AdminCoreApiConfugiration(createConfiguration(env.ADMIN_API_URL));
export const adminForumApiConfiguration: AdminForumApiConfugiration = new AdminForumApiConfugiration(createConfiguration(env.ADMIN_API_URL + "/forum"));

export const PROBLEM_API = new ProblemApi(adminCoreApiConfiguration);
export const CATEGORY_API = new CategoryApi(adminCoreApiConfiguration);
export const CONTEST_API = new ContestApi(adminCoreApiConfiguration);
export const CODE_RUNNER_API = new CodeRunnerApi(adminCoreApiConfiguration);

export const FORUM_CATEGORY_API = new ForumCategoryApi(adminForumApiConfiguration);
export const COMMUNITY_API = new CommunityApi(adminForumApiConfiguration);
export const POST_API = new PostApi(adminForumApiConfiguration);
export const COMMENT_API = new CommentApi(adminForumApiConfiguration);
