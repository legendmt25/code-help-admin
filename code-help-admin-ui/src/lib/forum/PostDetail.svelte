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

  .floating-actions {
    position: absolute;
    width: 100%;
    right: 0;
    top: 6rem;
    display: flex;
    flex-direction: column;
    align-items: end;
    gap: 2rem;
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
  <div class="floating-actions">
    <Button maxContent class="side-btn">
      <Icon src={AiOutlineClose} size={24} on:click={handleMarkPostAsInnapropriate} />
      <span>Mark as innapropriate</span>
    </Button>
  </div>
{:catch err}
  <AlertBox type="error" message="An error occured! ({err})" />
{/await}
