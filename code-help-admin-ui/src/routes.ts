import { type RouteDefinition } from "svelte-spa-router";
import { wrap } from "svelte-spa-router/wrap";
import Home from "./lib/Home.svelte";
import CategoriesOverview from "./lib/core/CategoriesOverview.svelte";
import ContestEditCreate from "./lib/core/ContestEditCreate.svelte";
import ContestsOverview from "./lib/core/ContestsOverview.svelte";
import ProblemsEditCreate from "./lib/core/ProblemsEditCreate.svelte";
import ProblemsOverview from "./lib/core/ProblemsOverview.svelte";
import CommunitiesOverview from "./lib/forum/CommunitiesOverview.svelte";
import CommunityEditCreate from "./lib/forum/CommunityEditCreate.svelte";
import { default as ForumCategoriesOverview } from "./lib/forum/CategoriesOverview.svelte";
import PostDetail from "./lib/forum/PostDetail.svelte";

export const Route = {
  index: "/",

  categories_overview: "/categories",

  problems_overview: "/problems",
  problems_create: "/problems/create",
  problems_edit: "/problems/edit/:id",
  contest_problems_create: "/problems/create?contestId=:contestId",

  contests_overview: "/contests",
  contests_create: "/contests/create",
  contests_edit: "/contests/edit/:id",

  communities_overview: "/communities",
  communities_create: "/communities/create",
  communities_edit: "/communities/edit/:name",
  post_edit: "/posts/edit/:uid",

  forum_categories_overview: "/forum-categories"
} as const;

export const routes: RouteDefinition = {
  [Route.index]: wrap({
    component: Home
  }),
  [Route.categories_overview]: wrap({
    component: CategoriesOverview
  }),
  [Route.problems_overview]: wrap({
    component: ProblemsOverview
  }),
  [Route.problems_edit]: wrap({
    component: ProblemsEditCreate
  }),
  [Route.problems_create]: wrap({
    component: ProblemsEditCreate
  }),
  [Route.contests_overview]: wrap({
    component: ContestsOverview
  }),
  [Route.contests_edit]: wrap({
    component: ContestEditCreate
  }),
  [Route.contests_create]: wrap({
    component: ContestEditCreate
  }),
  [Route.communities_overview]: wrap({
    component: CommunitiesOverview
  }),
  [Route.communities_create]: wrap({
    component: CommunityEditCreate
  }),
  [Route.communities_edit]: wrap({
    component: CommunityEditCreate
  }),
  [Route.forum_categories_overview]: wrap({
    component: ForumCategoriesOverview
  }),
  [Route.post_edit]: wrap({
    component: PostDetail
  })
};
