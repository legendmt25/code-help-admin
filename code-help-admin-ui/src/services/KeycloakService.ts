import Keycloak, { type KeycloakTokenParsed } from "keycloak-js";
import env from "../env";

export const ID_TOKEN_KEY = "id_token";

export type KeycloakType = Keycloak;

const onAuthSuccess = (keycloak: KeycloakType) => {
  localStorage.setItem(ID_TOKEN_KEY, keycloak.idToken!);
};

const onAuthLogout = (keycloak: KeycloakType) => {
  localStorage.removeItem(ID_TOKEN_KEY);
};

export const initKeycloak = async (): Promise<KeycloakType> => {
  const keycloak = new Keycloak({
    url: env.KEYCLOAK_URL,
    realm: env.KEYCLOAK_REALM,
    clientId: env.KEYCLOAK_CLIENTID
  });

  keycloak.onTokenExpired = () => {
    console.log("token expired");
    keycloak
      .updateToken(60)
      .then((refreshed) => {
        if (refreshed) {
          console.info("Token was successfully refreshed");
          onAuthSuccess(keycloak);
        } else {
          console.info("Token is still valid");
        }
      })
      .catch((err) => {
        const errorMessage = "Failed to refresh the token, or the session has expired";
        console.error(errorMessage, err);
        keycloak.clearToken();
        keycloak.logout();
      });
  };

  try {
    await keycloak.init({ onLoad: "login-required", checkLoginIframe: false });
    onAuthSuccess(keycloak);
  } catch {}

  return keycloak;
};

export const getIdToken = () => localStorage.getItem(ID_TOKEN_KEY);

export const isAuthenticated = (keycloak: KeycloakType) => true;

export const getPrefferedUsername = (keycloak: KeycloakType): string | undefined => {
  const idTokenParsed: KeycloakTokenParsed = keycloak.idTokenParsed ?? {};
  const usernameKey = "preferred_username";

  if (usernameKey in idTokenParsed) {
    return idTokenParsed[usernameKey];
  }

  return undefined;
};

export const logout = (keycloak: KeycloakType) => {
  keycloak?.logout().then(() => onAuthLogout(keycloak));
};

export const login = (keycloak: KeycloakType) => {
  keycloak.login().then(() => onAuthLogout(keycloak));
};
