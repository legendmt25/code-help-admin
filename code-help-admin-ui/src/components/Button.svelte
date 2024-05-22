<script lang="ts">
  import type { ButtonType } from "./types";

  export let fullwidth: boolean | undefined = undefined;
  export let active: boolean | undefined = undefined;
  export let type: ButtonType = "primary";

  export let clazz: string = "";
  export { clazz as class };
</script>

<style>
  .primary {
    border: 1px solid transparent;
    padding: 0.6em 1.2em;
    font-size: 1em;
    width: 100%;
    font-weight: 500;
    font-family: inherit;
    color: white;
    background-color: #18283b;
  }

  .primary:hover {
    background-color: #314257;
  }

  button {
    cursor: pointer;
    transition: border-color 0.25s;
    width: 100%;
    max-width: 200px;
    transition: all 200ms;
  }

  .primary:active {
    background-color: #354457;
  }

  .primary:focus,
  .primary:focus-visible {
    outline: 4px auto -webkit-focus-ring-color;
  }

  .btn-fullwidth {
    max-width: 100%;
  }

  .active.primary {
    background-color: #314257;
  }
</style>

<button {...$$restProps} on:click class:btn-fullwidth={fullwidth} class={type + ' ' + clazz} class:active>
  <slot />
</button>
