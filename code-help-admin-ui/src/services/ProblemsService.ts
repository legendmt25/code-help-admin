import { ProblemApi } from "../generated/core-api";

const PROBLEM_API = new ProblemApi();

export const getAllProblems = () => {
  return PROBLEM_API.getProblemEntries();
};

