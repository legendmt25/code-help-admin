import type { Middleware } from "../generated/admin-api";
import { getInstance } from "./KeycloakService";

export const authMiddleware: Middleware["pre"] = (requestContext) => {
  const keycloak = getInstance();

  if (keycloak?.token) {
    (requestContext.init.headers as any)["Authorization"] = "Bearer " + keycloak?.token;
  }
  return Promise.resolve();
};
