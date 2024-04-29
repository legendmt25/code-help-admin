<script lang="ts">
  export let renderElement: HTMLElement = document.body;

  let ref: HTMLElement | undefined = undefined;
  $: ref && renderElement.appendChild(ref);
</script>

<style>
</style>

<div bind:this={ref} class="portal">
  <slot></slot>
</div>
