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
    width: 100%;
    max-width: 500px;
    border: 1px solid #b6b7b8;
  }

  dialog[open] {
    display: flex;
    flex-direction: column;
  }


  .header {
    padding: 1rem 1rem;
    box-sizing: border-box;
    border-bottom: 1px solid #b6b7b8;
    display: flex;
    justify-content: space-between;
    border-radius: 0.1rem;
  }

  .close-icon {
    margin-left: auto;
  }

  .footer {
    padding: 1rem 1rem;
    box-sizing: border-box;
    border-top: 1px solid #b6b7b8;
    border-radius: 0.1rem;
    justify-content: end;
    display: flex;
  }

  .modal-dialog {
    background-color: white;
    color: black;
    border-radius: 0.3rem;
  }

  .modal-dialog-content {
    flex-grow: 1;
    padding-block: 1rem;
  }
</style>

<Portal>
  <dialog class="modal-dialog" bind:this={dialog} on:close>
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
    <div class="modal-dialog-content">
      <slot />
    </div>
    <div class="footer">
      <slot name="footer" />
    </div>
  </dialog>
</Portal>
