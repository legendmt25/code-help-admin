import { CategoryApi, ProblemApi } from "../generated/core-api";

const PROBLEM_API = new ProblemApi();
const CATEGORIES_API = new CategoryApi();

export const getAllProblems = () => {
  return PROBLEM_API.getProblemEntries();
};

export const getAllCategories = () => {
  return CATEGORIES_API.getCategories();
};

export const getProblemById = (id: number) => {
  return PROBLEM_API.getProblemEntry({ id });
};
