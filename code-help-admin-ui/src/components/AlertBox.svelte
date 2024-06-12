<script lang="ts">
  import { Icon } from "svelte-icons-pack";
  import type { MessageBoxType } from "./types";
  import { colors, messageBoxOutlineIcons } from "./utils";

  export let type: MessageBoxType = "info";
  export let message: string | undefined;
  export let wrapper: boolean = true;
</script>

<style>
  .alert-box {
    margin: auto;
    display: flex;
    flex-direction: column;
    box-sizing: border-box;
    text-align: center;
    justify-content: center;
    align-items: center;
    padding: 2rem 1rem;
    gap: 0.5rem;
    box-shadow: rgba(149, 157, 165, 0.2) 0px 8px 24px;

    width: 100%;
    max-width: 500px;
  }

  .info {
    border: 1px solid blue;
  }
  .success {
    border: 1px solid green;
  }
  .error {
    border: 1px solid red;
  }

  .wrapper {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100%;
    padding: 0.2rem;
  }
</style>

{#if wrapper}
  <div class="wrapper">
    <div class="alert-box {type}">
      <Icon src={messageBoxOutlineIcons[type]} color={colors[type]} size="48" />
      {message}
    </div>
  </div>
{:else}
  <div class="alert-box {type}">
    <Icon src={messageBoxOutlineIcons[type]} color={colors[type]} size="48" />
    {message}
  </div>
{/if}
