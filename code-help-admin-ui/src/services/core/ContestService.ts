import { type ContestEditRequest, type ContestRequest } from "../../generated/admin-core-api";
import { CONTEST_API } from "../api";

export const getAllContests = (page?: number, size?: number, sortBy?: string[]) =>
  CONTEST_API.getAllContests({ page, size, sortBy });
export const getContestById = (id: number) => CONTEST_API.getContest({ id });
export const createContest = (body: ContestRequest) => CONTEST_API.createContest({ contestRequest: body });
export const updateContest = (id: number, body: ContestEditRequest) =>
  CONTEST_API.updateContest({ id, contestEditRequest: body });
export const deleteContest = (id: number) => CONTEST_API.deleteContest({ id });
