import "./app.css";
import App from "./CodeHelpAdminApp.svelte";

const app = new App({
  target: document.getElementById("app") as Element
});

export default app;
