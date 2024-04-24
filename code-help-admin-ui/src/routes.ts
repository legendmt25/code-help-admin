import { type RouteDefinition } from "svelte-spa-router";
import { wrap } from "svelte-spa-router/wrap";
import ProblemsEditCreate from "./lib/ProblemsEditCreate.svelte";
import ProblemsOverview from "./lib/ProblemsOverview.svelte";
import ContestsOverview from "./lib/ContestsOverview.svelte";
import ContestEditCreate from "./lib/ContestEditCreate.svelte";
import CommunitiesOverview from "./lib/CommunitiesOverview.svelte";
import Home from "./lib/Home.svelte";

export const Route = {
  index: "/",

  problems_overview: "/problems",
  problems_create: "/problems/create",
  problems_edit: "/problems/edit/:id",

  contests_overview: "/contests",
  contests_create: "/contests/create",
  contests_edit: "/contests/edit/:id",
  communities_overview: "/communities"
} as const;

export const routes: RouteDefinition = {
  [Route.index]: wrap({
    component: Home
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
  })
};
