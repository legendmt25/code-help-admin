import { type KeycloakType } from "./services/KeycloakService";

declare global {
  interface Window {
    keycloak?: KeycloakType;
  }
}
