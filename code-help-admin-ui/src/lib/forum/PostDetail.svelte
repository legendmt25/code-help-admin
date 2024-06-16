<script lang="ts">
  import { Icon } from "svelte-icons-pack";
  import AlertBox from "../../components/AlertBox.svelte";
  import Spinner from "../../components/Spinner.svelte";
  import { getPost } from "../../services/forum/ForumService";
  import Button from "../../components/Button.svelte";
  import { AiOutlineClose } from "svelte-icons-pack/ai";
  import { formatDate } from "../../util";
  import CommentReplies from "./CommentReplies.svelte";

  export let params: { uid: string };

  const getPostPromise = getPost(params.uid);

  const handleMarkPostAsInnapropriate = () => undefined;
</script>

<style>
  .secondary-sm {
    font-size: 12px;
    color: gray;
  }

  section {
    padding: 3rem 2rem;
    display: flex;
    flex-direction: column;
    position: relative;
  }

  :global(.side-btn:nth-of-type(2)) {
    top: 150px;
  }

  :global(.side-btn) {
    position: absolute;
    right: 0;
    top: 80px;
    display: flex;
    align-items: center;

    overflow: hidden;
    transition: all 200ms;
    max-width: 70px !important;
    opacity: 0.5;
    white-space: nowrap;
  }

  :global(.side-btn:hover) {
    max-width: 250px !important;
    opacity: 1;
  }

  :global(.side-btn > span) {
    max-width: 0;
    overflow: hidden;
  }

  :global(.side-btn:hover > span) {
    max-width: 250px;
  }
</style>

{#await getPostPromise}
  <Spinner />
{:then postEntry}
  <section class="column gap-1">
    <div>
      <h2>Post detail</h2>
      <div class="secondary-sm">
        {postEntry.user.username}, {formatDate(postEntry.created)}
      </div>
    </div>

    <div>
      <h2>{postEntry.title}</h2>
      <div>{postEntry.content}</div>
    </div>
  </section>
  <section>
    <hr />

    <h2>Comments</h2>
    <CommentReplies replies={postEntry.comments} />
  </section>
  <Button maxContent class="side-btn">
    <Icon src={AiOutlineClose} size={24} on:click={handleMarkPostAsInnapropriate} />
    <span>Mark as innapropriate</span>
  </Button>
{:catch err}
  <AlertBox type="error" message="An error occured! ({err})" />
{/await}
