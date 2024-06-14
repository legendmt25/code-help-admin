<script lang="ts">
  import { onMount } from "svelte";
  import { initKeycloak, type KeycloakType } from "../services/KeycloakService";

  let keycloak: KeycloakType;

  onMount(() => {
    initKeycloak().then((initializedKeycloak) => (keycloak = window.keycloak = initializedKeycloak));
  });
</script>

{#if keycloak}
  <slot {keycloak} />
{/if}
