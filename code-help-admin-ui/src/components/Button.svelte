<script lang="ts">
  import { link } from "svelte-spa-router";
  import type { ButtonType } from "./types";

  export let fullWidth: boolean | undefined = undefined;
  export let active: boolean | undefined = undefined;
  export let maxContent: boolean | undefined = false;
  export let type: ButtonType = "primary";

  export let style = "";

  export let className: string = "";
  export { className as class };

  export let toggleButton = false;
  export let toggled = false;
  export let href: string | undefined = undefined;
</script>

<style>
  .btn-typography {
    font-size: 1em;
    font-weight: 500;
    font-family: inherit;
  }

  .primary {
    border: 1px solid transparent;
    background-color: #18283b;
    color: white;
  }

  .primary-outline {
    background-color: inherit;
    border: 1px solid #18283b;
    color: #18283b;
  }

  .primary-outline,
  .primary {
    padding: 0.6em 1.2em;
    border-radius: 0.3rem;
    width: 100%;
  }

  .primary:hover {
    background-color: #314257;
  }

  .primary-outline:hover {
    background-color: #314257;
    color: white;
  }

  .active.primary-outline {
    background-color: #18283b;
    color: white;
  }

  .btn {
    text-decoration: none;
    cursor: pointer;
    transition: border-color 0.25s;
    width: 100%;
    max-width: 200px;
    transition: all 200ms;
    display: inline-flex;
    align-items: center;
    gap: 0.3rem;
  }

  .primary:active {
    background-color: #354457;
  }

  .primary:focus,
  .primary:focus-visible {
    outline: 4px auto -webkit-focus-ring-color;
  }

  .btn-full-width {
    max-width: 100%;
  }

  .btn-max-content {
    max-width: max-content;
  }

  .active.primary {
    background-color: #314257;
  }

  .btn.toggle {
    outline: none;
  }

  .btn.toggle:not(.toggled) {
    background-color: #eaeaea;
    color: black;
  }

  .link {
    max-width: max-content;
    background-color: white;
    border: none;
    color: #18283b;
  }

  .link:hover {
    color: #314257;
  }

  .link:active {
    color: #354457;
  }
</style>

{#if href && href.trim().length > 0}
  <a
    use:link
    {...$$restProps}
    {href}
    on:click
    class:toggle={toggleButton}
    class:toggled
    class={type + " btn btn-typography " + className}
    class:btn-full-width={fullWidth}
    {style}
    class:active
    class:btn-max-content={maxContent}>
    <slot />
  </a>
{:else}
  <button
    {...$$restProps}
    on:click
    class:toggle={toggleButton}
    class:toggled
    class={type + " btn btn-typography " + className}
    class:btn-full-width={fullWidth}
    {style}
    class:active
    class:btn-max-content={maxContent}>
    <slot />
  </button>
{/if}
