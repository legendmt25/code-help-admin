<script lang="ts">
  import { javascript } from "@codemirror/lang-javascript";
  import { marked } from "marked";
  import { onMount } from "svelte";
  import CodeMirror from "svelte-codemirror-editor";
  import type { FormEventHandler } from "svelte/elements";
  import Button from "../components/Button.svelte";
  import Spinner from "../components/Spinner.svelte";
  import { Difficulty, type Category, type ProblemRequest } from "../generated/admin-api";
  import { createProblem, getAllCategories, getProblemById, updateProblem } from "../services/ProblemsService";
  export let params: { id?: string } = {};

  let testCaseSelected: number | undefined = undefined;

  let testCases: {
    in?: string;
    out?: string;
  }[] = [];

  let starterCode: string = "";
  let runnerCode: string = "";

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
  let formValue: Partial<ProblemRequest> = {};
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
          formValue = {
            category: problem.category,
            difficulty: problem.difficulty,
            markdown: problem.markdown,
            title: problem.title,
            starterCode: problem.starterCode,
            runnerCode: problem.runnerCode,
            testCases: problem.testCases
          };
        })
      );
    }

    Promise.all(promiseArray).finally(() => (loading = false));
  });

  const handleEditCreate: FormEventHandler<HTMLFormElement> = (event) => {
    event.preventDefault();

    const body = formValue as ProblemRequest;

    const id = Number(params.id);
    if (!Number.isNaN(id)) {
      updateProblem(id, body);
    } else {
      createProblem(body);
    }
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

  :global(.test-case-btn) {
    display: flex;
    justify-content: space-between;
  }
</style>

{#if loading}
  <Spinner />
{:else}
  <div class="container">
    <div class="form-editor-preview-wrapper">
      <div class="tab-group">
        {#each tabs as tabItem}
          <div
            class="tab"
            class:active-tab={activeTab === tabItem.key}
            tabindex="-1"
            role="button"
            on:keypress={undefined}
            on:click={() => (activeTab = tabItem.key)}>
            {tabItem.label}
          </div>
        {/each}
      </div>

      <form on:submit={handleEditCreate}>
        <div class="input-container">
          <label for="title">Title</label>
          <input id="title" name="title" bind:value={formValue.title} />
        </div>
        <div class="input-container">
          <label for="difficulty">Difficulty</label>
          <select id="difficulty" name="difficulty" bind:value={formValue.difficulty}>
            {#each Object.values(Difficulty) as difficultyValue}
              <option value={difficultyValue}>{difficultyValue}</option>
            {/each}
          </select>
        </div>
        <div class="input-container">
          <label for="category">Category</label>
          <select id="category" name="category" class="input-container" bind:value={formValue.category}>
            {#each categories as category}
              <option value={category}>{category.name}</option>
            {/each}
          </select>
        </div>
        <Button fullwidth>Create/Edit</Button>
      </form>
    </div>

    {#if activeTab === "markdown"}
      <section class="markdown-tab-section">
        <h2>Markdown Editor - Preview</h2>
        <div class="editor-preview-container">
          <textarea class="editor" style:resize="none" name="markdown" bind:value={formValue.markdown} />
          <div class="preview">{@html marked(formValue.markdown ?? "")}</div>
        </div>
      </section>
    {:else if activeTab === "tests"}
      <section class="tests-tab-section">
        <h2>Tests Editor</h2>
        <div class="tests-tab-container">
          <div class="tests-tab-sidebar">
            <Button fullwidth on:click={() => (testCases = [{}, ...testCases])}>+</Button>
            {#each testCases as _, index}
              <Button fullwidth on:click={() => (testCaseSelected = index)} class="test-case-btn"
                ><div>Test case {index}</div>
                <div
                  role="button"
                  on:keydown={undefined}
                  on:click={(event) => {
                    event.preventDefault();
                    event.stopPropagation();
                    testCaseSelected = undefined;
                    testCases = [...testCases.filter((_, arrIndex) => arrIndex !== index)];
                  }}>
                  x
                </div></Button>
            {/each}
          </div>
          {#if testCaseSelected != undefined}
            <div class="tests-editor-container">
              <div class="input-container">
                <label for="testcase-input">INPUT:</label>
                <textarea id="testcase-input" name="testcase-input" bind:value={testCases[testCaseSelected].in} />
              </div>
              <div class="input-container">
                <label for="testcase-output">OUTPUT:</label>
                <textarea id="testcase-output" name="testcase-input" bind:value={testCases[testCaseSelected].out} />
              </div>
            </div>
          {/if}
        </div>
        <div></div>
      </section>
    {:else if activeTab === "solution"}
      <section class="solution-tab-section">
        <h2>Starter code</h2>
        <CodeMirror bind:value={starterCode} lang={javascript()} />
        <h2>Runner code</h2>
        <CodeMirror bind:value={runnerCode} lang={javascript()} />
      </section>
    {/if}
  </div>
{/if}
