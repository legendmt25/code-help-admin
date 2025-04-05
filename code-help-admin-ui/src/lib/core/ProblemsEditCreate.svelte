<script lang="ts">
  import { javascript } from "@codemirror/lang-javascript";
  import { java } from "@codemirror/lang-java";
  import { marked } from "marked";
  import {type ComponentProps, onMount} from "svelte";
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
    type TestCase, type ProblemCodeResponse
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
  import { nonNullAssert } from "../../util";

  type ProblemFormData = Omit<Partial<ProblemRequest>, "testCases" | "codes"> & {
    testCases: Partial<TestCase>[];
    codes: Partial<ProblemCodeResponse>[];
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
    codes: [],
  };

  const languageParsers: Record<string, ComponentProps<CodeMirror>['lang']> = {
    javascript: javascript(),
    java: java()
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
            codes: problem.codes,
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
    if (Number.isNaN(codeSelected) || !formValue.codes.length || !formValue.testCases.length) {
      return;
    }

    const problemCode = formValue.codes[codeSelected!];

    runCode({
      code: problemCode.starter! + "\n" + problemCode.runner!,
      language: problemCode.language!,
      testCases: formValue.testCases as TestCase[]
    }).then((resp) => {
      runCodeMessage = resp.message;

      setTimeout(() => {
        runCodeMessage = undefined;
      }, 15000);
    });
  };

  const findNotInitializedLanguages = (codes: typeof formValue['codes'], codeSelected?: number) => {
    const languages = Object.keys(languageParsers).filter(key => !codes.find(code => code.language === key));

    if(Number.isInteger(codeSelected)) {
      languages.push(codes[codeSelected!].language!)
    }

    return languages;
  }
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
                const notInitializedLanguages = findNotInitializedLanguages(formValue.codes, codeSelected);
                if(Number.isInteger(codeSelected) && notInitializedLanguages.length <= 1 || !notInitializedLanguages.length) {
                  return;
                }

                formValue.codes = [...(formValue.codes ?? []), { language: notInitializedLanguages[0] }];
              }}>
              <span style="font-size: 1.4em; font-weight:bold; line-height: 90%">+</span>
            </Button>
            {#each formValue.codes ?? [] as code, index}
              <Button
                fullwidth
                on:click={() => (codeSelected = index)}
                class="test-case-btn"
                style="position: relative;"
                active={codeSelected === index}>
                <div>
                  {code.language ?? "No language"}
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
                    formValue.codes = [
                      ...(formValue.codes ?? []).filter((_, arrIndex) => arrIndex !== index)
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
              {#if Number.isInteger(testCaseSelected) && formValue.testCases.length}
                <div class="input-container w-100 h-100">
                  <label for="testcase-input">INPUT:</label>
                  <textarea
                    class="h-100"
                    id="testcase-input"
                    name="testcase-input"
                    placeholder="INPUT"
                    style="resize: none; max-height: 20rem; min-height: 18rem;"
                    bind:value={formValue.testCases[nonNullAssert(testCaseSelected)]._in} />
                </div>
                <div class="input-container w-100 h-100">
                  <label for="testcase-output">OUTPUT:</label>
                  <textarea
                    class="h-100"
                    id="testcase-output"
                    name="testcase-input"
                    placeholder="OUTPUT"
                    style="resize: none; max-height: 20rem; min-height: 18rem;"
                    bind:value={formValue.testCases[nonNullAssert(testCaseSelected)].out} />
                </div>
              {/if}
            </div>
            {#if formValue.testCases.length}
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
          {#if Number.isInteger(codeSelected)}
            {#if !previewEnabled}
              <select
                id="language"
                name="language"
                placeholder="Language"
                bind:value={formValue.codes[Number(codeSelected)].language}>
                {#each findNotInitializedLanguages(formValue.codes, codeSelected) as language}
                  <option value={language}>{language}</option>
                {/each}
              </select>
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
              {formValue.codes[Number(codeSelected)].language}
            {/if}
            {#if formValue.codes[Number(codeSelected)]?.language}
              {#if editCode === "starter-code" || previewEnabled}
                <div style="background-color: white; height: 100%">
                  <CodeMirror
                    bind:value={formValue.codes[Number(codeSelected)].starter}
                    readonly={previewEnabled}
                    placeholder="Starter code"
                    lang={languageParsers[nonNullAssert(formValue.codes[Number(codeSelected)].language)]} />
                </div>
              {:else}
                <div style="background-color: white; height: 100%">
                  <CodeMirror
                    bind:value={formValue.codes[Number(codeSelected)].runner}
                    placeholder="Runner code"
                    lang={languageParsers[nonNullAssert(formValue.codes[Number(codeSelected)].language)]} />
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
