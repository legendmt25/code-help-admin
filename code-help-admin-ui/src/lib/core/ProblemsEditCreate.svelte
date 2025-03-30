<script lang="ts">
  import { javascript } from "@codemirror/lang-javascript";
  import { java } from "@codemirror/lang-java";
  import { marked } from "marked";
  import { onMount } from "svelte";
  import CodeMirror from "svelte-codemirror-editor";
  import type { FormEventHandler } from "svelte/elements";
  import Button from "../../components/Button.svelte";
  import Spinner from "../../components/Spinner.svelte";
  import {
    Difficulty,
    type Category,
    type CategoryRequest,
    type ProblemCode,
    type ProblemRequest,
    type TestCase
  } from "../../generated/admin-core-api";
  import {
    createContestProblem,
    createProblem,
    getAllCategories,
    getProblemById,
    runCode,
    updateProblem
  } from "../../services/core/ProblemsService";
  import { Icon } from "svelte-icons-pack";
  import { BiSave } from "svelte-icons-pack/bi";
  import { VscPreview } from "svelte-icons-pack/vsc";
  import { push, querystring } from "svelte-spa-router";
  import { Route } from "../../routes";
  import MessageBox from "../../components/MessageBox.svelte";
  import { AiFillCaretRight } from "svelte-icons-pack/ai";

  type ProblemFormData = Omit<Partial<ProblemRequest>, "testCases" | "runnerCode" | "starterCode"> & {
    testCases: Partial<TestCase>[];
    runnerCode: Partial<ProblemCode>[];
    starterCode: Partial<ProblemCode>[];
  };

  export let params: { id?: string; contestId?: string } = {
    contestId: new URLSearchParams($querystring).get("contestId") ?? undefined
  };

  let previewEnabled: boolean = false;
  let testCaseSelected: number | undefined = undefined;
  let codeSelected: number | undefined = undefined;
  let editCode: "starter-code" | "runner-code" = "starter-code";

  let loading = false;
  let formValue: ProblemFormData = {
    testCases: [],
    runnerCode: [],
    starterCode: []
  };

  const languageParsers = {
    javascript: javascript(),
    java: java()
  };

  $: () => {
    if (codeSelected) {
      formValue.runnerCode[codeSelected].language = formValue.starterCode[codeSelected].language;
    }
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
        formValue.category = categories.find((category) => category.name === formValue.category?.name);
      })
      .finally(() => (loading = false));
  });

  const handleEditCreate: FormEventHandler<HTMLFormElement> = (event) => {
    event.preventDefault();

    const body: ProblemRequest = formValue as ProblemRequest;

    let editCreatePromise: Promise<void>;

    const id = Number(params.id);
    const contestId = Number(params.contestId);

    if (!Number.isNaN(contestId) && Number.isNaN(id)) {
      editCreatePromise = createContestProblem(body, contestId).then(() => {});
    } else if (!Number.isNaN(id)) {
      editCreatePromise = updateProblem(id, body).then(() => {});
    } else {
      editCreatePromise = createProblem(body)
        .finally(() => push(Route.problems_overview))
        .then(() => {});
    }

    if (!Number.isNaN(contestId)) {
      editCreatePromise.finally(() => push(Route.contests_edit.replace(":id", params.contestId!)));
    } else {
      editCreatePromise.finally(() => push(Route.problems_overview));
    }
  };

  let runCodeMessage: string | undefined = undefined;
  const handleTestProblem = () => {
    if (!formValue.starterCode || !formValue.runnerCode || !formValue.testCases || !formValue.testCases.length) {
      return;
    }

    // runCode({
    //   code: formValue.starterCode!,
    //   runnerCode: formValue.runnerCode!,
    //   testCases: formValue.testCases!
    // }).then((resp) => {
    //   runCodeMessage = resp.message;

    //   setTimeout(() => {
    //     runCodeMessage = undefined;
    //   }, 15000);
    // });
  };
</script>

