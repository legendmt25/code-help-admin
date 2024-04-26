import { CategoryApi, ProblemApi, type ProblemRequest } from "../generated/admin-api";
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

export const createProblem = (body: ProblemRequest) => {
  return PROBLEM_API.createProblem({ problemRequest: body });
};

export const updateProblem = (id: number, body: ProblemRequest) => {
  return PROBLEM_API.updateProblem({ id, problemRequest: body });
};
