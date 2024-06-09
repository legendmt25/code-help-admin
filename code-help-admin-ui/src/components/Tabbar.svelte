<script lang="ts">
  import { createEventDispatcher } from "svelte";
  import Button from "./Button.svelte";
  import type { TabOption } from "./types";

  const dispatch = createEventDispatcher();

  export let options: TabOption[];
  export let defaultActive: string | undefined = undefined;
  let activeTab: string = defaultActive ?? (options.length > 0 ? options[0].key : "");
</script>

<style>
  div {
    display: flex;
    gap: 1rem;
  }
</style>

<div>
  {#each options as option}
    <Button
      active={activeTab === option.key}
      on:click={() => {
        activeTab = option.key;
        dispatch("change", option.key);
      }}>{option.label}</Button>
  {/each}
</div>
