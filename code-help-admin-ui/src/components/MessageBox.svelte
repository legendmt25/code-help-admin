<script lang="ts">
  import { Icon } from "svelte-icons-pack";
  import type { MessageBoxType } from "./types";
  import { colors, messageBoxFillIcons } from "./utils";

  export let style: string | undefined = undefined;
  export let type: MessageBoxType = "info";
  export let fullwidth: boolean | undefined = undefined;
</script>

<style>
  div {
    display: flex;
    gap: 1rem;
    background-color: #eeeef0;
    padding: 1rem;

    max-width: 500px;
    width: 100%;
    position: absolute;
    right: 0;
    bottom: 0;
  }

  .fullwidth {
    max-width: 100%;
  }
</style>

<div {style} class:fullwidth>
  <Icon src={messageBoxFillIcons[type]} size="24" color={colors[type]} />
  <slot />
</div>
