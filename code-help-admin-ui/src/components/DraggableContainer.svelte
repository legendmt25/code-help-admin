<script lang="ts">
  import { onDestroy } from "svelte";
  import type { MouseEventHandler } from "svelte/elements";

  let containerRef: HTMLDivElement | undefined = undefined;
  let mouseDown: boolean = false;
  let focused: boolean = false;
  export let width: string | undefined = undefined;
  export let height: string | undefined = undefined;

  const handleMouseMove: MouseEventHandler<HTMLDivElement> = (event) => {
    if (!mouseDown || !containerRef) {
      return;
    }

    containerRef.style.left = (containerRef.offsetLeft + event.movementX).toString() + "px";
    containerRef.style.top = (containerRef.offsetTop + event.movementY).toString() + "px";
  };

  const handler: MouseEventHandler<any> = (event) => {
    if (containerRef !== event.target && !containerRef?.contains(event.target as Node)) {
      focused = false;
    }
  };

  $: focused,
    (() => {
      if (focused) {
        window.addEventListener("mousedown", handler);
      }
    })();

  onDestroy(() => {
    window.removeEventListener("mousedown", handler);
  });
</script>

<style>
  div {
    position: absolute;
    resize: both;
    overflow: auto;
    background-color: #e9e9ede3;
    border: 1px solid #8392a5;
    border-radius: 0.625rem;
    padding: 1rem;
    box-shadow: rgba(0, 0, 0, 0.05) 0 0 2px 0px;
  }

  .focused {
    z-index: 9999;
  }
</style>

<!-- svelte-ignore a11y-no-static-element-interactions -->
<div
  style:width
  style:height
  on:mousedown={() => {
    focused = true;
    mouseDown = true;
  }}
  on:mouseup={() => (mouseDown = false)}
  on:mousemove={handleMouseMove}
  class:focused
  bind:this={containerRef}>
  <slot />
</div>
