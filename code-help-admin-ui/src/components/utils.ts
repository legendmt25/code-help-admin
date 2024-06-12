import type { IconType } from "svelte-icons-pack";
import type { MessageBoxType } from "./types";
import {
  AiFillCheckCircle,
  AiFillCloseCircle,
  AiFillExclamationCircle,
  AiOutlineCheckCircle,
  AiOutlineCloseCircle,
  AiOutlineExclamationCircle
} from "svelte-icons-pack/ai";

export const colors: Record<MessageBoxType, string> = {
  error: "red",
  success: "green",
  info: "blue"
};

export const messageBoxOutlineIcons: Record<MessageBoxType, IconType> = {
  info: AiOutlineExclamationCircle,
  error: AiOutlineCloseCircle,
  success: AiOutlineCheckCircle
};

export const messageBoxFillIcons: Record<MessageBoxType, IconType> = {
  info: AiFillExclamationCircle,
  error: AiFillCloseCircle,
  success: AiFillCheckCircle
};
