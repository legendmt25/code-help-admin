import { type RouteDefinition } from "svelte-spa-router";
import { wrap } from "svelte-spa-router/wrap";
import ProblemsEditCreate from "./lib/ProblemsEditCreate.svelte";
import ProblemsOverview from "./lib/ProblemsOverview.svelte";

export const Route = {
  problems_overview: "/problems",
  problems_create: "/problems/create",
  problems_edit: "/problems/edit/:id"
} as const;

export const routes: RouteDefinition = {
  [Route.problems_overview]: wrap({
    component: ProblemsOverview
  }),
  [Route.problems_edit]: wrap({
    component: ProblemsEditCreate
  }),
  [Route.problems_create]: wrap({
    component: ProblemsEditCreate
  })
};