<style>
  /* Extra small devices (phones, 600px and down) */
  @media only screen and (max-width: 600px) {
    .sm-column,
    .sm-column * {
      flex-direction: column;
      width: 100% !important;
    }
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
    gap: 0.5rem;
    max-width: 10rem;
    height: 20rem;
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

  .floating-actions {
    position: absolute;
    width: 20rem;
    max-width: 80%;
    right: 0;
    top: 6rem;
    display: flex;
    flex-direction: column;
    align-items: end;
    gap: 2rem;
  }

  .tab-control {
    display: flex;
  }

  .button-close-icon {
    padding-inline: 0.75rem;
    /* font-weight: bold; */
    height: 100%;
    position: absolute;
    right: 0;
    top: 0;
    display: flex;
    align-items: center;
    font-size: 1.5em;
    padding-bottom: 0.1em;
  }

  .content-container {
    background: #fdfdfd;
  }
</style>

{#if loading}
  <Spinner />
{:else}
  <div class="column h-100 p-1 gap-1 content-container">
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
        {#if !previewEnabled}
          <section>
            <Button
              class="test-case-add-btn"
              fullwidth
              on:click={() => {
                formValue.runnerCode = [...(formValue.runnerCode ?? []), { language: Object.keys(languageParsers)[0] }];
                formValue.starterCode = [
                  ...(formValue.starterCode ?? []),
                  { language: Object.keys(languageParsers)[0] }
                ];
              }}>
              <span style="font-size: 1.4em; font-weight:bold; line-height: 90%">+</span>
            </Button>
            {#each formValue.runnerCode ?? [] as runnerCode, index}
              <Button
                fullwidth
                on:click={() => (codeSelected = index)}
                class="test-case-btn"
                style="position: relative;"
                active={codeSelected === index}>
                <div>
                  {runnerCode.language ?? "No language"}
                </div>
                <!-- svelte-ignore a11y-interactive-supports-focus -->
                <div
                  role="button"
                  class="button-close-icon"
                  on:keydown={undefined}
                  on:click={(event) => {
                    event.preventDefault();
                    event.stopPropagation();
                    codeSelected = undefined;
                    formValue.runnerCode = [
                      ...(formValue.runnerCode ?? []).filter((_, arrIndex) => arrIndex !== index)
                    ];
                    formValue.starterCode = [
                      ...(formValue.runnerCode ?? []).filter((_, arrIndex) => arrIndex !== index)
                    ];
                  }}>
                  ×
                </div>
              </Button>
            {/each}
          </section>
          <section>
            <h2 style="padding-bottom: 0.4rem ;">Tests Editor</h2>
            <div class="row gap-1 w-100 h-50">
              <div class="tests-tab-sidebar column h-100 w-100">
                <Button
                  class="test-case-add-btn"
                  fullwidth
                  on:click={() => (formValue.testCases = [...(formValue.testCases ?? []), {}])}>
                  <span style="font-size: 1.4em; font-weight:bold; line-height: 90%">+</span>
                </Button>
                {#each formValue.testCases ?? [] as _, index}
                  <Button
                    fullwidth
                    on:click={() => (testCaseSelected = index)}
                    class="test-case-btn"
                    style="position: relative;"
                    active={testCaseSelected === index}>
                    <div>Test case {index}</div>
                    <!-- svelte-ignore a11y-interactive-supports-focus -->
                    <div
                      role="button"
                      class="button-close-icon"
                      on:keydown={undefined}
                      on:click={(event) => {
                        event.preventDefault();
                        event.stopPropagation();
                        testCaseSelected = undefined;
                        formValue.testCases = [
                          ...(formValue.testCases ?? []).filter((_, arrIndex) => arrIndex !== index)
                        ];
                      }}>
                      ×
                    </div>
                  </Button>
                {/each}
              </div>
              {#if testCaseSelected != undefined && formValue.testCases?.length}
                <div class="input-container w-100 h-100">
                  <label for="testcase-input">INPUT:</label>
                  <textarea
                    class="h-100"
                    id="testcase-input"
                    name="testcase-input"
                    placeholder="INPUT"
                    style="resize: none; max-height: 20rem; min-height: 18rem;"
                    bind:value={formValue.testCases[testCaseSelected]._in} />
                </div>
                <div class="input-container w-100 h-100">
                  <label for="testcase-output">OUTPUT:</label>
                  <textarea
                    class="h-100"
                    id="testcase-output"
                    name="testcase-input"
                    placeholder="OUTPUT"
                    style="resize: none; max-height: 20rem; min-height: 18rem;"
                    bind:value={formValue.testCases[testCaseSelected].out} />
                </div>
              {/if}
            </div>
            {#if formValue.testCases?.length}
              <div class="column py-1 gap-0">
                {#if runCodeMessage}
                  <MessageBox type="info" fullwidth>{runCodeMessage}</MessageBox>
                {/if}
              </div>
            {/if}
          </section>
        {/if}
      </div>
      <hr />
      <div class="w-50 column">
        <section class="column gap-1 h-50">
          {#if codeSelected != undefined}
            {#if !previewEnabled}
              <select
                id="language"
                name="language"
                placeholder="Language"
                bind:value={formValue.starterCode[codeSelected].language}>
                {#each Object.keys(languageParsers) as language}
                  <option value={language}>{language}</option>
                {/each}</select>
              <div class="tab-control">
                <Button
                  fullWidth
                  toggleButton
                  toggled={editCode === "starter-code"}
                  style={"border-top-right-radius: 0; border-bottom-right-radius: 0;"}
                  on:click={() => (editCode = "starter-code")}>Edit {"Starter code"}</Button>
                <Button
                  fullWidth
                  toggleButton
                  toggled={editCode === "runner-code"}
                  on:click={() => (editCode = "runner-code")}
                  style={"border-top-left-radius: 0; border-bottom-left-radius: 0;"}>Edit {"Runner code"}</Button>
              </div>
            {:else}
              {formValue.starterCode[codeSelected].language}
            {/if}
            {#if formValue.starterCode[codeSelected].language && formValue.runnerCode[codeSelected].language}
              {#if editCode === "starter-code" || previewEnabled}
                <div style="background-color: white; height: 100%">
                  <CodeMirror
                    bind:value={formValue.starterCode[codeSelected].code}
                    readonly={previewEnabled}
                    placeholder="Starter code"
                    lang={languageParsers[formValue.starterCode[codeSelected].language]} />
                </div>
              {:else}
                <div style="background-color: white; height: 100%">
                  <CodeMirror
                    bind:value={formValue.runnerCode[codeSelected].code}
                    placeholder="Runner code"
                    lang={languageParsers[formValue.runnerCode[codeSelected].language]} />
                </div>
              {/if}
            {/if}
          {/if}
        </section>
      </div>
    </div>

    <div class="floating-actions">
      <Button class="side-btn" form="edit-problem-form">
        <Icon src={BiSave} size={24} />
        <span>Save</span>
      </Button>
      <Button class="side-btn" on:click={handleTestProblem}>
        <Icon src={AiFillCaretRight} size={24} />
        <span>Test problem</span>
      </Button>
      <Button class="side-btn" on:click={() => (previewEnabled = !previewEnabled)}>
        <Icon src={VscPreview} size={24} />
        <span>Preview</span></Button>
    </div>
  </div>
{/if}
