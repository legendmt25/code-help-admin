<script lang="ts">
  import { Icon } from "svelte-icons-pack";
  import { AiFillExclamationCircle } from "svelte-icons-pack/ai";
  import type { MessageBoxType } from "./types";

  export let type: MessageBoxType = "info";

  const colors: Record<MessageBoxType, string> = {
    error: "red",
    success: "green",
    info: "blue"
  };
</script>

<style>
  div {
    display: flex;
    gap: 1rem;
    background-color: #eeeef0;
    padding: 1rem;

    max-width: 500px;
    width: 100%;
  }
</style>

<div>
  <Icon src={AiFillExclamationCircle} size="24" color={colors[type]} />
  <slot />
</div>
