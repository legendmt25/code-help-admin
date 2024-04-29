import { CategoryApi, ProblemApi, type CreateCategoryRequest, type ProblemRequest } from "../generated/admin-api";
import { baseConfiguration } from "./api";

const PROBLEM_API = new ProblemApi(baseConfiguration);
const CATEGORIES_API = new CategoryApi(baseConfiguration);

export const getAllProblems = () => {
  return PROBLEM_API.getAllProblems();
};

export const getAllCategories = () => {
  return CATEGORIES_API.getAllCategories();
};

export const getProblemById = (id: number) => {
  return PROBLEM_API.getProblem({ id });
};

export const createCategory = (body: CreateCategoryRequest) => {
  return CATEGORIES_API.createCategory({ category: body.category });
};

export const updateCategory = (oldCategory: string, body: CreateCategoryRequest) => {
  return CATEGORIES_API.updateCategory({
    name: oldCategory,
    category: body.category
  });
};

export const deleteProblem = (id: number) => {
  return PROBLEM_API.deleteProblem({ id });
};

export const createProblem = (body: ProblemRequest) => {
  return PROBLEM_API.createProblem({ problemRequest: body });
};

export const updateProblem = (id: number, body: ProblemRequest) => {
  return PROBLEM_API.updateProblem({ id, problemRequest: body });
};
