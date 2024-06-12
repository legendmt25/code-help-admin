<script lang="ts">
  import { Icon } from "svelte-icons-pack";
  import { SlArrowUp } from "svelte-icons-pack/sl";
  import { SlArrowDown } from "svelte-icons-pack/sl";

  export let opened: boolean = false;
</script>

<style>
  .accordion > div:first-of-type {
    display: flex;
    gap: 0.5rem;
    justify-content: space-between;
  }

  .accordion > div:last-of-type {
    max-height: 0;
    overflow: hidden;

    transition: all 200ms;
  }

  .opened {
    max-height: 300px !important;
    margin-left: 0.25rem;
  }

  .title-container {
    display: flex;
    align-items: center;
    justify-content: space-between;

    cursor: pointer;

    padding-right: 1rem;

    padding-bottom: 1rem;
    padding-top: 1rem;
  }
</style>

<div class="accordion">
  <!-- svelte-ignore a11y-no-static-element-interactions -->
  <div class="title-container" on:click={() => (opened = !opened)} on:keydown={undefined}>
    <slot name="title"></slot>
    {#if opened}
      <Icon src={SlArrowUp} />
    {:else}
      <Icon src={SlArrowDown} />
    {/if}
  </div>
  <hr />
  <div class:opened>
    <slot></slot>
  </div>
</div>
