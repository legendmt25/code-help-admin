import env from "../env";
import { Configuration } from "../generated/admin-api";
import { authMiddleware } from "./middlewares";

export const baseConfiguration: Configuration = new Configuration({
  basePath: env.ADMIN_API_URL,
  middleware: [
    {
      pre: authMiddleware
    }
  ]
});
