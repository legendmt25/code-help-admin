<script lang="ts">
  import { marked } from "marked";
  import { Difficulty, type EditProblemEntryRequest, type Category } from "../generated/core-api";
  import { onMount } from "svelte";
  import { getAllCategories, getProblemById } from "../services/ProblemsService";
  import Spinner from "../components/Spinner.svelte";

  export let params: { id?: string } = {};

  let activeTab = "markdown";
  const tabs = [
    {
      key: "markdown",
      label: "Markdown"
    },
    {
      key: "tests",
      label: "Tests"
    },
    {
      key: "solution",
      label: "Solution"
    }
  ];

  let loading = false;
  let value: Partial<EditProblemEntryRequest> = {};
  let categories: Category[] = [];
  onMount(() => {
    loading = true;

    const promiseArray: Promise<void>[] = [];

    promiseArray.push(
      getAllCategories().then((resp) => {
        categories = resp.categories;
      })
    );

    if (!Number.isNaN(Number(params.id))) {
      promiseArray.push(
        getProblemById(Number(params.id)).then((problem) => {
          value = {
            category: problem.category,
            difficulty: problem.difficulty,
            markdown: problem.markdown,
            id: problem.id,
            title: problem.title,
            starterCode: new Blob([problem.starterCode], { type: "plain/text" })
          };
        })
      );
    }

    Promise.all(promiseArray).finally(() => (loading = false));
  });
</script>

<style>
  .tab-group {
    display: flex;
  }

  .tab {
    border-bottom: 3px solid transparent;
    padding: 0.5rem 0.8rem;
    text-align: center;
    min-width: 90px;
    user-select: none;

    transition: all 200ms;

    cursor: pointer;
  }

  .tab:hover:not(.active-tab) {
    border-bottom-color: #ccc;
  }

  .active-tab {
    border-bottom-color: black;
  }

  .container {
    display: flex;
    padding: 1rem;
    gap: 1rem;
    height: 100%;
  }

  input,
  select {
    width: 100%;
    max-width: 300px;
    padding: 5px;
    font-size: 1.05rem;
    outline: none;
  }

  .input-container {
    display: flex;
    flex-direction: column;
  }

  form {
    margin-top: 1rem;
    display: flex;
    flex-direction: column;
    gap: 0.5rem;
  }

  .editor {
    height: 100%;
    width: 50%;
  }

  .preview {
    width: 50%;
  }

  .editor-preview-container {
    display: flex;
    gap: 1rem;
    width: 100%;
    height: 100%;
  }

  .editor-preview-section {
    display: flex;
    flex-direction: column;
    gap: 1rem;
    width: 100%;
  }
</style>

{#if loading}
  <Spinner />
{:else}
  <div class="container">
    <div class="form-editor-preview-wrapper">
      <div class="tab-group">
        {#each tabs as tabItem}
          <div class="tab" class:active-tab={activeTab === tabItem.key} on:click={() => (activeTab = tabItem.key)}>
            {tabItem.label}
          </div>
        {/each}
      </div>

      <form>
        <div class="input-container">
          <label for="name">Name</label>
          <input id="name" bind:value={value.title} />
        </div>
        <div>
          <label for="difficulty">Difficulty</label>
          <select id="difficulty" class="input-container" bind:value={value.difficulty}>
            {#each Object.values(Difficulty) as difficultyValue}
              <option value={difficultyValue}>{difficultyValue}</option>
            {/each}
          </select>
        </div>
        <div>
          <label for="category">Category</label>
          <select id="category" class="input-container" bind:value={value.category}>
            {#each categories as category}
              <option value={category}>{category.name}</option>
            {/each}
          </select>
        </div>
      </form>
    </div>
    <section class="editor-preview-section">
      <h2>Editor - Preview</h2>
      <div class="editor-preview-container">
        <textarea class="editor" style="resize: none;" name="markdown" bind:value={value.markdown}></textarea>
        <div class="preview">{@html marked(value.markdown ?? "")}</div>
      </div>
    </section>
  </div>
{/if}
