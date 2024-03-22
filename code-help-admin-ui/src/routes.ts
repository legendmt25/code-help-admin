import { type RouteDefinition } from "svelte-spa-router";
import { wrap } from "svelte-spa-router/wrap";
import ProblemsEdit from "./lib/ProblemsEdit.svelte";
import ProblemsOverview from "./lib/ProblemsOverview.svelte";

export const Route = {
  problems_overview: "/",
  problems_edit: "/problems/edit/:id"
} as const;

export const routes: RouteDefinition = {
  [Route.problems_overview]: wrap({
    component: ProblemsOverview
  }),
  [Route.problems_edit]: wrap({
    component: ProblemsEdit
  })
};
