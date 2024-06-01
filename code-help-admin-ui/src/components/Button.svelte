<script lang="ts">
  import type { ButtonType } from "./types";

  export let fullWidth: boolean | undefined = undefined;
  export let active: boolean | undefined = undefined;
  export let type: ButtonType = "primary";


  export let style = "";

  export let className: string = "";
  export { className as class };

  export let toggleButton = false;
  export let toggled = false;
</script>

<style>
  .primary {
    border: 1px solid transparent;
    padding: 0.6em 1.2em;
    font-size: 1em;
    border-radius: 0.3rem;
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

  .btn-full-width {
    max-width: 100%;
  }

  .active.primary {
    background-color: #314257;
  }

  button.toggle {
    outline: none;
  }

  button.toggle:not(.toggled) {
    background-color: #eaeaea;
    color: black;
  }
</style>

<button {...$$restProps} on:click class:toggle={toggleButton} class:toggled class={type + ' ' + className} class:btn-full-width={fullWidth} {style}  class:active>
  <slot />
</button>
