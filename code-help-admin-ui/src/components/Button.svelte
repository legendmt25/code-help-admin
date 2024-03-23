<script lang="ts">
  export let onClick: (() => void) | undefined = undefined;
  export let fullwidth: boolean | undefined = undefined;

  export let clazz: string = "";
  export { clazz as class };
</script>

<style>
  button {
    border: 1px solid transparent;
    padding: 0.6em 1.2em;
    font-size: 1em;
    width: 100%;
    font-weight: 500;
    font-family: inherit;
    color: white;
    background-color: #1a1a1a;
    cursor: pointer;
    transition: border-color 0.25s;
    width: 100%;
    max-width: 200px;
  }

  button:hover {
    border-color: #646cff;
  }

  button:focus,
  button:focus-visible {
    outline: 4px auto -webkit-focus-ring-color;
  }

  .btn-fullwidth {
    max-width: 100%;
  }
</style>

<button on:click={onClick} class:btn-fullwidth={fullwidth} class={clazz}>
  <slot />
</button>
