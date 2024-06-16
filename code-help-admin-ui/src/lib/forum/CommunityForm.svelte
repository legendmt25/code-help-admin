<script lang="ts">
  import type { FormEventHandler } from "svelte/elements";
  import Button from "../../components/Button.svelte";
  import MessageBox from "../../components/MessageBox.svelte";
  import { FormSubmitStatus } from "../../types";
  import type { Community, CommunityRequest, ForumCategory } from "../../generated/admin-api";
  import { createCommunity, updateCommunity } from "../../services/forum/ForumService";
  import { createEventDispatcher, onDestroy } from "svelte";

  const dispatch = createEventDispatcher<{
    submit: {
      name: string;
      edit: boolean;
    };
  }>();

  export let name: string | undefined = undefined;
  export let value: Partial<CommunityRequest> & {
    categories: { uids: string[] };
  } = {
    categories: {
      uids: []
    }
  };
  export let categories: ForumCategory[] = [];

  let formSubmitStatus: FormSubmitStatus | undefined;
  let clearFormSubmitStatusTimeout: (() => void) | undefined = undefined;

  const handleFormSubmit: FormEventHandler<HTMLFormElement> = (event) => {
    event.preventDefault();

    formSubmitStatus = FormSubmitStatus.SUBMITTING;
    let editCreatePromise: Promise<Community>;

    if (name) {
      editCreatePromise = updateCommunity(name, value as CommunityRequest);
    } else {
      editCreatePromise = createCommunity(value as CommunityRequest);
    }

    editCreatePromise
      .then(() => (formSubmitStatus = FormSubmitStatus.SUCCESS))
      .then(() => dispatch("submit", { name: value.name!, edit: Boolean(name) }))
      .catch(() => (formSubmitStatus = FormSubmitStatus.ERROR));

    clearFormSubmitStatusTimeout?.();
    const formSubmitTimeout = setTimeout(() => {
      formSubmitStatus = undefined;
    }, 2000);

    clearFormSubmitStatusTimeout = () => {
      clearTimeout(formSubmitTimeout);
    };
  };

  onDestroy(() => {
    clearFormSubmitStatusTimeout?.();
  });
</script>

<style>
  .form-modifier {
    display: grid;
    gap: 2rem;
  }

  .right {
    display: flex;
    justify-content: end;
  }

  .input-container > select > option {
    padding: 0.5rem 1rem;
    color: black;
    font-weight: bold;
    margin-bottom: 0.1rem;
    font-size: 1.2em;
  }

  select[multiple]:focus option:checked {
    background: #dfe1f5 linear-gradient(0deg, #dfe1f5 0%, #dfe1f5 100%);
    background-color: #dfe1f5;
    color: black;
    -webkit-text-fill-color: black;
  }

  select[multiple]:not(:focus) option:checked {
    background-color: #dfe1f5;
    color: black;
  }

  label {
    font-size: 0.8em;
    font-weight: bold;
  }

  form {
    display: flex;
    flex-direction: column;
    gap: 0.5rem;
  }
</style>

<div class="form-container">
  <form on:submit={handleFormSubmit} class="form form-modifier">
    <div class="input-container">
      <label for="name">Name</label>
      <input id="name" name="name" bind:value={value.name} />
    </div>
    <div class="input-container">
      <label for="description">Description</label>
      <input id="description" name="description" bind:value={value.description} />
    </div>
    <div class="input-container">
      <label for="categories">Categories:</label>
      <select
        name="categories"
        multiple
        on:change={(event) =>
          (value.categories.uids = Array.from(event.currentTarget.selectedOptions).map((x) => x.value))}>
        {#each categories as category}
          <option value={category.uid} selected={!!value.categories.uids.find((uid) => category.uid === uid)}>
            {category.name}
          </option>
        {/each}
      </select>
    </div>
    <div class="right">
      <Button maxContent>Submit</Button>
    </div>
    {#if formSubmitStatus === FormSubmitStatus.SUCCESS}
      <MessageBox type={FormSubmitStatus.SUCCESS}>Saved successfully!</MessageBox>
    {:else if formSubmitStatus === FormSubmitStatus.ERROR}
      <MessageBox type={FormSubmitStatus.ERROR}>Error happened while saving</MessageBox>
    {/if}
  </form>
</div>
