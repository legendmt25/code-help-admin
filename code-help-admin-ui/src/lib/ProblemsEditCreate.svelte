<script lang="ts">
  import { javascript } from "@codemirror/lang-javascript";
  import { marked } from "marked";
  import { onMount } from "svelte";
  import CodeMirror from "svelte-codemirror-editor";
  import type { FormEventHandler } from "svelte/elements";
  import Button from "../components/Button.svelte";
  import Spinner from "../components/Spinner.svelte";
  import { Difficulty, type Category, type CategoryRequest, type ProblemRequest } from "../generated/admin-api";
  import { createProblem, getAllCategories, getProblemById, updateProblem } from "../services/ProblemsService";
  import { Icon } from "svelte-icons-pack";
  import { BiSave } from "svelte-icons-pack/bi";
  import { VscPreview } from "svelte-icons-pack/vsc";
  export let params: { id?: string } = {};

  let previewEnabled: boolean = false;
  let testCaseSelected: number | undefined = undefined;
  let editCode: "starter-code" | "runner-code" = "starter-code";

  let loading = false;
  let formValue: Partial<ProblemRequest> = {
    testCases: []
  };

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
            category: problem.category as CategoryRequest,
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

    Promise.all(promiseArray)
      .then(() => {
        formValue.category = categories.find(category => category.name === formValue.category?.name);
      })
      .finally(() => (loading = false));
  });

  const handleEditCreate: FormEventHandler<HTMLFormElement> = (event) => {
    event.preventDefault();

    const body: ProblemRequest = formValue as ProblemRequest;

    const id = Number(params.id);
    if (!Number.isNaN(id)) {
      updateProblem(id, body);
    } else {
      createProblem(body);
    }
  };
</script>

<style>
  .column {
    display: flex;
    flex-direction: column;
    height: 100%;
  }

  .row {
    display: flex;
  }

  .p-1 {
    padding: 1rem;
  }

  .h-100 {
    height: 100%;
  }

  .w-100 {
    width: 100%;
  }

  .w-50 {
    width: 50%;
  }

  .gap-1 {
    gap: 1rem;
  }

  /* Extra small devices (phones, 600px and down) */
  @media only screen and (max-width: 600px) {
    .sm-column,
    .sm-column * {
      flex-direction: column;
      width: 100% !important;
    }
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

  .tests-tab-sidebar {
    gap: 0.3rem;
    max-width: 250px;
    max-height: 300px;
    overflow-y: auto;
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

  :global(.test-case-add-btn) {
    position: sticky;
    bottom: 0;
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
  }

  :global(.side-btn:hover) {
    max-width: 200px !important;
    opacity: 1;
  }

  :global(.side-btn > span) {
    max-width: 0;
    overflow: hidden;
  }

  :global(.side-btn:hover > span) {
    max-width: 100px;
  }
</style>

{#if loading}
  <Spinner />
{:else}
  <div class="column h-100 p-1 gap-1">
    <div class="row sm-column h-100 gap-1">
      <div class="column h-100 w-50 gap-1">
        <form id="edit-problem-form" class="form" on:submit={handleEditCreate}>
          <div class="input-container">
            {#if previewEnabled}
              {formValue.title}
            {:else}
              <input id="title" name="title" placeholder="Title" bind:value={formValue.title} />
            {/if}
          </div>
          <div class="row gap-1">
            <div class="input-container">
              {#if previewEnabled}
                {formValue.difficulty}
              {:else}
                <label for="difficulty">Difficulty</label>
                <select id="difficulty" name="difficulty" bind:value={formValue.difficulty}>
                  {#each Object.values(Difficulty) as difficultyValue}
                    <option value={difficultyValue}>{difficultyValue}</option>
                  {/each}
                </select>
              {/if}
            </div>
            <div class="input-container">
              {#if previewEnabled}
                {formValue.category?.name}
              {:else}
                <label for="category">Category</label>
                <select id="category" name="category" class="input-container" bind:value={formValue.category}>
                  {#each categories as category}
                    <option value={category}>{category.name}</option>
                  {/each}
                </select>
              {/if}
            </div>
          </div>
        </form>

        <section class="h-100 w-100">
          {#if previewEnabled}
            <div class="h-100 w-100">{@html marked(formValue.markdown ?? "")}</div>
          {:else}
            <textarea
              class="h-100 w-100"
              style:resize="none"
              name="markdown"
              placeholder="Markdown"
              bind:value={formValue.markdown} />
          {/if}
        </section>
      </div>
      <hr />
      <div class="w-50 column">
        <section class="column gap-1 h-50">
          {#if !previewEnabled}
            <Button fullwidth on:click={() => (editCode = editCode === "runner-code" ? "starter-code" : "runner-code")}
              >Edit {editCode === "runner-code" ? "Runner code" : "Starter code"}</Button>
          {/if}
          {#if editCode === "runner-code" || previewEnabled}
            <CodeMirror
              bind:value={formValue.starterCode}
              readonly={previewEnabled}
              placeholder="Starter code"
              lang={javascript()} />
          {:else}
            <CodeMirror bind:value={formValue.runnerCode} placeholder="Runner code" lang={javascript()} />
          {/if}
        </section>
        {#if !previewEnabled}
          <section>
            <h2>Tests Editor</h2>
            <div class="row gap-1 w-100 h-50">
              <div class="tests-tab-sidebar column w-100">
                {#each formValue.testCases ?? [] as _, index}
                  <Button
                    fullwidth
                    on:click={() => (testCaseSelected = index)}
                    class="test-case-btn"
                    active={testCaseSelected === index}
                    ><div>Test case {index}</div>
                    <!-- svelte-ignore a11y-interactive-supports-focus -->
                    <div
                      role="button"
                      on:keydown={undefined}
                      on:click={(event) => {
                        event.preventDefault();
                        event.stopPropagation();
                        testCaseSelected = undefined;
                        formValue.testCases = [
                          ...(formValue.testCases ?? []).filter((_, arrIndex) => arrIndex !== index)
                        ];
                      }}>
                      x
                    </div></Button>
                {/each}
                <Button
                  class="test-case-add-btn"
                  fullwidth
                  on:click={() => (formValue.testCases = [...(formValue.testCases ?? []), {}])}>+</Button>
              </div>
              {#if testCaseSelected != undefined && formValue.testCases}
                <div class="input-container w-100 h-100">
                  <label for="testcase-input">INPUT:</label>
                  <textarea
                    class="h-100"
                    id="testcase-input"
                    name="testcase-input"
                    placeholder="INPUT"
                    bind:value={formValue.testCases[testCaseSelected]._in} />
                </div>
                <div class="input-container w-100 h-100">
                  <label for="testcase-output">OUTPUT:</label>
                  <textarea
                    class="h-100"
                    id="testcase-output"
                    name="testcase-input"
                    placeholder="OUTPUT"
                    bind:value={formValue.testCases[testCaseSelected].out} />
                </div>
              {/if}
            </div>
          </section>
        {/if}
      </div>
    </div>
    <Button class="side-btn" form="edit-problem-form"><Icon src={BiSave} size={24} /> <span>Save</span></Button>
    <Button class="side-btn" on:click={() => (previewEnabled = !previewEnabled)}
      ><Icon src={VscPreview} size={24} /> <span>Preview</span></Button>
  </div>
{/if}
