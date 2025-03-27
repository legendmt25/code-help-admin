<script lang="ts">
  import Button from "../../components/Button.svelte";
  import Spinner from "../../components/Spinner.svelte";
  import type { Comment } from "../../generated/admin-forum-api";
  import { deleteComment, getCommentReplies } from "../../services/forum/ForumService";
  import { formatDate } from "../../util";
  import CommentReplies from "./CommentReplies.svelte";

  export let comment: Comment;
  export let replies: Comment[] | undefined = undefined;
  export let loading: boolean = false;

  const handleDeleteComment = (uid: string) => deleteComment(uid);
  const handleMarkCommentAsInnapropriate = (uid: string) => undefined;

  const handleShowMore = () => {
    loading = true;
    getCommentReplies(comment.uid)
      .then((resp) => (replies = resp.comments ?? []))
      .finally(() => (loading = false));
  };
</script>

<style>
  .comment {
    padding-top: 0.5rem;
  }
</style>

<div class="comment">
  <span>{comment.user.username}</span>
  <span class="secondary-sm">({formatDate(comment.created)})</span>
  <span>commented:</span>
</div>
<div>{comment.content}</div>

<div>
  <Button type="link" on:click={() => handleDeleteComment(comment.uid)}>Delete</Button>
  <span>|</span>
  <Button type="link" on:click={() => handleMarkCommentAsInnapropriate(comment.uid)}>Mark as innapropriate</Button>
  {#if !replies}
    <span>|</span>
    <Button type="link" on:click={handleShowMore}>Show more</Button>
  {:else if loading}
    <Spinner />
  {/if}
</div>

{#if replies}
  <CommentReplies {replies} />
{/if}
