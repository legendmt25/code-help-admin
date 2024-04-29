<script lang="ts">
  import { link } from "svelte-spa-router";

  export let fullwidth: boolean | undefined = undefined;
  export let active: boolean | undefined = undefined;

  export let clazz: string = "";
  export { clazz as class };
</script>

<style>
  a {
    border: 1px solid transparent;
    padding: 0.6em 1.2em;
    font-size: 1em;
    width: 100%;
    font-weight: 500;
    font-family: inherit;
    color: white;
    background-color: #18283b;
    cursor: pointer;
    transition: border-color 0.25s;
    width: 100%;
    max-width: 200px;
    transition: all 200ms;
  }

  a:hover {
    background-color: #314257;
  }

  a:active {
    background-color: #354457;
  }

  a:focus,
  a:focus-visible {
    outline: 4px auto -webkit-focus-ring-color;
  }

  .btn-fullwidth {
    max-width: 100%;
  }

  .active {
    background-color: #314257;
  }
</style>

<a {...$$restProps} use:link on:click class:btn-fullwidth={fullwidth} class={clazz} class:active>
  <slot />
</a>
