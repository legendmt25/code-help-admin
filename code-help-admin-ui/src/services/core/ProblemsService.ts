import {
  type CodeRunnerRequest,
  type CreateCategoryRequest,
  type ProblemRequest
} from "../../generated/admin-core-api";
import { CATEGORY_API, CODE_RUNNER_API, PROBLEM_API } from "../api";

const DEFAULT_CODE_LANGUAGE = "javascript";

export const getAllProblems = (page?: number, size?: number, sortBy?: string[]) =>
  PROBLEM_API.getAllProblems({ page, size, sortBy });

export const getAllCategories = () => CATEGORY_API.getAllCategories();

export const getProblemById = (id: number) => PROBLEM_API.getProblem({ id });

export const createCategory = (body: CreateCategoryRequest) =>
  CATEGORY_API.createCategory({ categoryRequest: body.categoryRequest });

export const updateCategory = (id: number, body: CreateCategoryRequest) => {
  return CATEGORY_API.updateCategory({
    id,
    categoryRequest: body.categoryRequest
  });
};

export const deleteCategory = (id: number) => {
  return CATEGORY_API.deleteCategory({ id });
};

export const deleteProblem = (id: number) => {
  return PROBLEM_API.deleteProblem({ id });
};

export const createProblem = (body: ProblemRequest) => {
  return PROBLEM_API.createProblem({ problemRequest: body });
};

export const createContestProblem = (body: ProblemRequest, contestId: number) => {
  return PROBLEM_API.createProblem({ problemRequest: body, contestId });
};

export const updateProblem = (id: number, body: ProblemRequest) => {
  return PROBLEM_API.updateProblem({ id, problemRequest: body });
};

export const runCode = (body: Omit<CodeRunnerRequest, "language">) => {
  return CODE_RUNNER_API.runCode({ codeRunnerRequest: { ...body, language: DEFAULT_CODE_LANGUAGE } });
};
