import type { Category, Contest, ForumCategory, Problem, ShortCommunity, ShortPost, User } from "./generated/admin-api";

type FilterFunction<T> = (search: string | undefined, items: T[]) => T[];

export const filterCategories: FilterFunction<Category> = (search, categories) => {
  const searchLowerCase = search?.toLowerCase();

  return categories.filter(
    (category) =>
      !searchLowerCase ||
      category.id.toString().toLowerCase().includes(searchLowerCase) ||
      category.name.toLowerCase().includes(searchLowerCase)
  );
};

export const filterContests: FilterFunction<Contest> = (search, contests) => {
  const searchLowerCase = search?.toLowerCase();

  return contests.filter(
    (contest) =>
      !searchLowerCase ||
      contest.name.toLowerCase().includes(searchLowerCase) ||
      contest.duration.toLowerCase().includes(searchLowerCase) ||
      contest.id.toString().toLowerCase().includes(searchLowerCase)
  );
};

export const filterProblems: FilterFunction<Problem> = (search, problems) => {
  const searchLowerCase = search?.toLowerCase();

  return problems.filter(
    (problem) =>
      !searchLowerCase ||
      problem.title.toLowerCase().includes(searchLowerCase) ||
      problem.difficulty.toLowerCase().includes(searchLowerCase) ||
      problem.id.toString().toLowerCase().includes(searchLowerCase) ||
      problem.category?.name.toLowerCase().includes(searchLowerCase)
  );
};

export const filterCommunities: FilterFunction<ShortCommunity> = (search, communities) => {
  const searchLowerCase = search?.toLowerCase();

  return communities.filter(
    (community) =>
      !searchLowerCase ||
      community.name.toLowerCase().includes(searchLowerCase) ||
      community.description.toLowerCase().includes(searchLowerCase)
  );
};

export const filterForumCategories: FilterFunction<ForumCategory> = (search, categories) => {
  const searchLowerCase = search?.toLowerCase();

  return categories.filter(
    (category) =>
      !searchLowerCase ||
      category.name.toLowerCase().includes(searchLowerCase) ||
      category.uid.toLowerCase().includes(searchLowerCase)
  );
};

export const filterPosts: FilterFunction<ShortPost> = (search, posts) => {
  const searchLowerCase = search?.toLowerCase();

  return posts.filter(
    (post) =>
      !searchLowerCase ||
      post.title.toLowerCase().includes(searchLowerCase) ||
      post.uid.toLowerCase().includes(searchLowerCase) ||
      post.user.username.toLowerCase().includes(searchLowerCase)
  );
};

export const filterCommunityModerators: FilterFunction<User> = (search, communityModerators) => {
  const searchLowerCase = search?.toLowerCase();

  return communityModerators.filter(
    (communityModerator) => !searchLowerCase || communityModerator.username.toLowerCase().includes(searchLowerCase)
  );
};

export const formatDate = (date: Date) => `${date.toLocaleDateString()} ${date.toLocaleTimeString()}`;
