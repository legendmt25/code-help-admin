import type { Middleware } from "../generated/admin-core-api";
import { getIdToken } from "./KeycloakService";

export const authMiddleware: Middleware["pre"] = (requestContext) => {
  const idToken = getIdToken();

  if (idToken) {
    (requestContext.init.headers as any)["Authorization"] = "Bearer " + idToken;
  }
  return Promise.resolve();
};
