<script lang="ts">
  export let wrapper: boolean = true;
</script>

<style>
  .loader {
    margin: auto;
    width: 48px;
    height: 48px;
    border: 5px solid #fff;
    border-bottom-color: #ff3d00;
    border-radius: 50%;
    display: inline-block;
    box-sizing: border-box;
    animation: rotation 1s linear infinite;
  }

  @keyframes rotation {
    0% {
      transform: rotate(0deg);
    }
    100% {
      transform: rotate(360deg);
    }
  }

  .spinner-wrapper {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100%;
  }
</style>

{#if wrapper}
  <div class="spinner-wrapper">
    <span class="loader"></span>
  </div>
{:else}
  <span class="loader"></span>
{/if}
