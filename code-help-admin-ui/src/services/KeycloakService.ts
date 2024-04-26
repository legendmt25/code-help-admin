import Keycloak from "keycloak-js";
import env from "../env";

export const KEYCLOAK_KEY = "keycloak";

export const initKeycloak = async () => {
  const keycloak = new Keycloak({
    url: env.KEYCLOAK_URL,
    realm: env.KEYCLOAK_REALM,
    clientId: env.KEYCLOAK_CLIENTID
  });

  await keycloak.init({ onLoad: "login-required" });

  return keycloak;
};

export const getInstance = (): Keycloak | undefined =>
  localStorage.getItem("keycloak") ? (JSON.parse(localStorage.getItem("keycloak")!) as Keycloak) : undefined;
