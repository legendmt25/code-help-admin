<script lang="ts">
  import { Icon } from "svelte-icons-pack";
  import { AiOutlineClose } from "svelte-icons-pack/ai";
  import Portal from "./Portal.svelte";

  export let dialog: HTMLDialogElement | undefined = undefined;
  export let title: string | undefined = undefined;
</script>

<style>
  dialog {
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    border: none;
    width: 100%;
    max-width: 500px;
  }

  .header {
    padding: 1rem 0;
    margin: 0 0.5rem;
    border-bottom: 1px solid black;
    display: flex;
    justify-content: space-between;
  }

  .close-icon {
    margin-left: auto;
  }

  .footer {
    padding: 1rem 0;
    margin: 0 0.5rem;
    border-top: 1px solid black;
  }
</style>

<Portal>
  <dialog bind:this={dialog} on:close>
    <!-- svelte-ignore a11y-click-events-have-key-events -->
    <!-- svelte-ignore a11y-no-static-element-interactions -->
    <div class="header p-1">
      {#if title}
        <span>{title}</span>
      {/if}
      <div class="close-icon" on:click={() => dialog?.close()} on:keydown={undefined}>
        <Icon src={AiOutlineClose} size="20px" />
      </div>
    </div>
    <slot />
    <div class="footer">
      <slot name="footer" />
    </div>
  </dialog>
</Portal>
