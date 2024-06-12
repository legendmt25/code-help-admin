<script lang="ts">
  import type { Comment } from "../../generated/admin-api";
  import CommentDetail from "./CommentDetail.svelte";

  export let replies: Comment[] = [];
</script>

<style>
  .replies {
    padding-left: 1rem;
  }
</style>

{#if replies.length > 0}
  <ul class="replies">
    {#each replies as reply}
      <li><CommentDetail comment={reply} /></li>
    {/each}
  </ul>
{/if}
