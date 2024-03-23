<script lang="ts">
  import { marked } from "marked";
  import { Difficulty, type EditProblemEntryRequest, type Category } from "../generated/core-api";
  import { onMount } from "svelte";
  import { getAllCategories, getProblemById } from "../services/ProblemsService";
  import Spinner from "../components/Spinner.svelte";
  import Button from "../components/Button.svelte";
  import type { FormEventHandler } from "svelte/elements";

  export let params: { id?: string } = {};

  let testCaseSelected: number | undefined = undefined;

  let testCases: {
    in?: string;
    out?: string;
  }[] = [];

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

  const handleEditCreate: FormEventHandler<HTMLFormElement> = (event) => {
    event.preventDefault();

    const trimmedTestCases = testCases.filter((testCase) => !!testCase.in && !!testCase.out);
    console.log(value, trimmedTestCases);
  };
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
    width: 50%;
    height: 98%;
  }

  .preview {
    width: 50%;
  }

  .editor-preview-container {
    display: flex;
    gap: 1rem;
    width: 100%;
    height: 97%;
  }

  .markdown-tab-section,
  .tests-tab-section,
  .solution-tab-section {
    width: 100%;
    height: 100%;
  }

  .tests-tab-sidebar {
    display: flex;
    flex-direction: column;
    gap: 0.3rem;
    width: 100%;
    max-width: 250px;
    max-height: 84.5vh;
    overflow-y: auto;
  }

  .tests-tab-container {
    width: 100%;
    display: flex;
    gap: 1rem;
  }

  .tests-tab-container .input-container {
    height: 300px;
    resize: vertical;
  }

  .input-container #testcase-output,
  .input-container #testcase-input {
    height: 100%;
    width: 100%;
  }

  .tests-editor-container {
    height: 100%;
    width: 100%;
    display: flex;
    flex-direction: column;
    gap: 1rem;
    padding: 1rem;
  }

  :global(.testcase-save-btn) {
    display: block;
    margin-left: auto;
    margin-top: auto;
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

      <form on:submit={handleEditCreate}>
        <div class="input-container">
          <label for="name">Name</label>
          <input id="name" bind:value={value.title} />
        </div>
        <div class="input-container">
          <label for="difficulty">Difficulty</label>
          <select id="difficulty" bind:value={value.difficulty}>
            {#each Object.values(Difficulty) as difficultyValue}
              <option value={difficultyValue}>{difficultyValue}</option>
            {/each}
          </select>
        </div>
        <div class="input-container">
          <label for="category">Category</label>
          <select id="category" class="input-container" bind:value={value.category}>
            {#each categories as category}
              <option value={category}>{category.name}</option>
            {/each}
          </select>
        </div>
        <Button>Create/Edit</Button>
      </form>
    </div>

    {#if activeTab === "markdown"}
      <section class="markdown-tab-section">
        <h2>Markdown Editor - Preview</h2>
        <div class="editor-preview-container">
          <textarea class="editor" style="resize: none;" name="markdown" bind:value={value.markdown}></textarea>
          <div class="preview">{@html marked(value.markdown ?? "")}</div>
        </div>
      </section>
    {:else if activeTab === "tests"}
      <section class="tests-tab-section">
        <h2>Tests Editor</h2>
        <div class="tests-tab-container">
          <div class="tests-tab-sidebar">
            <Button fullwidth onClick={() => (testCases = [{}, ...testCases])}>+</Button>
            {#each testCases as _, index}
              <Button fullwidth onClick={() => (testCaseSelected = index)}>Test case {index}</Button>
            {/each}
          </div>
          {#if testCaseSelected != undefined}
            <div class="tests-editor-container">
              <div class="input-container">
                <label for="testcase-input">INPUT:</label>
                <textarea id="testcase-input" name="testcase-input" bind:value={testCases[testCaseSelected].in}
                ></textarea>
              </div>
              <div class="input-container">
                <label for="testcase-output">OUTPUT:</label>
                <textarea id="testcase-output" name="testcase-input" bind:value={testCases[testCaseSelected].out}
                ></textarea>
              </div>
            </div>
          {/if}
        </div>
        <div></div>
      </section>
    {:else if activeTab === "solution"}
      <section class="solution-tab-section"></section>
    {/if}
  </div>
{/if}
