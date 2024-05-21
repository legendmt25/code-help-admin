import { type CreateCategoryRequest, type ProblemRequest } from "../../generated/admin-api";
import { CATEGORY_API, PROBLEM_API } from "../api";

export const getAllProblems = () => {
  return PROBLEM_API.getAllProblems();
};

export const getAllCategories = () => {
  return CATEGORY_API.getAllCategories();
};

export const getProblemById = (id: number) => {
  return PROBLEM_API.getProblem({ id });
};

export const createCategory = (body: CreateCategoryRequest) => {
  return CATEGORY_API.createCategory({ categoryRequest: body.categoryRequest });
};

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
